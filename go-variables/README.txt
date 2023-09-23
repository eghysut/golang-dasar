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

### Deklarasi Variabel Type Inference (Inferensi Tipe):
Type inference adalah kemampuan bahasa Go untuk secara otomatis mengidentifikasi
tipe variabel berdasarkan nilai yang diberikan pada saat inisialisasi.
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

### Penggunaan Blank Identifier:
Blank identifier  (`_`) adalah tanda pengganti yang digunakan untuk mengabaikan nilai yang dikembalikan oleh suatu ekspresi atau fungsi.

Contoh:
 _, err := SomeFunction()   Mengabaikan nilai yang dikembalikan, hanya memeriksa error.
