Fungsi (functions) adalah blok kode yang berisi serangkaian pernyataan yang dapat dipanggil (invoked) dengan memberikan argumen (input), menjalankan operasi tertentu, dan mengembalikan hasil (output). Dalam bahasa pemrograman Go (Golang), fungsi sangat penting karena digunakan untuk mengorganisasi dan mengelompokkan kode menjadi unit-unit yang lebih kecil dan dapat digunakan kembali.

Berikut adalah beberapa konsep dasar tentang fungsi di Go:

### Deklarasi Fungsi
Anda dapat mendeklarasikan fungsi dengan menggunakan kata kunci `func` diikuti oleh nama fungsi, parameter (input), tipe data kembalian (jika ada), dan blok kode fungsi.

Contoh:
func namaFungsi(parameter1 tipeData1, parameter2 tipeData2) tipeDataKembalian {
    // Blok kode fungsi
    return nilaiKembalian
}

Contoh:
func tambah(a, b int) int {
    hasil := a + b
    return hasil
}

### Pemanggilan Fungsi
Fungsi dapat dipanggil dengan memberikan argumen yang sesuai dengan tipe data parameter dan menangkap nilai kembalian jika ada.

Contoh:
hasil := tambah(5, 3)
fmt.Println(hasil) 
// Output: 8


### Fungsi return
Fungsi dapat mengembalikan satu atau lebih nilai kembalian. Anda dapat mengembalikan nilai menggunakan pernyataan `return`. Tipe data kembalian harus sesuai dengan yang dideklarasikan dalam deklarasi fungsi.

Contoh:
func kuadrat(x int) int {
    return x * x
}

### Fungsi returning multiple values
Mengembalikan beberapa nilai dari fungsi Go memungkinkan fungsi untuk mengembalikan lebih dari satu nilai. Ini sangat berguna ketika Anda perlu mengembalikan sejumlah data terkait dalam satu panggilan fungsi.

Contoh 1:
func getData() (string, int) {
    return "Golang", 14
}


Contoh 2:
func bagi(x int) (int, int) {
    hasil1 := x / 2
    hasil2 := x / 3
    return hasil1, hasil2
}

func main() {
    contoh 1:
    x, y := getData()

    contoh 2:
    hasil1, hasil2 := bagi(10)
    fmt.Println(hasil1, hasil2)


### Parameter Variadic
Anda dapat menggunakan parameter variadic dengan menggunakan tanda elipsis `...` di depan tipe data parameter. Ini memungkinkan fungsi menerima jumlah argumen yang bervariasi.

Contoh:
func jumlah(angka ...int) int {
    total := 0
    for _, nilai := range angka {
        total += nilai
    }
    return total
}

### Fungsi Anonim (Anonymous Function)
Anda dapat membuat fungsi tanpa nama (anonim) yang dapat digunakan secara langsung atau disimpan dalam variabel.

Contoh:
fungsiAnonim := func(a, b int) int {
    return a * b
}
// Memanggil fungsi anonim
hasil := fungsiAnonim(5, 3) 


### Fungsi Sebagai Tipe Data
Fungsi juga dapat digunakan sebagai tipe data. Anda dapat menyimpan fungsi dalam variabel atau menggunakannya sebagai parameter dalam fungsi lain.

Contoh:
type FungsiPengali func(int, int) int

func kali(a, b int) int {
    return a * b
}

var fungsiPengali FungsiPengali
fungsiPengali = kali
// Memanggil fungsi yang disimpan dalam variabel
hasil := fungsiPengali(5, 3) 

### Fungsi Method (Metode):
Fungsi metode adalah fungsi yang terkait dengan tipe data tertentu. Mereka digunakan dalam pemrograman berorientasi objek dan dipanggil pada instance tipe data tersebut.

Contoh:
type PersegiPanjang struct {
    Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
    return p.Panjang * p.Lebar
}

### Fungsi sebagai Nilai (Function Values)
Dalam Go, Anda dapat menyimpan fungsi dalam variabel seperti tipe data lainnya. Ini memungkinkan Anda untuk mengirim fungsi sebagai argumen atau mengembalikan fungsi dari fungsi lain. 

Contoh:
package main

import "fmt"

// Fungsi yang menerima dua int dan fungsi lain sebagai argumen
func operasi(a, b int, fungsiOperasi func(int, int) int) int {
    return fungsiOperasi(a, b)
}

func main() {
    // Mendefinisikan fungsi tambah
    tambah := func(x, y int) int {
        return x + y
    }

    // Menggunakan fungsi sebagai nilai
    hasilTambah := operasi(5, 3, tambah)
    fmt.Println("Hasil Penambahan:", hasilTambah)
}


### Fungsi Closure (clouser) adalah konsep yang penting dalam bahasa pemrograman Go (Golang).             Ini mengacu pada kemampuan fungsi untuk mengakses variabel dari lingkup luar di mana fungsi tersebut didefinisikan,                                            bahkan setelah lingkup luar sudah selesai dieksekusi. Dalam Go, closure sering digunakan dalam situasi di mana Anda ingin menyimpan dan mengakses state di dala
m fungsi.
                                                     Contoh:
package main

import "fmt"

func main() {
    x := 10

    // Closure: fungsi yang mengakses variabel x
    fungsiClosure := func() {
        fmt.Println(x)
    }

    x = 20 // Mengubah nilai x

    // Memanggil closure
    fungsiClosure() // Output: 20
}                                                                                                         Manfaat utama dari closure adalah kemampuannya untuk membuat fungsi yang memiliki akses ke state internal dan dapat mempertahankan state tersebut di antara panggilan-panggilan fungsi. Ini sangat berguna dalam berbagai situasi, termasuk dalam pembuatan fungsi-fungsi yang digunakan sebagai goroutine (concurrency), pembuatan penutupan (closures) dalam pengolahan data, dan banyak lagi.

Closure adalah salah satu fitur yang kuat di Go yang memungkinkan Anda untuk menulis kode yang lebih fleksibel dan dapat digunakan kembali dengan cara yang jelas dan aman.
