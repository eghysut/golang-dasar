package main


import "fmt"

func getNama(nama string) string {
    return "Hello " + nama

}

func main() {
    funcVal := getNama
    fmt.Println(funcVal("Golang"))
    // Output:
    // Hello Golang
}
