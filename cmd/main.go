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
	// use full path in order to make the resultant command/tool usable from any file path
	err := godotenv.Load("/Users/neeraj/Desktop/Personal/Projects/weather_info_cli/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	// load apiKey value and set location
	apiKey := goDotEnvVariable("API_KEY")
	home := "Raipur"

	numArgs := len(os.Args)
	if numArgs == 2 {
		home = os.Args[1]
	} else if numArgs == 3 {
		home = os.Args[1] + "%20" + os.Args[2]
	} else if numArgs > 2 {
		log.Fatalf("Illegal Use! Kindly enter a valid city name.")
	}

	// try to call the api (with apikey) and get response data
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?q=" + home + "&days=1&key=" + apiKey)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var apiError models.ApiError
		_ = json.Unmarshal(body, &apiError)
		log.Fatalln(apiError.Error.Message)
	}

	// read API call response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var weather models.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalln(err)
	}

	location, current, forecast := weather.Location, weather.Current, weather.Forecast

	color.Green("\nWeather Update (%v, %v)\n", location.Name, location.Country)
	headline := fmt.Sprintf("%v  -  %4.1f° C, %3v%% Humidity, %3v%% Chance of Rain, %v\n\n",
		time.Now().Format("15:04"),
		current.Temperature,
		current.Humidity,
		current.ChanceOfRain,
		strings.Trim(current.Condition.Text, " "))

	color.White(headline)

	for _, hour := range forecast.ForecastDay[0].Hour {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		output := fmt.Sprintf(
			"%v  -  %4.1f° C, %3v%% Humidity, %3v%% Chance of Rain, %v",
			date.Format("15:04"),
			hour.Temp,
			hour.Humidity,
			hour.ChanceOfRain,
			strings.Trim(hour.Condition.Text, " "))

		// color coding based on temperature
		// Temp >= 25, Color = Yellow
		// Temp  < 25, Color = Cyan
		if hour.Temp < 25 {
			color.Cyan(output)
		} else {
			color.Yellow(output)
		}
		time.Sleep(time.Duration(150) * time.Millisecond)
	}
}
