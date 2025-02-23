package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/GOLANG-NINJA/crud-app/examplepb" // Импортируем сгенерированный код

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Реализация сервиса Greeter
type server struct {
	pb.UnimplementedGreeterServer
}

// Реализация метода SayHello
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Получен запрос: %s", req.Name) // Логируем запрос
	message := fmt.Sprintf("Привет, %s!", req.Name)
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	// Создаём TCP-сервер
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	// Создаём gRPC-сервер
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterGreeterServer(grpcServer, &server{}) // Регистрируем сервер

	log.Println("gRPC-сервер запущен на порту 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при работе сервера: %v", err)
	}
}
