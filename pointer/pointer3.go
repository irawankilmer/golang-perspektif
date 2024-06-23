package pointer

/**
 * Contoh contoh penggunaan pointer di function
 */

import (
	"fmt"
)

// Contoh 1
func Update(nilai *int, num int) {
	*nilai = num
}

func Panggil() {
	nomor := 55

	// Sebelum di update
	fmt.Println(nomor)

	// Sesudah di update
	Update(&nomor, 78)
	fmt.Println(nomor)
}

// Contoh 2
func Tarik(value *int) {
	*value = 76
}

func OkDah() {
	values := 89

	fmt.Println(values) // sebelum di update

	Tarik(&values)
	fmt.Println(values) // sesudah di update
}

/**
 * Contoh 3
 * return pointer dari function
 */
func Hellos() *string {
	pesan := "Irawan Kilmer"

	return &pesan
}

func YoAyo() {
	// memanggil function
	hehelo := Hellos()
	fmt.Println("Alamat memori", hehelo)
	fmt.Println("Value", *hehelo)
}

/**
 * Contoh call by value dan reference
 */
func CallByValue(value int) {
	value = 30

	fmt.Println(value)
}

func CallByReference(value *int) {
	*value = 56

	fmt.Println(*value)
}

func PaFer() {
	var number int

	// Call By Value
	CallByValue(number)

	// Call By Reference
	CallByReference(&number)
}
