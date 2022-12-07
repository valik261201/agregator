package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func postProducer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)

	// add orders to the end of the queue
	ordersProducer.Enqueue(order)

	json.NewEncoder(w).Encode(&order)

	ord, err := PrettyStruct(order)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\nAgregator recieved the data:\n", ord)

	///go performPostRequestToProducer(order)
}

func postConsumer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)

	// add orders to the end of the queue
	ordersConsumer.Enqueue(order)

	json.NewEncoder(w).Encode(&order)

	ord, err := PrettyStruct(order)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\nAggregator recieved the data 1:\n", ord)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/aggregator", postProducer).Methods("POST")
	router.HandleFunc("/aggregator", postConsumer).Methods("POST")

	go SendToConsumer()
	go SendToProducer()

	http.ListenAndServe(":8080", router)
}
