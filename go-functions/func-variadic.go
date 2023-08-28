package main

import "fmt"

// Fungsi dengan nilai Variadic 


// Contoh 1:
func jumlah(angka ...int) int {
    total := 0

    for _, val := range angka {
	total += val
    }

    return total
}

// Contoh 2:
func getUser(nama string, angka ...int) (string, int) {
    total := 0
    for _, val := range angka {
	total += val
    }
    return nama, total
}

func main() {
    // nilai parameter menggunakan tipe data int
    hasil1 := jumlah(10, 20, 30)
    fmt.Println("val:", hasil1)
    // Output:
    // val: 60

    // nilai parameter menggunakan tipe data slice
    slice := []int{50, 60, 70}
    hasil2 := jumlah(slice...)
    fmt.Println("val:", hasil2)
    // Output:
    // val: 180

    // contoh 2:
    // nilai parameter dengan beberapa tipe data
    user, score := getUser("Alice", slice...)
    fmt.Println("User:", user)
    fmt.Println("Score:", score)
    // Output:
    // User: Alice
    // Score: 180

}
