package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/adapters/controllers"
	handler "github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/adapters/repositories"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	gw "github.com/sadensmol/article_my_clean_architecture_go_application/api/proto/v1"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/services"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	err := godotenv.Load("infrastructure/local/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var connectString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	accountRepository := handler.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	accountUsecases := usecases.NewAccountUsecases(accountService)
	accountHandler := controllers.NewAccountHandler(accountUsecases)

	gw.RegisterAccountServiceServer(s, accountHandler)

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
