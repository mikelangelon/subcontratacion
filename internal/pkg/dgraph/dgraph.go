package dgraph

import (
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

func New() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	// ctx := context.Background()
	//
	// // Perform login call. If the Dgraph cluster does not have ACL and
	// // enterprise features enabled, this call should be skipped.
	// for {
	// 	// Keep retrying until we succeed or receive a non-retriable error.
	// 	err = dg.Login(ctx, "groot", "password")
	// 	if err == nil || !strings.Contains(err.Error(), "Please retry") {
	// 		break
	// 	}
	// 	time.Sleep(time.Second)
	// }
	// if err != nil {
	// 	log.Fatalf("While trying to login %v", err.Error())
	// }

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}
