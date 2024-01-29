package models

/*
It is interesting to note that spacing might screw things up in Go
For example -

	Temperature float64 `json: "temp_c"`
	--> would result in default value because of the space in between `json: "temp_c"`

	Temperature float64 `json:"temp_c"`
	--> Remove the space, and volla, the issue is resolved!
*/

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		Temperature float64 `json:"temp_c"`
		Condition   struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity     float64 `json:"humidity"`
		ChanceOfRain float64 `json:"chance_of_rain"`
	} `json:"current"`

	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				Temp      float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				Humidity     float64 `json:"humidity"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
