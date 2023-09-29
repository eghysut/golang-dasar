package mahasiswa

type Mahasiswa struct{
    Nama string // Exported, bisa diakses dari luar paket
    usia int    // Unexported, hanya bisa diakses dari dalam paket / internal (bisa disebut di dalam file ini)
}

// Method dengan receiver value atau disebut (pass by value)
func (m Mahasiswa) GetUsia() int {
    return m.usia
}

// Method dengan receiver pointer atau disebut (pass by reference)
func (m *Mahasiswa) SetUsia(usia_baru int) {
    m.usia = usia_baru
}
