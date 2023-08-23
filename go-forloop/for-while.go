package main

import "fmt"

// For Loop Sebagai While Loop
// Perintah for adalah "while"-nya Go
// Dengan cara ini anda bisa menghilangkan perintah awal dan akhir,
// menggunakan hanya ekpresi kondisi sehingga for menjadi seperti while pada bahasa C.

// Contoh:
// inisialisasi := 1
//for inisialisasi <= 5 {
//    fmt.Println(inisialisasi)
//    inisialisasi++
//}

func main() {
    
    init := 1
    for init <= 3 {
	fmt.Println("for:", init)
	init++
    }
    fmt.Println("end:", init)
    // Output:
    //for: 1
    //for: 2
    //for: 3
    //end: 4

    i := init
    for i <= 10 {
	i++
	fmt.Println("for:", i)
    }
    fmt.Println("end:", i)
    // Output:
    // for: 5
    // for: 6
    // for: 7
    // for: 8                                            // for: 9
    // for: 10
    // for: 11
    // end: 11

}
