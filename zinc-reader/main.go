package main

import (
	"flag"
	"io"
	"log"
	"net/mail"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strings"

	"zincreader/model"
	"zincreader/wrapper"
)

var (
	dataRootPath      *string
	bulkSize          *int
	maxMailsToProcess *int
	profile           *bool
)

func init() {
	dataRootPath = flag.String("data_path", "data", "Root of emails to load")
	bulkSize = flag.Int("bulk_size", 1000, "Size of ZincSearch upload bulk")
	maxMailsToProcess = flag.Int("max_mails", -1, "Limit of emails to upload")
	profile = flag.Bool("profile", false, "write cpu profile to `file`")
	flag.Parse()
}

func main() {

	if *profile {
		log.Print("Profiling enabled. Starting...")
		file, err := os.Create("cpuprofile")
		if err != nil {
			log.Fatal("could not create file cpuprofile: ", err)
		}
		defer file.Close()

		err = pprof.StartCPUProfile(file)
		if err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}

	}

	mailsPaths := collectMailsPaths(*dataRootPath, *maxMailsToProcess)
	wrapper.CheckAndCreateIndex()
	ProcessFilesBatch(mailsPaths, *bulkSize)

	if *profile {
		log.Print("Stopping Profiling...")
		defer pprof.StopCPUProfile()
	}
}

func collectMailsPaths(rootPath string, maxMailsToProcess int) []string {

	log.Print("Starting collection of mails paths")
	mailsPaths := []string{}
	mailsPathsLen := 0
	err := filepath.Walk(rootPath,
		func(path string, fileInfo os.FileInfo, err error) error {
			// Use only files with no extension
			if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == "" {
				mailsPaths = append(mailsPaths, path)
				mailsPathsLen++
				if mailsPathsLen == maxMailsToProcess {
					log.Printf("Collect has reached length limit: %d", mailsPathsLen)
					return io.EOF
				}
			}
			return nil
		})
	if err != nil && err != io.EOF {
		log.Printf("Error while collecting paths: %s", err)
	}

	log.Printf("Mails path collection finalized. Total collected: %d", len(mailsPaths))
	return mailsPaths
}

func ProcessFilesBatch(mailsPaths []string, bulkSize int) {
	total := len(mailsPaths)
	log.Printf("Preparing to proccess files.  Bulk size: %v. Total records: %v", bulkSize, total)

	var bulk []model.Email
	for i, item := range mailsPaths {
		parsedEmail, err := parseEmailFromPath(item)
		if err != nil { // Skip if error while parsing email
			continue
		}
		bulk = append(bulk, parsedEmail)

		if (i+1)%bulkSize == 0 { // Upload bulk and start over
			log.Printf("Uploading bulk %v / %v", i+1, total)
			wrapper.SaveBulk(bulk)
			bulk = nil
		} else if bulk != nil && (i+1) == total { // Upload last bulk
			log.Printf("Uploading bulk %v / %v", i+1, total)
			wrapper.SaveBulk(bulk)
		}
	}
}

func parseEmailFromPath(path string) (model.Email, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Print("Error reading file with path", path, err)
		return model.Email{}, err
	}
	r := strings.NewReader(string(fileContent))
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Print("Error parsing file with path", path, err)
		return model.Email{}, err
	}
	header := m.Header
	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Print("Error reading body of email with path", path, err)
		return model.Email{}, err
	}

	recipients := strings.Split(header.Get("To"), ", ")
	mail := model.Email{
		MessageId: header.Get("Message-ID"),
		Date:      header.Get("Date"),
		From:      header.Get("From"),
		To:        recipients,
		Subject:   header.Get("Subject"),
		Content:   string(body),
	}

	return mail, nil
}
