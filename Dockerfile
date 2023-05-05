FROM golang:1.19-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

ARG DEFAULT_PORT=8080
EXPOSE 8080

RUN go build *.go
CMD ["/app/main"]
