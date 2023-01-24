# EMAIL LOOKUP - Parse, index, search emails from the Enron Database

This repository contains three aplications: `enron-indexer`, `e-lookup-be`, and `e-lookup-fe`.

* You can donwload the lastest version of ZincSearch from its [repo](https://github.com/zinclabs/zinc/releases).
* And the Enron Emails Database [here](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) (423MB).

## Building with `docker compose up`

For `docker compose up` to work, put your ZincSearch folder inside `/zinc-server` an rename it to `zinc`.
Your file structure should look like this:
  ```
  email-lookup
  |- e-lookup-be
  |- e-lookup-fe
  |- zinc-reader
  |- zinc-server
    |- Dockerfile
    |- zinc          <- The folder you download from ZincSearch
      |- LICENSE
      |- README.md
      |- zinc        <- zinc executable
  ```

Running `docker compose up` will set up the ZincSearch server, the backend of this app (e-lookup-be) and the frontend (e-lookup-fe) each one in a docker container, binding ports `6001`, `6002` and `6003` respectively. **This will NOT index the emails nor upload them to the ZincSearch server, that has to be done manually. Indexing the emails with `enron-indexer` is explained below.**

## **enron-indexer**

This app parses the raw files from the Enron Database and upload the emails to a ZincZearch index.

### Build

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

### Usage

```bash
  #Example
  ./zincreader --data_path "data/enron_mail_20110402" --bulk_size 500 --max_mails 15000
  ```

* `data_path` Is the root directory that contains all the emails. Default is `data/`
* `bulk_size` Is the number of emails uploaded to the ZincSearch server in a single request. Default and recomended is `1000`
* `max_mails` Is the total number of emails to be uploaded. Do not pass this argument if you wish to load all the emails.
* `profile` Passing this flag will generate a `cpuprofile` file that can be readed with `go tool pprof cpuprofile` for performance info.

*Make sure your ZincSearch server is running and the credentials provided match.*

## **zinc-server**

Contains a Dockerfile for the ZincSearch binary. 

### Build

Accepts two args `user` and `pass` to define credentials.

```bash
  #pwd zinc-server
  docker build --build-arg user=employee --build-arg pass=hyperComplexpass@321 .
```
#### ZincSearch server default credentials
* `ZINC_FIRST_ADMIN_USER=admin`
* `ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123`


### Usage

The ZincSearch API is exposed on port `4080`. Bind the port as you need.

## **e-lookup-be**

Small backend to wrap part of the ZincSearch API.

### Build

Accepts one arg `zinc_host`. Defaults to `http://localhost:4080/`
```bash
  #pwd e-lookup-be
  docker build --build-arg zinc_host=http://localhost:4080/ .
```

### Usage

Endpoints are exposed on port `3000`. Bind the port as you need.

* You can check if the server is running with
  ```bash
  curl 'http://localhost:3000/api/v1/ping'
  ```

* Query indexed emails on the ZincSearch server and get paginated resutls with
  ```bash
  curl -u <zinc_username:zinc_password> 'http://localhost:3000/api/v1/lookup?word=hired&page=2&max_per_page=10&index=enron'
  ```

* Get a list of indexes in the ZincSearch server
  ```bash
  curl -u <zinc_username:zinc_password> 'http://localhost:3000/api/v1/indexes
  ```

*Make sure your ZincSearch server is running and the credentials match.*

## **e-lookup-fe**

A visualizer for the API wrapper. Build with Vue 3 and Tailwind, on top of Vite 3.

### Build

Accepts one arg `elookupbe_host`. Defaults to `http://localhost:6002/`

```bash
  #pwd e-lookup-fe
  docker build --build-arg elookupbe_host=http://localhost:6002/ .
```
