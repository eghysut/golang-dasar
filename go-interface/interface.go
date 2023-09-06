package main

import "fmt"


// Definisikan sebuah interface dengan kontrak method GetNama
type Method interface {
    GetNama() string
}

// Definisikan struct Mahasiswa yang memenuhi kontrak 
type Mahasiswa struct {
    Nama string
}

// Implementasikan method GetNama untuk struct Mahasiswa
func (mhs Mahasiswa) GetNama() string {
	return "Hello, " + mhs.Nama
}

// Fungsi yang menerima parameter tipe data yang memenuhi kontrak Method
func DataNama(method Method) {
	fmt.Println(method.GetNama())
}


func main() {
	val := Mahasiswa{Nama: "Golang"}
	fmt.Println(val.GetNama())
	DataNama(val)
}
