package main

import "fmt"

// Kondisi Else                                     Anda dapat menggunakan blok `else` untuk mengeksekusi perintah ketika kondisi `if` tidak benar (false).

// if kondisi {
    // kondisi boolean bernilai false,
    // kode program ke statement else
    // } else {
	// kode program di jalankan
//}

func main() {

    x := 5
    if x > 10 {
	fmt.Println("on")
    } else {
	fmt.Println("off")
    }
    // Output: off
}
