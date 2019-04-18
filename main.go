package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			fmt.Println("You talk too much. Quit signal")
			return
		}
	}
}

func boring(msg string, quit chan bool) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		quit <- true
	}()
	return c // Return the channel to the caller.
}
