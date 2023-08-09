package repository

import (
	"context"
	"corn-weather/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WeatherRepository interface {
	StoreData(data *models.WeatherData) error
	GetLatestData() (*models.WeatherData, error)
}

type weatherRepository struct {
	collection *mongo.Collection
}

func NewWeatherRepository(collection *mongo.Collection) WeatherRepository {
	return &weatherRepository{
		collection: collection,
	}
}

func (r *weatherRepository) StoreData(data *models.WeatherData) error {
	_, err := r.collection.InsertOne(context.Background(), data)
	return err
}

func (r *weatherRepository) GetLatestData() (*models.WeatherData, error) {
	opts := options.Find().SetLimit(1)
	cursor, err := r.collection.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var data models.WeatherData
		err := cursor.Decode(&data)
		if err != nil {
			return nil, err
		}

		return &data, nil
	}

	return nil, nil
}
