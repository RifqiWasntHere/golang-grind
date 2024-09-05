package go_routine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	defer fmt.Println("Channel's Closed")

	go func() { //Anonymous function (hekerrrr)
		time.Sleep(2 * time.Second)
		channel <- "This is the payload"
		fmt.Println("Data transfer completed, Cihuy")
	}()

	data := <-channel //Fetching payload from channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func RunHello(channel chan string, name string) {
	time.Sleep(2 * time.Second)
	channel <- string(name)
}

func TestWithParameter(t *testing.T) {
	channel := make(chan string)
	// OR, we can make Buffered channel :
	// channel := make(chan string, 3)

	defer close(channel)

	go RunHello(channel, "rifqi")

	result := <-channel

	fmt.Println(result)

	time.Sleep(5 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai -")
}

func SelectChannelHelper(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Ini Payload"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go SelectChannelHelper(channel2)
	go SelectChannelHelper(channel1)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Received data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Received data from channel 2", data)
			counter++
		default:
			fmt.Println("Waiting for data...")
			time.Sleep(10 * time.Millisecond)
		}

		if counter == 2 {
			break
		}
	}

}
