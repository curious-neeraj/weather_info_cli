# Weather Info and Forecast

![Project Logo - Weather CLI Command](/images/logo.jpeg)

It's always helpful to know weather forecast for the day as it allows us to plan our coffee breaks according to the nature's mood outside!

This project basically does excatly that by letting the users know the weather forecast of an location on their terminal.

## Expectations :

- [x] ~~Basic CLI tool/command(s) to get weather forecase of current date | User should be allowed to pass location!~~
- [x] ~~Read data from making call to Weather API~~
- [x] ~~Organize data as per liking to get necessary info displayed when called~~
- [x] ~~Beautify the result a bit (optional)~~

## Go to Have Add-Ons :

- [x] Better Error Handling and Logging
- [ ] Add Unit Tests (umm, someday, maybe!)

## Project Structure
```
.
├── LICENSE
├── README.md
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── images
│   ├── logo.jpeg
│   ├── screenshot-1.png
│   ├── screenshot-2.png
│   └── screenshot-3.png
├── models
│   ├── api_error.go
│   └── weather.go
└── weather
```

## Usage Instructions

> If you're looking to just use the tool, simply place the executable file { `weather` } in the `usr/local/bin` folder to make it available as a terminal command. 

Currently, the default location points to **Raipur, India**.

- In case you'd like to get weather info about any other city, pass the city name as an argument to the executable : `weather <city>`

- Also, if the city name is common, you can pass the state/country name to pin-point the intentional location : `weather <city> <country>`

## Resultant Examples

```
$ weather
```
![Weather Info - Raipur](/images/screenshot-1.png)

```
$ weather barcelona
```
![Weather Info - Barcelona](/images/screenshot-2.png)

```
$ weather delhi india 
```
![Weather Info - Delhi, India](/images/screenshot-3.png)
