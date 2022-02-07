FROM golang:latest

WORKDIR /server

COPY ./ /server

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /server

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
