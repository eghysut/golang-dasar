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
