FROM golang:1.8-alpine
ADD . /go/src/api
RUN go install api

FROM alpine:latest
COPY --from=0 /go/bin/api .
ENV PORT 8080
CMD ["./api"]