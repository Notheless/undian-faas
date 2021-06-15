
FROM alpine:latest as os
FROM golang:1.14-alpine AS build
#RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
ENV GO111MODULE=auto

WORKDIR /app/src/api
COPY ./api/go.mod .
COPY ./api/go.sum .

RUN go mod tidy -v
RUN go mod download -x

COPY ./api/. .

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app

RUN app
EXPOSE 8080