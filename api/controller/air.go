package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	. "main/utils"
	"os"
  route "main/route"
)

func Air(r chi.Router) {
	OpenWeatherMap, err := NewAPISourceMiddleware(APIConfig{
		URL: "https://api.openweathermap.org/data/2.5/air_pollution?lat={{ .Lat }}&lon={{ .Lon }}&appid={{ .Key }}",
		Key: os.Getenv("OPEN_WEATHER_MAP_API_KEY"),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("Error creating OpenWeatherMap middleware: %v", err))
		return
	}
	r.Use(OpenWeatherMap)

  r.Get("/", route.Air.Get)
}

