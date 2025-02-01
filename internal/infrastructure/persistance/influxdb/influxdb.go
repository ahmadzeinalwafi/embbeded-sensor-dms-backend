package influxdb

import (
	"context"
	"dms/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InfluxDBConnector() (influxdb2.Client, error) {
	cfg := config.LoadConfig()

	client := influxdb2.NewClient(cfg.GetString("INFLUXDB_URL"), cfg.GetString("INFLUXDB_KEY"))

	// validate client connection health
	_, err := client.Health(context.Background())

	return client, err
}
