tipe variabel, inisialisasi, type inference, dan blank identifier

### Tipe Variabel:
Tipe variabel adalah jenis data yang menentukan jenis nilai yang dapat disimpan dalam variabel tersebut.

Contoh:
var usia int        Variabel dengan tipe int
var nama string     Variabel dengan tipe string
var status bool     Variabel dengan tipe boolean


### Inisialisasi Variabel:
Inisialisasi variabel adalah memberikan nilai awal pada variabel saat deklarasi.

Contoh:
var score int = 95
var greeting string = "Hello, world!"

### Deklarasi Variabel dengan Pendekatan Singkat (Short Variable Declaration):
Ini adalah cara singkat untuk mendeklarasikan dan menginisialisasi variabel di dalam fungsi. 
Ini menggunakan operator := dan hanya dapat digunakan di dalam fungsi.

Contoh:
usia := 25      Tipe variabel usia diinferensi menjadi int
status := true  Tipe variabel status diinferensi menjadi bool
pesan := "Hi"   Tipe variabel pesan diinferensi menjadi string

### Deklarasi Variabel Konstan (Constant Variables): 
Variabel konstan adalah variabel yang nilainya tetap dan tidak dapat diubah 
setelah dideklarasikan. Mereka dideklarasikan dengan kata kunci const.

Contoh:
const pi = 3.14

### Deklarasi Variabel dengan Zero Value: 
Jika variabel dideklarasikan tanpa inisialisasi, maka secara otomatis akan 
memiliki nilai nol sesuai dengan tipe datanya. 
Misalnya, int akan memiliki nilai nol 0, dan string akan memiliki string kosong "".

Contoh:
var angka int    // Variabel angka akan memiliki nilai 0
var nama string // Variabel nama akan memiliki string kosong

### Deklarasi multiple variabel dalam satu pernyataan. 
Jenis-jenis multiple variabel yang umum digunakan termasuk:

Multiple Variabel dengan Tipe Data Yang Sama:
Anda dapat mendeklarasikan beberapa variabel dengan tipe data yang sama dalam satu pernyataan.
Contoh:
var a, b, c int    // Mendeklarasikan tiga variabel bertipe int.
var nama1, nama2 string // Mendeklarasikan dua variabel bertipe string.

Multiple Variabel dengan Inisialisasi Pendek:
Anda dapat menggunakan pendekatan inisialisasi pendek (:=) untuk mendeklarasikan dan menginisialisasi beberapa variabel dengan tipe data yang sama dalam satu pernyataan. Ini umum digunakan dalam fungsi.
Contoh:
x, y, z := 10, 20, 30 // Mendeklarasikan dan menginisialisasi tiga variabel bertipe int.
nama1, nama2 := "Alice", "Bob" // Mendeklarasikan dan menginisialisasi dua variabel bertipe string.

### Penggunaan Blank Identifier:
Blank identifier  (`_`) adalah tanda pengganti yang digunakan untuk mengabaikan nilai yang dikembalikan oleh suatu ekspresi atau fungsi.

Contoh:
 _, err := SomeFunction()   Mengabaikan nilai yang dikembalikan, hanya memeriksa error.
