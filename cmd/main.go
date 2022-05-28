package main

import (
	"fmt"
	"htt/httbackend"
	"htt/httbackend/graph"
	"htt/httbackend/graph/generated"
	"htt/httbackend/models"
	"htt/httbackend/repositories"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type RouteBuilder struct {
	router *mux.Router
}

func (routeBuilder *RouteBuilder) MakeRoute(path string, f func(RouteBuilder, *mux.Router)) *RouteBuilder {
	apiRouter := routeBuilder.router.PathPrefix(path).Subrouter()
	f(RouteBuilder{router: apiRouter}, apiRouter)
	return routeBuilder
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Some error occurred. Err: %s", err)
	}

	app := &httbackend.ApplicationHandler{}

	app.PrepareConfig()

	app.InitDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		SermonRepo:       repositories.New(app.DB, &models.Sermon{}),
		GalleryRepo:      repositories.New(app.DB, &models.Gallery{}),
		ContactRepo:      repositories.New(app.DB, &models.Contact{}),
		SubscriptionRepo: repositories.New(app.DB, &models.NewsletterSubscription{}),
	}}))

	mainRouteBuilder := &RouteBuilder{router: mux.NewRouter()}
	mainRouteBuilder.MakeRoute("/", func(apiRouteBuilder RouteBuilder, router *mux.Router) {
		router.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))
		router.Handle("/graphql", srv).Methods("GET", "POST")
	})

	host, _ := os.Hostname()
	fmt.Printf("Starting the server at http://%s:%v\n", host, app.Config.AppPort)

	log.Fatal(http.ListenAndServe(":"+app.Config.AppPort, mainRouteBuilder.router))
}
