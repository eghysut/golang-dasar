package main

import "fmt"

// Error
// nama := "hello world"

func main() {

    // variabel tipe inference hanya dapat digunakan di dalam fungsi
    s := "Hello World"
    fmt.Printf("%s bertipe %T\n", s, s)
    // Output: Hello World bertipe string
}
