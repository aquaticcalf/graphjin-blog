package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/dosco/graphjin/core"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	DBType   string `yaml:"DB_TYPE"`
	DB       string `yaml:"DB"`
	LogLevel string `yaml:"log_level"`
	Tables   []struct {
		Name   string `yaml:"name"`
		Schema string `yaml:"schema"`
	} `yaml:"tables"`
}

func main() {

	config_data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error reading config.yaml: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(config_data, &cfg); err != nil {
		log.Fatalf("error parsing config.yaml: %v", err)
	}

	db, err := sql.Open("pgx", cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	gj, err := core.NewGraphJin(nil, db)
	if err != nil {
		log.Fatal(err)
	}

	// example
	query := `
        query getPosts {
            posts {
                id
                title
                content
                created_at
            }
        }`


	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, request *http.Request) {
		ctx := context.WithValue(request.Context(), core.UserIDKey, 1)
		res, err := gj.GraphQL(ctx, query, nil, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res.Data)
	})

	log.Println("go server started on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
