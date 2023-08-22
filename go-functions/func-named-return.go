package main

import "fmt"

// Fungsi Named return value "README.txt"
// func nama_fungsi() (Param1 TypeData, Param2 TypeData) {
// kode program 
// return
//}

func getNamaUsia() (nama string, usia int) {
    nama = "Alice"
    usia = 23
    return
}

func getData() (nama, alamat, email string, usia int) {
    nama = "carl"
    usia = 20
    alamat = "Jakarta"
    email = "carl@gmail.com"
    return

}

func main() {

    x, y := getNamaUsia()
    fmt.Println("Nama:", x)
    fmt.Println("Usia:", y)
    // Output:
    // Nama: Alice
    // Usia: 23

    // membuat variabel sesuaikan urutan dari Return_Type
    nama, alamat, email, usia := getData()
    fmt.Println("Nama:", nama)
    fmt.Println("Usia:", usia)
    fmt.Println("Alamat:", alamat)
    fmt.Println("Email:", email)
    // Output:
    // Nama: carl
    // Usia: 20
    // Alamat: Jakarta
    // Email: carl@gmail.com

}
