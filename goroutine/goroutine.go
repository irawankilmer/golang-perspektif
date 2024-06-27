package goroutine

import (
	"fmt"
	"time"
)

func RoutineSatu() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "Fungsi satu")
		time.Sleep(200 * time.Millisecond)
	}
}

func RoutineDua() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c Fungsi dua\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}
