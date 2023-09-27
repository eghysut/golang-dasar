package main

import "fmt"

// Tipe Data Struktur (Structs): 
// Anda mendefinisikan tipe data struktur(struct) untuk merepresentasikan data. 
// Structs adalah kumpulan field(atribut) yang dapat memiliki tipe data yang berbeda-beda. 
// Ini analogi dengan objek dalam bahasa pemrograman lain.

// Mendefinisikan struct Mahasiswa dengan sekumpulan atribut(setiap atribut memiliki tipe)
// Atribut juga disebut properti, field(bidang), anggota data.
type Mahasiswa struct{
    Nama string
    Umur int
}


func main() {

    // Membuat instance objek dari struct Mahasiswa
    var mhs1 Mahasiswa = Mahasiswa{Nama: "Alice", Umur: 20}
    fmt.Println(mhs1)
    // Output: {Alice 20}

    // Membuat instance objek secara implisit
    mhs2 := Mahasiswa{Nama: "Carl", Umur: 22}
    fmt.Println(mhs2)
    // Output: {Carl 22}

    // Memanggil fields Nama menggunakan notasi titik(.)
    fmt.Println(mhs1.Nama)
    // Output: Alice

    fmt.Println(mhs2.Umur)
    // Output: 22

}
