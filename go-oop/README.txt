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
