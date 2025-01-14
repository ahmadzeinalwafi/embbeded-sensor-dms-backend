package repository

import (
	"context"
	"fmt"
	repository "golang_api/internal/domain/historical_device_records"
	entity "golang_api/internal/domain/historical_device_records/entities"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type HistoricalDeviceRecordsRepositoryImpl struct {
	client       influxdb2.Client
	writeAPI     api.WriteAPIBlocking
	queryAPI     api.QueryAPI
	bucket       string
	organization string
}

func NewHistoricalDeviceRecordsRepository(client influxdb2.Client, bucket, organization string) repository.DeviceConfigRepository {
	return &HistoricalDeviceRecordsRepositoryImpl{
		client:       client,
		writeAPI:     client.WriteAPIBlocking(organization, bucket),
		queryAPI:     client.QueryAPI(organization),
		bucket:       bucket,
		organization: organization,
	}
}

func (r *HistoricalDeviceRecordsRepositoryImpl) CreateMeasurement(ctx context.Context, config *entity.DeviceConfigRecords) error {
	// Prepare initial data with the correct zero values based on the field types
	initialFields := make(map[string]interface{})
	for field, fieldType := range config.Fields {
		switch fieldType {
		case "float64":
			initialFields[field] = float64(0) // zero value for float64
		case "int8", "int16", "int32", "int64":
			initialFields[field] = int64(0) // zero value for integer types (use int64 for InfluxDB)
		default:
			return fmt.Errorf("unsupported field type for %s", field)
		}
	}

	// Write initial data to InfluxDB
	p := influxdb2.NewPoint(config.DeviceID,
		map[string]string{"device_id": config.DeviceID},
		initialFields,
		time.Now()) // Set the current timestamp

	// Write the point to InfluxDB
	err := r.writeAPI.WritePoint(ctx, p)
	if err != nil {
		return fmt.Errorf("failed to create measurement: %v", err)
	}
	return nil
}

func (r *HistoricalDeviceRecordsRepositoryImpl) WriteData(ctx context.Context, config *entity.DeviceConfigRecords) error {
	// Write data to the measurement
	p := influxdb2.NewPoint(config.DeviceID,
		map[string]string{"device_id": config.DeviceID},
		config.Fields,
		time.Now())
	for field, value := range config.Fields {
		p.AddField(field, value)
	}

	err := r.writeAPI.WritePoint(ctx, p)
	if err != nil {
		return fmt.Errorf("failed to write data: %v", err)
	}
	return nil
}

func (r *HistoricalDeviceRecordsRepositoryImpl) ReadData(ctx context.Context, deviceID string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -30d)
		|> filter(fn: (r) => r._measurement == "%s")
	`, r.bucket, deviceID)

	result, err := r.queryAPI.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %v", err)
	}

	data := []map[string]interface{}{}
	for result.Next() {
		record := result.Record()
		data = append(data, map[string]interface{}{
			"time":   record.Time(),
			"field":  record.Field(),
			"value":  record.Value(),
			"device": record.Measurement(),
		})
	}

	if result.Err() != nil {
		return nil, fmt.Errorf("query error: %v", result.Err())
	}
	return data, nil
}

func (r *HistoricalDeviceRecordsRepositoryImpl) DeleteMeasurement(ctx context.Context, deviceID string) error {
	// InfluxDB does not support direct measurement deletion via API.
	// The alternative is to delete all data from the bucket matching the measurement.
	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -inf)
		|> filter(fn: (r) => r._measurement == "%s")
		|> drop(columns: ["_value", "_field", "_measurement"])
	`, r.bucket, deviceID)

	_, err := r.queryAPI.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to delete measurement: %v", err)
	}
	return nil
}
