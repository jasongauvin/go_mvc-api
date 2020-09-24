FROM golang:latest

WORKDIR /api

COPY ./ /api

RUN go mod download -x

RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -exclude-dir=.git --build="go build api/main.go" --command=./main