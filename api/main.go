package main

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"net/http"
  controller "main/controller"
  "github.com/go-chi/chi/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		errors.Wrap(err, "Error loading .env file")
	}

  r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/air", controller.AirController)
		r.Route("/water", controller.WaterController)
	})


	http.ListenAndServe(":8080", r)
}
