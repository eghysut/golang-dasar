package main

import "fmt"

// For Loop dengan keyword range

// array := [3]string{"Alice", "Bob", "Carl"}
// for idx, arr := range array {
//     fmt.Println(idx, val)
// }

func main() {
    // perulangan dengan data collection(koleksi) 
    // seperti array, slice, map, chanel dll.
    arr := [3]string{"Bear", "Coffe", "Drink"}
    for index, value := range arr {
	fmt.Printf("idx: %d val: %s\n", index, value)
    }
    // Output:
    // idx: 0 val: Bear
    // idx: 1 val: Coffe
    // idx: 2 val: Drink

    // gunakan karakter underscore _ untuk mengabaikan nilai dari sebuah data koleksi
    for _, val := range arr {
	fmt.Println("val:", val)
    }
    // Output:
    // val: Bear
    // val: Coffe
    // val: Drink

    // tipe data map
    data := make(map[string]string)
    data["nama"] = "alice"
    data["usia"] = "23"
    for key, val := range data {
	fmt.Println("key:", key, "val:", val)
    }
    // Output:
    // key: nama val: alice
    // key: usia val: 23

    // tipe data string
    text := "Golang"
    for _, txt := range text {
	if l := "l"; l == string(txt) {
	    break
	}
	fmt.Printf("%s", string(txt))

    }

}
