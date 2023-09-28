Menggunakan Go Modules untuk mengelola dependensi dan package dalam 
proyek Go adalah praktik yang sangat disarankan.
Dengan Go Modules, Anda dapat mengontrol versi paket, 
memastikan kompatibilitas, dan menghindari masalah terkait dependensi.
Berikut adalah langkah-langkah untuk menggunakan Go Modules:

### Langkah 1: Inisialisasi Go Module

1. Buat direktori/folder baru untuk proyek Anda (jika belum ada).
2. Di dalam direktori/folder proyek, jalankan perintah berikut untuk menginisialisasi Go Module:

go mod init nama-modul-anda

Contoh:
go mod init Myapp

Ini akan membuat sebuah file "go.mod" yang mencatat dependensi proyek Anda.

### Langkah 2: Membuat Kode Anda
Mulailah dengan menulis kode Go Anda dalam proyek yang baru dibuat.

### Langkah 3: Import Paket-Paket yang Dibutuhkan
Import paket-paket eksternal atau dari standard library sesuai kebutuhan proyek Anda. 
Anda dapat menggunakan perintah "import" seperti yang biasa, 
dan Go Modules akan meresolusi dan mengelola dependensi secara otomatis.

### Langkah 4: Mereferensikan Paket-Paket Dalam Kode Anda
Gunakan paket-paket yang diimpor dalam kode Anda sesuai kebutuhan.

### Langkah 5: Menambahkan Dependensi Baru
Jika Anda perlu menambahkan dependensi baru, gunakan perintah "go get" 
untuk mengunduh dan menambahkannya ke modul Anda. 

Misalnya:

go get nama-paket

Contoh:
go get github.com/gin-gonic/gin


### Langkah 6: Merekam Dependensi
Setelah menambahkan dependensi baru, Anda perlu merekamnya ke dalam file "go.mod". 
Jalankan perintah berikut untuk melakukannya:

go mod tidy

Ini akan memperbarui file `go.mod` dengan dependensi yang baru ditambahkan.

### Langkah 7: Build dan Jalankan Proyek Anda
Sekarang Anda dapat membangun dan menjalankan proyek Anda seperti biasa:

go build ./nama-modul-anda


### Langkah 8: Versioning dan Penyimpanan Modul
Go Modules secara otomatis mengelola versi paket yang digunakan oleh proyek Anda.
Anda dapat melihat versi yang digunakan dalam file "go.mod". 
Selain itu, Go Modules mendukung penyimpanan modul yang aman dan terpercaya,
yang memastikan bahwa dependensi Anda dapat diunduh dari sumber yang aman.

### Catatan Tambahan:
- Pastikan untuk mengkomitkan file "go.mod" dan "go.sum" ke repositori version 
  control Anda agar anggota tim lain dapat menggunakan Go Modules dengan mudah.

- Untuk pengembangan bersama yang lebih baik, bagikan file "go.mod" dan "go.sum" 
  kepada anggota tim Anda agar semua dapat menggunakannya.

- Jika Anda ingin memperbarui semua dependensi ke versi terbaru yang kompatibel,
  Anda dapat menggunakan perintah `go get -u=patch`:

go get -u=patch

- Go Modules sangat mendukung manajemen versi, 
  sehingga Anda dapat dengan aman menambahkan dependensi baru dan memperbarui yang 
  sudah ada tanpa khawatir mengenai konflik dependensi.
