package main

import "fmt"

// Short statement adalah fitur dalam bahasa Go (Golang) yang memungkinkan Anda untuk mendeklarasikan dan menginisialisasi variabel hanya untuk lingkup dari sebuah blok `if` atau `switch, dan variabel tersebut hanya dapat digunakan di dalam blok tersebut. 
// Ini adalah fitur yang berguna untuk menginisialisasi variabel sementara saat Anda memerlukan nilai sementara dalam kondisi atau ekspresi.

func main() {

    // membuat short statement di if 
    if x := 5; x > 10 {
	fmt.Println("if")
	fmt.Println("x:", x)
    } else if x < 10 {
	fmt.Println("else if")
	fmt.Println("x:", x)
    } else {
	fmt.Println("else")
	fmt.Println("x:", x)
    }
    // akses variabel x diluar block statement
    // menampilkan pesan error, undefined: x
    //fmt.Println("x:", x)

    switch y := 5; y {
    case 1:
	fmt.Println("satu")
	fmt.Println("y:", y)
    case 5:
	fmt.Println("lima")
	fmt.Println("y:", y)
    default:
	fmt.Println("tidak ada")
    }
}
