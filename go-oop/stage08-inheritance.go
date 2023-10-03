/*package main

import "fmt"

// single type embedding 
// Anda dapat embedding(menyematkan) sebuah tipe ke dalam tipe lain. 

// Kita memiliki tipe struct Manusia dengan dua field:
type Manusia struct{
    Nama string
    Umur int
}

// Method display() untuk menampilkan data
func (m Manusia) display() {
    fmt.Println(m.Nama, m.Umur)
}

// Seorang Mahasiswa dan Karyawan, keduanya adalah manusia; 
// kita bisa menyematkan tipe Manusia ke dalam kedua tipe tersebut:
type Mahasiswa struct{
    Manusia
    Jurusan string
}


type Karyawan struct{
    Manusia
    Gaji float64
}

func main() {
    // Membuat instance objek Mahasiswa
    mhs := Mahasiswa{
        Manusia: Manusia{Nama: "Alice", Umur: 20},
        Jurusan: "Teknik Komputer",
    }

    // Membuat instance objek Karyawan
    kar := Karyawan{
        Manusia: Manusia{Nama: "Carl", Umur: 30},
        Gaji: 25.4,
    }

    // Memanggil method display()
    mhs.display()
    // Output: Alice 20

    kar.display()
    // Output: Carl 30

}*/
