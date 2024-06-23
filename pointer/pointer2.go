package pointer

import "fmt"

/**
 * Contoh penggunaan pointer dengan struct
 */

type Siswa struct {
	nama, kelas string
	nilai       int
}

func PointerLima() {
	var siswa1 Siswa = Siswa{ // bisa juga seperti ini siswa1 := Siswa{}
		nama:  "Irawan",
		kelas: "XII RPL 1",
		nilai: 89,
	}

	var siswa2 *Siswa = &siswa1
	// Bisa juga seperti ini
	siswa3 := &siswa1

	fmt.Println("Menampilkan value", siswa1) // {Irawan XII RPL 1 89}
	fmt.Println("Menampilkan value", siswa2) // &{Irawan XII RPL 1 89}
	fmt.Println("Menampilkan value", siswa3) // &{Irawan XII RPL 1 89}
}

func PointerEnam() {
	var siswa1 Siswa = Siswa{ // bisa juga seperti ini siswa1 := Siswa{}
		nama:  "Irawan",
		kelas: "XII RPL 1",
		nilai: 89,
	}

	var siswa2 *Siswa = &siswa1
	// Bisa juga seperti ini
	siswa3 := &siswa1

	fmt.Println("Menampilkan value", siswa1) // {Irawan XII RPL 1 89}
	fmt.Println("Menampilkan value", siswa2) // &{Irawan XII RPL 1 89}
	fmt.Println("Menampilkan value", siswa3) // &{Irawan XII RPL 1 89}

	fmt.Println("Mengganti data dari variable siswa2")
	siswa2.nama = "Hikamul Albi"
	siswa2.kelas = "XII RPL 4"
	siswa2.nilai = 91

	fmt.Println("Menampilkan value", siswa1) // {Irawan XII RPL 1 89}
	fmt.Println("Menampilkan value", siswa2) // &{Irawan XII RPL 1 89}
	fmt.Println("Menampilkan value", siswa3) // &{Irawan XII RPL 1 89}
}

/**
 * Mengganti nilai langsung dari struct
 */
func PointerTujuh() {
	var siswa1 Siswa = Siswa{ // bisa juga seperti ini siswa1 := Siswa{}
		nama:  "Irawan",
		kelas: "XII RPL 1",
		nilai: 89,
	}

	var siswa2 *Siswa = &siswa1

	// // Merubah value yang salah, ini akan masuk ke alamat memori yang baru
	// siswa2 = &Siswa{"Ghania Jilbi", "X RPL", 98}

	// /**
	//  * value siswa1 masih tetap yang tadi
	//  * yang berubah hanya siswa2
	//  */
	// fmt.Println("Nilai siswa 1", siswa1) // {Irawan XII RPL 1 89}
	// fmt.Println("Nilai siswa 2", siswa2) // &{Ghania Jilbi X RPL 94}

	/**
	 * Jika ingin merujuk ke alamat memori yang sama
	 * harus menambahkan tanda bintang sebelum variable siswa2
	 * dan menghilangkan operator & sebelum struct Siswa
	 */
	*siswa2 = Siswa{"Rhania Syaranopa", "XI RPL 7", 90}

	/**
	 * Sekarang data baru dari siswa 2
	 * akan dimasukan kealamat memori yang sama dengan siswa 1
	 */
	fmt.Println("Nilai siswa 1", siswa1) // {Rhania Syaranopa XI RPL 7 90}
	fmt.Println("Nilai siswa 2", siswa2) // &{Rhania Syaranopa XI RPL 7 90}
}

/**
 * Kesimpulan ganti value satu per satu
 * berdasarkan properti dari struct
 */
func PointerDelapan() {
	var siswa1 Siswa = Siswa{"Irawan", "XII RPL", 87}
	var siswa2 *Siswa = &siswa1 // Membuat pointer

	fmt.Println("===Sebelum ada perubahan nilai===")
	fmt.Println("Siswa 1", siswa1)
	fmt.Println("Siswa 2", siswa2)

	siswa2.nama = "Rani Syaranopa"
	siswa2.kelas = "XII RPL 5"
	siswa2.nilai = 93
	fmt.Println()

	fmt.Println("===Sesudah ada perubahan nilai===")
	fmt.Println("Siswa 1", siswa1)
	fmt.Println("Siswa 2", siswa2)
}

/**
 * Kesimpulan ganti value langsung dari struct
 */
func PointerSembilan() {
	var siswa1 Siswa = Siswa{"Irawan", "XII RPL", 87}
	var siswa2 *Siswa = &siswa1 // Membuat pointer

	fmt.Println("===Sebelum ada perubahan nilai===")
	fmt.Println("Siswa 1", siswa1)
	fmt.Println("Siswa 2", siswa2)

	// Sebelum nama variable harus pakai bintang
	*siswa2 = Siswa{"Ghania Jilbi", "XI RPL 4", 97}
	fmt.Println()

	fmt.Println("===Sesudah ada perubahan nilai===")
	fmt.Println("Siswa 1", siswa1)
	fmt.Println("Siswa 2", siswa2)
}

/**
 * Keyword new
 * selain menggunakan operator &
 * kita juga bisa menggunakan keyword new
 */

func PointerSepuluh() {
	var siswa = new(string)
	*siswa = "Hello World"

	fmt.Println("Alamat memori", siswa)
	fmt.Println("Value", *siswa)
}
