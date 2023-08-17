package main

import (
    "fmt"
    "unsafe"
)

//contoh penggunaan tipe data bool dalam bahasa pemrograman Go (Golang).

func main() {
    // tipe variabel dan inisialisasi
    var b bool = true
    fmt.Printf("b adalah %v bertipe %T\n", b, b)

    // mengubah nilai variabel b menjadi true
    b = false
    fmt.Printf("b adalah %v bertipe %T\n", b, b)


    // nilai boolean dengan nilai kosong
    var status bool
    fmt.Println("status:", status)

    // menampilkan ukuran dari boolean
    x := true
    y := false
    fmt.Printf("\nnilai x:%d nilai y:%d", unsafe.Sizeof(x), unsafe.Sizeof(y))
}
