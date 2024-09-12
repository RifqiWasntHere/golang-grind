package go_routine

import (
	"fmt"
	"time"
)

func consumer(id int, done <-chan struct{}) {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop() // Ensure timer is stopped

	select {
	case <-timer.C: // Timer expired
		fmt.Printf("Consumer %d timed out waiting for data.\n", id)
		return
	case <-done: // Producer signaled that data is ready
		fmt.Printf("Consumer %d is processing shared data after being signaled.\n", id)
		return
	}
}

func producer(done chan<- struct{}) {
	// time.Sleep(5 * time.Second) // Simulate a delay before data becomes available
	close(done) // Signal all consumers that data is available
}

// func main() {
// 	done := make(chan struct{}) // Channel to signal consumers

// 	go consumer(1, done)
// 	go consumer(2, done)

// 	// Simulate some work by the producer
// 	go producer(done)

// 	time.Sleep(6 * time.Second) // Give enough time for all operations to complete
// }
