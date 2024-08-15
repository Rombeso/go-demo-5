package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrWrongFormat = errors.New("wrong format")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format <= 0 || format > 4 {
		return "", ErrWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", errors.New("error parsing url")
	}
	params := url.Values{}

	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return "", errors.New("error creating request")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return "", errors.New("error getting weather")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error reading response")
	}
	return string(body), nil
}
