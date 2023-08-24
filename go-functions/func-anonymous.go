package main

import "fmt"

// Fungsi Anonymous
func main() {

    fungsiAnonim := func(a, b int) int {
        return a * b
    }

    hasil := fungsiAnonim(2, 3)
    fmt.Printf("val: %d bertipe %T\n", hasil, hasil)
    // fmt.Println(hasil)

}
