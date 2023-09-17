package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	Route "main/route"
	. "main/utils"
	"os"
)

func WaterController(r chi.Router) {
	WaterServices, err := NewAPISourceMiddleware(APIConfig{
		URL: "https://waterservices.usgs.gov/nwis/iv/?format=json&bBox={{ .West }},{{ .South }},{{ .East }},{{ .North }}",
		Key: os.Getenv("WATER_SERVICES_API_KEY"),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("Error creating WaterServices middleware: %v", err))
		return
	}

	r.Use(WaterServices)

	r.Get("/", Route.Water.Get)
}
