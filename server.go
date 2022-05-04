package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DavidHODs/TechHUB-goGraph/auth"
	"github.com/DavidHODs/TechHUB-goGraph/graph"
	"github.com/DavidHODs/TechHUB-goGraph/graph/generated"
	database "github.com/DavidHODs/TechHUB-goGraph/postgres"
	"github.com/DavidHODs/TechHUB-goGraph/utils"
	"github.com/gorilla/mux"
)


func main() {
	port, host, _, _, _, _, _, _ := utils.LoadEnv()

	router := mux.NewRouter()
	router.Use(auth.AuthMiddleware())

	database.ConnectAndMigrate()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
