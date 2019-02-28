package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expgenproto/userpb"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:18080", "Server address")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func do() error {
	flag.Parse()

	option := grpc.WithInsecure()
	conn, err := grpc.Dial(*addr, option)
	if err != nil {
		return err
	}

	client := userpb.NewUserServiceClient(conn)

	got, err := client.RegisterUser(
		context.Background(),
		&userpb.RegisterUserRequest{
			UserId:       fmt.Sprintf("UserID%d", rand.Int63()),
			EmailAddress: fmt.Sprintf("Email%d", rand.Int63()),
			Password:     "my-pw",
		},
	)
	if err != nil {
		return err
	}

	log.Println("auth token: ", got)
	return nil
}

func main() {
	if err := do(); err != nil {
		log.Printf("err=%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
