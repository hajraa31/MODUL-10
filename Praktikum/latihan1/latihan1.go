package main

import "fmt"

// Definisi tipe data array dengan ukuran tetap
type arrInt [2023]int

//Fungsi untuk mencari indeks nilai terkecil dalam array
func terkecil_2(tabInt arrInt, n int) int {
	// Inisialisasi indeks nilai terkecil dengan indeks pertama
	var idx int = 0 // idx menunjukkan indeks elemen minimum sementara
	var j int = 1   // Mulai pencarian dari elemen kedua (indeks 1)

	//Iterasi untuk menemukan indeks nilai terkecil
	for j < n {
		if tabInt[idx] > tabInt[j] { // Jika elemen di indeks idx lebih besar dari elemen di indeks j
			idx = j // Update indeks nilai terkecil dengan j
		}
		j = j + 1 // Lanjutkan ke elemen berikutnya
	}
	return idx // Kembalikan indeks dari elemen terkecil
}

func main() {
	//Contoh penggunaan
	var data arrInt                                             // Array dengan kapasitas maksimum 202
	data[0], data[1], data[2], data[3], data[4] = 5, 2, 8, 4, 3 // Inisialisasi elemen
	n := 5

	// Cetak array dan indeks nilai terkecilnya
	fmt.Println("Array:", data[:n])                            // Tampilan elemen array yang terpakai
	fmt.Println("Indeks nilai terkecil:", terkecil_2(data, n)) // Tampilan indeks terkecil
}
