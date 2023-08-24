package main

import "fmt"

// Fungsi Rekursif (fungsi yang memanggil dirinya sendiri)

//func rekursif() string {
//    fmt.Println("Hello World")

    // memanggil fungsi 
//    return rekursif()
//}

func faktorial(n int) int {
    if n == 0 {
	return 1
    }
    return n * faktorial(n - 1)
}

func main() {

    // rekursif()
    // Output:
    // Hello World
    // Hello World
    // Hello World
    // Infinite Loop

    f := faktorial(5)
    fmt.Println("val:", f)
    // Output:
    // val: 120
}
