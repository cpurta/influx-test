#!/bin/bash

go version 1>/dev/null

if [ $? -ne 0 ]; then
    echo 'Go is not install or your $PATH does not contain the go executable, please install or set up correct'
    exit 1
fi

deps=(github.com/influxdata/influxdb/client/v2)

for dep in $deps ; do
    echo "installing $dep..."
    go get $dep
done

echo "building influxtest..."

go build -o influx-test .

echo "build influxtest docker image"
docker build -t influxtest:latest .

echo "Done!"
