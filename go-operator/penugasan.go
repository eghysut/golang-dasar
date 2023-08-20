package main

import "fmt"

// operator penugasan/assaigment
// a += 10	a = a + 10
// a -= 10	a = a - 10
// a *= 10	a = a * 10
// a /= 10	a = a / 10
// a %= 10	a = a % 10

func main() {

    x := 10
    fmt.Println("Nilai:", x)

    x += 10
    fmt.Println("Penjumlahan:", x)

    x -= 10
    fmt.Println("Pengurangan:", x)

    x *= 2
    fmt.Println("Perkalian:", x)

    x /= 5
    fmt.Println("Pembagian:", x)

    x %= 2
    fmt.Println("Modulus:", x)



    
}
