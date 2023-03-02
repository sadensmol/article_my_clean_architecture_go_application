package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	handler "github.com/sadensmol/article_my_clean_architecture_go_application/internal/account"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/sadensmol/article_my_clean_architecture_go_application/api/proto/v1"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	// here we construct handlers along with requirements like account usecases etc
	gw.RegisterAccountServiceServer(s, &handler.AccountHandler{
		AccountUsecases: new(usecases.AccountUsecases),
	})

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:9090")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = gw.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

}

[]
fasdfasdfasdfas (fasfasd fasdfas, fasdfasd fasdfasd ,fasdfasdfas) {

}