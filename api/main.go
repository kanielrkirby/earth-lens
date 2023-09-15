package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"net/http"
	"text/template"
  "bytes"
  "io"
  "os"
  "fmt"
  "strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		errors.Wrap(err, "Error loading .env file")
	}
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
    r.Route("/air", func(r chi.Router) {
      OpenWeatherMap, err := NewAPISourceMiddleware(APIConfig{
				URL:     "http://api.openweathermap.org/data/2.5/air_pollution?lat={{ .Lat }}&lon={{ .Lon }}&appid={{ .ApiKey }}",
				Key: os.Getenv("OPEN_WEATHER_MAP_API_KEY"),
			})
      if err != nil {
        fmt.Println(fmt.Errorf("Error creating OpenWeatherMap middleware: %v", err))
        return
      }

			r.Use(OpenWeatherMap)

      r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				apiSource := r.Context().Value("apiSource").(*APIConfig)
        lat, lon, err := parseLatLon(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        tmpl, err := template.New("url").Parse(apiSource.URL)
        if err != nil {
            http.Error(w, "Error parsing URL template", http.StatusInternalServerError)
            return
        }
        var buffer bytes.Buffer
        err = tmpl.Execute(&buffer, struct{
          Lat string
          Lon string
          ApiKey string
        }{
          Lat: fmt.Sprintf("%f", lat),
          Lon: fmt.Sprintf("%f", lon),
          ApiKey: apiSource.Key,
        })
        if err != nil {
            http.Error(w, "Error executing URL template", http.StatusInternalServerError)
            return
        }
        url := buffer.String()
        err = apiRequest(url, w)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
      })
    })

    r.Route("/water", func(r chi.Router) {
      WaterServices, err := NewAPISourceMiddleware(APIConfig{
        URL: "https://waterservices.usgs.gov/nwis/iv/?format=json&bBox={{ .West }},{{ .South }},{{ .East }},{{ .North }}",
        Key: os.Getenv("WATER_SERVICES_API_KEY"),
      })
      if err != nil {
        fmt.Println(fmt.Errorf("Error creating WaterServices middleware: %v", err))
        return
      }

      r.Use(WaterServices)

      r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        apiSource := r.Context().Value("apiSource").(*APIConfig)
        lat, lon, err := parseLatLon(r)
        if err != nil {
          http.Error(w, err.Error(), http.StatusBadRequest)
          return
        }

        west := lon - float64(1) / 6
        east := lon + float64(1) / 6
        south := lat - float64(1) / 6
        north := lat + float64(1) / 6

        tmpl, err := template.New("url").Parse(apiSource.URL)
        if err != nil {
          http.Error(w, "Error parsing URL template", http.StatusInternalServerError)
          return
        }

        var buffer bytes.Buffer
        err = tmpl.Execute(&buffer, struct{
          West string
          East string
          South string
          North string
        }{
          West: fmt.Sprintf("%f", west),
          East: fmt.Sprintf("%f", east),
          South: fmt.Sprintf("%f", south),
          North: fmt.Sprintf("%f", north),
        })
        if err != nil {
          http.Error(w, "Error executing URL template", http.StatusInternalServerError)
          return
        }
        url := buffer.String()

        err = apiRequest(url, w)
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
      })
    })
  })

    http.ListenAndServe(":8080", r)
  }

type APIConfig struct {
	URL     string
	Key string
}

func NewAPISourceMiddleware(config APIConfig) (func(next http.Handler) http.Handler, error) {
  if config.URL == "" {
    return nil, errors.New("NewAPISourceMiddleware: URL is required")
  }
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "apiSource", &config)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}, nil
}

func parseLatLon(r *http.Request) (float64, float64, error) {
    latQuery := r.URL.Query().Get("lat")
    lonQuery := r.URL.Query().Get("lon")
    if latQuery == "" || lonQuery == "" {
        return 0, 0, errors.New("Latitude and Longitude are required")
    }
    lat, err := strconv.ParseFloat(latQuery, 64)
    if err != nil {
        return 0, 0, errors.New("Latitude must be a float")
    }
    lon, err := strconv.ParseFloat(lonQuery, 64)
    if err != nil {
        return 0, 0, errors.New("Longitude must be a float")
    }
    return lat, lon, nil
}

func apiRequest(url string, w http.ResponseWriter) error {
    resp, err := http.Get(url)
    if err != nil {
        return errors.New("Issue with the API request")
    }
    defer resp.Body.Close()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(resp.StatusCode)
    _, err = io.Copy(w, resp.Body)
    if err != nil {
        return errors.New("Error streaming response")
    }
    return nil
}

