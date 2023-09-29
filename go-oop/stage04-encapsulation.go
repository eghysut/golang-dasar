package main

import (
	"fmt"
	"go-oop/mahasiswa"
)

func main() {
    // Membuat instance objek
    mhs := mahasiswa.Mahasiswa{Nama: "alice"}

    // field Nama adalah huruf besar(exported) bisa diakses dari luar
    fmt.Println(mhs.Nama)
    // Output: alice

    // field usia adalah huruf kecil(unexported), 
    // sehingga hanya dapat diakses melalui method GetUsia dan SetUsia
    fmt.Println(mhs.GetUsia())
    // Output: 0

    // mengubah nilai usia
    mhs.SetUsia(20)
    fmt.Println(mhs.GetUsia())
    // Output: 20

}
