package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
YYYY - MMMM - DD
*/

type WeatherClouds struct {
}

const LATITUDE = ""
const LONGITUDE = ""
const WEATHER_KEY = "03f3ae71c48847e7a7e2b0077bf35a76"
const WEATHER_URL = "https://api.weatherbit.io/v2.0/history/hourly?lat={}&lon={}&start_date={}&end_date={}&tz=local&key={}"

func MakingDays() []string {

	daysString := make([]string, 0)

	return daysString
}

func ConnectAPI() {

	response, err := http.Get("https://api.weatherbit.io/v2.0/history/hourly?lat=-2.335017&lon=-80.229769&start_date=2020-10-21&end_date=2020-10-22&tz=local&key=03f3ae71c48847e7a7e2b0077bf35a76")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseData)

}

func main() {

	ConnectAPI()
}
