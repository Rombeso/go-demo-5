package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange
	city := "London"
	expected := geo.GeoData{City: "London"}
	// Act
	got, err := geo.GetMyLocation(city)
	// Assert
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидали %s получено %s", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonwdwd"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидали %s получено %s", geo.ErrNoCity, err)
	}
}
