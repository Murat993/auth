package main

import (
	"context"
	"github.com/Murat993/auth/pkg/user_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:5000"
	noteID  = 12
)

func main() {
	// Используется для подключения Dial. insecure.NewCredentials - подключение не секъюрное
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := user_v1.NewUserV1Client(conn) // Подключаем клиент. user_v1 - это сгенерированный пакет pb

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Get(ctx, &user_v1.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get note by id: %v", err)
	}

	log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", r.GetName()))
}
