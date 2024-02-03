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

If you're looking to use the tool, you'd need to do the following -

1. Get an API key for `https://api.weatherapi.com` and create a `.env` file with `API_KEY=<key-value>`

2. Now, in the `main.go` file, provide the correct value for the `projectPath` variable.

3. Replace the value of `home` variable with city of your choice. { Currently, the default location points to **Raipur, India** }

4. Build an executable file using the command `go build -o weather cmd/main.go`

5. Place the executable file { `weather` } in `usr/local/bin`, using the command `cp weather usr/local/bin`

> Congratulations! Now you can use the command `weather` from any location to get the weather information and forecast.

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
