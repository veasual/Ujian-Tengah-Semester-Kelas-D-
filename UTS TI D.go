package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct untuk Buku
type Buku struct {
	NamaBuku string
	Jumlah   int
}

// Struct untuk Peminjaman
type Peminjaman struct {
	Username string
	Buku     Buku
	Jumlah   int
}

// Inisialisasi data awal buku
var daftarBuku = []Buku{
	{"Pemrograman", 10},
	{"Film", 5},
	{"Printing", 20},
}

// Array untuk menyimpan histori peminjaman
var historiPeminjaman []Peminjaman

// Variabel untuk menyimpan username dan password (NPM)
var username string = "Thorieq"
var password string = "2406441184"
var NPM string = "2406441184"
var jeniskelamin string = "Laki-Laki"
var makanan string = "mie ayam"
var minuman string = "jus alpukat"

// Fungsi untuk membaca input dari pengguna
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi Lihat Informasi Pengguna Program
func LihatInformasiPenggunaProgram() {
	fmt.Println("=========================================")
	fmt.Println("Informasi Pengguna Program")
	fmt.Println("Username:", username)
	fmt.Println("NPM", NPM)
	fmt.Println("Jenis Kelamin", jeniskelamin)
	fmt.Println("Makanan Favorit", makanan)
	fmt.Println("Minuman Favorit", minuman)
	fmt.Println("=========================================")
}

// Fungsi Lihat Daftar Buku
func LihatDaftarBuku() {
	fmt.Println("=========================================")
	fmt.Println("Daftar Buku yang Tersedia")
	for _, buku := range daftarBuku {
		fmt.Printf("Nama Buku: %s, Jumlah: %d\n", buku.NamaBuku, buku.Jumlah)
	}
	fmt.Println("=========================================")
}

// Fungsi Tambah Daftar Buku
func TambahDaftarBuku() {
	namaBuku := input("Masukkan nama buku yang ingin ditambahkan: ")
	jumlah := input("Masukkan jumlah buku: ")

	// Cek apakah jumlah yang dimasukkan valid
	var jumlahBuku int
	fmt.Sscanf(jumlah, "%d", &jumlahBuku)

	if jumlahBuku <= 0 {
		fmt.Println("Jumlah buku harus lebih dari 0")
		return
	}

	// Tambah buku ke daftar
	for i, buku := range daftarBuku {
		if buku.NamaBuku == namaBuku {
			daftarBuku[i].Jumlah += jumlahBuku
			fmt.Println("Buku berhasil ditambahkan!")
			return
		}
	}
	daftarBuku = append(daftarBuku, Buku{NamaBuku: namaBuku, Jumlah: jumlahBuku})
	fmt.Println("Buku baru berhasil ditambahkan!")
}

// Fungsi Tambah Peminjaman Buku
func TambahPeminjamanBuku() {
	LihatDaftarBuku()
	namaBuku := input("Masukkan nama buku yang ingin dipinjam: ")
	jumlah := input("Masukkan jumlah buku yang ingin dipinjam: ")

	var jumlahBuku int
	fmt.Sscanf(jumlah, "%d", &jumlahBuku)

	if jumlahBuku <= 0 {
		fmt.Println("Jumlah peminjaman tidak valid.")
		return
	}

	for i, buku := range daftarBuku {
		if buku.NamaBuku == namaBuku {
			if buku.Jumlah < jumlahBuku {
				fmt.Println("Jumlah buku tidak mencukupi.")
				return
			}
			// Kurangi stok buku
			daftarBuku[i].Jumlah -= jumlahBuku
			// Tambahkan ke histori peminjaman
			historiPeminjaman = append(historiPeminjaman, Peminjaman{
				Username: username,
				Buku:     buku,
				Jumlah:   jumlahBuku,
			})
			fmt.Println("Buku berhasil dipinjam!")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan.")
}

// Fungsi Histori Peminjaman Buku
func HistoriPeminjamanBuku() {
	fmt.Println("=========================================")
	fmt.Println("Histori Peminjaman Buku")
	for _, peminjaman := range historiPeminjaman {
		fmt.Printf("Username: %s, Buku: %s, Jumlah: %d\n", peminjaman.Username, peminjaman.Buku.NamaBuku, peminjaman.Jumlah)
	}
	fmt.Println("=========================================")
}

// Fungsi Keluar dari Program
func KeluarDariProgram() {
	fmt.Println("Terima kasih telah menggunakan program ini. Sampai Jumpa!")
	os.Exit(0)
}

// Fungsi Utama (main)
func main() {
	fmt.Println("=========================================")
	fmt.Println("= Selamat Datang di Perpustakaan Vokasi =")
	fmt.Println("=========================================")

	// Verifikasi login
	inputUsername := input("Silahkan Input Username: ")
	inputPassword := input("Silahkan Input Password: ")

	if inputUsername != username || inputPassword != password {
		fmt.Println("Username atau Password salah!")
		return
	}

	fmt.Println("=========================================")
	fmt.Println("Login Sukses!")
	fmt.Println("=========================================")

	// Menu utama
	for {
		fmt.Println("\nMenu Program")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar Dari Program")

		choice := input("Masukkan pilihan menu: ")

		switch choice {
		case "1":
			LihatInformasiPenggunaProgram()
		case "2":
			LihatDaftarBuku()
		case "3":
			TambahDaftarBuku()
		case "4":
			TambahPeminjamanBuku()
		case "5":
			HistoriPeminjamanBuku()
		case "6":
			KeluarDariProgram()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
