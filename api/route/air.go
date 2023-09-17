package route

import (
	"bytes"
	"fmt"
	. "main/utils"
	"net/http"
	"text/template"
)

type AirT struct{}

var Air AirT

func (c AirT) Get(w http.ResponseWriter, r *http.Request) {
		apiSource := r.Context().Value("apiSource").(*APIConfig)
		lat, lon, err := ParseLatLon(r)
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
		err = tmpl.Execute(&buffer, struct {
			Lat string
			Lon string
			Key string
		}{
			Lat: fmt.Sprintf("%f", lat),
			Lon: fmt.Sprintf("%f", lon),
			Key: apiSource.Key,
		})
		if err != nil {
			http.Error(w, "Error executing URL template", http.StatusInternalServerError)
			return
		}
		url := buffer.String()

		err = ApiRequest(url, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
}
