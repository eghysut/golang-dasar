package main

import "fmt"

// interface juga dapat digunakan sebagai fields

type MyInterface interface{
    GetNama() string
}

type Mahasiswa struct{
    Nama string
    Alamat MyInterface
}

func (m Mahasiswa) GetNama() string {
    return m.Nama
}


func main() {
}
