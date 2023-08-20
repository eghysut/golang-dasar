package main

import "fmt"

// operator perbandingan
// >	    lebih besar
// <	    kurang dari
// >=	    lebih besar sama dengan
// <=	    kurang dari sama dengan
// !=	    tidak sama dengan
func main() {
    value1 := 100
    value2 := 200

    fmt.Println(value1, value2)

    fmt.Printf("%d > %d = %t\n", value1, value2, value1 > value2)
    fmt.Printf("%d < %d = %t\n", value1, value2, value1 < value2)
    fmt.Printf("%d >= %d = %t\n", value1, value2, value1 >= value2)
    fmt.Printf("%d <= %d = %t\n", value1, value2, value1 <= value2)
    fmt.Printf("%d != %d = %t\n", value1, value2, value1 != value2)
}
