package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/Knight-Jo/learning-go-by-examples/go-rest-api/pkg/swagger/server/restapi"
	"github.com/Knight-Jo/learning-go-by-examples/go-rest-api/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/loads"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
