package handlers

import (
	"context"
	"encoding/json"
	"log"
	"topup-service/dto"

	"github.com/segmentio/kafka-go"
)

func TopupPublish(ctx context.Context, data *dto.RequestPayload) error {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("host.docker.internal:29092"),
		Topic:                  "topup",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	d, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("topup"),
			Value: d,
		},
	)
	if err != nil {
		log.Println("failed to write messages:", err)
		return err
	}

	if err := w.Close(); err != nil {
		log.Println("failed to close writer:", err)
		return err
	}

	return nil
}
