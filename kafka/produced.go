package kafka

import (
	"context"
	"encoding/json"

	"github.com/safayet-shawn/Simple-Kafka-Delivery/model"
	"github.com/segmentio/kafka-go"
)

func ProduceOrder(order model.Order) error {
	writer := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "Orders",
	}
	data, _ := json.Marshal(order)
	return writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: data,
		},
	)
}
