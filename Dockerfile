# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

# Go compile main.go with the output name: go-docker-go
RUN go build -o /go-docker-go

CMD [ "/go-docker-go" ]