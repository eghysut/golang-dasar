package main

import "fmt"


// Kondisi If
// Struktur dasar dari "if" di Go adalah sebagai berikut:

//if kondisi {                                             // Perintah yang akan dijalankan jika kondisi benar (true)                                       
//}
func main() {
    // membuat if statement
    x := 5
    if x == 5 {
	fmt.Println("ok")
    }

    // jika kondisi boolean bernilai false
    if x > 10 {
	// kode program ini tidak di jalankan
	fmt.Println("True")
    }
    // akhir dari program
    fmt.Println("selesai")

}
