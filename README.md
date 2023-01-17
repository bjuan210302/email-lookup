# EMAIL LOOKUP - Parse, index, search emails from the Enron Database

This repository contains three aplications: `zinc-reader`, `e-lookup-be`, and `e-lookup-fe`

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
    

## **e-lookup-be**

Small backend to wrap the ZincSearch API.

## Build

Uses the sames steps of `zinc-reader`
Note: `main.go` is located inside the `web` directory. Use `go run web/main.go` to run

## Usage

* To get up and running use `./lookupbe --port {PORT}`. Port defaults to 3000 if not specified.

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

* Provide `.env` with backend link
    ```
    VITE_BACKEND_QUERY_ENDPOINT=http://localhost:3000/api/v1/lookup?
    ```

* Run dev server
    ```bash
    npm run i
    npm run dev
    ```