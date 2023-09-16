package main

import "fmt"

func main() {
    // Mendefinisikan struct Anonymous
    Mahasiswa := struct{
        Nama string
        Umur int
    }{"Alice", 20}
    fmt.Println(Mahasiswa)
    // Output: {Alice 20}

    x := struct{
        Nama string
        Umur int
    }{
        Nama: "Carl",
        Umur: 22,
    }
    fmt.Println("Nama:", x.Nama)
    fmt.Println("Umur:", x.Umur)
    // Output: 
    // Nama: Carl
    // Umur: 22

    var y struct{
        Nama string
        Umur int
    }

    y.Nama = "Eliot"
    y.Umur = 23
    fmt.Println("Nama:", y.Nama)
    fmt.Println("Umur:", y.Umur)
    // Output:
    // Nama: Eliot
    // Umur: 23

}
