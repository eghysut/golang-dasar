package main

import (
    "fmt"
    "strings"
)

// parameter sebagai fungsi
func getText(nama string, role func(string) string) {
    rvRole := role(nama)
    fmt.Println("sebelum:", nama)
    fmt.Println("sesudah:", rvRole)
}

// fungsi untuk huruf besar
func setText(nama string) string {
    return strings.ToUpper(nama) 
}

func main() {
    getText("Alice", setText)
    // Output:
    // sebelum: Alice
    // sesudah: ALICE
}
