# EMAIL LOOKUP - Parse, index, search emails from the Enron Database

This repository contains three aplications: `zinc-reader`, `e-lookup-be`, and `e-lookup-fe`. The instructions below explain how to configurate and run each one separately, however if you just want to get the projects running quickly, you can execute `justrunman.sh` to run everything with default values.

* Project depends on `Go >= 1.9` and `Node >= 16`
* You can donwload the lastest version of ZincSearch from its [repo](https://github.com/zinclabs/zinc/releases)
* And the Enron Emails Database [here](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) (423MB)

## `justrunman.sh`

For `justrunman.sh` to work, put your ZincSearch folder in the root of this project:
  ```
  email-lookup
  |- e-lookup-be
  |- e-lookup-fe
  |- zinc-reader
  |- zinc          <- The folder you download from ZincSearch
    |- LICENSE
    |- README.md
    |- zinc        <- zinc executable
  ```

Running `justrunman.sh` will set up the ZincSearch server, the backend of this app (e-lookup-be) and the frontend (e-lookup-fe),
however it will NOT index the emails nor upload them to the ZincSearch server, that has to be done manually. Indexing the emails with `zinc-reader` is explained below.

## **zinc-reader**

This app parses the raw files from the Enron Database and upload the emails to a ZincZearch index.

## Build

* Environment variables
    ```bash
    export ZINC_SEARCH_SERVER_URL="http://localhost:4080/api/"
    export ZINC_SEARCH_USER="admin"
    export ZINC_SEARCH_PASSWORD="Complexpass#123"
    export ZINC_SEARCH_INDEX_NAME="enron-index"
    ```
* Get dependencies
    ```bash
    go mod download
    ```

* Build app
    ```bash
    go build -o zincreader
    ```

## Usage

```bash
  #Example
  ./zincreader --data_path "data/enron_mail_20110402" --bulk_size 500 --max_mails 15000
  ```

* `data_path` Is the root directory that contains all the emails. Default is `data/`
* `bulk_size` Is the number of emails uploaded to the ZincSearch server in a single request. Default and recomended is `1000`
* `max_mails` Is the total number of emails to be uploaded. Do not pass this argument if you wish to load all the emails.

*Make sure your ZincSearch server is running and the credentials provided match.*

## **e-lookup-be**

Small backend to wrap the ZincSearch API.

## Build

Uses the same variables of `zinc-reader`
* As `main.go` is not on root folder
    ```bash
    #pwd e-lookup-be
    cd web
    go build -o elookup
    ```

## Usage

To get up and running use `./elookup --port {PORT}`. Port defaults to 3000 if not specified.

You can check if the server is running with
  ```bash
  curl --location --request GET 'http://localhost:3000/api/v1/ping'
  ```
    
Query indexed emails on the ZincSearch server and get paginated resutls with
  ```bash
  curl --location --request GET 'http://localhost:3000/api/v1/lookup?word=money&page=2&max_per_page=10'
  ```

*Make sure your ZincSearch server is running and the credentials provided match.*

## **e-lookup-fe**

A visualizer for the API wrapper. Build with Vue 3 and Tailwind. Uses Vite 3.

## Run

Provide `.env` with backend link
  ```
  VITE_ELOOKUP_BACKEND_QUERY_URL=http://localhost:3000/api/v1/lookup?
  ```

Run dev server
  ```bash
  npm i
  npm run dev
  ```