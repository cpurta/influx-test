FROM golang:latest

MAINTAINER Chris Purta <cpurta@gmail.com>

RUN apt-get update && \
    mkdir /app

ADD . /app

WORKDIR /app

RUN go get github.com/influxdata/influxdb/client/v2 && \
    go build -o influx-test .

ENTRYPOINT ["/app/influx-test"]
