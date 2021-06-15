
# FROM alpine:latest as os
# FROM golang:1.14-alpine AS build
# #RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
# ENV GO111MODULE=auto

# WORKDIR /app/src/api
# COPY ./api/go.mod .
# COPY ./api/go.sum .

# RUN go mod tidy -v
# RUN go mod download -x

# COPY ./api/. .

# RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app

# RUN ["app"]
# EXPOSE 8080

## We specify the base image we need for our
## go application
FROM golang:1.12.0-alpine3.9
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]