# INTRODUCTION
## 💸 TRA-Pinjaman: Aplikasi Simpan Pinjam Online [Golang-CLI] 
**TraPinjaman** adalah aplikasi simulasi pinjaman sederhana berbasis terminal yang dibuat menggunakan bahasa pemrograman **Golang**.
Aplikasi ini dirancang untuk meniru alur sistem pinjaman dan kredit secara virtual, lengkap dengan fitur login, registrasi, pengajuan pinjaman, pelunasan cicilan, hingga pencatatan riwayat pembayaran. Seluruh data disimpan secara lokal menggunakan struktur data array, tanpa menggunakan database eksternal.
Tujuan utama proyek ini adalah sebagai latihan dan pembuktian kemampuan dalam mengimplementasikan:
 - Konsep dasar pemrograman Go
 - Struktur data statis (`array`)
 - Modularisasi fungsi
 - Penerapan logika **CRUD** (Create, Read, Update, Delete)

## ⛓️ FITUR TRA-Pinjaman
| Fitur              | Keterangan  |
------------------------------------
-  **Register & Login** || Pengguna dapat mendaftar dan masuk dengan nama, email, dan password unik
-  **Ajukan Pinjaman**  || Simulasi pengajuan pinjaman sesuai nominal dan tenor (hari/bulan)
-  **Pelunasan**        || Bayar cicilan sesuai jumlah angsuran, otomatis tercatat
-  **Lihat Profil**     || Menampilkan informasi pengguna dan status pinjaman
-  **Riwayat Pelunasan** || Menyimpan histori transaksi (tanggal, jumlah bayar, metode, sisa)
-  **ResetPinjaman**     || Menghapus pinjaman aktif jika dibutuhkan

## 🛠 TEKNOLOGI & TOOLS
- **Golang** `v1.xx`
-  Visual Studio Code (VSCode)
-  CLI Terminal
-  Git & GitHub

## 👨‍💻 CARA MENJALANKAN APLIKASI
1. Pastikan Go ver 1.xx sudah terinstall di sistem Anda (OS : Windows, macOS, atau Linux).
2. Kloning Repsositori atau mengekstrak file `.zip` terlebih dahulu.
3. Lalu jalankan program dengan
   ```
    go run . go
   ```
   
## 📁 STRUKTUR FILE
1. **main1.go** Fungsi utama & alur program
2. **auth.go** Login & register
3. **db[1].go** Struktur data & database array
4. **history.go** Pelunasan dan data riwayat setoran
5. **seed.go** Data awal pengguna & pinjaman
6. **go.mod** file konfigurasi modul Go
7. **READ-ME** Dokumentasi proyek

## 📂 PENJELASAN MODUL PER FILE
### 🔗 **main1.go**
> _File utama yang berfungsi sebagai entry point aplikasi_
### Fungsi :
  * `main()` : fungsi utama aplikasi. Fungsi yang menampilkan menu awal (login, regrister, exit)
  * `menuAwal()` : menampilkan menu login & regrister
  * `menuUtama(userAktif *Pengguna)` : menampilkan menu utama pengguna (Ajukan Pinjaman, Pelunasan, Profil Saya, Reset Pinjaman, Keluar)
  * `ajukanPinjaman()` : fungsi untuk mengajukan pinjaman bagi pengguna
  * `pelunasan()` : digunakan untuk membayar pinjaman pengguna
  * `tampilkanProfil()` : untuk menampilkan data pengguna dan status pinjaman, serta riwayat pelunasan
  *  * `hapusPinajaman()` : digunakan untuk menghapus data pinjaman hanya jika sudah lunas
  * `hitungBunga(nominal int) float64` : menentukan tingkat bunga sesuai jumlah pinjaman
  * `pilihanTenor(nominal int)` : menampilkan pilihan tenor & cicilan berdasarkan jumlah pinjaman
 
### 🔗 **auth.go**
>_Modul yang menangani proses autentikasi pengguna_
#### Fungsi :
* `masuk(db *[100]Pengguna)*Pengguna` : fungsi untuk melakukan login dengan mencocokkan nama dan password dari input pengguna dengann data yang tersimpan di dbPengguna
* `daftar(db *[100]Pengguna)*Pengguna` : fungsi untuk mendaftarkan pengguna baru dengan inputan nama, email, password dan akan disimpan di slot kosong idPengguna == 0


### 🔗 **db[1].go**
>_Database pengguna dan pinjaman dalam bentuk array_
#### Tipe Data :
* `type Pinjaman` : dengan field `idPeminjam`, `jumlahPinjaman`, `tenor`, `bunga`, `statusLunas`, dll.
* `type Pengguna` : dengan field `idPengguna`, `nama`, `email`, `password`
#### Variabel :
* `var dbPengguna [100]Pengguna` : "database" pengguna
* `var dbDataPeminjam [100]Pinjaman` : "database" pinjaman
#### Fungsi :
* `indeksKosongDbPeminjam(db *[100]Pinjaman)` : fungsi untuk mencari ideks kosong dalam array pinjaman

#### 🔗 **history.go**
>_Modul yang menangani pencatatan dan tampilan riwayar pinjaman pengguna_
#### Tipe Data :
* `type Pelunasan` : dengan field `nama`, `tanggal`, `jumlahBayar`, `metode`, `sisaPinjaman`
#### Variabel :
* `var riwayatPelunasan [100]Pelunasan`
* `var totalRiwayat int`
#### Fungsi :
* `tambahRiwayat (...)` : fungsi yang mencatat pelunasan pinjaman ke dalam array dengan tanggal saat ini `(time.Now())`
* `tampilkanRiwayat(nama string) : menampilkan daftar pelunasan pinjaman pengguna berdasarkan nama

### 🔗 **seed.go**
>_File berisi data dummy atau data awal yang digunakan agar program dapat dijalankan_
#### Fungsi :
* `seed()` : fungsi untuk mengisi `dbPengguna` dan `dbDataPeminjam`




   
