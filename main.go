package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dosco/graphjin/core"
	"github.com/dosco/graphjin/serv"
)

func main() {
	// Create a new GraphJin config
	conf := core.Config{
		DBType:     "postgres",
		DBHost:     "localhost",
		DBPort:     5432,
		DBName:     "blog_db",
		DBUser:     "your_username",
		DBPass:     "your_password",
		DBSchema:   "public",
		LogLevel:   "debug",
		Production: false,
	}

	// Create a new GraphJin instance
	gj, err := core.NewGraphJin(&conf, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new GraphJin server
	srv, err := serv.NewGraphJinServer(gj)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the HTTP handler
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		srv.GraphQLHandler(w, r)
	})

	// Start the server
	log.Println("Server running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
