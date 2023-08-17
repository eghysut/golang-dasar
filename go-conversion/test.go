package main

import (
    "fmt"
)

func main() {

    // konversi int8 ke int32
    var nilai8 int8 = 127 
    var nilai32 int32 = int32(nilai8)
    fmt.Printf("%d bertipe %T\n", nilai32, nilai32)

    // konversi nilai int 144 ke alphabet
    var nilai int64 = 114
    var nama string = string(nilai)
    fmt.Printf("nilai: %d alphabet: %s\n", nilai, nama)
    
    // konversi byte ke alphabet
    fname := "Hello World"
    indeks := string(fname[6])
    fmt.Println(fname)
    fmt.Println("Indeks ke 6 adalah:", indeks)

}
