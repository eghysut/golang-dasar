package main

import "fmt"

// Interface dipenuhi secara implisit
// Sebuah tipe mengimplementasikan sebuah interface dengan mengimplementasikan method-methodnya. 
// Tidak ada deklarasi eksplisit; tidak ada perintah "implements".

// Interface implisit memisahkan definisi dari sebuah interface dari implementasinya, yang bisa saja muncul dalam paket manapun tanpa adanya penataan sebelumnya.

type Kontrak interface{
    Foo()
}

type Mahasiswa struct{
    Nama string
}


// Method berikut berarti type Mahasiswa  mengimplementasikan interface Kontrak,
// tapi kita tidak perlu secara eksplisit mendeklarasikannya.
func (mhs Mahasiswa) Foo() {
    fmt.Println(mhs.Nama)
}

func main() {
    // deklarasi eksplisit
    var x Kontrak = Mahasiswa{Nama: "Alice"}
    x.Foo()
    // Output: Alice

    // deklarasi implisit 
    y := Mahasiswa{Nama: "Carl"}
    y.Foo()
    // Output: Carl

}
