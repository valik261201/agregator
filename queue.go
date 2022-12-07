package main

// Queue represents a queue that holds a slice
var ordersProducer Queue
var ordersConsumer Queue

type Queue struct {
	elements []Order
}

func (q *Queue) Enqueue(order Order) {
	q.elements = append(q.elements, order)
}

func (q *Queue) isEmpty() bool {
	return len(q.elements) == 0
}

func (ol *Queue) getSize() int {
	return len(ol.elements)
}

func (q *Queue) Dequeue() *Order {
	if q.isEmpty() {
		return nil
	}
	order := q.elements[0]
	if q.getSize() == 1 {
		q.elements = nil
		return &order
	}

	// discard top element
	q.elements = q.elements[1:]
	return &order
}
