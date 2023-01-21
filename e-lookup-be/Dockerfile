FROM golang:1.19

WORKDIR "/e-lookup-be"

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /e-lookup-be/build/elookup ./web/main.go

ENV ZINC_SERVER_HOST=http://localhost:4080/
EXPOSE 3000

CMD ["elookup --port=3000"]