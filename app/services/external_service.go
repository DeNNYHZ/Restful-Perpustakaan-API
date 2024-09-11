package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// ExternalService provides methods for interacting with external services
type ExternalService struct{}

// NewExternalService creates a new ExternalService instance
func NewExternalService() *ExternalService {
	return &ExternalService{}
}

// GetWeatherData mengambil data cuaca dari API OpenWeatherMap berdasarkan nama kota
func (es *ExternalService) GetWeatherData(city string) (map[string]interface{}, error) {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY") // Ambil API key dari environment variable
	if apiKey == "" {
		return nil, errors.New("OPENWEATHERMAP_API_KEY environment variable not set")
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
