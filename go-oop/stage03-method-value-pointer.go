/*package main

import "fmt"


// Membuat struct Mahasiswa yang memiliki fields Nilai
type Mahasiswa struct{
    Nilai int
}

// Method dengan Pointer receiver
func (m *Mahasiswa) IncrementPointer() {
    m.Nilai++
}

// Method dengan Values receiver
func (m Mahasiswa) IncrementValue() {
    m.Nilai++

}

func main() {
    obj := Mahasiswa{Nilai: 20}

    // Memanggil method dengan value receiver
    // Nilai tidak berubah
    obj.IncrementValue()
    fmt.Println("Method dengan Value receiver:", obj.Nilai)
    // Output: 
    // Method dengan Value receiver: 20

    // Memanggil method dengan pointer receiver
    // Nilai berubah
    obj.IncrementPointer()
    fmt.Println("Method dengan Pointer receiver:", obj.Nilai)
    // Output:
    // Method dengan Pointer receiver: 21

}*/
