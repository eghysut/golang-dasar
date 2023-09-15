package main

import "fmt"

func PassByValue(angka int) {
    angka = angka * 2
    fmt.Println("PassByValue:", angka)
}

func PassByReference(angka *int) {
    *angka = *angka * 2
    fmt.Println("PassByReference:", *angka)
}

func main() {
    var i int = 100
    fmt.Println(i)
    // Output: 100
    
    PassByValue(i)
    fmt.Println("Value i:", i)

    PassByReference(&i)
    fmt.Println("Value i:", i)

}
