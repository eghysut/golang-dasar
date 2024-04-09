//contoh penggunaan berbagai cara mendefinisikan variabel dalam bahasa pemrograman Go (Golang).

package main

import "fmt"

//Berikut adalah contoh lengkap yang menggabungkan konsep-konsep variabel dari "README.txt"
var hello = "variabel hello diluar fungsi main"

func main() {
    fmt.Println("hello:",hello)
    // Tipe Variabel dan Inisialisasi
    var nama string = "John"
    var usia int = 30
    var status bool = true

    // Type Inference
    score := 85
    pesan := "Hello"

    // Penggunaan Blank Identifier
    _, err := SomeFunction() // Mengabaikan nilai yang dikembalikan, hanya memeriksa error

    // Menampilkan nilai variabel
    fmt.Printf("Nama: %s\n", nama)
    fmt.Printf("Usia: %d\n", usia)
    fmt.Printf("Status: %v\n", status)
    fmt.Printf("Score: %d\n", score)
    fmt.Printf("Pesan: %s\n", pesan)
    fmt.Printf("Error: %v\n", err)

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

    nama4 := "alice"
    usia4 := 20
    fmt.Printf("%s bertipe %T\n", nama4, nama4)
    // Output: alice bertipe string
    fmt.Printf("%d bertipe %T\n", usia4, usia4)
    // Output: 20 bertipe int

    // mengubah isi variabel harus bertipe yang sama
    nama4 = "carl"
    usia4 = 25
    fmt.Println(nama4, usia4)
    // Output: carl 25

    // variabel tanpa nilai
    var kosong1 int
    var kosong2 string
    fmt.Println(kosong1)
    fmt.Println(kosong2)
    // Output: 0
    // Output: 

}


func SomeFunction() (int, error) {
	// Fungsi ini mengembalikan nilai dan error
	return 42, nil
}

// Catatan: nama variabel unik(jika ada nama variabel yang sama maka akan error).
// nilai variabel bisa di ubah dengan tipe data yang sama(string dengan string, int dengan int).
// variabel di golang tidak mesti menggunakan nama keyword tipe data, secara otomatis golang akan memberikan nama tipe data tersebut sesuai sistem os anda, jika anda membuat variabel yang nilainya int:
// var number = 10 (os x_86) int32
// var number = 10 (os x_64) int64
