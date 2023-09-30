### Go Workspaces(Ruang kerja)
Ruang kerja di Go 1.18 memungkinkan Anda bekerja pada beberapa modul secara bersamaan 
tanpa harus mengedit file go.mod untuk setiap modul. 
Setiap modul di dalam workspace diperlakukan sebagai modul root ketika menyelesaikan dependencies(ketergantungan).

Go Workspaces(Ruang kerja) hanyalah sebuah direktori yang berisi semua kode untuk proyek tertentu. 
Ini termasuk semua kode sumber, serta semua binari atau executable yang telah di compiled(kompilasi)
dari kode sumber tersebut.

Manfaat menggunakan Go workspace ada dua. 
1. memungkinkan Anda untuk mengatur kode Anda dengan cara yang masuk akal untuk proyek Anda. 
2. ini memungkinkan Anda menggunakan golang. 


Cara menggunakan Go Workspaces:

Saya memiliki directories/folder go-example "adalah folder utama"
dan myApp-1 dan myApp-2 adalah sub dari foleder go-example

CATATAN: saat anda ingin membuat go.work, anda harus membuat go module (go.mod) di folder myApp-1 dan myApp-2

/go-example
├── myApp-1
│   ├── go.mod
│   ├── hello-world.go
└── myApp-2
    ├── go.mod
    ├── test.go


perintah go workspace:
(ketik di terminal)

go work init            (secara otomatis go membuat file go.work)
go work use myApp-1     
go work use myApp-2

isi file dari go.work :
go 1.21.0       (versi saat ini saya pakai)

use (
	./myApp-1
	./myApp-2
)