package main

import "time"

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	for {
		select {
		case <-ch:
			println("got message")
		case <-timeout:
			println("time out")
			return
		default:
			println("yawn")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("sent and closed")
}
