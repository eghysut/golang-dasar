package main

import (
    "fmt"
)

// Ruang Lingkup Variabel Golang yang Didefinisikan oleh Tanda Kurung Penjepit
// Dalam bahasa Go, ruang lingkup variabel mengacu pada wilayah kode di mana variabel tersebut dapat diakses.
// Ruang lingkup variabel ditentukan oleh tempat variabel tersebut dideklarasikan.

// Variabel yang dideklarasikan di dalam blok (misalnya, di dalam fungsi atau struktur kontrol seperti if, for, atau switch) 
// hanya dapat diakses di dalam blok tersebut dan blok-blok bersarang. 
// Setelah eksekusi meninggalkan blok, variabel tersebut keluar dari cakupan dan tidak dapat diakses lagi.

var s = "Hello"

func main() {
    fmt.Println(s)
    var x = true
    if x {
        y := 100
        if x != false {
            fmt.Println("s:", s)
            fmt.Println("x:", x)
            fmt.Println("y:", y)
            var x = "BOM"
            fmt.Println("x:", x)
        }
        fmt.Println("y:", y)
    } 
    fmt.Println("x:", x) 
}
