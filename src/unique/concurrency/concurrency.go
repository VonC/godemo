package main

import "time"
import "fmt"

func main() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	fmt.Printf("Sending ball %d\n", Ball)
	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int) {
	for {
		ball := <-table
		fmt.Printf("Receive ball %d\n", ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Send back ball %d\n", ball)
		table <- ball
	}
}
