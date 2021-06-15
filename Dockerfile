
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

RUN rm .env
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app
RUN rm -rf /app/src

FROM os
WORKDIR /app
COPY --from=build /app .
#ENTRYPOINT ["./api"]

EXPOSE 8080
RUN echo $(ls -R ./main)