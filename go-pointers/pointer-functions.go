package main

import "fmt"

// Go (Golang) menggunakan konsep "pass by value" untuk sebagian besar jenis data, 
// tetapi ada beberapa pengecualian yang menggunakan "pass by reference."

// Pass by Value:
// Dalam "pass by value," saat Anda mengirimkan suatu nilai ke dalam sebuah fungsi,
// yang sebenarnya dikirimkan adalah salinan (copy) dari nilai tersebut.
// Ini berarti jika Anda mengubah nilai tersebut di dalam fungsi, itu tidak akan mempengaruhi nilai aslinya di luar fungsi. Tipe data seperti angka, string, dan array dalam Go dioperasikan secara "pass by value."

// Contoh "pass by value" dalam Go:
func zeroval(ival int) {
    ival = 0
}


// Pass by Reference (Pointer):
// Namun, dalam beberapa kasus, Anda ingin mengubah nilai variabel asli yang dilewatkan ke dalam fungsi.
// Untuk ini, Anda dapat menggunakan pointer. 
// Pointer adalah tipe data yang menyimpan alamat memori variabel lain,
// dan dengan menggunakan pointer, Anda dapat mengubah nilai variabel yang ditunjuk oleh pointer tersebut.
// Ini disebut "pass by reference."

// Contoh "pass by reference" dengan menggunakan pointer dalam Go:
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    // i := 100
    var i int = 100
    fmt.Println("initial:", i, "address:", &i)
    // Output:
    // initial: 100 address: 0xc000092068
    
    zeroval(i)
    fmt.Println("zeroval:", i, "address:", &i)
    // Output:
    // zeroval: 100 address: 0xc000092068

    zeroptr(&i)
    fmt.Println("zeroptr:", i, "address:", &i)
    // Output:
    // zeroptr: 0 address: 0xc000092068

}
