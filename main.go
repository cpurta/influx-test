package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	db       = "test_database"
	username = "admin"
	password = "Adm1nPa$$W0rd"
)

func main() {
	addr := os.Getenv("INFLUX_PORT_8086_TCP_ADDR")
	port := os.Getenv("INFLUX_PORT_8086_TCP_PORT")

	fmt.Printf("Influx url: http://%s:%s\n", addr, port)

	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%s", addr, port),
		Username: username,
		Password: password,
	})

	makeDatabase(c)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	for {
		// Create a point and add to batch
		tags := map[string]string{"cpu": "cpu-total"}
		fields := map[string]interface{}{
			"idle":   rand.Float64() * 100,
			"system": rand.Float64() * 100,
			"user":   rand.Float64() * 100,
		}
		pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())

		if err != nil {
			log.Fatalln("Error: ", err)
		}

		bp.AddPoint(pt)

		// Write the batch
		c.Write(bp)
	}
}

func makeDatabase(c client.Client) {
	_, err := queryDB(c, fmt.Sprintf("CREATE DATABASE %s", db))
	if err != nil {
		log.Fatal(err)
	}
}

// queryDB convenience function to query the database
func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: db,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
