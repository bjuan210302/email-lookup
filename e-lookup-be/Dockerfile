FROM golang:1.19

WORKDIR "/e-lookup-be"

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/elookup ./web/main.go

ARG zinc_host=http://localhost:4080/
ENV ZINC_SERVER_HOST=$zinc_host
EXPOSE 3000

CMD ["/usr/local/bin/elookup"]