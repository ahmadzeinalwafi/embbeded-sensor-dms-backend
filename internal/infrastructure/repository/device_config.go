package repository

import (
	"context"
	"errors"
	"fmt"
	entity "golang_api/internal/domain/device_config/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB implementation of the SensorRepository
type SensorRepositoryImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewSensorRepository creates a new instance of SensorRepositoryImpl
func NewSensorRepository(client *mongo.Client, dbName, collectionName string) *SensorRepositoryImpl {
	collection := client.Database(dbName).Collection(collectionName)
	return &SensorRepositoryImpl{
		client:     client,
		collection: collection,
	}
}

// Insert saves a sensor to the database
func (r *SensorRepositoryImpl) Insert(ctx context.Context, sensor *entity.SensorConfig) error {
	_, err := r.collection.InsertOne(ctx, sensor)
	if err != nil {
		return fmt.Errorf("error inserting sensor: %v", err)
	}
	return nil
}

// FindByID retrieves a sensor by its sensorID
func (r *SensorRepositoryImpl) FindByID(ctx context.Context, sensorID string) (*entity.SensorConfig, error) {
	var sensor entity.SensorConfig
	err := r.collection.FindOne(ctx, bson.M{"sensorid": sensorID}).Decode(&sensor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Not found
		}
		return nil, fmt.Errorf("error finding sensor: %v", err)
	}
	return &sensor, nil
}

// DeleteByDeviceId deletes a sensor document by its Device_Id
func (r *SensorRepositoryImpl) DeleteByDeviceId(ctx context.Context, deviceID string) error {
	// Use `sensorid` as the identifier field
	filter := bson.M{"sensorid": deviceID}

	result, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no document found with the given sensorid")
	}
	return nil
}
