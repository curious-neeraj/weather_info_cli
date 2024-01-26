package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
func handleApiError(err string) {
	log.Println(err)
}

func main() {

	// load apiKey value and set location
	apiKey := goDotEnvVariable("API_KEY")
	location := "Raipur"

	// try to call the api (with apikey) and get response data
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?q=" + location + "&days=1&key=" + apiKey)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(res)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		handleApiError(res.Status)
	}

}
