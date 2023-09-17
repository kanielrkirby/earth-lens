package route

import (
	"bytes"
	"fmt"
	. "main/utils"
	"net/http"
	"text/template"
)

type WaterT struct{}

var Water WaterT

func (c WaterT) Get(w http.ResponseWriter, r *http.Request) {
	apiSource := r.Context().Value("apiSource").(*APIConfig)
	lat, lon, err := ParseLatLon(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	west := lon - float64(1)/6
	east := lon + float64(1)/6
	south := lat - float64(1)/6
	north := lat + float64(1)/6

	tmpl, err := template.New("url").Parse(apiSource.URL)
	if err != nil {
		http.Error(w, "Error parsing URL template", http.StatusInternalServerError)
		return
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, struct {
		West  string
		East  string
		South string
		North string
	}{
		West:  fmt.Sprintf("%f", west),
		East:  fmt.Sprintf("%f", east),
		South: fmt.Sprintf("%f", south),
		North: fmt.Sprintf("%f", north),
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
