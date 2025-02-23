package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/GOLANG-NINJA/crud-app/examplepb" // Импортируем сгенерированный код

	"google.golang.org/grpc"
)

func main() {
	// Подключаемся к серверу gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn) // Создаём gRPC-клиент

	// Создаём контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Отправляем запрос
	req := &pb.HelloRequest{Name: "Egori"}
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Ошибка при вызове SayHello: %v", err)
	}

	// Выводим ответ от сервера
	fmt.Println("Ответ сервера:", resp.Message)
}
