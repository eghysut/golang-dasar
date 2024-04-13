package main

import "fmt"

// For Loop Sederhana

//for inisialisasi; kondisi; perubahan {
    // Perintah yang akan dijalankan selama kondisi benar (true)
//}

func main() {
    
    for init := 0; init < 5; init++ {
	fmt.Println("for:", init)
    }

    // Output:
    // for: 0
    // for: 1
    // for: 2
    // for: 3
    // for: 4

    for i := 1; i <= 3; i++ {
	fmt.Printf("Hello World %d kali\n", i)
    }

    // perintah awal dan akhir adalah opsional
    total := 1
    for ; total < 5; {
        total += total
        fmt.Println("for:", total)                           // Output:
        // for: 2
        // for: 4                                            // for: 8

    }
    fmt.Println("total:", total)
    // Output:
    // total: 8
    
    // tanpa perintah tengah
    for k := 0; ; k++ {
        fmt.Println(k)
        if k == 10 {
            break
        }
    }

    for k := 0; {
        fmt.Println(k)
    }

}
