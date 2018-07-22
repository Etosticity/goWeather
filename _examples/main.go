// Copyright 2018 Etosticity Rammington. All rights reserved.
// Use of this source code is governed by a AGPL-3.0 license
// that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"log"

	weather "goWeather"
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
