package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	models "github.com/curious-neeraj/weather_info_cli/models"
	"github.com/joho/godotenv"
)

// type ApiError struct {
// 	code    int
// 	message string
// }

// use godotenv package to load .env file and return value of key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// handle the error encountered while making the API Call
func handleError(err string) {
	log.Println(err)
}

func main() {

	// load apiKey value and set location
	apiKey := goDotEnvVariable("API_KEY")
	home := "Raipur"

	// try to call the api (with apikey) and get response data
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?q=" + home + "&days=1&key=" + apiKey)
	if err != nil {
		handleError(res.Status)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		handleError(res.Status)
	}

	// read API call response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Print(string(body))

	var weather models.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println(err)
	}

	location, current, _ := weather.Location, weather.Current, weather.Forecast

	fmt.Printf("%v, %v, %.2f C, %v, %v%% Humidity \n",
		location.Name,
		location.Country,
		current.Temperature,
		current.Condition.Text,
		current.Humidity)
}
