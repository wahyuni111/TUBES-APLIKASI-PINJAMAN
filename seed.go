package main

func seed() {
	// Isi data pengguna
	dbPengguna[0] = Pengguna{1, "Agha", "agha@example.com", "password123"}
	dbPengguna[1] = Pengguna{2, "Elfan", "elfangamtenk@example.com", "elfan123"}
	dbPengguna[2] = Pengguna{3, "Citra", "citra@example.com", "citra456"}
	dbPengguna[3] = Pengguna{4, "Farel", "farel67@example.com", "farel456"}
	dbPengguna[4] = Pengguna{5, "Geby", "geby22@example.com", "geby456"}
	dbPengguna[5] = Pengguna{6, "Caca", "cacaboom@example.com", "caca456"}
	dbPengguna[6] = Pengguna{7, "Geebry", "geebry@example.com", "1234567"}
	// Isi data pinjaman
	dbDataPeminjam[0] = Pinjaman{1, 1000000, 12, 5, 12, false}
	dbDataPeminjam[1] = Pinjaman{2, 2000000, 24, 4, 12, true}
	dbDataPeminjam[2] = Pinjaman{3, 1500000, 18, 6, 12, false}
	dbDataPeminjam[3] = Pinjaman{4, 2000000, 18, 6, 12, false}
	dbDataPeminjam[4] = Pinjaman{5, 2500000, 18, 6, 12, false}
	dbDataPeminjam[5] = Pinjaman{6, 1000000, 12, 5, 12, true}
	dbDataPeminjam[6] = Pinjaman{7, 1200000, 12, 5, 12, false}
}
// mengii data awal ke sistem . antara id peminjam dan id pengguna ada relasinya 
// misal id pengguna 2 memuat id peminjam 2 artinya di di 2 memuat pinjaman elfan