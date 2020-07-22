package main

import (
	"dataflow/db/mysql"
	"dataflow/pkg/api/middleware"
	tx "dataflow/pkg/infrastructure/mysql"
	"dataflow/pkg/pb"
	"fmt"
	"log"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	// DBインスタンスの作成
	dbInstance := mysql.CreateSQLInstance()
	defer dbInstance.Close()

	// トランザクションマネージャーの作成
	masterTxManager := tx.NewDBMasterTxManager(dbInstance)

	// APIインスタンスの作成
	userAPI := InitUserAPI(masterTxManager)

	// firebase middlewareの作成
	firebaseClient := middleware.CreateFirebaseInstance()

	// gRPC: 8080
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Run server port: %d", port)

	// gRPC Server Option Set
	ops := make([]grpc.ServerOption, 0)
	ops = append(ops, grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(firebaseClient.MiddlewareFunc()),
			middleware.UnaryErrorHandling(),
		)),
	)
	ops = append(ops, grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(firebaseClient.MiddlewareFunc()),
			middleware.StreamErrorHandling(),
		)),
	)
	ops = append(ops, grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    5 * time.Second,
		Timeout: 5 * time.Hour,
	}))
	grpcServer := grpc.NewServer(
		ops...,
	)

	// User Service
	pb.RegisterUserServiceServer(grpcServer, &userAPI)

	// Serve
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
