package main

import "fmt"

// Fungsi mengembalikan beberapa nilai "README.txt"


// contoh 1:
// fungsi biasa dengan return default
func getVersion() (string, float32) {
    return "Golang", 1.21
}

// contoh 2:
// fumgsi dengan parameter dan beberapa return
func getSplit(x int) (int, int) {
    a := x / 2
    b := x / 3
    return a, b
}

func main() {

    // contoh 1:
    lang, vers := getVersion()
    fmt.Println("language:", lang)
    fmt.Println("version:", vers)
    // Output:
    // language: Golang
    // version: 1.21

    // contoh 2:
    div2, div3:= getSplit(10)
    fmt.Println("div2:", div2)
    fmt.Println("div3:", div3)
    // Output:
    // div2: 5
    // div3: 3

    // mengabaikan nilai return dengan karakter undescore _
    language, _ := getVersion()
    fmt.Println("language:", language)
    // Output:
    // langauage: Golang

    _, version := getVersion()
    fmt.Println("version:", version)
    // Output:
    // version: 1.21



}
