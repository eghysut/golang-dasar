tipe variabel, inisialisasi, type inference, dan blank identifier

1. Tipe Variabel:
   Tipe variabel adalah jenis data yang menentukan jenis nilai yang dapat disimpan dalam variabel tersebut.

Contoh:
var usia int        Variabel dengan tipe int
var nama string     Variabel dengan tipe string
var status bool     Variabel dengan tipe boolean


2. Inisialisasi Variabel:
Inisialisasi variabel adalah memberikan nilai awal pada variabel saat deklarasi.

Contoh:
var score int = 95
var greeting string = "Hello, world!"

3. Type Inference (Inferensi Tipe):
Type inference adalah kemampuan bahasa Go untuk secara otomatis mengidentifikasi tipe variabel berdasarkan nilai yang diberikan pada saat inisialisasi.

Contoh:
usia := 25  // Tipe variabel usia diinferensi menjadi int
status := true  // Tipe variabel status diinferensi menjadi bool
pesan := "Hi" // Tipe variabel pesan diinferensi menjadi string


4. Penggunaan Blank Identifier:
Blank identifier  (`_`) adalah tanda pengganti yang digunakan untuk mengabaikan nilai yang dikembalikan oleh suatu ekspresi atau fungsi.

Contoh:
 _, err := SomeFunction()  // Mengabaikan nilai yang dikembalikan, hanya memeriksa error.
