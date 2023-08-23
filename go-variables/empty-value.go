package main

import "fmt"

// Nilai kosong
// Variabel yang dideklarasikan tanpa nilai awal diberikan nilai kosong.

// Nilai kosong adalah:
// 0 untuk tipe numerik,
// false untuk tipe boolean, dan
// "" (string kosong) untuk string.

func main() {
    var i int
    var f float64
    var b bool
    var s string

    fmt.Println("int:", i)
    fmt.Println("float:", f)
    fmt.Println("bool:", b)
    fmt.Printf("string: %q\n", s)
    // Output:
    // int: 0
    // float: 0
    // bool: false
    //string: ""
}
