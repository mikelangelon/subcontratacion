package main

import (
	"fmt"
	"log"
	"mikelangelon/m/v2/api/resource"
	"mikelangelon/m/v2/api/rest"
	"mikelangelon/m/v2/api/rest/operation"
	"net/http"
	"os"

	"github.com/go-openapi/loads"
	"github.com/gomodule/redigo/redis"
	"github.com/syllabix/swagserver"
	"github.com/syllabix/swagserver/option"
	"github.com/syllabix/swagserver/theme"
)

func main() {
	fmt.Println("Starting testing server...")

	fmt.Println("Connect to redis...")
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.Do("SET", "k1", 1)
	n, _ := redis.Int(conn.Do("GET", "k1"))
	fmt.Printf("%#v\n", n)
	n, _ = redis.Int(conn.Do("INCR", "k1"))
	fmt.Printf("%#v\n", n)

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
