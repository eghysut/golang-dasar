package main

import "fmt"

// Membuat struct Mahasiswa yang memiliki fields Nama
type Mahasiswa struct{
    Nama string
}

// Membuat method CetakNama untuk struct Mahasiswa
func (mhs Mahasiswa) CetakNama() {
    fmt.Println(mhs.Nama)
}

// Membuat method "pointer" SetNama untuk struct Mahasiswa
func (mhs *Mahasiswa) SetNama(nama_baru string) {
    mhs.Nama = nama_baru
}

// Membuat fungsi FunCetakNama untuk struct Mahasiswa
func FunCetakNama(mhs Mahasiswa) {
    fmt.Println(mhs.Nama)
}

// Membuat fungsi FuncSetNama untuk struct Mahasiswa
func FuncSetNama(mhs *Mahasiswa) {
    mhs.Nama = mhs.Nama
}

func main() {
    // Membuat instance struct Mahasiswa
    var val Mahasiswa = Mahasiswa{Nama: "Alice"}
    // Memanggil fields Nama menggunakan notasi titik(.)
    fmt.Println(val.Nama)
    // Output: Alice

    // Memanggil Method CetakNama() pada struct Mahasiswa
    val.CetakNama()
    // Output: Alice

    // Memanggil Method SetNama() pada struct Mahasiswa
    // Mengubah nilai dari variabel val
    val.SetNama("Hello World")
    val.CetakNama()
    // Output: Hello World

    fmt.Println(val.Nama)
    // Output: Hello World

    // Pointer ke variabel val
    var ptr *Mahasiswa = &val
    // Mengubah nilai dari variabel val
    ptr.Nama = "Hello Test"
    FuncSetNama(ptr) // Memanggil fungsi FuncSetNama untuk mengubah nilai
    FunCetakNama(*ptr) // Memanggil fungsi FunCetakNama untuk menampilkan nilai
    // Output: Hello Test
}
