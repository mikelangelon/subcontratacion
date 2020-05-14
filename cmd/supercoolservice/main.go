package main

import (
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
	redisURI := goDotEnvVariable("redis_hostname", "localhost")
	redisPort := goDotEnvVariable("redis_host", "6379")

	fmt.Println("Starting testing server...")

	fmt.Println(fmt.Sprintf("connecting to redis with  %s:%s", redisURI, redisPort))

	client := pkg.New(fmt.Sprintf("%s:%s", redisURI, redisPort))

	err := client.SetPair("key", "testingValue")
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.GetPair("key")
	fmt.Println(fmt.Sprintf("val %s with err %v", val, err))

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

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key, defaultOption string) string {

	value, exists := os.LookupEnv(key)

	if exists {

		return value
	}
	return defaultOption
}
