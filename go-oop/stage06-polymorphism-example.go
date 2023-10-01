/*package main

import "fmt"

// Polymorphism berarti "Nama yang sama dengan banyak bentuk".

// Deklarasi struct Mahasiswa dengan field "Nama" dan "Ipk"
type Mahasiswa struct{
    Nama string
    Ipk float64
}

// Deklarasi struct Karyawan dengan field "Nama" dan "Gaji" 
type Karyawan struct{
    Nama string
    Gaji float64
}

// Method display()
func (m Mahasiswa) display() {
    fmt.Printf("nama: %s ipk: %.2f\n", m.Nama, m.Ipk)
}

// Method display() 
func (k Karyawan) display() {
    fmt.Printf("nama: %s gaji: $%.2f\n", k.Nama, k.Gaji)
}

// Deklarasi interface DataInf dengan kontrak method display()
type DataInf interface{
    display()
}

func main() {

    // Membuat instance objek Mahasiswa
    mhs := Mahasiswa{Nama: "alice", Ipk: 2.6}
    // Membuat instance objek Karyawan
    kar := Karyawan{Nama: "carl", Gaji: 25}

    var koleksi []DataInf = []DataInf{mhs, kar}
    for _, v := range koleksi {
        v.display()
    }
    // Output:
    // nama: alice ipk: 2.60
    // nama: carl gaji: $25.00

    var mhsInf DataInf = DataInf(mhs)
    var karInf DataInf = DataInf(kar)
    mhsInf.display() // Output: nama: alice ipk: 2.60
    karInf.display() // Output: nama: carl gaji: $25.00

}
