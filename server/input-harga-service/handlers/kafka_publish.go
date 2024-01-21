package handlers

import (
	"context"
	"encoding/json"
	"input-harga-service/dto"
	"log"

	"github.com/segmentio/kafka-go"
)

func HargaPublish(data *dto.RequestPayload) error {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "input-harga",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	d, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("harga"),
			Value: d,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return err
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
		return err
	}

	return nil
}
