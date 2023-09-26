package main

import (
    "fmt"
)

func Cetak[T any](data T) {
    fmt.Println(data)
}

func main() {

    Cetak[string]("Hello")
    // Output: Hello 

    Cetak("World")
    // Output: World

}
