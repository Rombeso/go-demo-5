package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("NO_CITY")
var ErrNo200 = errors.New("NOT_200")

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		if !checkCity(city) {
			return nil, ErrNoCity
		}
		return &GeoData{
			City: city,
		}, nil
	}
	url := "https://ipapi.co/json/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, ErrNo200
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	json.Unmarshal(body, &geo)
	fmt.Println(&geo)
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	url := "https://countriesnow.space/api/v0.1/countries/population/cities"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("content-type", "application/json")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if resp.StatusCode != 200 {
		fmt.Println(errors.New("NOT_200"))
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	var population CityPopulationResponse
	json.Unmarshal(body, &population)
	return !population.Error
}
