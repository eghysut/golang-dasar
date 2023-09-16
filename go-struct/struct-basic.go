package main

import "fmt"

// Mendefinisikan struct Mahasiswa yang memiliki dua field: Nama dan Umur
type Mahasiswa struct{
    // untuk membuat field biasanya huruf pertama harus besar
    Nama string
    Umur int
}

func main() {
    // Membuat instance struct secara eksplisit
    var x Mahasiswa = Mahasiswa{Nama: "Alice", Umur: 20}
    fmt.Println(x)
    // Output: {Alice 20}

    // Membuat instance struct secara implisit
    y := Mahasiswa{Nama: "Carl", Umur: 22}
    fmt.Println(y)
    // Output: {Carl 22}

    // Mengakses struct fields menggunakan notasi titik (.)
    fmt.Println(y.Nama) // Output: Carl
    fmt.Println(y.Umur) // Output: 22
}
