package main

import "fmt"

// Membuat fungsi Foo() untuk menampilkan nilai
func Foo(e string) {
    fmt.Println(e)
}

// Membuat struct Mahasiswa
type Mahasiswa struct{
    Nama string
}

// Membuat fungsi Bar() untuk struct Mahasiswa
func Bar(mhs Mahasiswa) {
    fmt.Println(mhs.Nama)
}

func main() {

    // Contoh dengan fungsi Foo()
    x := "Hello"
    // Memanggil fungsi Foo()
    Foo(x)
    // Output: Hello

    // Pointer ke variabel x
    ptr := &x
    // Mengubah nilai dengan pointer
    *ptr = "World"
    Foo(*ptr)  // Output: World
    Foo(x)     // Output: World

    // Contoh dengan fungsi Bar()
    y := Mahasiswa{Nama: "Alice"}
    // Memanggil fungsi Bar()
    Bar(y)
    // Output: Alice
    
    // Pointer ke variabel y
    z := &y
    // Mengubah nilai dengan pointer ke struct Mahasiswa
    z.Nama = "Carl"
    Bar(*z)
    Bar(y)
    // Output:
    // Carl
    // Carl
}
