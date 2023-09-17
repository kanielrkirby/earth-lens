package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	. "main/utils"
	"os"
  route "main/route"
)

func Transportation(r chi.Router) {
	API, err := NewAPISourceMiddleware(APIConfig{
		URL: "",
		Key: os.Getenv(""),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("Error creating API middleware: %v", err))
		return
	}
	r.Use(API)

  r.Get("/", route.Transportation.Get)
}




