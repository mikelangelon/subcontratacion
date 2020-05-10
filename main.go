package main

import (
	"fmt"
	"mikelangelon/m/v2/api/resource"
	"mikelangelon/m/v2/api/rest"
	"mikelangelon/m/v2/api/rest/operation"
	"net/http"
	"os"

	"github.com/go-openapi/loads"
	"github.com/syllabix/swagserver"
	"github.com/syllabix/swagserver/option"
	"github.com/syllabix/swagserver/theme"
)

func main() {
	fmt.Println("Starting testing server...")

	specs, err := loads.Analyzed(rest.SwaggerJSON, "")
	if err != nil {
		os.Exit(1)
	}
	api := operation.NewCoolappAPI(specs)

	r := new(resource.Resource)
	r.Register(api)

	s := http.NewServeMux()
	s.HandleFunc("/", handler)
	swaggerAPI := api.Serve(nil)
	s.Handle("/v1/", swaggerAPI)
	s.Handle("/swagger.json", swaggerAPI)

	s.Handle("/internal/v1/", swagserver.NewHandler(
		option.Path("/internal/v1"),
		option.SwaggerSpecURL("/swagger.json"),
		option.Theme(theme.Muted),
	))

	http.ListenAndServe(":8080", s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
