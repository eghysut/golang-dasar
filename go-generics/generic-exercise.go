/*package main

import (
	"fmt"
)

type TypeData interface {
    ~int | ~float64 | ~rune
}

// fungsi mencari nilai tertinggi
func max[TD TypeData](s []TD) TD {
    var angkaMax TD
    for _, v := range s {
        if v > angkaMax {
            angkaMax = v
        }
    }
    return angkaMax
}

// fungsi mencari nilai terendah
func min[TD TypeData](s []TD) TD {
    var angkaMin = s[0]
    for _, v := range s {
        if v < angkaMin {
            angkaMin = v
        }
    }
    return angkaMin
}

func main() {

    arr := []int{10, 20, 100, 30}
    //fmt.Println(index(arr, 30))
    // Output: 3

    fmt.Println(max(arr))
    // Output: 100

    fmt.Println(min(arr))
    // Output: 10
    
}*/
