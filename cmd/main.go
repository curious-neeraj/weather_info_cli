package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

func main() {
	fmt.Println("So we're building a basic weather info cli tool using Golang. Let's GO!!")

	// load apiKey value and set location
	apiKey := goDotEnvVariable("API_KEY")
	location := "Raipur"

	// try to call the api (with apikey) and get response data
	body, _ := http.Get("https://api.weatherapi.com/v1/forecast.json?q=" + location + "&days=1&key=" + apiKey)
	fmt.Println(body.Status)
}
