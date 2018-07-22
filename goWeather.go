// Copyright 2018 Etosticity Rammington. All rights reserved.
// Use of this source code is governed by a AGPL-3.0 license
// that can be found in the LICENSE file.

package weather

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func queryWeatherInformation(degreeType string, searchString string) ([]byte, error) {
	resp, err := http.Get("http://weather.service.msn.com/find.aspx?src=outlook&weadegreetype=" + degreeType + "&culture=en-US&weasearchstr=" + searchString)
	if err != nil {
		return nil, fmt.Errorf("[%s] Error calling API: %s", programName, err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[%s] Error reading data: %s", programName, err)
	}

	return bytes, nil
}

// GetWeatherInC Gets Weather Information in Celsius Degrees
func GetWeatherInC(searchString string) ([]byte, error) {
	if searchString == "" {
		return nil, fmt.Errorf("[%s] searchString must not be empty", programName)
	}

	bytes, err := queryWeatherInformation("C", searchString)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(bytes, &information); err != nil {
		return nil, fmt.Errorf("[%s] Error unmarshalling data: %s", programName, err)
	}

	jsonData, err := json.Marshal(information)
	if err != nil {
		return nil, fmt.Errorf("[%s] Error marshalling data: %s", programName, err)
	}

	return jsonData, nil
}

// GetWeatherInF Gets Weather Information in Fahrenheit Degrees
func GetWeatherInF(searchString string) ([]byte, error) {
	if searchString == "" {
		return nil, fmt.Errorf("[%s] searchString must not be empty", programName)
	}

	bytes, err := queryWeatherInformation("F", searchString)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(bytes, &information); err != nil {
		return nil, fmt.Errorf("[%s] Error unmarshalling data: %s", programName, err)
	}

	jsonData, err := json.Marshal(information)
	if err != nil {
		return nil, fmt.Errorf("[%s] Error marshalling data: %s", programName, err)
	}

	return jsonData, nil
}
