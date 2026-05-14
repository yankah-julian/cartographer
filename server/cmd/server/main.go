package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	http.Handle("/", playground.Handler("Cartographer", "/query"))
	http.Handle("/query", srv)

	log.Printf("Cartographer running at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
