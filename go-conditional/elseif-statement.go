package main

import "fmt"

// Kondisi Else If                                  `else if` digunakan ketika Anda memiliki beberapa kondisi yang ingin diuji secara berurutan.

//if kondisi {
    // kodisi boolean bernilai false
//} else if {
    // jika kondisi boolean if bernilai false
    // kode program ini di jalankan
//} else {
    // jila semua kondisi di atas tidak terpenuhi
    // kode program ini di jalankan
//}

func main() {
    x := 5
    if x > 10 {
	fmt.Println("if")
    } else if x < 10 {
	fmt.Println("else if")
    } else {
	fmt.Println("else")
    }
    // Output: else if
}
