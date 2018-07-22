// Copyright 2018 Etosticity Rammington. All rights reserved.
// Use of this source code is governed by a AGPL-3.0 license
// that can be found in the LICENSE file.

package weather

type data struct {
	Day              string `xml:"day,attr"`
	Date             string `xml:"date,attr"`
	SkyText          string `xml:"skytext,attr"`
	ShortDay         string `xml:"shortday,attr"`
	Humidity         string `xml:"humidity,attr"`
	WindSpeed        string `xml:"windspeed,attr"`
	FeelsLike        string `xml:"feelslike,attr"`
	Temperature      string `xml:"temperature,attr"`
	WindDisplay      string `xml:"winddisplay,attr"`
	ObservationTime  string `xml:"observationtime,attr"`
	ObservationPoint string `xml:"observationpoint,attr"`
}

type info struct {
	Data                []data `xml:"current"`
	Location            string `xml:"weatherlocationname,attr"`
	Provider            string `xml:"provider,attr"`
	Latitude            string `xml:"lat,attr"`
	Longitude           string `xml:"long,attr"`
	DegreeType          string `xml:"degreetype,attr"`
	EncodedLocationName string `xml:"encodedlocationname,attr"`
}

type query struct {
	Info []info `xml:"weather"`
}
