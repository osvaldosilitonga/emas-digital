package main

import (
	"context"
	"encoding/json"
	"fmt"
	"input-harga-storage/configs"
	"input-harga-storage/dto"
	"input-harga-storage/repository"
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

	hargaRepository := repository.NewHargaRepository(db)

	conf := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "input-harga",
		GroupID:     "input-harga",
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
				harga := dto.Harga{}

				_ = json.Unmarshal(message.Value, &harga)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()

				err := hargaRepository.Save(ctx, &harga)
				if err == nil {
					log.Printf("Data [%s] inserted successfully\n", harga.ID)
					break
				}

				log.Printf("Inserting data [%s] failed\n", harga.ID)
				log.Println("[ERROR]:", err.Error())
				log.Println("Retriying...")
				time.Sleep(time.Second * 3)
			}
		}
	}
}
