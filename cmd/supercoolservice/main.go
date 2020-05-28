package main

import (
	"fmt"
	"mikelangelon/m/v2/api/resource"
	"mikelangelon/m/v2/api/rest"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/internal/app/attachmentrepo"
	"mikelangelon/m/v2/internal/app/user"
	"mikelangelon/m/v2/internal/pkg"
	"mikelangelon/m/v2/internal/pkg/dgraph"
	"mikelangelon/m/v2/internal/pkg/minioclient"
	"net/http"
	"os"

	"github.com/go-openapi/loads"
	"github.com/minio/minio-go/v6"
	"github.com/syllabix/swagserver"
	"github.com/syllabix/swagserver/option"
	"github.com/syllabix/swagserver/theme"
)

func main() {

	r := redis()
	c := setupMinio()
	setupDGraph()

	// ###### API ######
	server := api(c, r)

	fmt.Println("Listening in 8080...")
	http.ListenAndServe(":8080", server)
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

func redis() pkg.RedisClient {
	redisURI := goDotEnvVariable("redis_hostname", "localhost")
	redisPort := goDotEnvVariable("redis_host", "6379")

	fmt.Println(fmt.Sprintf("connecting to redis with  %s:%s", redisURI, redisPort))

	client := pkg.New(fmt.Sprintf("%s:%s", redisURI, redisPort))

	return client
}

func setupMinio() attachmentrepo.Store {
	client, err := minio.New("localhost:9000", "minioadmin", "minioadmin", false)
	if err != nil {
		panic(err)
	}
	r := minioclient.New(client)
	err = r.Prepare()
	if err != nil {
		panic(err)
	}
	return r
}

func setupDGraph() {
	dgraph.New()
}

func api(store attachmentrepo.Store, redis user.Redis) *http.ServeMux {
	specs, err := loads.Analyzed(rest.SwaggerJSON, "")
	if err != nil {
		os.Exit(1)
	}
	api := operation.NewCoolappAPI(specs)

	resource.Register(api, resource.Dependencies{
		Store: store,
		Redis: redis,
	})

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
	return s
}
