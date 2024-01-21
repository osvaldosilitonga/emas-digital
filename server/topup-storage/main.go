package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"topup-storage/configs"
	"topup-storage/dto"
	"topup-storage/handlers"
	"topup-storage/repository"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to load .env file")
	}
}

func main() {
	db := configs.InitDB()
	defer db.Close()

	transactionRepository := repository.NewTransactionRepository(db)

	topupHandler := handlers.NewTopupHandler(transactionRepository)

	conf := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "topup",
		GroupID:     "topup",
		StartOffset: kafka.LastOffset,
		// MaxBytes: 10, // the maximum batch size that the consumer will accept, default 1MB
	}

	reader := kafka.NewReader(conf)

	for {
		message, err := reader.ReadMessage(context.Background())
		if err == nil {
			fmt.Printf("Received message: %s\n", message.Value)

			timeRetries := time.Now().Add(time.Hour * 2).Unix()
			timeNow := time.Now().Unix()
			for timeNow < timeRetries {
				data := dto.Topup{}

				_ = json.Unmarshal(message.Value, &data)

				err := topupHandler.Topup(&data)
				if err == nil {
					log.Printf("Data [%s] inserted successfully\n", data.ReffID)
					break
				}

				log.Printf("Inserting data [%s] failed\n", data.ReffID)
				log.Println("[ERROR]:", err.Error())
				log.Println("Retriying...")
				time.Sleep(time.Second * 3)
			}
		}
	}
}