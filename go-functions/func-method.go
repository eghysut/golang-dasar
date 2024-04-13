package main

import "fmt"

// Fungsi Method(metode)

type PersegiPanjang struct {
    Panjang int
    Lebar int
}

func (p PersegiPanjang) HitungLuas() int {
    return p.Panjang * p.Lebar
}

func main() {

    val := PersegiPanjang{Panjang: 2, Lebar: 3}
    fmt.Println("Persegi panjang:", val)
    // Output:
    // Persegi panjang: {2 3}

    luas := val.HitungLuas()
    fmt.Println("luas:", luas)
    // Output:
    // luas: 6

    

}
