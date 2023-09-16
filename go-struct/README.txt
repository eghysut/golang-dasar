### Struct (Struktur):
Struktur adalah tipe data yang memungkinkan Anda untuk menggabungkan 
beberapa tipe data berbeda ke dalam satu entitas yang lebih besar.
Contoh: 
type Mahasiswa struct{ 
    Nama string
    Umur int 
} 
adalah deklarasi struktur yang memiliki dua bidang: Name (tipe string) dan Age (tipe int).


### Anonymous Struct (Struktur Anonymous):
Struktur anonim adalah struktur yang tidak memiliki nama. 
Mereka sering digunakan dalam situasi di mana Anda hanya perlu menyimpan 
data tanpa perlu mendefinisikan tipe struktur secara eksplisit.
Contoh: 
data := struct{ 
    Nama string
    Umur int 
}{"Alice", 30} 
adalah contoh penggunaan struktur anonim untuk membuat data dengan dua bidang(field).

### Embedded Struct (Struktur yang Tersisip):
Struktur yang tersisip adalah konsep yang memungkinkan Anda untuk 
memasukkan satu struktur ke dalam struktur lain.
Ini mirip dengan konsep pewarisan (inheritance) dalam bahasa pemrograman lain.
Contoh:
type Alamat struct {
    Kota  string
    Negara string
}

type Person struct {
    Nama   string
    Umur     int
    Alamat  // Struktur Alamat Embedded(tersisip)
}

### Struct Method:
Method adalah fungsi yang terkait dengan sebuah tipe data tertentu (tipe yang didefinisikan oleh pengguna, termasuk struct).
Method memungkinkan Anda untuk menambahkan perilaku khusus yang terkait dengan tipe data tersebut.
Method didefinisikan dengan menempelkan fungsi ke tipe data menggunakan receiver.
Contoh:
type Mahasiswa struct{
    Nama string
}

func (m Mahasiswa) GetNama() {
    return m.Nama
}
