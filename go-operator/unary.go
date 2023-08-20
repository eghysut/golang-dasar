package main

import "fmt"

// operator unary
// ++	a = a + 1
// --	a = a - 1
// -	negatif
// +	positif
// !	boolean kebalikan

func main() {
    x := 5
    fmt.Println("Nilai:", x)

    x++
    fmt.Println("++ ", x)

    x--
    fmt.Println("--", x)

    a := -10
    b := 10
    fmt.Println(a)
    fmt.Println(b)

    var status bool
    fmt.Println(status)
    
    fmt.Println("!status:", !status)
    
}
