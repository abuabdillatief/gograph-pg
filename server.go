package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/abuabdillatief/gograph-tutorial/database"
	"github.com/abuabdillatief/gograph-tutorial/graph/generated"
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
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
	//AddQueryHook requires a type of QueryHook, which is an interface
	//in package database, we define a struct called DBLogger
	//in this struct we implement 2 methods, in order to implement the QueryHook interface

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	config := &generated.Config{Resolvers: &resolvers.Resolver{
		MeetupsRepo: database.MeetupsRepo{DB: DB},
		UsersRepo:   database.UsersRepo{DB: DB},
	}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(*config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", model.DataloaderMiddlewareDB(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
