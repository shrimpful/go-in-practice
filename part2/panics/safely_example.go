package main

import (
	"errors"
	"github.com/Masterminds/cookoo/safely"
	"time"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(1000*time.Millisecond)
}