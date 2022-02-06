package main

import (
	"cmoore/chore-board/db"
	"cmoore/chore-board/graph"
	"cmoore/chore-board/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.ConnectDB()
	defer db.GlobalInstance.Close()

	frontend := http.FileServer(http.Dir("frontend/build/"))
	http.Handle("/", frontend)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for choreboard", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
