### Class
Ini adalah entitas pemrograman yang ditentukan pengguna
Class mendefinisikan sekumpulan atribut (setiap atribut memiliki tipe)
Atribut juga disebut properti, bidang, anggota data

### Method
Method mengacu pada fungsi yang terkait dengan tipe data tertentu. 
Method adalah cara untuk menambahkan perilaku khusus kepada 
tipe data (struktur atau tipe data definisi sendiri) yang Anda buat.
Dalam banyak bahasa pemrograman berorientasi objek, 
ini sering disebut sebagai "metode" atau "fungsi anggota" 
yang terkait dengan kelas atau objek.
Nama method harus unik(uniq) didalam kumpulan method

Dalam Go, penggunaan method sedikit berbeda dari bahasa berorientasi objek lainnya.

Berikut adalah beberapa hal penting tentang method di Go:

Method pada Tipe Data Struktur (Struct): 
Anda dapat mendefinisikan method pada tipe data struktur (struct) yang Anda buat.
Ini memungkinkan Anda untuk menambahkan perilaku khusus untuk tipe data Anda.

Receiver: 
Setiap method memiliki receiver, yang adalah parameter pertama dari method dan menghubungkannya dengan tipe data tertentu.
Receiver ini mirip dengan "this" atau "self" di bahasa lain, 
tetapi di Go, Anda dapat memberinya nama apa pun yang Anda inginkan.

Penggunaan Pointer: 
Anda dapat menggunakan pointer sebagai receiver. 
Ini memungkinkan method untuk mengubah data dalam tipe data tersebut secara langsung.
Jika Anda tidak menggunakan pointer, tipe data akan disalin (pass by value) dan perubahan pada tipe data tersebut tidak akan mempengaruhi tipe data asli.

Penamaan Konvensi: 
Pada umumnya, nama receiver diawali dengan huruf kapital jika menggunakan 
pointer dan nama receiver yang pendek.

***
Method dengan pointer receiver dapat mengambil pointer ATAU value sebagai receiver
Method dengan value receiver dapat mengambil pointer ATAU value sebagai receiver
***

### Encapsulation
Encapsulation adalah salah satu konsep utama dalam pemrograman berorientasi objek 
yang menunjukkan cara membatasi akses ke data dan method dalam sebuah objek. 
Dalam encapsulasi, atribut/field/property dan method dalam sebuah kelas dapat dibuat private, protected, atau public.
Private dan protected hanya dapat diakses secara internal oleh objek itu sendiri atau turunannya, 
sedangkan public dapat diakses dari luar objek.
untuk mengakses dan mengubah variabel private diluar kelas kita butuh method yang perilakunya bisa mengakses fields
private dan protected.
Golang menyediakan enkapsulasi pada tingkat package(paket). 
Go tidak memiliki kata kunci public, private, atau protected. 
Satu-satunya mekanisme untuk mengontrol visibilitas adalah dengan menggunakan format huruf besar dan huruf kecil.

Huruf Besar (Exported):
identifier yang diawali dengan huruf besar(misalnya, Nama, Usia, Fungsi()) 
dianggap sebagai identifier yang dapat diekspor(exported) dan 
dapat diakses dari luar paket(package).

Huruf Kecil (Unexported):
identifier yang diawali dengan huruf kecil(misalnya, nama, usia, fungsi()) 
dianggap sebagai identifier yang tidak dapat diekspor(unexported) dan 
hanya dapat diakses dari dalam/internal paket(package) yang sama.

### Polymorphism
Polymorphism berarti "Nama yang sama dengan banyak bentuk". 
Kata-kata ini berasal dari istilah poly(yang berarti beberapa) dan "morphism" yang mendesain bentuk, bentuk sesuatu.
Polymorphism dianggap sebagai salah satu fitur penting dari Object Oriented Programming. 
Polymorphism memungkinkan kita untuk melakukan satu tindakan dengan cara yang berbeda. 
Dengan kata lain, Polymorphism memungkinkan Anda untuk mendefinisikan satu antarmuka dan memiliki beberapa implementasi.
