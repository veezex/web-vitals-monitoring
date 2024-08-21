package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/config"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/metric"
	"log"
	"net/http"
	"text/template"
)

/*
  - Resquest example:
    fetch('http://localhost:6510/metric', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json'
    },
    body: JSON.stringify({id: "id-page", client: "mobile", uri: "/best", name: 'name v', value: 1.1, target: 'target t', rating: 'rating r'})
    })
*/
func createHandleMetric() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method == "POST" {
			// check the content type
			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
				return
			}

			// read the body
			var data map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			m, err := metric.Parse(data)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			errSave := dbInstance.SaveMetric(m)
			if errSave != nil {
				log.Printf("Failed to save metric: %s", errSave)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			return
		}

		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}
}

func createHandleScript(config config.Config) http.HandlerFunc {
	tmpl, err := template.ParseFiles("templates/client.js")
	if err != nil {
		log.Fatalf("Failed to parse template: %s", err)
		return nil
	}

	data := struct {
		Domain   string
		Protocol string
		Port     int
	}{
		Domain:   config.GetDomain(),
		Port:     config.GetPort(),
		Protocol: protocol(config),
	}

	var tplBuffer bytes.Buffer
	err = tmpl.Execute(&tplBuffer, data)
	if err != nil {
		panic(err)
	}

	// Создаём минификатор
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	// Минифицируем результат
	var minifiedJS bytes.Buffer
	err = m.Minify("text/javascript", &minifiedJS, &tplBuffer)
	if err != nil {
		panic(err)
	}

	// Сохраняем минифицированный результат в переменной
	renderedJS := minifiedJS.Bytes()

	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method != "GET" {
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("Cache-Control", "public, max-age=600")

		_, err = fmt.Fprint(w, string(renderedJS))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
