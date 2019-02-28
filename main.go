package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/lima1909/graphql-go-example/hello"
)

// Schemaer interface for create the Schema
type Schemaer interface {
	Schema() *graphql.Schema
}

func main() {
	example := flag.String("e", "hello", "-e hello")
	flag.Parse()
	fmt.Println("Example:", *example)

	var schema Schemaer
	switch *example {
	case "hello":
		schema = &hello.Query{}
	default:
		fmt.Println("Not supported example:", *example)
		return
	}

	startServer(schema)
}

func startServer(s Schemaer) {
	http.Handle("/query", &relay.Handler{Schema: s.Schema()})
	fmt.Println("http://localhost:8080/query")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
