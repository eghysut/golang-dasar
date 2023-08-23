package main

import (
    "fmt"
    "time"
)

// switch tanpa kondisi
// Perintah switch tanpa sebuah kondisi sama seperti switch true.
// Konstruksi ini merupakan cara yang bersih untuk menulis rantaian if-then-else yang panjang.
func main() {
    user := "c"
    switch {
    case user == "y":
	fmt.Println("yes")
    case user == "n":
	fmt.Println("no")
    case user == "c":
	fmt.Prinln("continue")
    default:
	fmt.Println("nil")
    }
}
