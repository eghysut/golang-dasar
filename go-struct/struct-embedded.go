package main

import "fmt"


type Alamat struct{
    Kota string
    Negara string
}

type Karyawan struct{
    Nama string
    Umur int
    // struct Alamat Embedded(Tersisip)
    //Alamat Alamat // Deklarasi secara eksplisit
    Alamat // Deklarasi secara anonymous
}


func main() {
    // Membuat instance struct secara eksplisit
    var x Karyawan = Karyawan{
        Nama: "Alice",
        Umur: 22,
        Alamat: Alamat{Kota: "Jakarta", Negara: "Indonesia"},
    }

    // Mengakses struct fields
    fmt.Println("Nama:", x.Nama)
    fmt.Println("Umur:", x.Umur)
    fmt.Println("Kota:", x.Kota)
    fmt.Println("Negara:", x.Negara)
    // Output:
    // Nama: Alice
    // Umur: 22
    // Kota: Jakarta
    // Negara: Indonesia

    // Membuat instance struct secara implisit
    y := Karyawan{
        Nama: "Carl",
        Umur: 25,
        Alamat: Alamat{Kota: "Bandung", Negara: "Indonesia"},
    }

    // Mengakses struct fields
    fmt.Println("Nama:", y.Nama)
    fmt.Println("Umur:", y.Umur)
    fmt.Println("Kota:", y.Kota)
    fmt.Println("Negara:", y.Negara)
    // Output:
    // Nama: Carl
    // Umur: 25
    // Kota: Bandung
    // Negara: Indonesia

}
