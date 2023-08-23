Dasar-dasar perulangan (`for loop`) dalam bahasa pemrograman Go (Golang). 
Perulangan digunakan untuk menjalankan serangkaian perintah berulang kali hingga kondisi tertentu terpenuhi.
Go hanya memiliki satu perintah pengulangan yaitu for.

### For Loop Sederhana
Struktur dasar dari `for loop` dalam Go adalah sebagai berikut:

for inisialisasi; kondisi; perubahan {
    // Perintah yang akan dijalankan selama kondisi benar (true)
}

Contoh:
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}

### For Loop Sebagai While Loop
Anda dapat menggunakan `for` tanpa inisialisasi dan perubahan untuk membuatnya berfungsi seperti `while loop`:

Contoh:
inisialisasi := 1
for inisialisasi <= 5 {
    fmt.Println(inisialisasi)
    inisialisasi++
}

### For Loop Tanpa kondisi(awal tengah dan akhir)
Anda juga dapat menggunakan `for` tanpa kondisi, yang akan menjalankan perulangan hingga dihentikan secara manual:

Contoh:
i := 1
for {
    fmt.Println(i)
    i++
    if i > 5 {
        break
    }
}

Pengulangan selamanya
Jika anda mengosongkan kondisi maka pengulangan akan berjalan selamanya, dengan ini pengulangan tanpa henti dapat diekspresikan dengan singkat.

Untuk keluar dari pengulangan anda bisa menggunakan perintah "break" atau "return" bergantung kepada kondisi yang dibutuhkan pada program.
for {
    fmt.Println("Hello world")
    // untuk menghentikan perulangan gunakan
    // return atau break
}

### Penggunaan `range` dalam For Loop
Untuk mengulang melalui koleksi seperti slice, map, atau channel, Anda dapat menggunakan perulangan `for` dengan `range`:

Contoh:
slice := []string{"apel", "jeruk", "pisang"}
for indeks, nilai := range slice {
    fmt.Printf("Indeks: %d, Nilai: %s\n", indeks, nilai)
}

### Perulangan Bersarang
Anda dapat menyusun perulangan bersarang untuk situasi yang lebih kompleks:

Contoh:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}

Dalam contoh di atas, perulangan luar akan berjalan tiga kali, dan setiap kali itu perulangan dalam akan berjalan tiga kali pula.

Perulangan (`for loop`) adalah struktur dasar yang sangat penting dalam pemrograman. Anda dapat menggunakannya untuk menjalankan serangkaian perintah berulang kali sesuai dengan kebutuhan aplikasi Anda.
