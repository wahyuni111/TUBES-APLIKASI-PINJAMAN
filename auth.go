package main

import "fmt"

func masuk(db *[100]Pengguna) *Pengguna {
	var (
		nNama     string
		nPassword string
	)

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nNama)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&nPassword)

	for i := 0; i < len(db); i++ {
		if nNama == db[i].nama && nPassword == db[i].password {
			return &db[i]
		}
	}

	return nil
}
func daftar(db *[100]Pengguna)*Pengguna{
	var nama, email, password string

	fmt.Println("\n===== Register ====")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)

	fmt.Print("Masukkan password dengan panjang 6 karakter: ")
	fmt.Scan(&password)

	//seacrh indeks yg kosong

	for i:=0; i<len(db); i++{
		if db[i].idPengguna ==0{
			db[i]= Pengguna{
				idPengguna : i+1,
				nama : nama,
				email: email,
				password: password,
			}
			fmt.Println("Pendaftaran Berhasil! Silahkan login.")
			return &db[i] // harus return pointer pengguna yg baru didaftarkan
		}
	}
	fmt.Println("Gagal daftar. Coba lagi!.")
	return nil
}
//fitur login 
//