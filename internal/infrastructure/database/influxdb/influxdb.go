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

	client := influxdb2.NewClient("http://localhost:8086", "Pf89iGH1BikOhQ6RFEHp2SbPHtWCdG1DbIi6KKurmP9iikaLWmC79UQ6xHm40Tuxi3RFBrj19tojkoA5bO738w==")

	// validate client connection health
	_, err := client.Health(context.Background())

	return client, err
}
