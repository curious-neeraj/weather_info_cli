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
│   └── screenshot-2.png
├── models
│   ├── api_error.go
│   └── weather.go
└── weather
```

## Instructions

> If you're looking to just use the tool, simply place the executable file named (`weather`) in the `usr/local/bin` folder to make it available as a terminal command. 


## Result

```
$ weather
```
![Weather Info - Raipur](/images/screenshot-1.png)

```
$ weather barcelona
```
![Weather Info - Barcelona](/images/screenshot-2.png)