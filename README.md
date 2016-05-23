# influx-test

Simple project that will spin up an instance of InfluxDB, Grafana to send "CPU Data" to.
Just wanted to see if this was possible and to kill some time. 

## Pre-requisites

- [Golang](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/engine/installation/)
- [Compose](https://docs.docker.com/compose/install/)

To set everything up to run it is as simple as executing the bootstrap:

```
$ ./bootstrap.sh
```

This will make sure that you have Go installed, build the executable and build the influx-test
docker image that is needed for compose to link everyhting together and run properly.

## Running the program

Simple as running a command:

```
$ docker-compose up
```

### Troubleshooting

Make sure that the docker daemon is running.

Ensure that your $GOPATH is set up correctly or that is project was cloned in a place that
your $GOPATH can reach.
