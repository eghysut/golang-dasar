package main

import "fmt"

// Pointer indirection(pengarahan) adalah konsep yang memungkinkan Anda mengakses 
// nilai yang disimpan di alamat memori yang ditunjuk oleh pointer.
// Ini adalah salah satu fitur yang kuat dalam bahasa pemrograman Go (Golang),
// yang memungkinkan Anda bekerja dengan data yang lebih efisien dan fleksibel.
// Pointer indirection terjadi saat Anda menggunakan operator asterisk (*) untuk mengakses nilai yang ditunjuk oleh pointer.

func main() {
    // Deklarasi variabel integer 
    var angka int = 10
    fmt.Println("angka:", angka)
    //fmt.Println("address:", &angka)

    // Pointer ke variabel angka
    var x *int = &angka
    *x = 20
    fmt.Println("nilai x:", *x)
    //fmt.Println("address angka:", x)
    //fmt.Println("address x:", &x)

    // Pointer ke variabel x
    var y **int = &x
    **y = 30
    fmt.Println("nilai y:", **y)
    //fmt.Println("address angka:", *y)
    //fmt.Println("address x:", y)
    //fmt.Println("address y:", &y)

    // Pointer ke variabel y
    var z ***int = &y
    ***z = 40
    fmt.Println("nilai z:", ***z)
    //fmt.Println("address angka:", **z)
    //fmt.Println("address x:", *z)
    //fmt.Println("address y:", z)
    //fmt.Println("address z:", &z)
}
