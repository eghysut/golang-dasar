package main

import "fmt"


// fungsi pass by value
func PassByValue(angka int) {
    angka = angka * 2
    fmt.Println("PassByValue:", angka)
}

// fungsi pass by reference (pointer)
func PassByReference(angka *int) {
    *angka = *angka * 2
    fmt.Println("PassByReference:", *angka)
}

func main() {
    var i int = 100
    fmt.Println("Value i:", i)
    // Output: Value i: 100
    
    PassByValue(i)
    fmt.Println("Value i:", i)
    // Output:
    // PassByValue: 200
    // Value i: 100

    PassByReference(&i)
    fmt.Println("Value i:", i)
    // Output:
    // PassByReference: 200
    // Value i: 200

}
