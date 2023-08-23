package main

import "fmt"

// Switch Statement
// Go juga mendukung `switch` statement yang digunakan untuk memeriksa beberapa nilai yang berbeda. 
//Ini adalah contoh penggunaan `switch`:                                                                
//nilai := 2
//switch nilai {
//case 1:
//    jika kondisi variabel nilai = 1
//case 2:
//    jika kondisi variabel nilai = 2
//default:
//    jika semua kondisi di atas tidak terpenuhi

func main() {

    //nilai := 1
    //nilai := 2
    nilai := 0
    switch nilai {
    case 1: // jika variabel nilai = 1
	fmt.Println("ON")
    case 2: // jika variabel nilai = 2
	fmt.Println("OFF")
    default: // jika variabel nilai = 0
	fmt.Println("NIL")
    }
    // Output: NIL


    input := "y"
    //input := "n"
    //input := ""
    switch input {
    case "y": // jika variabel input = y
	fmt.Println("yes")
    case "n": // jika variabel input = n
	fmt.Println("no")
    default: // jika variabel input = ""
	fmt.Println("NIL")
    }
    // Output: yes

    // switch tanpa kondisi
    password := "golang"
    panjang := len(password)
    switch {
    case panjang > 5:
	fmt.Println("ok")

    case panjang > 8:
	fmt.Println("good")

    case panjang > 10:
	fmt.Println("perfect")
    
    default:
	fmt.Println("password is not long enough")
    }
}
