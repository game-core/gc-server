# Build
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/gen/main.go

RUN apk update && apk add bash protobuf-dev

# Dev
FROM golang:1.21-alpine AS local

RUN export PATH="$PATH:$(go env GOPATH)/bin"
ENV PATH=$PATH:$GOPATH/bin
ENV CGO_ENABLED=0
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
ENV GO111MODULE=on
COPY . .
RUN apk upgrade --update && apk add bash && apk --no-cache add git && apk add bash protobuf-dev
