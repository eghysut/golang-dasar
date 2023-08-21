Daftar dari format-verb string yang digunakan dalam fungsi-fungsi seperti 
`Sprintf` dan `Printf` dalam bahasa pemrograman Go (Golang) untuk mengatur 
cara nilai-nilai akan diformat saat ditampilkan dalam bentuk string. 

Berikut adalah penjelasan untuk beberapa format-verb yang tercantum dalam tabel tersebut:

Format-verb untuk semua tipe data:
- %v: Menghasilkan nilai dalam format default.
- %#v: Menghasilkan representasi Go-syntax dari nilai.
- %T: Menghasilkan representasi Go-syntax dari tipe nilai.
- %%: Mencetak tanda persen secara literal, tanpa memerlukan nilai.

Format-verb untuk tipe data boolean:
- %t: Menghasilkan kata "true" atau "false" tergantung pada nilai boolean.

Format-verb untuk tipe data integer:
- %b: Menghasilkan nilai dalam basis 2 (biner).
- %c: Menghasilkan karakter yang sesuai dengan kode Unicode yang diberikan.
- %d: Menghasilkan nilai dalam basis 10 (desimal).
- %o: Menghasilkan nilai dalam basis 8 (oktal).
- %q: Menghasilkan karakter tunggal yang diapit oleh tanda kutip satu dan diekspresikan dengan sintaks Go yang aman.
- %x: Menghasilkan nilai dalam basis 16 (heksadesimal) dengan huruf kecil a-f.
- %X: Menghasilkan nilai dalam basis 16 (heksadesimal) dengan huruf besar A-F.
- %U: Menghasilkan representasi Unicode dalam format U+1234 (serupa dengan "U+%04X").

Format-verb untuk tipe data floating-point dan kompleks:
- %b: Notasi ilmiah tanpa angka desimal dengan eksponen berbasis dua.
- %e: Notasi ilmiah, misalnya -1.234456e+78.
- %E: Notasi ilmiah dengan huruf besar, misalnya -1.234456E+78.
- %f: Titik desimal tanpa eksponen, misalnya 123.456.
- %F: Sinonim untuk %f.
- %g: Notasi %e untuk eksponen besar, %f untuk eksponen lainnya.
- %G: Notasi %E untuk eksponen besar, %F untuk eksponen lainnya.

Format-verb untuk tipe data string dan slice of bytes:
- %s: Byte dari string atau slice tanpa interpretasi.
- %q: String dalam tanda kutip ganda yang aman dengan sintaks Go.
- %x: Basis 16 dengan huruf kecil, dua karakter per byte.
- %X: Basis 16 dengan huruf besar, dua karakter per byte.

Format-verb untuk tipe data pointer:
- %p: Notasi basis 16 dengan awalan "0x".

Tabel ini memberikan petunjuk tentang berbagai cara memformat nilai-nilai dalam bahasa Go ketika ingin mencetak atau menghasilkan string dengan pola yang diinginkan
