package main

import (
	"fmt"
	"homework/5/event"
	"time"
)

func generator(ch chan int) {
	count := 0
	for {
		time.Sleep(2 * time.Second)
		ch <- count
		count++
	}
}

func multiplier(input chan int, output chan interface{}) {
	for {
		val := <-input
		output <- val * 2
	}
}

func spammer(ch chan interface{}) {
	for {
		time.Sleep(1 * time.Second)
		ch <- "spam"
	}
}

func printer(ch chan interface{}) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	// 1
	// ch1 := make(chan int)
	// ch2 := make(chan interface{})

	// go generator(ch1)
	// go multiplier(ch1, ch2)
	// go spammer(ch2)
	// printer(ch2)

	// 2
	eventHandler := event.NewEventHandler()
	eventHandler.Start()

	eventHandler.RegisterHandler(event.SystemStarted, func() {
		msg := "-> System started! (start time: %s) \n"
		fmt.Printf(msg, time.Now().String())
	})
	eventHandler.RegisterHandler(event.Connected, func() {
		fmt.Println("-> Connection done!")
	})

	eventHandler.Send(event.SystemStarted)
	time.Sleep(time.Second * 2)
	eventHandler.Send(event.Connected)

	eventHandler.Stop()
}
