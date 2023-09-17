package utils

import (
  "context"
  "errors"
  "net/http"
  "strconv"
  "io"
)

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

func ParseLatLon(r *http.Request) (float64, float64, error) {
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

func ApiRequest(url string, w http.ResponseWriter) error {
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

type APIConfig struct {
	URL     string
	Key string
}

