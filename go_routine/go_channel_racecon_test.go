package go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Race condition occures, supposed value: 1000000
func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				x = x + 1
			}
		}()

	}
	time.Sleep(1 * time.Second)
	fmt.Println("Iterasi :", x)
}

// Race condition Handled!, supposed value: 1000000
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock() //Write Lock : only one go routine is allowed to obtain the writer lock.
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Iterasi :", x)
}

// //////////////////////////////////////////////////////////// //
type Counter struct {
	mu    sync.RWMutex
	value int
}

// Increment increments the counter value
func (c *Counter) Increment() {
	c.mu.Lock()         // Lock for writing (exclusive access)
	defer c.mu.Unlock() // Ensure Unlock is called after writing
	c.value++
}

// Value returns the current value of the counter
func (c *Counter) Value() int {
	c.mu.RLock()         // Lock for reading (allows multiple readers)
	defer c.mu.RUnlock() // Ensure RUnlock is called after reading
	return c.value
}

func TestRWMutex(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Start 5 reader goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 100 * time.Millisecond) // Stagger the readers
			fmt.Printf("Reader %d: Counter value: %d\n", id, counter.Value())
		}(i)
	}

	// Introduce a writer in the middle of the reader goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(300 * time.Millisecond) // Make the writer wait so that some readers run first
			counter.Increment()
		}()
	}

	// Start more readers after the writer has started
	for i := 5; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 100 * time.Millisecond) // Stagger the readers
			fmt.Printf("Reader %d: Counter value: %d\n", id, counter.Value())
		}(i)
	}

	time.Sleep(5 * time.Second)
	counter.Increment()

	fmt.Printf("Counter value: %d\n", counter.Value())
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All readers and writer are done")
}
