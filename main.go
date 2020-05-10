package main

import (
	"fmt"
	"mikelangelon/m/v2/api/resource"
	"mikelangelon/m/v2/api/rest"
	"mikelangelon/m/v2/api/rest/operation"
	"net/http"
	"os"

	"github.com/go-openapi/loads"
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

	http.ListenAndServe(":8080", api.Serve(nil))
}
