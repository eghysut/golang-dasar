//contoh penggunaan berbagai cara mendefinisikan variabel dalam bahasa pemrograman Go (Golang).

package main

import "fmt"

// Catatan: nama variabel unik(jika ada nama variabel yang sama maka akan error).
// nilai variabel bisa di ubah dengan tipe data yang sama(string dengan string, int dengan int).
// variabel di golang tidak mesti menggunakan nama keyword tipe data, secara otomatis golang akan memberikan nama tipe data tersebut sesuai sistem os anda, jika anda membuat variabel yang nilainya int:
// var number = 10 (os x_86) int32
// var number = 10 (os x_64) int64


func main() {
    // variabel dengan keyword tipe data string
    var nama1 string = "Hello"
    fmt.Println(nama1)
    // Output: Hello

    // variabel tanpa keyword tipe data 
    var nama2 = "World"
    fmt.Printf("%s bertipe %T", nama2, nama2)
    // Output: World bertipe string

    fmt.Println("\n")

    // variabel dengan grouping var
    var (
	x = "Grouping x"
	y = "Grouping y"
    )

    fmt.Println(x)  //Output: Grouping x
    fmt.Println(y)  //Output: Grouping y

    // variabel inisialisasi dengan nilai kosong
    var nama3 string
    nama3 = "Inisialisasi dengan nilai kosong"
    fmt.Println(nama3)
    // Output: Inisialisasi dengan nilai kosong

    // variabel dengan keyword var tidak wajib 
    // anda bisa menggunakan keyword := tanpa harus
    // inisialisasi var dan nama tipe data. 
    // nama_var := "nilai_var"

    nama := "alice"
    usia := 20
    fmt.Printf("%s bertipe %T\n", nama, nama)
    // Output: alice bertipe string
    fmt.Printf("%d bertipe %T", usia, usia)
    // Output: 20 bertipe int

}
