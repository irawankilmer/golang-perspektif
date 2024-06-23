package pointer

import "fmt"

/**
 * Ketika kita membut suatu variable di go
 * Maka variable tersebut akan mempunyai alamat memori untuk menyimpan nilainya
 * Alamat memori tersebut berupa bilangan acak
 * Untuk mengetahui alamat memori suatu variable
 * Kita bisa menggunakan operator &
 */
func PointerSatu() {
	var nama string = "Irawan"

	fmt.Println("Menampilkan value ", nama)          // Irawan
	fmt.Println("Menampilkan alamat memori ", &nama) // 0xc000026080
}

/**
 * Diatas kita sudah bisa mengetahui alamat memori suatu variable
 * menggunakan operator &
 * Kita juga bisa mengakses alamat memori dari variable lain
 * menggunakan tanda *
 */
func PointerDua() {
	var nama string = "Irawan"
	var nami *string = &nama // Membuat pointer

	/**
	 * Sekarang variable nama dan nami mempunyai alamat memori yang sama
	 * yang sebenarnya variable nami merujuk ke alamat memori variable nama
	 */
	fmt.Println("Menampilkan alamat memori", &nama) // menampilkan alamat memori menggunakan &
	fmt.Println("Menampilkan alamat memori", nami)
}

/**
 * Sebelumnya kita sudah bisa mengakses alamat memori dari variable lain
 * untuk mengakses nilainya, pada saat pemanggilan variable diawali dengan tanda *
 */
func PointerTiga() {
	var nama string = "Irawan"
	var nami *string = &nama // Membuat pointer

	fmt.Println("Menampilkan value", nama)
	fmt.Println("Menampilkan value", *nami) // Menampilkan value menggunakan tanda *
}

/**
 * Sekarang kita coba bagaimana suatu value variable akan berubah
 * mengikuti variable lain yang valuenya diganti
 */
func PointerEmpat() {
	var nama string = "Irawan"
	var nami *string = &nama // Membuat pointer

	fmt.Println("Menampilkan value", nama)
	fmt.Println("Menampilkan value", *nami) // Menampilkan value menggunakan tanda *

	nama = "Hikamul Albi"                         // Mengganti value variable nama
	fmt.Println("Menampilkan value nama ", nama)  // Hikamul Albi
	fmt.Println("Menampilkan value nami ", *nami) // Hikamul Albi

	/**
	 * Untuk merubah value dari variable pointer agar masuk ke alamat memori variable nami
	 * kita harus menggunakan tanda * seperti contoh dibawah
	 */
	*nami = "Ghania Jilbi" // Mengganti value variable nami
	fmt.Println("Menampilkan value nama", nama)
	fmt.Println("Menampilkan value nami", *nami)
}
