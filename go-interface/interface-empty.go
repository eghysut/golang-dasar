package main

import "fmt"

// Dalam bahasa pemrograman Go (Golang), "interface kosong" mengacu pada sebuah antarmuka (interface) yang tidak memiliki metode. 
// Dalam Go, sebuah tipe data yang memenuhi (implements) antarmuka kosong ini disebut sebagai "tipe data kosong" atau "interface kosong." Interface kosong sering digunakan untuk mengatasi tipe data yang tidak diketahui atau beragam secara dinamis.

func PrintData(data interface{}) {
    fmt.Printf("%v bertipe %T\n", data, data)
}

func main() {
    // Menggunakan interface kosong dengan berbagai tipe data
    PrintData(14)	// Integer
    PrintData("Golang") // String
    PrintData(3.14)	// Float
    PrintData([]int{1,2,4}) // Slice
}
