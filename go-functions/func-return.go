package main

import "fmt"

// Fungsi return value (mengembalikan nilai)
// contoh 1
// func nama_fungsi(Param1 TypeData, Param2 TypeData) Return_Type {
// block kode fungsi
//}

// contoh 2
// func nama_fungsi(Param1, Param2, TypeData) Return_Type {
// block kode fungsi
//}

// menggunakan cara contoh 1
func getAlamat(alamat string) string {
    return alamat
    // kode block di bawah return tidak di eksekusi
    // akhir dari program
}

// menggunakan cara contoh 2
func tambah(a, b int) int {
    return a + b
    // kode block di bawah return tidak di eksekusi
    // akhir dari program
}

func main() {
    x := getAlamat("Jakarta")
    fmt.Println("get:", x)
    // Output:
    //get: Jakarta 

    y := tambah(2, 3)
    fmt.Println("add:", y)
    // Output: 
    // add: 5
}
