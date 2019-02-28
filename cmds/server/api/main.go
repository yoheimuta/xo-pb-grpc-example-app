package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expgenproto/userproductpb"

	"github.com/yoheimuta/xo-pb-grpc-example-app/transport/grpc/services/userproductservice"

	"google.golang.org/grpc"

	"github.com/yoheimuta/xo-pb-grpc-example-app/app/userapp"
	"github.com/yoheimuta/xo-pb-grpc-example-app/app/userproductapp"
	"github.com/yoheimuta/xo-pb-grpc-example-app/domain/authtoken"
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/expgenproto/userpb"
	"github.com/yoheimuta/xo-pb-grpc-example-app/infra/exptime"
	"github.com/yoheimuta/xo-pb-grpc-example-app/repository/expdb/expmysql"
	"github.com/yoheimuta/xo-pb-grpc-example-app/transport/grpc/services/userservice"
)

var (
	mysqlSourceName   = flag.String("mysql_source_name", "root:my-pw@tcp(127.0.0.1:3306)/test-xo-db", "MySQL DataSourceName")
	serverAddr        = flag.String("server_addr", ":18080", "Server Local address")
	authTokenSecret   = flag.String("auth_token_secret", "test_secret", "Secret for generating an auth token")
	authTokenLifetime = flag.Duration("auth_token_lifetime", 1*time.Minute, "Lifetime of an auth token")
)

func do() error {
	flag.Parse()

	// Setup dependencies. {
	db, err := expmysql.NewClient(*mysqlSourceName)
	if err != nil {
		return err
	}

	clock := exptime.NewClock()

	authTokenGenerator, err := authtoken.NewGenerator(
		[]byte(*authTokenSecret),
		*authTokenLifetime,
	)
	if err != nil {
		return err
	}
	// }

	// Creates apps. {
	userApp := userapp.NewApp(
		db,
		clock,
		authTokenGenerator,
	)
	userProductApp := userproductapp.NewApp(
		db,
	)
	// }

	// Creates gRPC services. {
	userService := userservice.NewService(
		userApp,
	)
	userProductService := userproductservice.NewService(
		userProductApp,
	)
	// }

	// Creates a gRPC server. {
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userService)
	userproductpb.RegisterUserProductServiceServer(grpcServer, userProductService)
	// }

	// Listens and serves requests. {
	listener, err := net.Listen("tcp", *serverAddr)
	if err != nil {
		return err
	}
	log.Printf("Listening on %v", listener.Addr().(*net.TCPAddr))
	return grpcServer.Serve(listener)
	// }
}

func main() {
	if err := do(); err != nil {
		log.Printf("err=%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
