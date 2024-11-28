package main

import "fmt"

// Definisi struktur untuk menyimpan data mahasiswa
type mahasiswa struct {
	nama, nim, kelas, jurusan string  // Informasi mahasiswa: nama, nim, kelas, jurusan
	ipk                       float64 // Nilai IPK mahasiswa
}

// Definisi array untuk menampung data mahasiswa
type arrMhs [2023]mahasiswa // Array statis dengan maksimum 2023 mahasiswa

// Fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
func IPK_2(T arrMhs, n int) int {
	// Inisialisasi indeks nilai IPK tertinggi dengan elemen pertama
	var idx int = 0 // idx menunjukkan indeks mahasiswa dengan IPK tertinggi sementara
	var j int = 1   // Pencarian dimulai dari mahasiswa kedua (indeks 1)

	// Iterasi untuk membandingkan IPK mahasiswa
	for j < n { // Selama j kurang dari jumlah mahasiswa n
		if T[idx].ipk < T[j].ipk { // Jika IPK pada idx lebih kecil dari IPK pada indeks j
			idx = j // Update indeks mahasiswa dengan IPK tertinggi
		}
		j = j + 1 // Lanjutkan ke mahasiswa berikutnya
	}
	return idx // Kembalikan indeks mahasiswa dengan IPK tertinggi
}
func main() {
	// Contoh data mahasiswa
	var data arrMhs
	data[0] = mahasiswa{"Andi", "12345", "TI-1", "Teknik Informatika", 3.5}
	data[1] = mahasiswa{"Budi", "67890", "TI-2", "Teknik Informatika", 3.9}
	data[2] = mahasiswa{"Citra", "11223", "TI-3", "Teknik Informatika", 3.7}
	data[3] = mahasiswa{"Dina", "44556", "TI-4", "Teknik Informatika", 4.0}
	data[4] = mahasiswa{"Eka", "77889", "TI-5", "Teknik Informatika", 3.8}

	n := 5 // Jumlah mahasiswa dalam array

	// Cari indeks mahasiswa dengan IPK tertinggi
	idxTertinggi := IPK_2(data, n)

	// Cetak hasil
	fmt.Println("Mahasiswa dengan IPK tertinggi:")
	fmt.Printf("Nama: %s, NIM: %s, Kelas: %s, Jurusan: %s, IPK: %.2f\n",
		data[idxTertinggi].nama, data[idxTertinggi].nim, data[idxTertinggi].kelas,
		data[idxTertinggi].jurusan, data[idxTertinggi].ipk)
}
