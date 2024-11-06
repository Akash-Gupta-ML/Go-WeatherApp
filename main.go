package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type WeatherResponse struct {
	Name    string `json:"name"`
	Main    struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

type WeatherDescription struct {
	Description string `json:"description"`
}

type WeatherData struct {
	Error   string
	Name    string
	Main    struct {
		Temp     float64
		Humidity int
	}
}

const apiKey = "0aa9e4c62edc0525abd44d283f00d4e0"

func getWeather(city string) (WeatherResponse, error) {
	var weather WeatherResponse
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return weather, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&weather)
	return weather, err
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	var weatherData WeatherData

	if r.Method == "POST" {
		city := r.FormValue("city")
		if city != "" {
			weather, err := getWeather(city)
			if err != nil {
				weatherData.Error = "Failed to get weather data. Try again."
			} else {
				weatherData.Name = weather.Name
				weatherData.Main.Temp = float64(int(weather.Main.Temp))
				weatherData.Main.Humidity = weather.Main.Humidity
			}
		} else {
			weatherData.Error = "Please enter a city."
		}
	}

	tmpl := template.Must(template.ParseFiles("/app/templates/index.html"))
	err := tmpl.Execute(w, weatherData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static"))))
	http.HandleFunc("/", weatherHandler)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
