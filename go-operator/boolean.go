package main

import "fmt"

// operator boolean
// &&	    dan
// ||	    atau
// !	    kebalikan

func main() {
    value1 := 100
    value2 := 200

    fmt.Println(value1 > 80 && value2 < 250)
    fmt.Println(value1 < 80 || value2 < 250)

    fmt.Println(!(value1 > 80 && value2 < 250))
    fmt.Println(!(value1 < 80 || value2 < 250))

    hasil := value1 > value2
    fmt.Println(!hasil)
}


