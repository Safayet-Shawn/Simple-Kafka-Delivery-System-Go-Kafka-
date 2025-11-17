package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/safayet-shawn/Simple-Kafka-Delivery/model"
	"github.com/segmentio/kafka-go"
)

func StartConsumer(workerQueue chan model.Order) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "Orders",
		GroupID:        "G-test",
		MinBytes:       1,
		MaxBytes:       10e6,
		CommitInterval: time.Second, // automatically commits offsets every second
	})
	for { // FIX: LOOP
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Failed to read message:", err)
			continue
		}
		var val model.Order
		json.Unmarshal(msg.Value, &val)
		workerQueue <- val
	}
}
