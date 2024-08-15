package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestWeather(t *testing.T) {
	expect := "Moscow"
	geoData := geo.GeoData{City: expect}
	format := 3
	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	if !strings.Contains(result, expect) {
		t.Errorf("Ожидали %s получено %s", expect, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expect := "Moscow"
			geoData := geo.GeoData{City: expect}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидали %s получено %s", weather.ErrWrongFormat, err)
			}
		})
	}
}
