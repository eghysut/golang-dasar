package main

//penggunaan dan sintaksis constants(konstanta)
//dalam bahasa pemrograman Go (Golang). 
//Berikut adalah penjelasan rinci:

//constants (Konstanta) Dasar:

//const dalam Golang bisa berupa karakter, string, boolean, atau nilai numerik.
//const dapat dideklarasikan seperti variabel, tetapi nilai mereka tidak dapat diubah setelah diberikan nilai awal.

import (
    "fmt"
    "math"
)

//first_namr adalah const  string yang diekspor (publik).
const first_name string = "Alice"

//constant dapat dikelompokkan dalam satu blok const declaration.
const (
    //usia adalah constant int yang tidak diekspor(package private).
    usia int = 20
    // Output: 20
    // kondisi adalah constant boolean
    kondisi bool = true
    // Output: true
)


func main() {
    //const dapat digunakan seperti variabel
    const angka = 5
    fmt.Printf("angka %d bertipe %T\n", angka, angka)
    // Output: angka 5 bertipe int

    const x float64 = 2.456
    fmt.Println(x)
    // Output: 2.456


    //const untyped (tanpa tipe) mengambil tipe yang dibutuhkan oleh konteksnya. 
    //Sebagai contoh, math.Sin mengharapkan tipe float64, sehingga konstanta y dengan nilai 10 mengambil tipe tersebut.
    const y = 10
    fmt.Println(math.Sin(y))
    // Output: -0.5440211108893699


    // variabel dengan keyword const tanpa digunakan dia tidak akan error
    const nama = "Alice"
    const usia = 23
    const ipk = 2.8
}
