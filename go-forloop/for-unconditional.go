package main

import "fmt"

//For Loop Tanpa Kondisi

// i := 1
// for {
// fmt.Println(i)
// i++
// if i > 5 {
//     break
// }
// }

func main() {
    init := 1
    for {
	fmt.Println("for:", init)
	init++
	if init >= 5 {
	    fmt.Println("out:", init)
	    break
	}
    }
    fmt.Println("end:", init)
    // Output:
    // for: 1
    // for: 2                                               for: 3
    // for: 4
    // out: 5
    // end: 5

    // menggunakan perintah break
    text := "golang"
    i := 0
    for {
	txt := string(text[i])
	fmt.Println("txt:", txt)
	i++

	if i >= len(text) {
	    fmt.Println("out:", text)
	    break
	}
    }
    fmt.Println("end:", i)
    // Output:
    // txt: g
    // txt: o
    // txt: l
    // txt: a
    // txt: n
    // txt: g
    // out: golang
    // end: 6

    // menggunakan perintah return
    for {
	fmt.Println("Hello Golang")
	return
    }
    // Output:
    // Hello Golang

}
