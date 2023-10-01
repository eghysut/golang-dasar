/*package main

import "fmt"


// Mendefinisikan interface sebagai pencapaian abstraksi
type Luas interface{
    luas() float64
}

// Deklarasi struct Persegi
type Persegi struct{
    panjang, lebar float64
}

// Deklarasi struct Lingkaran
type Lingkaran struct{
    radius float64
}

// Rumus luas persegi panjang:
// Luas = panjang (p) x lebar (l)
// Method untuk jenis "Persegi"
func (p Persegi) luas() float64 {
    return p.panjang * p.lebar
}


// Rumus luas lingkaran:
// π = pi = 3.14
// r = jari-jari lingkaran
// Luas lingkaran = π x r² atau
// Luas lingkaran = pi x r x r
// Method untuk jenis "Lingkaran"
func (l Lingkaran) luas() float64 {
    return  3.14 * (l.radius * l.radius)
}

func main() {

    p := Persegi{panjang: 5, lebar: 10}
    l := Lingkaran{radius: 5}

    // Di sini kita memberitahu program kita untuk memanggil method luas().
    // Anda tidak perlu mengetahui bagaimana method tersebut diimplementasikan 
    // (bagaimana kode di dalamnya).

    // Deklarasi variabel x dari tipe Interface
    var x Luas

    x = p
    fmt.Println("luas persegi:", x.luas())
    // Output:
    // luas persegi: 50

    x = l
    fmt.Println("luas lingkaran:", x.luas())
    // Output:
    // luas lingkaran: 78.5

}
