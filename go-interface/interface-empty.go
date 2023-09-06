package main

import "fmt"

// Interface kosong
// Tipe interface yang tidak memiliki method dikenal juga dengan interface kosong:

// interface{}
// Sebuah interface kosong bisa menyimpan nilai dari tipe apapun. 
// (Setiap tipe mengimplementasikan paling tidak nol method.)

// Interface kosong digunakan pada kode yang menangani nilai yang tidak diketahui. 
// Sebagai contohnya, fungsi PrintData mengambil sejumlah argumen dengan tipe interface{}.

func PrintData(data interface{}) {
    fmt.Printf("%v bertipe %T\n", data, data)
}

func main() {
    // Menggunakan interface kosong dengan berbagai tipe data
    PrintData(14)	// Integer
    PrintData("Golang") // String
    PrintData(3.14)	// Float
    PrintData([]int{1,2,4}) // Slice
}
