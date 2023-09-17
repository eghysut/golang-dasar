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

### Deklarasi Method Struct:

Method struct dideklarasikan dengan menambahkan receiver (penerima) pada fungsi.
Receiver adalah tipe data yang digunakan untuk menghubungkan fungsi dengan tipe struct yang bersangkutan.
Method ini berada dalam lingkup (scope) yang sama dengan tipe struct yang didefinisikan.
Artinya, Anda dapat mengakses bidang-bidang(fields) struct tersebut di dalam method.

1. Receiver(Penerima) Method:

Penerima method bisa berupa tipe data apa saja, termasuk tipe data pointer (*T) atau non-pointer (T).
Jika Anda menggunakan tipe data non-pointer sebagai penerima, maka method akan bekerja pada salinan nilai struct, dan perubahan yang dilakukan di dalam method tidak akan mempengaruhi struct asli.
Jika Anda menggunakan tipe data pointer sebagai penerima, maka method akan bekerja pada instansi struct yang sebenarnya, dan perubahan yang dilakukan di dalam method akan mempengaruhi struct asli.

Contoh:
type Mahasiswa struct{
    Nama string
}

// m: adalah variabel receiver(penerima) dari struct Mahasiswa
func (m Mahasiswa) CetakNama() {
    fmt.Println(m.Nama)
}

func (m Mahasiswa) SetNama() string {
    return m.Mahasiswa
}

2. Pemanggilan Method:

Method dapat dipanggil pada variabel yang memiliki tipe struct tersebut.
Pemanggilan method dilakukan menggunakan operator dot (.) setelah variabel.

Contoh:
mhs := Mahasiswa{Nama: "Alice"}
// memanggil method CetakNama() pada struct Mahasiswa
mhs.CetakNama()

3. Method Pointer vs. Method Non-Pointer:

Pemilihan antara method pointer atau method non-pointer tergantung pada apakah Anda 
ingin mengubah nilai asli struct atau hanya bekerja dengan salinan nilai.
Jika Anda ingin mengubah nilai asli, gunakan method pointer.
Jika Anda hanya ingin membaca nilai struct, gunakan method non-pointer.
