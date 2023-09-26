package main

import (
	"fmt"
)


// Deklarasi Interface TypeData:
// Mendefinisikan antarmuka (interface) TypeData, 
// yang memiliki constraint (batasan) generik. 
// Interface ini mengizinkan tipe data generik (K dan V) 
// yang bisa berupa int, int32, float32, float64, atau string.
type TypeData interface{
    ~int | ~int32 | ~float32 | ~float64 | ~string
}


// Deklarasi Tipe Data Generik GenericMap:
// Mendefinisikan tipe data struktur generik GenericMap yang memiliki 
// dua parameter tipe generik, K (untuk kunci) dan V (untuk nilai).
// Tipe data ini mewakili map(peta) dengan kunci bertipe K dan nilai bertipe V.
type GenericMap[K TypeData, V TypeData] map[K]V


// Method sum():
// Mendefinisikan method sum() yang digunakan untuk menghitung 
// jumlah nilai dalam peta generik. 
// Method ini mengiterasi melalui semua nilai dalam map(peta) dan 
// mengakumulasikan mereka ke dalam variabel angka.
func (g GenericMap[K, V]) sum() V {
    var angka V
    for _, v := range g {
        angka += v
    }
    return angka
}

// Method values():
// Mendefinisikan method values() yang digunakan 
// untuk mencetak semua nilai dalam peta generik.
// Method ini melakukan iterasi melalui semua nilai dalam map(peta) 
// dan mencetak setiap nilai.
func (g GenericMap[_, V]) values() {
    for _, val := range g {
        fmt.Println(val)
    }
}


// Method keys():
// Mendefinisikan method keys() yang digunakan 
// untuk mencetak semua kunci dalam peta generik.
// Method ini melakukan iterasi melalui semua kunci dalam map(peta) 
// dan mencetak setiap kunci.
func (g GenericMap[K, _]) keys() {
    for key, _ := range g {
        fmt.Println(key)
    }
}

// 
func (g GenericMap[K, V]) get(s K) V {
    var val V

    for i, v := range g {
        if i == s {
            return v
        }
    }
    return val
}

func main() {

    x := GenericMap[string, int]{"foo": 10, "bar": 20}
    fmt.Println(x.sum())
    // Output: 30

    x.values()
    // Output: 
    // 10
    // 20

    x.keys()
    // Output:
    // foo
    // bar

    fmt.Println(x.get("hello"))
    // Output: 0

    fmt.Println(x.get("foo"))
    // Output: 10

}
