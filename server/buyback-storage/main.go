package main

import (
	"buyback-storage/configs"
	"buyback-storage/dto"
	"buyback-storage/handlers"
	"buyback-storage/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	buybackHandler := handlers.NewBuybackHandler(transactionRepository)

	conf := kafka.ReaderConfig{
		Brokers:     []string{"host.docker.internal:29092"},
		Topic:       "buyback",
		GroupID:     "buyback",
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
				data := dto.Buyback{}

				_ = json.Unmarshal(message.Value, &data)

				err := buybackHandler.Buyback(&data)
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
