package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	. "main/utils"
	"os"
  route "main/route"
)

func LandUse(r chi.Router) {
	API, err := NewAPISourceMiddleware(APIConfig{
    URL: "https://api.nasa.gov/planetary/earth/imagery?lon={{ .Lon }}&lat={{ .Lat }}&dim=0.15&api_key={{ .Key }}",
		Key: os.Getenv("NASA_API_KEY"),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("Error creating API middleware: %v", err))
		return
	}
	r.Use(API)

  r.Get("/", route.LandUse.Get)
}




