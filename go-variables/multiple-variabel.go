package main

import "fmt"

// Deklarasi variabel dapat "difaktorkan" ke dalam blok
var (
    nama string = "Golang"
    umur int = 15
    
)

func main() {
    // Mendeklarasikan 2 variabel bertipe int.
    var a, b int
    a, b = 10, 20
    fmt.Println(a, b)
    // Output: 10 20

    // Mendeklarasikan 3 variabel bertipe string
    var x, y, z string = "Alice", "Carl", "Eliot"
    fmt.Println(x, y, z)
    // Output: Alice Carl Eliot

    fmt.Println(nama, umur)
    // Output: Golang 15

    var (
        p int = 100
        q int = 200
    )
    fmt.Println(p, q)
    // Output: 100, 200

}
