package repository

import (
	"context"
	entity "dms/internal/domain/entities"
	repository "dms/internal/domain/repositories"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceConfigRepositoryImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDeviceConfigRepository(client *mongo.Client, dbName, collectionName string) repository.DeviceConfigRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &DeviceConfigRepositoryImpl{
		client:     client,
		collection: collection,
	}
}

func (r *DeviceConfigRepositoryImpl) Insert(ctx context.Context, device_config entity.DeviceConfig) error {
	_, err := r.collection.InsertOne(ctx, device_config)
	if err != nil {
		return fmt.Errorf("error inserting device configuration: %v", err)
	}
	return nil
}

func (r *DeviceConfigRepositoryImpl) FindByDeviceId(ctx context.Context, device_id string) (entity.DeviceConfig, error) {
	var sensor entity.DeviceConfig
	err := r.collection.FindOne(ctx, bson.M{"device_id": device_id}).Decode(&sensor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.DeviceConfig{}, nil
		}
		return entity.DeviceConfig{}, fmt.Errorf("error finding device configuration: %v", err)
	}
	return sensor, nil
}

func (r *DeviceConfigRepositoryImpl) DeleteByDeviceId(ctx context.Context, device_id string) error {
	filter := bson.M{"device_id": device_id}

	result, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no document found with the given device id")
	}
	return nil
}
