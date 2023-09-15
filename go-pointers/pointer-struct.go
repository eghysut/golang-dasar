package main

import "fmt"

// Di Go, pointer ke struktur(struct) adalah salah satu konsep yang penting. 
// Pointer digunakan untuk mengakses dan memanipulasi data yang disimpan dalam struktur secara efisien.
// Mari kita bahas apa itu pointer ke struct di Go.

// Pointer ke Struct:
// Sebuah pointer ke struct adalah variabel yang menyimpan alamat memori dari instance struktur(struct) di dalamnya. 
// Ini memungkinkan kita untuk mengakses, mengubah, dan memanipulasi data yang terkandung dalam struct tanpa 
// perlu membuat salinan struktur tersebut. 
// Penggunaan pointer ke struct sangat berguna ketika Anda bekerja dengan data yang besar atau ketika 
// Anda ingin memodifikasi data asli dari luar fungsi.

// Mendefinisikan struct Mahasiswa
type Mahasiswa struct{
    Nama string
}

func main() {
    // Membuat instance struct
    var mhs Mahasiswa = Mahasiswa{Nama: "Alice"}
    fmt.Println("Nama:", mhs.Nama)
    // Output: Nama: Alice

    // Deklarasi pointer secara eksplisit
    // Membuat pointer ke struct
    var x *Mahasiswa = &mhs

    // Mengubah data dalam struct melalui pointer
    (*x).Nama = "Carl"
    fmt.Println("Mengubah Nama:", mhs.Nama)
    // Output: Mengubah Nama: Carl

    // Deklarasi pointer secara implisit
    // Membuat pointer ke struct 
    y := &mhs

    // Mengubah data dalam struct melalui pointer
    y.Nama = "Eliot"
    fmt.Println("Mengubah Nama:", mhs.Nama)
    // Output: Mengubah Nama: Eliot

}
