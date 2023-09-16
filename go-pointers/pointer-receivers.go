package main

import "fmt"

// Pointer receivers(penerima) memungkinkan Anda untuk mengubah nilai asli dari tipe data dalam method,
// alih-alih hanya menerima salinan nilai tersebut.

// Dalam Go, method adalah fungsi yang terkait dengan tipe data tertentu, 
// dan mereka memungkinkan Anda untuk menambahkan perilaku atau operasi yang berkaitan dengan tipe data tersebut.
// Pointer receivers(penerima) adalah salah satu cara untuk memodifikasi nilai dari tipe data yang ada melalui method.

// Secara umum, Anda dapat memiliki dua jenis receivers dalam metode Go:
// Value Receivers (Non-Pointer Receivers): 
// Method dengan receiver jenis ini mengoperasikan salinan nilai dari tipe data. 
// Perubahan yang dilakukan pada nilai di dalam method tidak memengaruhi nilai asli yang ada di luar method.

// Pointer Receivers: 
// Method dengan receiver jenis ini mengoperasikan variabel pointer ke tipe data. 
// Perubahan yang dilakukan pada variabel yang ditunjuk oleh pointer dalam method akan memengaruhi nilai asli tipe data
// yang ada di luar method.

// Mendefinisikan struct Mahasiswa
type Mahasiswa struct{
    Nama string
}

// Method dengan value receivers (Non-Pointer receivers)
func (mhs Mahasiswa) SetNamaVal(nama_baru string) {
    mhs.Nama = nama_baru
}

// Method dengan pointer receivers
func (mhs *Mahasiswa) SetNamaPtr(nama_baru string) {
    mhs.Nama = nama_baru
}

func main() {
    // Membuat instance struct 
    val := Mahasiswa{Nama: "Hello"}
    fmt.Println(val.Nama)
    // Output: Hello

    // Panggil method dengan pointer receivers dan ubah nilainya
    val.SetNamaPtr("World")
    fmt.Println(val.Nama)
    // Output: World

    // Panggil method dengan Non-pointer receivers 
    val.SetNamaVal("Test")
    fmt.Println(val.Nama)
    // Output: World

}
