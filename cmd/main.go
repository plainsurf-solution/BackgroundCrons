package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"corn-weather/app/controllers"
	"corn-weather/app/services"
	"corn-weather/repository"

	"github.com/robfig/cron/v3"
)

const (
	mongoURI   = "mongodb://localhost:27017"
	database   = "weatherdb"
	collection = "weatherdata"
)

func main() {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	coll := client.Database(database).Collection(collection)

	repo := repository.NewWeatherRepository(coll)
	service := services.NewWeatherService(repo)
	controller := controllers.NewWeatherController(service)

	router := http.NewServeMux()
	router.HandleFunc("/get", controller.GetLatestWeatherHandler)

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		err := service.FetchAndStoreWeather()
		if err != nil {
			fmt.Println("Error fetching and storing data:", err)
		}
	})
	c.Start()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on :8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
