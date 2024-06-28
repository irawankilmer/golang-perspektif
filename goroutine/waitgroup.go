package goroutine

import (
	"fmt"
	"sync"
	"time"
)

/**
Penggunaan goroutine sebelumnya memang sudah benar.
Tapi kendalanya di penetapan waktu
kalau kita salah dalam menetapkan waktu malah membuat aplikasi menjadi lambat
golang menyediakan satu fungsi untuk membuat penetapan waktu menjadi otomatis
semuanya sudah dilakukan oleh go itu sendiri
nama fungsinya yaitu sync.WaitGroup()
mari kita coba
*/

func WgNumber() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		/**
		Disini saya menggunakan time untuk contoh saja
		supaya terlihat beberapa fungsi dipanggil secara bersamaan oleh goroutine
		harusnya tidak perlu pakai time seperti ini
		*/
		time.Sleep(50 * time.Millisecond)
	}
}

func WgHuruf() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
		/**
		Disini saya menggunakan time untuk contoh saja
		supaya terlihat beberapa fungsi dipanggil secara bersamaan oleh goroutine
		harusnya tidak perlu pakai time seperti ini
		*/
		time.Sleep(50 * time.Millisecond)
	}
}

func WgPanggilSatu() {
	/**
	Disini kita akan menambahkan fungsi ke goroutine satu per satu
	*/
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		WgNumber()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		WgHuruf()
	}()

	wg.Wait()
}

/**
Diatas kita sudah berhasil mengatasi permasalahan pada saat pengunaan time secara manual pada goroutine
Tapi di masing - masing functionnya masih menggunakan tim
itu hanya untuk contoh saja, untuk melihat bahwa beberapa function benar - benar di panggil bersamaan
sekarang kita coba tanpa menggunakan time sama sekali, bahkan didalam function
*/

func WgNumb() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func WgStr() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
	}
}

/*
*
Kalau kita jalankan program ini
kelihatannya 2 fungsi tidak dipanggil secara bersamaan menggunakan goroutine
yang terlihat adalah fungsi dua dulu(WgStr) yang dipanggil
kemudian fungsi satu(WgNumb)
Tapi pada kenyataannya kedua fungsi tersebut dipanggil secara bersamaan
*/
func WgPanggilDua() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		WgNumb()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		WgStr()
	}()

	wg.Wait()
}

/**
Penggunaan waitGroup diatas memang sudah berhasil memecahkan masalah penggunaan time secara manual
tapi kalau kita perhatikan, ketika kita ingin menjalankan goroutine
kita harus memanggil function nya satu persatu
bagaimana kalau misalkan kita ingin menjalankan banyak function sekaligus?
Pastinya agak ribet juga.
Nahh alternatif sementara nya, kenapa sementara?
karena ada alternatif yang jauh lebih bagus dengan menggunakan tambahan package go yang lain
tapi itu akan kita bahas di commit berikutnya
Ide nya, disini saya akan menggunakan looping
untuk memasukan fungsi kedalam goroutine
*/

func WgNumbDua() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, "Func ke tiga")
	}
}

func WgPanggilTiga() {
	var wg sync.WaitGroup

	functions := []func(){
		WgNumb,
		WgStr,
		WgNumbDua,
	}

	for _, function := range functions {
		wg.Add(1)
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(function)
	}

	wg.Wait()
}
