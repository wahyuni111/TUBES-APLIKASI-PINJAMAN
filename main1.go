package main

import "fmt"

func main (){
	seed()
	userAktif := menuAwal()
	if userAktif == nil{
		fmt.Println("Layanan selesai.")
		return
	}
	menuUtama(userAktif)
}
func menuAwal()*Pengguna{
	var pilihan int
	var userAktif *Pengguna

	for {
		fmt.Println("\n=== Selamat Datang di Aplikasi TraPinjaman ===")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan{
		case 1:
			maxCoba := 3 // batasan login
			for i:=0; i<maxCoba; i++{
				user:= masuk(&dbPengguna)
				if user != nil{
					fmt.Printf("\nHalo %s!\n", user.nama)
					fmt.Println("Selamat Datang di TraPinjaman Online!")
					userAktif = user
					return userAktif
				}else{
					fmt.Println("Username atau password anda salah!")
				}
			}
			fmt.Println("Anda sudah terlalu banyak melakukan percobaan.Keluar dan kembali setelah 1 jam kemudian!")
			return nil

		case 2:
			daftar(&dbPengguna)
		case 0:
			fmt.Println("Terima kasih.Sampai jumpa!")
			return nil

		default:
			fmt.Println("Pilihan anda tidak valid!")
		}
		return nil
		}
	}

func menuUtama(userAktif *Pengguna) {
	for {
		var pilihan int
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("[1] Ajukan Pinjaman")
		fmt.Println("[2] Pelunasan")
		fmt.Println("[3] Profil Saya")
		fmt.Println("[4] Reset Pinjaman saya")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			ajukanPinjaman(userAktif, &dbDataPeminjam)
		case 2:
			pelunasan(userAktif, &dbDataPeminjam)
		case 3:
			tampilkanProfil(userAktif, &dbDataPeminjam)
		case 4 :
			hapusPinjaman(userAktif, &dbDataPeminjam)
		case 0:
			fmt.Println("Terima kasih, sampai jumpa lagi di TraPinjaman Online!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
//func main() {
	//var (
		//pilihan     int
		//statusMasuk bool
		//userAktif   *Pengguna
	//)
	//seed() //isi dlu data awal dr file seed.go
//login
	//}
	//userAktif = masuk(&dbPengguna)
	//statusMasuk = userAktif != nil

	//if !statusMasuk {
		//fmt.Println("Masuk gagal")
		//return
	//}

	//fmt.Println("[1] Ajukan Pinjaman") //nominal, tenor
	//fmt.Println("[2] Pelunasan")       //pilih setoran, sontok setoran
	//fmt.Println("[3] Urutkan Data Peminjam")
	//fmt.Println("[3] Profil saya")
	//fmt.Println("[0] Kembali")
	//fmt.Print("Pilih: ")
	//fmt.Scan(&pilihan)

	//switch pilihan {
	//case 1:
		//ajukanPinjaman(userAktif, &dbDataPeminjam)
	//case 2:
		//pelunasan(userAktif, &dbDataPeminjam)/ fungsi tampil pinjaman
	//case 3:
		//
	//case 0:
		//fmt.Println("Terima kasih, sampai jumpa lagi di TraPinjaman Online!") 
		//return
	//default:
		//fmt.Println("Pilihan tidak valid.")
	//}
func ajukanPinjaman(userAktif *Pengguna, db *[100]Pinjaman) {
	var (
		nNominal       int
		nTenor         int
		nAngsuran      int
		nBunga         float64
		indeksTersedia int
	)

	fmt.Print("Masukkan nominal Dengan ketentuan tanpa menggunakan titik, contoh 10000000: ")
	fmt.Scan(&nNominal)

	pilihanTenor(nNominal)

	fmt.Print("Pilih tenor: ")
	fmt.Scan(&nTenor)
	fmt.Print("Pilih angsuran: ")
	fmt.Scan(&nAngsuran)

	nBunga = hitungBunga(nNominal)
	//insert ke userAktif.idPengguna == db.idPeminjam
	indeksTersedia = indeksKosongDbPeminjam(db)
	if indeksTersedia == -1 {
	fmt.Println("Maaf. Gagal mengajukan pinjaman: database sudah penuh.")
	return
}

	db[indeksTersedia] = Pinjaman{
		idPeminjam:     userAktif.idPengguna,
		jumlahPinjaman: nNominal,
		tenor:          nTenor,
		bunga:          nBunga,
		jumlahAngsuran: nAngsuran,
		statusLunas:    false,
	}

}
func pelunasan(userAktif *Pengguna, db *[100]Pinjaman) {
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == userAktif.idPengguna && !db[i].statusLunas {
			fmt.Printf("\nSisa Pinjaman Anda: Rp%d\n", db[i].jumlahPinjaman)
			var bayar int
			fmt.Print("Masukkan jumlah yang ingin dibayar: ")
			fmt.Scan(&bayar)
			db[i].jumlahPinjaman -= bayar

			if db[i].jumlahPinjaman <= 0 {
				db[i].statusLunas = true
				db[i].jumlahPinjaman = 0
				fmt.Println("Selamat, pinjaman Anda sudah lunas!")
			} else {
				fmt.Printf("Sisa pinjaman: Rp%d\n", db[i].jumlahPinjaman)
			}
			return
		}
	}
	fmt.Println("Anda tidak memiliki pinjaman aktif.")
}

func hitungBunga(nominal int) float64 {
	switch {
	case nominal < 10000000:
		return 0.05
	case nominal <= 50000000:
		return 0.08
	case nominal <= 100000000:
		return 0.15
	default:
		return 0.18
	}
}
func pilihanTenor(nominal int) {
	fmt.Printf("Nominal: Rp%d\n", nominal)
	switch {
	case nominal <= 300000:
		fmt.Println("Pilihan tenor:7 Hari,14 Hari, 30 Hari")
		fmt.Println("Cicilan: 1x, 2x, 3x")
	case nominal <= 1000000:
		fmt.Println("Pilihan tenor: 30 Hari, 60 Hari, 90 Hari")
		fmt.Println("Cicilan: 1x, 2x, 3x")
	case nominal <= 3000000:
		fmt.Println("Pilihan tenor: 3 Bulan, 4 Bulan, 6 Bulan")
		fmt.Println("Cicilan: 3x, 4x, 6x")
	case nominal <= 5000000:
		fmt.Println("Pilihan tenor: 3 Bulan, 6 Bulan, 9 Bulan")
		fmt.Println("Cicilan: 3x, 6x, 9x")
	case nominal <= 10000000:
		fmt.Println("Pilihan tenor: 6 Bulan, 9 Bulan, 12 Bulan")
		fmt.Println("Cicilan: 6x, 9x, 12x")
	case nominal <= 20000000:
		fmt.Println("Pilihan tenor: 12 Bulan, 18 Bulan, 24 Bulan")
		fmt.Println("Cicilan: 12x, 18x, 24x")
	case nominal <= 30000000:
		fmt.Println("Pilihan tenor: 24 Bulan, 30 Bulan, 36 Bulan")
		fmt.Println("Cicilan: 24x, 30x, 36x")
	case nominal <= 50000000:
		fmt.Println("Pilihan tenor: 24 Bulan, 30 Bulan, 36 Bulan")
		fmt.Println("Cicilan: 24x, 30x, 36x")
	default:
		fmt.Println("Nominal melebihi batas dukungan.")
	}
}
func tampilkanSemuaPinjaman(db *[100]Pinjaman) {
	fmt.Println("\n=== DAFTAR PINJAMAN TERURUT ===")
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam != 0 {
			fmt.Printf("- Pinjaman ke-%d\n", i+1)
			fmt.Printf("  Jumlah: Rp%d\n", db[i].jumlahPinjaman)
			fmt.Printf("  Tenor : %d bulan\n", db[i].tenor)
			fmt.Printf("  Bunga : %.2f%%\n", db[i].bunga*100)
			fmt.Println("-----------------------------")
		}
	}
}
func tampilkanProfil(userAktif *Pengguna, db *[100]Pinjaman) {
	fmt.Println("\n=== Profil Anda ===")
	fmt.Printf("Nama    : %s\n", userAktif.nama)
	fmt.Printf("Email   : %s\n", userAktif.email)

	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == userAktif.idPengguna {
			fmt.Printf("Pinjaman: Rp%d\n", db[i].jumlahPinjaman)
			fmt.Printf("Tenor   : %d bulan\n", db[i].tenor)
			fmt.Printf("Status  : %s\n", 
				func() string {
					if db[i].statusLunas {
						return "Lunas"
					}
					return "Belum Lunas"
				}())
			return
		}
	}
	fmt.Println("Anda belum memiliki pinjaman.")
}
func hapusPinjaman(userAktif *Pengguna, db *[100]Pinjaman) {
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == userAktif.idPengguna {
			db[i] = Pinjaman{} // kosongkan struct
			fmt.Println("Pinjaman Anda berhasil dihapus.")
			return
		}
	}
	fmt.Println("Tidak ada pinjaman untuk dihapus.")
}
