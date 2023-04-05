package http_server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/RaraykinValery/l0/internal/cache"
	"github.com/RaraykinValery/l0/internal/models"
)

var templates_paths = []string{
	filepath.Join("web", "templates", "index.html"),
	filepath.Join("web", "templates", "order.html"),
}

var templates = template.Must(template.ParseFiles(templates_paths...))

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "index", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		var order models.Order

		uuid := r.FormValue("uuid")

		order, ok := cache.GetOrderFromCache(uuid)
		if ok == false {
			log.Printf("The order is not in cache")
		}

		templates.ExecuteTemplate(w, "index", order)
	}

}

func StartHTTPServer(port string) error {
	fileServer := http.FileServer(http.Dir("web/static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", orderHandler)

	log.Fatal(http.ListenAndServe(port, nil))

	return nil
}
