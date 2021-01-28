package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/abuabdillatief/gograph-tutorial/database"
	"github.com/abuabdillatief/gograph-tutorial/graph/generated"
	"github.com/abuabdillatief/gograph-tutorial/graph/resolvers"
	"github.com/go-pg/pg/v9"
)

const defaultPort = "8080"

func main() {
	DB := database.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "meetmeup_dev",
	})

	defer DB.Close()
	DB.AddQueryHook(database.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		MeetupsRepo: database.MeetupsRepo{DB: DB},
		UsersRepo:   database.UsersRepo{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
