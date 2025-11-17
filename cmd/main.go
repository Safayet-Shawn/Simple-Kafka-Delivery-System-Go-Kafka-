package main

import (
	"github.com/safayet-shawn/Simple-Kafka-Delivery/kafka"
	"github.com/safayet-shawn/Simple-Kafka-Delivery/model"
	"github.com/safayet-shawn/Simple-Kafka-Delivery/worker"
)

func main() {
	workerQueue := make(chan model.Order, 100)

	// produce first
	orders := []model.Order{
		{Id: 1, Order: []string{"burger"}, Address: "Dhaka"},
		{Id: 2, Order: []string{"pizza"}, Address: "Chittagong"},
		{Id: 3, Order: []string{"pasta"}, Address: "Khulna"},
		{Id: 4, Order: []string{"sandwich"}, Address: "Rajshahi"},
		{Id: 5, Order: []string{"fries"}, Address: "Sylhet"},
		{Id: 6, Order: []string{"Kacci"}, Address: "Narsingdi"},
	}
	//Start kafka producer , which-> writing /inserting all order to the kafka
	for _, o := range orders {
		kafka.ProduceOrder(o)
	}
	//Start kafka consumer parallally, which ->Read message/order from kafka by kafka topic
	// and then send to the WorkerQueue
	go kafka.StartConsumer(workerQueue)
	// Here Number of Worker is: 5, and workerqueue contain the job/order
	//which need to done/processed ,all worker is individual go routine that work parallay
	worker.StartWorker(5, workerQueue)
	select {}
}
