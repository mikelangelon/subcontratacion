package main

import (
	"flag"
	"fmt"
	"mikelangelon/m/v2/api/resource"
	"mikelangelon/m/v2/api/rest"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/internal/pkg"
	"net/http"
	"os"

	"github.com/go-openapi/loads"
	"github.com/syllabix/swagserver"
	"github.com/syllabix/swagserver/option"
	"github.com/syllabix/swagserver/theme"
)

func main() {
	var redisURL = flag.String("redisURL", "redis", "help message for flagname")
	flag.Parse()

	fmt.Println("Starting testing server...")

	fmt.Println("Connect to redis2...")

	client := pkg.New(fmt.Sprintf("%s:6379", *redisURL))

	client.SetPair("Something", "sooooomething")
	fmt.Println(client.GetPair("Something"))

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
	fmt.Println("Listening in 8080...")
	http.ListenAndServe(":8080", s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
