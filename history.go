package main 

import(
	"fmt"
	"time"
)

type Pelunasan struct {
	nama string 
	tanggal string 
	jumlahBayar int
	metode string 
	sisaPinjaman int 
}

var riwayatPelunasan [100]Pelunasan
var totalRiwayat int

func tambahRiwayat(nama string, jumlah int, metode string, sisa int){
	tanggalSekarang :=time.Now().Format("2006-01-02")
	riwayatPelunasan[totalRiwayat]=Pelunasan{
		nama: nama,
		tanggal: tanggalSekarang,
		jumlahBayar: jumlah,
		metode : metode,
		sisaPinjaman: sisa,
	}
	totalRiwayat++
}
func tampilkanRiwayat(nama string){
	fmt.Println("\n -------Riwayat Peluasaan----")
	ketemu := false

	for i:=0; i<totalRiwayat; i++{
		if riwayatPelunasan[i].nama==nama{
			ketemu = true
			fmt.Printf("Tanggal : %s\n", riwayatPelunasan[i].tanggal)
			fmt.Printf("Jumlah  : Rp%d\n", riwayatPelunasan[i].jumlahBayar)
			fmt.Printf("Metode  : %s\n", riwayatPelunasan[i].metode)
			fmt.Printf("Sisa    : Rp%d\n", riwayatPelunasan[i].sisaPinjaman)
			fmt.Println("------------------------------")
		}
	}
	if !ketemu{
		fmt.Println("Belum ada riwayat Plenasan")
	}
}
