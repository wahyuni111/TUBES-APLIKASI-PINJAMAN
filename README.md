# INTRODUCTION
## TRA-Pinjaman: Aplikasi Simpan Pinjam Online [Golang-CLI]
**TraPinjaman** adalah aplikasi simulasi pinjaman sederhana berbasis terminal yang dibuat menggunakan bahasa pemrograman **Golang**.
Aplikasi ini dirancang untuk meniru alur sistem pinjaman dan kredit secara virtual, lengkap dengan fitur login, registrasi, pengajuan pinjaman, pelunasan cicilan, hingga pencatatan riwayat pembayaran. Seluruh data disimpan secara lokal menggunakan struktur data array, tanpa menggunakan database eksternal.
Tujuan utama proyek ini adalah sebagai latihan dan pembuktian kemampuan dalam mengimplementasikan:
 - Konsep dasar pemrograman Go
 - Struktur data statis (`array`)
 - Modularisasi fungsi
 - Penerapan logika **CRUD** (Create, Read, Update, Delete)

## Fitur TRA-Pinjaman
| Fitur              | Keterangan  |
------------------------------------
-  **Register & Login** || Pengguna dapat mendaftar dan masuk dengan nama, email, dan password unik
-  **Ajukan Pinjaman**  || Simulasi pengajuan pinjaman sesuai nominal dan tenor (hari/bulan)
-  **Pelunasan**        || Bayar cicilan sesuai jumlah angsuran, otomatis tercatat
-  **Lihat Profil**     ||Menampilkan informasi pengguna dan status pinjaman
-  **Riwayat Pelunasan** ||Menyimpan histori transaksi (tanggal, jumlah bayar, metode, sisa)
-  **ResetPinjaman**     ||Menghapus pinjaman aktif jika dibutuhkan

## 🛠 TEKNOLOGI & TOOLS
- **Golang** `v1.xx`
-  Visual Studio Code (VSCode)
-  CLI Terminal
-  Git & GitHub

## STRUKTUR FILE
1. **main.go** Fungsi utama & alur program
2. **auth.go** Login & register
3. **db.go** Struktur data & database array
4. **history.go** Pelunasan dan data riwayat setoran
5. **seed.go** Data dummy awal pengguna & pinjaman
6. **READ-ME**

7. 
