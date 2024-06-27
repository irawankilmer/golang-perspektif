package main

import (
	"goperspektif/goroutine"
	"time"
)

func main() {
	go goroutine.RoutineSatu()
	go goroutine.RoutineDua()

	time.Sleep(2 * time.Second)
}
