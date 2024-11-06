# Go-WeatherApp

A simple weather application built using **Go (Golang)** that allows users to check the current weather conditions of any city. This application uses the **OpenWeatherMap API** to fetch weather data and presents it in a clean web interface.

## Features

- Enter a city name to fetch weather information.
- Displays current temperature, humidity, and other weather details.
- Simple, user-friendly web interface with a responsive layout.
- Deployed on Kubernetes, containerized using Docker.

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS (Simple design using `style.css`)
- **API**: OpenWeatherMap API for weather data
- **Containerization**: Docker
- **Orchestration**: Kubernetes (for deployment)
  
## Build and Deploy

### Steps:

```bash
docker build -t go-weatherapp .
docker push username/go-weatherapp
kubectl apply -f kubernetes/app.yml
kubectl apply -f kubernetes/svc.yml


