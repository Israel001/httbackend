package main

import (
	"fmt"
	"htt/httbackend/graph"
	"htt/httbackend/graph/generated"
	httbackend "htt/httbackend/init"
	"htt/httbackend/models"
	"htt/httbackend/repositories"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

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
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
                return true
            },
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
        },
	})

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: strings.Split(app.Config.AllowedOrigins, ","),
		AllowCredentials: true,
		AllowedHeaders: []string{"*"},
		Debug: true,
	}).Handler)

	router.Handle("/", playground.Handler("GraphQL Playground", "/graphql"))
	router.Handle("/graphql", srv)

	host, _ := os.Hostname()
	fmt.Printf("Starting the server at http://%s:%v\n", host, app.Config.AppPort)

	log.Fatal(http.ListenAndServe(":"+app.Config.AppPort, router))
}
