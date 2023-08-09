package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"corn-weather/app/models"
	"corn-weather/repository"
)

type WeatherService interface {
	FetchAndStoreWeather() error
	GetLatestWeather() (*models.WeatherData, error)
}

type weatherService struct {
	repo repository.WeatherRepository
}

func NewWeatherService(repo repository.WeatherRepository) WeatherService {
	return &weatherService{
		repo: repo,
	}
}

func (s *weatherService) FetchAndStoreWeather() error {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=9daf63b4dcf557be93dbe5384d968076")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var data models.WeatherData
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return err
	}
	fmt.Println("corn called API")
	err = s.repo.StoreData(&data)
	if err != nil {
		return err
	}

	return nil
}

func (s *weatherService) GetLatestWeather() (*models.WeatherData, error) {
	return s.repo.GetLatestData()
}
