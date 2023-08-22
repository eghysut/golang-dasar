package main

import "fmt"

// Fungsi dengan parameter
// func nama_fungsi(Param1 Typedata, Param2 TypeData) {
// block kode fungsi

//}

// fungsi dengan 1 parameter
func getNama(nama string) {
    fmt.Println("Nama:", nama)
}

// fungsi dengan 2 parameter
func getInfo(nama string, alamat string) {
    fmt.Println("Nama:", nama)
    fmt.Println("Alamat:", alamat)

}

// jika tipe data dari nilai parameter sama
func getInfo2(nama, alamat  string) {
    fmt.Println("Nama:", nama)
    fmt.Println("Alamat:", alamat)
}

// fungsi dengan 2 parameter dari tipe data yang berbeda
func getInfo3(nama string, usia int) {
    fmt.Println("Nama:", nama)
    fmt.Println("Usia:", usia)
}

func main() {

    // panggil fungsi dan nilai parameternya
    getNama("Alice")
    // Output:
    // Nama: Alice

    getInfo("Carl", "Jakarta")
    // Output:
    // Nama: Carl
    // Alamat: Jakarta


    getInfo2("eliot", "Bali")
    // Output:
    // Nama: eliot
    // Alamat: Bali

    getInfo3("Golang", 14)
    // Output:
    // Nama: Golang
    // Usia: 14

}
