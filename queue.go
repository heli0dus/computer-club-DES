package main

import "fmt"

type ClientQueue struct {
	queue []string
	size  int
}

func (q *ClientQueue) Push(client string) {
	q.queue = append(q.queue, client)
	q.size++
}

func (q *ClientQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ClientQueue) Pop() (string, error) {
	if q.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	q.size--
	return result, nil
}

func (q *ClientQueue) Remove(client string) bool {
	for ind, val := range q.queue {
		if val == client {
			q.queue = append(q.queue[:ind], q.queue[ind+1:]...)
			q.size--
			return true
		}
	}
	return false
}
