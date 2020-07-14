package main

import (
	"fmt"
	"mikelangelon/m/v2/internal/app/attachmentrepo"
	"mikelangelon/m/v2/internal/pkg"
	"mikelangelon/m/v2/internal/pkg/minioclient"
	. "os"

	"github.com/minio/minio-go/v6"
)

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

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key, defaultOption string) string {

	value, exists := LookupEnv(key)

	if exists {
		return value
	}
	return defaultOption
}
