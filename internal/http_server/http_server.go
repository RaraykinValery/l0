package http_server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/RaraykinValery/l0/internal/cache"
	"github.com/RaraykinValery/l0/internal/models"
)

var server http.Server

var templates_root string = filepath.Join("web", "templates")
var templates_paths = []string{
	filepath.Join(templates_root, "index.html"),
	filepath.Join(templates_root, "order.html"),
}
var templates = template.Must(template.ParseFiles(templates_paths...))

func orderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := templates.ExecuteTemplate(w, "index", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		var order models.Order

		uuid := r.FormValue("uuid")

		order, ok := cache.GetOrder(uuid)
		if ok == false {
			log.Printf("Requested order is not in cache")
		}

		templates.ExecuteTemplate(w, "index", order)
	}
}

func Start(port string) error {
	server = http.Server{Addr: port}

	fileServer := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", orderHandler)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return nil
}

func Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
