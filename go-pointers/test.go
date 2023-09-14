package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    // i := 100
    var i int = 100
    fmt.Println("initial:", i, "address:", &i)

    zeroval(i)
    fmt.Println("zeroval:", i, "address:", &i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i, "address:", &i)

}
