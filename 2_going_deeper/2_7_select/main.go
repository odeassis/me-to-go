package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(total chan int, exit chan bool) {
	value := rand.Intn(10)

	for {
		select {
		case total <- value:
			value += rand.Intn(10)
		case <- exit:
			fmt.Println("Exit :(")
			return
		}
	}
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		channel1 <- "Mensagem #1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "Mensagem #2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- channel1:
			fmt.Println(msg1)
		case msg2 := <- channel2:
			fmt.Println(msg2)
		}
	}

	// ################################################

	//rand.NewSource(time.Now().UnixNano())
	total := make(chan int)
	exit := make(chan bool)

	
	go func()  {
		for i := 0; i < 5; i++ {
			fmt.Println("Received: ", <-total)
			time.Sleep(1 * time.Second)
		}
		
		exit <- true
		}()
		
	go sum(total, exit)

	time.Sleep(12 * time.Second)

}