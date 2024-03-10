FROM golang:alpine

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go fmt ./...
RUN go vet ./...

COPY . .
RUN go mod tidy