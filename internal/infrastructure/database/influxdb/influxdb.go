package influxdb

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InfluxDBConnector() (influxdb2.Client, error) {

	// dbToken := os.Getenv("INFLUXDB_TOKEN")
	// if dbToken == "" {
	//     return nil, errors.New("INFLUXDB_TOKEN must be set")
	// }

	// dbURL := os.Getenv("INFLUXDB_URL")
	// if dbURL == "" {
	//     return nil, errors.New("INFLUXDB_URL must be set")
	// }

	client := influxdb2.NewClient("http://localhost:8086", "QfAga1dLkl5nciEq8kmA1wqNImfi48FYNpncgYic7LSg1a-JyeR-2pmIwy_8_u_k06b2x5RmcdyzOy3H_3FSog==")

	// validate client connection health
	_, err := client.Health(context.Background())

	return client, err
}
