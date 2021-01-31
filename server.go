package main

import (
	"fmt"
	"github.com/j-kk/go-graphql/graph/dtb"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/j-kk/go-graphql/graph"
	"github.com/j-kk/go-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	// Parse
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		fmt.Fprintf(os.Stderr, "database_url unset")
		os.Exit(1)
	}

	// Set db
	db, err := dtb.InitDB(dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect to dtb")
		os.Exit(1)
	}
	fmt.Printf("Connected to database\n")
	dtb.GlobalDB = db

	defer dtb.GlobalDB.CloseDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
