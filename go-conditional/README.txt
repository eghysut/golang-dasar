Dasar-dasar kondisional dalam bahasa pemrograman Go (Golang). 
Kondisional memungkinkan Anda untuk mengontrol alur program berdasarkan kondisi tertentu. Di Go, kita memiliki struktur kontrol `if`, `else if`, dan `else` untuk melakukan hal ini.

### Kondisi If
Struktur dasar dari `if` di Go adalah sebagai berikut:

if kondisi {
    // Perintah yang akan dijalankan jika kondisi benar (true)
}

Contoh:
umur := 18
if umur >= 18 {
    fmt.Println("Anda adalah seorang dewasa.")
}


### Kondisi Else
Anda dapat menggunakan blok `else` untuk mengeksekusi perintah ketika kondisi `if` tidak benar (false). Contoh:

umur := 15
if umur >= 18 {
    fmt.Println("Anda adalah seorang dewasa.")
} else {
    fmt.Println("Anda adalah seorang anak-anak atau remaja.")
}

### Kondisi Else If
`else if` digunakan ketika Anda memiliki beberapa kondisi yang ingin diuji secara berurutan. Contoh:

umur := 25
if umur >= 18 {
    fmt.Println("Anda adalah seorang dewasa.")
} else if umur >= 13 {
    fmt.Println("Anda adalah seorang remaja.")
} else {
    fmt.Println("Anda adalah seorang anak-anak.")
}

### Switch Statement
Go juga mendukung `switch` statement yang digunakan untuk memeriksa beberapa nilai yang berbeda. Ini adalah contoh penggunaan `switch`:

nilai := 3
switch nilai {
case 1:
    fmt.Println("Satu")
case 2:
    fmt.Println("Dua")
case 3:
    fmt.Println("Tiga")
default:
    fmt.Println("Nilai tidak ditemukan")
}

### Switch tanpa kondisi/ekspresi
membuat kode yang lebih bersih dan mudah dibaca saat Anda perlu mengevaluasi beberapa kondisi dengan kasus-kasus yang berbeda.

nilai := 3                                           switch {
case nilai > 5: 
    fmt.Println("nilai kurang dari 5")
case nilai > 1:
    fmt.Println("nilai lebih besar dari 1")           default:                                                 fmt.Println("Nil")            
}

Penting untuk diingat bahwa dalam Go, `if`, `else if`, `else`, dan `switch` hanya mengevaluasi satu kondisi atau nilai pada saat yang sama. Jadi, ketika salah satu kondisi terpenuhi, tidak ada pengevaluasian lanjutan. Selain itu, Anda dapat menggunakan pernyataan `if` bersarang atau `switch` bersarang untuk kasus yang lebih kompleks.

Short statement adalah fitur dalam bahasa Go (Golang) yang memungkinkan Anda untuk mendeklarasikan dan menginisialisasi variabel hanya untuk lingkup dari sebuah blok `if` atau `switch`, dan variabel tersebut hanya dapat digunakan di dalam blok tersebut. Ini adalah fitur yang berguna untuk menginisialisasi variabel sementara saat Anda memerlukan nilai sementara dalam kondisi atau ekspresi. 

Fitur ini umumnya digunakan bersama dengan pernyataan kondisional untuk membuat kode lebih ringkas dan lebih mudah dibaca.

Berikut adalah contoh penggunaan short statement dalam blok `if`:

if x := 10; x > 5 {
    fmt.Println("x lebih besar dari 5")
}
// x tidak dapat digunakan di sini karena hanya terbatas pada blok if

Dalam contoh di atas, kita mendeklarasikan dan menginisialisasi variabel `x` dengan nilai 10 hanya untuk blok `if`. Variabel `x` hanya dapat digunakan di dalam blok `if`, dan tidak dapat diakses di luarnya.

Anda juga dapat menggunakan short statement dalam blok `switch`:

switch i := 2; i {
case 1:
    fmt.Println("Satu")
case 2:
    fmt.Println("Dua")
}
// i tidak dapat digunakan di sini karena hanya terbatas pada blok switch


Dalam kasus ini, variabel `i` hanya dapat digunakan di dalam blok `switch`.

Short statement berguna ketika Anda perlu membuat variabel sementara untuk digunakan dalam sebuah kondisi atau ekspresi, dan Anda tidak ingin variabel tersebut memengaruhi lingkup di luar kondisi tersebut. Ini membantu dalam menghindari konflik nama variabel dan menjaga kode lebih bersih dan terstruktur.
