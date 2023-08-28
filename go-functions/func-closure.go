package main

import "fmt"

// Fungsi Closure

func main() {
    nama := "Hello"

    // Closure: fungsi yang mengakses variabel x
    fungsiClosure := func() {
        fmt.Println(nama)
    }

    // Mengubah nama
    nama = "Golang" 

    // Memanggil closure
    fungsiClosure() 
    // Output: Golang

    // contoh sederhana
    jumlah := 1
    next := func() {
	fmt.Println(jumlah)
    }

    for jumlah <= 3 {
	next()
	jumlah++
    }

    fmt.Println("global:", jumlah)
}
