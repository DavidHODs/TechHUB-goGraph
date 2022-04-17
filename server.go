package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DavidHODs/TechHUB-goGraph/graph"
	"github.com/DavidHODs/TechHUB-goGraph/graph/generated"
	database "github.com/DavidHODs/TechHUB-goGraph/postgres"
)


func main() {
	port, host, _, _, _, _, _ := database.LoadEnv()

	database.ConnectAndMigrate()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
