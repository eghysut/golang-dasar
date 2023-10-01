/*package main

import "fmt"

// Polymorphism berarti "Nama yang sama dengan banyak bentuk".

// Deklarasi struct Mahasiswa dengan field "Nama" dan "Ipk"
type Mahasiswa struct{
    Nama string
    Ipk float64
}

// Deklarasi struct Mahasiswa dengan field "Nama" dan "Gaji"
type Karyawan struct{
    Nama string
    Gaji float64
}

// Method display()
func (m Mahasiswa) display() {
    fmt.Printf("nama: %s ipk: %.1f\n", m.Nama, m.Ipk)
}

// Method display() 
func (k Karyawan) display() {
    fmt.Printf("nama: %s gaji: $%.2f\n", k.Nama, k.Gaji)
}


func main() {
    // Membuat instance objek Mahasiswa
    mhs := Mahasiswa{Nama: "Alice", Ipk: 2.8}

    // Membuat instance objek Karyawan
    kar := Karyawan{Nama: "Carl", Gaji: 22.5}

    // Memanggil method display()
    mhs.display()
    // Output:
    // nama: Alice ipk: 2.8

    // Memanggil method display()
    kar.display()
    // Output:
    // nama: Carl gaji: $22.50

}
