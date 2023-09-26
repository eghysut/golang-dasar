package main

import (
    "fmt"
)

// Fungsi generik Cetak() untuk mencetak data
func Cetak[T any](data T) {
    fmt.Println(data)
}

// Fungsi generik Index() untuk mencari index di array/slice
func Index[T comparable](s []T, target T) int {
    for i, v := range s {
        // jika ada kembalikan nilai index
        if v == target {
            return i
        }
    }
    // jika tidak ada kembalikan nilai -1
    return -1
}

func main() {
    // Mendeklarasikan Tipe Data secara Eksplisit (String)
    Cetak[string]("Hello")
    // Output: Hello 

    // Menggunakan Fungsi Generik Cetak() tanpa 
    // Mendeklarasikan Tipe Data(Type Inference)
    Cetak("World")
    // Output: World

    Cetak(100)
    // Output: 100
 
    // Mendeklarasikan Tipe Data secara Eksplisit (Integer)
    Cetak[int](200)
    // Output: 200

    Cetak(1.25)
    // Output: 1.25

    // Mendeklarasikan Tipe Data secara Eksplisit (Floating)
    Cetak[float64](2.55)
    // Output: 2.55

    arr := []int{20, 50, 30, 10, 40}
    fmt.Println(Index(arr, 10))
    // Output: 3

    fmt.Println(arr[3])
    // Output: 10

    fmt.Println(arr[Index(arr, 10)])
    // Output: 10

}
