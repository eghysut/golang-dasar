package main

import "fmt"

func main() {
    i, j := 10, 1024
    p := &i // address memory
    fmt.Println("address i:", &i)
    fmt.Println("address p:", p)
    // akses nilai i lewat pointer
    fmt.Println("get i:", *p)
    // ubah nilai i lewat pointer
    *p = 20
    fmt.Println("set i:", i)
    
    // p menunjuk ke j
    p = &j
    fmt.Println("address j:", &j)
    fmt.Println("address p:", p)
    // akses nilai j lewat pointer
    fmt.Println("get j:", *p)
    // ubah nilai j lewat pointer
    *p = 100
    fmt.Println("set j:", j)

}
