package main

import "fmt"

// interface juga dapat digunakan sebagai fields

// Deklarasi interface dengan kontrak method KalkulasiHarga() dan Deskripsi()
type PerilakuProduk interface{
    KalkulasiHarga() float64
    Deskripsi() string
}

// Deklarasi struct Produk 
type Produk struct{
    Nama string
    Harga float64
    Perilaku PerilakuProduk // Fields bertipe interface
}

// Deklarasi struct ProdukFisik
type ProdukFisik struct{
    Berat float64
}

// Method KalkulasiHarga() untuk struct ProdukFisik
func (pf ProdukFisik) KalkulasiHarga() float64 {
    return pf.Berat * 2.5 // Harga berdasarkan berat
}

// Method Deskripsi() untuk struct ProdukFisik
func (pf ProdukFisik) Deskripsi() string {
    return fmt.Sprintf("Produk Fisik (berat: %.2f kg)", pf.Berat)
}

func main() {
    // Membuat instance struct
    var buku Produk = Produk{
        Nama: "Buku", 
        Harga: 4.5, 
        Perilaku: ProdukFisik{Berat: 0.5},
    }

    fmt.Printf("Harga: %.2f %s", buku.Perilaku.KalkulasiHarga(), buku.Perilaku.Deskripsi())
    // Output:
    // Harga: 1.25 Produk Fisik (berat: 0.50 kg)
}
