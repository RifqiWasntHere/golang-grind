package go_routine

import (
	"fmt"
	"time"
)

func main() {
	TestGoRoutine()
}

func RunHelloRifqi() {
	fmt.Println("Hello Rifqi")
}

func TestGoRoutine() {
	go RunHelloRifqi()
	fmt.Println("This is the end of the program")

	time.Sleep(1 * time.Second) // Sleep is added in order to give goroutine some time frame to execute, or else. it will left unfinished after the whole program ends
}
