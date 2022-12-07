package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/semaphore"
)

type Order struct {
	Id         int   `json:"id"`
	Items      []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max-wait"`
	PickUpTime int   `json:"pick-up-time"`
}

func SendToProducer() {
	for {
		for !ordersConsumer.isEmpty() {
			go performPostRequestToProducer()
			time.Sleep(time.Second * 1)
		}
	}
}

func SendToConsumer() {
	for {
		for !ordersProducer.isEmpty() {
			go performPostRequestToConsumer()
			time.Sleep(time.Second * 1)
		}
	}
}

func performPostRequestToProducer() {
	const myUrl = "http://localhost:3030/producer"

	// nr of items we can send before the aggregator will block
	// nr of concurrent items (10)
	sem := semaphore.NewWeighted(10)

	if err := sem.Acquire(context.Background(), 1); err != nil {
		log.Fatal(err)
	}
	// return the first order form the queue
	order := ordersConsumer.Dequeue()

	defer sem.Release(1)

	var requestBody, _ = json.Marshal(order)

	time.Sleep(time.Second * 3)

	fmt.Printf("\nData %v was sent to the Producer\n", string(requestBody))
	response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	time.Sleep(time.Second * 1)
}

func performPostRequestToConsumer() {
	const myUrl = "http://localhost:5050/consumer"

	// nr of items we can send before the aggregator will block
	// nr of concurrent items (10)
	sem := semaphore.NewWeighted(10)

	if err := sem.Acquire(context.Background(), 1); err != nil {
		log.Fatal(err)
	}
	// return the first order form the queue
	order := ordersProducer.Dequeue()

	defer sem.Release(1)

	var requestBody, _ = json.Marshal(order)
	time.Sleep(time.Second * 1)

	fmt.Printf("\nOrder %v was sent to the Consumer\n", string(requestBody))
	response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
}
