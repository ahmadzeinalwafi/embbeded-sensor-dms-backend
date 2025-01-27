package influxdb

import (
	"context"
"dms/config"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InfluxDBConnector() (influxdb2.Client, error) {
	cfg := config.LoadConfig()

	client := influxdb2.NewClient(cfg.GetString("INFLUXDB_URL"), "QfAga1dLkl5nciEq8kmA1wqNImfi48FYNpncgYic7LSg1a-JyeR-2pmIwy_8_u_k06b2x5RmcdyzOy3H_3FSog==")

	// validate client connection health
	_, err := client.Health(context.Background())

	return client, err
}
