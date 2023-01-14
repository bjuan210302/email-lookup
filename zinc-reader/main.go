package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"zincreader/model"
	"zincreader/wrapper"
)

func main() {
	rootPath, bulkSize, maxMailsToProcess := parseArgs()
	mailsPaths := collectMailsPaths(rootPath, maxMailsToProcess)
	wrapper.CheckAndCreateIndex()
	ProcessFilesBatch(mailsPaths, bulkSize)
}

func parseArgs() (string, int, int) {
	if len(os.Args) < 3 {
		panic("Two arguments required: root of mails (path) and bulk size (int)")
	}

	rootPath := os.Args[1]
	pathInfo, err := os.Stat(rootPath)
	if err != nil || !pathInfo.IsDir() {
		if err != nil {
			panic(fmt.Sprintf("Error while reading directory %s: %v", rootPath, err))
		}
		panic("Error: " + rootPath + " not is a directory.")
	}

	bulkSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(fmt.Sprintf("Error parsing second arg %d: %v", bulkSize, err))
	}

	// Optional args
	var maxMailsToProcess int
	if len(os.Args) > 3 {
		maxMailsToProcess, err = strconv.Atoi(os.Args[3])
	}
	if err != nil {
		maxMailsToProcess = -1
	}

	return rootPath, bulkSize, maxMailsToProcess
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
		log.Fatalf("Error reading file with path: %v", path)
		log.Fatalln(err)
		return model.Email{}, err
	}
	r := strings.NewReader(string(fileContent))
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatalf("Error reading file with path: %v", path)
		log.Fatalln(err)
		return model.Email{}, err
	}
	header := m.Header
	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatalf("Error reading body of email with path: %v", path)
		log.Fatalln(err)
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
