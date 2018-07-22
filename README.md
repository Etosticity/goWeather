# goWeather
A **Golang** package for obtaining weather information for free

Package `goWeather` uses `weather.service.msn.com` as the source of information

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/etosticity/goWeather
```

---

## Examples

Let's start by initializing `structs` to extract data out from `JSON` output.

```go
type WeatherData struct {
	Day              string `json:"Day"`
	Date             string `json:"Date"`
	SkyText          string `json:"SkyText"`
	ShortDay         string `json:"ShortDay"`
	Humidity         string `json:"Humidity"`
	WindSpeed        string `json:"WindSpeed"`
	FeelsLike        string `json:"FeelsLike"`
	Temperature      string `json:"Temperature"`
	WindDisplay      string `json:"WindDisplay"`
	ObservationTime  string `json:"ObservationTime"`
	ObservationPoint string `json:"ObservationPoint"`
}

type WeatherInfo struct {
	Data                []WeatherData `json:"Data"`
	Location            string        `json:"Location"`
	Provider            string        `json:"Provider"`
	Latitude            string        `json:"Latitude"`
	Longitude           string        `json:"Longitude"`
	DegreeType          string        `json:"DegreeType"`
	EncodedLocationName string        `json:"EncodedLocationName"`
}

type WeatherQuery struct {
	Info []WeatherInfo `json:"Info"`
}
```

We setup 3 `structs` for the extraction of `JSON` output before calling the Weather API. To call the API:

```go
data, err := weather.GetWeatherInC("Canada")
if err != nil {
	log.Fatalln(err)
}
```

Here we are calling the API to find Weather Data for a country / state / city. When calling the API, there are 2 methods.

`GetWeatherInC` calls for the Weather API to return data corresponding to Celsius Degrees.

`GetWeatherInF` calls for the Weather API to return data corresponding to Fahrenheit Degrees.

```go
var result WeatherQuery

err := json.Unmarshal(t, &result)
if err != nil {
	log.Fatalln(err)
}
```

After calling the API, data returned is practically useless unless `unmarshalled`. We are telling `JSON` to unpack the data and store it in a variable.

```go
for _, v := range result.Info {
  fmt.Println("Location: " + v.Location)
  fmt.Println("Provider: " + v.Provider)
  fmt.Println("Latitude: " + v.Latitude)
  fmt.Println("Longitude: " + v.Longitude)
  fmt.Println("DegreeType: " + v.DegreeType)
  fmt.Println("Encoded Location Name: " + v.EncodedLocationName)

  for _, v := range v.Data {
    fmt.Println("Day: " + v.Day)
    fmt.Println("Date: " + v.Date)
    fmt.Println("Sky Text: " + v.SkyText)
    fmt.Println("Humidity: " + v.Humidity)
    fmt.Println("Wind Speed: " + v.WindSpeed)
    fmt.Println("Feels Like: " + v.FeelsLike)
    fmt.Println("Temperature: " + v.Temperature)
    fmt.Println("Wind Display: " + v.WindDisplay)
    fmt.Println("Observation Time: " + v.ObservationTime)
    fmt.Println("Observation Point: " + v.ObservationPoint)
  }

  fmt.Println()
}
```

Here we are iterating through all unpacked data and presenting it to the user.

---

## Full Example

Here's the complete set on how to use `goWeather`

```go
package main

import (
  "log"
  "fmt"
  "encoding/json"

  weather "github.com/etosticity/goWeather"
)

var result WeatherQuery

type WeatherData struct {
	Day              string `json:"Day"`
	Date             string `json:"Date"`
	SkyText          string `json:"SkyText"`
	ShortDay         string `json:"ShortDay"`
	Humidity         string `json:"Humidity"`
	WindSpeed        string `json:"WindSpeed"`
	FeelsLike        string `json:"FeelsLike"`
	Temperature      string `json:"Temperature"`
	WindDisplay      string `json:"WindDisplay"`
	ObservationTime  string `json:"ObservationTime"`
	ObservationPoint string `json:"ObservationPoint"`
}

type WeatherInfo struct {
	Data                []WeatherData `json:"Data"`
	Location            string        `json:"Location"`
	Provider            string        `json:"Provider"`
	Latitude            string        `json:"Latitude"`
	Longitude           string        `json:"Longitude"`
	DegreeType          string        `json:"DegreeType"`
	EncodedLocationName string        `json:"EncodedLocationName"`
}

type WeatherQuery struct {
	Info []WeatherInfo `json:"Info"`
}

func main() {
	t, err := weather.GetWeatherInC("Canada")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(t, &result)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range result.Info {
		fmt.Println("Location: " + v.Location)
		fmt.Println("Provider: " + v.Provider)
		fmt.Println("Latitude: " + v.Latitude)
		fmt.Println("Longitude: " + v.Longitude)
		fmt.Println("DegreeType: " + v.DegreeType)
		fmt.Println("Encoded Location Name: " + v.EncodedLocationName)

		for _, v := range v.Data {
			fmt.Println("Day: " + v.Day)
			fmt.Println("Date: " + v.Date)
			fmt.Println("Sky Text: " + v.SkyText)
			fmt.Println("Humidity: " + v.Humidity)
			fmt.Println("Wind Speed: " + v.WindSpeed)
			fmt.Println("Feels Like: " + v.FeelsLike)
			fmt.Println("Temperature: " + v.Temperature)
			fmt.Println("Wind Display: " + v.WindDisplay)
			fmt.Println("Observation Time: " + v.ObservationTime)
			fmt.Println("Observation Point: " + v.ObservationPoint)
		}

		fmt.Println()
	}
}
```

---

## License

AGPL-3.0 licensed. See LICENSE file for details.