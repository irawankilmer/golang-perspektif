package goroutine

/**
Contoh penerapan goroutine menggunakan time
dengan menggunakan time kita harus menentukan waktu di tiap funciton
secara manual
*/
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
		time.Sleep(300 * time.Millisecond)
	}
}

func Rpanggil() {
	go RoutineSatu()
	go RoutineDua()

	time.Sleep(2 * time.Second)
}
