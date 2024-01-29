package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	models "github.com/curious-neeraj/weather_info_cli/models"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

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

	location, current, forecast := weather.Location, weather.Current, weather.Forecast

	headline := fmt.Sprintf("\nCurrent Weather at %v, %v:\n%v, %.1f° C, %v, %v%% Humidity, %v%% Chance of Rain\n\n",
		location.Name,
		location.Country,
		time.Now().Format("15:04"),
		current.Temperature,
		strings.Trim(current.Condition.Text, " "),
		current.Humidity,
		current.ChanceOfRain)

	color.White(headline)

	for _, hour := range forecast.ForecastDay[0].Hour {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		output := fmt.Sprintf(
			"%v, %.1f° C, %v, %v%% Humidity, %v%% Chance of Rain",
			date.Format("15:04"),
			hour.Temp,
			strings.Trim(hour.Condition.Text, " "),
			hour.Humidity,
			hour.ChanceOfRain)

		// color coding based on temperature
		// Temp >= 25, Color = Yellow
		// Temp  < 25, Color = Cyan
		if hour.Temp < 25 {
			color.Cyan(output)
		} else {
			color.Yellow(output)
		}
	}
}
