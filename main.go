package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dosco/graphjin/core"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gopkg.in/yaml.v3"
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

type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
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

	router := chi.NewRouter()
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var req GraphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		ctx := context.WithValue(r.Context(), core.UserIDKey, 1)
	
		var variables_raw json.RawMessage
		if req.Variables != nil {
			variables_raw, err = json.Marshal(req.Variables)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	
		res, err := gj.GraphQL(ctx, req.Query, variables_raw, nil)
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
