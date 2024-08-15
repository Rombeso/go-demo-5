package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "User city")
	format := flag.Int("format", 1, "Output weather format")
	flag.Parse()
	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData, _ := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
