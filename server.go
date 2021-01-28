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
	customMiddleware "github.com/abuabdillatief/gograph-tutorial/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := database.New(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
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
	userRepo := database.UsersRepo{DB: DB}
	
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	config := &generated.Config{
		Resolvers: &resolvers.Resolver{
			MeetupsRepo: database.MeetupsRepo{DB: DB},
			UsersRepo:   userRepo,
		},
		/**
		 * generated.Config takes 3 arguments:
		 * 		Resovlers
		 * 		DirectiveRoot
		 * 		ComplexityRoot
		 * those 3 are structs.
		 * Here we only define resolvers, becauase at the moment
		 * that's what we're defining
		 *
		 */
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(*config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", model.DataloaderMiddlewareDB(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
