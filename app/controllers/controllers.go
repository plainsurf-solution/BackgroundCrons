package controllers

import (
	"encoding/json"
	"net/http"

	"corn-weather/app/services"
)

type WeatherController struct {
	service services.WeatherService
}

func NewWeatherController(service services.WeatherService) *WeatherController {
	return &WeatherController{
		service: service,
	}
}


func (c *WeatherController) GetLatestWeatherHandler(w http.ResponseWriter, r *http.Request) {
	data, err := c.service.GetLatestWeather()
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		return
	}

	if data == nil {
		http.Error(w, "No data available", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error marshaling data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
