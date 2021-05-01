package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Retrieving the project, Weather predictions from 2020

type WheaterCloud struct {
	CityName string      `json:"city_name"`
	Data     []DataCloud `json:"data"`
}

type DataCloud struct {
	RelativeHumidity float64     `json:"rh"`
	Weather          WeatherData `json:"weather"`
	Temperature      float64     `json:"temp"`
	Precipitation    float64     `json:"precip"`
	Clouds           int         `json:"clouds"`
}

type WeatherData struct {
	Icon        string `json:"icon"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

const LATITUDE = "-2.335017"
const LONGITUDE = "-80.229769"
const WEATHER_KEY = "03f3ae71c48847e7a7e2b0077bf35a76"

func MakingDays() []string {

	daysString := make([]string, 0)

	return daysString
}

func BuildingUrl(timeStart string, timeEnd string) string {

	return fmt.Sprintf("https://api.weatherbit.io/v2.0/history/hourly?lat=%v&lon=%v&start_date=%v&end_date=%v&tz=local&key=%v", LATITUDE, LONGITUDE, timeStart, timeEnd, WEATHER_KEY)

}

func ConnectAPI() {

	var responseObject WheaterCloud
	response, err := http.Get(BuildingUrl("2020-10-21", "2020-10-22"))
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Data)

}

func MakingCSVFiles() error {

	recordFile, err := os.Create("./Overcast_clouds.csv")

	if err != nil {
		return errors.New("file already exists")
	}
	defer recordFile.Close()

	// Initialize the writer
	writer := csv.NewWriter(recordFile)

	var csvData = [][]string{
		{"SuperHero Name", "Power", "Weakness"},
		{"Batman", "Wealth", "Human"},
		{"Superman", "Strength", "Kryptonite"},
	}

	err = writer.WriteAll(csvData)
	if err != nil {
		return errors.New("something happend while creating the file")
	}

	return nil
}

func main() {

	ConnectAPI()
	MakingCSVFiles()
}
