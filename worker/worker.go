package worker

import (
	"fmt"
	"time"

	"github.com/safayet-shawn/Simple-Kafka-Delivery/model"
)

func StartWorker(workerCount int, jobs chan model.Order) {
	for i := 1; i <= workerCount; i++ {
		go func(id int) { // this is the worker
			for job := range jobs {
				ProcessOrder(id, job)
			}
		}(i)
	}
}
func ProcessOrder(WorkerId int, Orders model.Order) {
	fmt.Printf("Worker %v Processing Order %v \n", WorkerId, Orders.Id)
	time.Sleep(2 * time.Second)
	if Orders.Address == "" {
		fmt.Printf("Worker %v Failed to process Order %v ", WorkerId, Orders.Id)
	}
	fmt.Printf("Worker %v Successfully processed Order %v which is: %v at %v\n", WorkerId, Orders.Id, Orders.Order, Orders.Address)

}
