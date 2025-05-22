package main

import (
	"fmt"
)

type Pengeluaran struct {
	Kategori string
	Jumlah   float64
}

var data []Pengeluaran
var totalBudget float64

func tambahData() {
	var kategori string
	var jumlah float64

	fmt.Print("Masukkan kategori: ")
	fmt.Scanln(&kategori)
	fmt.Print("Masukkan jumlah: ")
	fmt.Scanln(&jumlah)

	data = append(data, Pengeluaran{Kategori: kategori, Jumlah: jumlah})
	fmt.Println("✅ Data berhasil ditambahkan.")
}

func tampilkanData() {
	if len(data) == 0 {
		fmt.Println("❌ Tidak ada data pengeluaran.")
		return
	}

	fmt.Println("📋 Daftar Pengeluaran:")
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d. %s - Rp %.2f\n", i+1, data[i].Kategori, data[i].Jumlah)
	}
}

func ubahData() {
	if len(data) == 0 {
		fmt.Println("❌ Tidak ada data untuk diubah.")
		return
	}

	tampilkanData()
	var i int
	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scanln(&i)

	fmt.Print("Masukkan kategori baru: ")
	fmt.Scanln(&data[i-1].Kategori)
	fmt.Print("Masukkan jumlah baru: ")
	fmt.Scanln(&data[i-1].Jumlah)
	fmt.Println("✅ Data berhasil diubah.")
}

func hapusData() {
	if len(data) == 0 {
		fmt.Println("❌ Tidak ada data untuk dihapus.")
		return
	}

	tampilkanData()
	var i int
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scanln(&i)

	for j := i - 1; j < len(data)-1; j++ {
		data[j] = data[j+1]
	}
	data = data[:len(data)-1]
	fmt.Println("🗑 Data berhasil dihapus.")
}

func anggaran() {
	var total float64
	for i := 0; i < len(data); i++ {
		total += data[i].Jumlah
	}
	fmt.Printf("💰 Total Pengeluaran: Rp %.2f\n", total)
	fmt.Printf("🎯 Budget Awal: Rp %.2f\n", totalBudget)

	if total <= totalBudget {
		fmt.Printf("✅ Sisa Budget: Rp %.2f\n", totalBudget-total)
	} else {
		fmt.Printf("⚠ Melebihi Budget: Rp %.2f\n", total-totalBudget)
	}
}

func cariDataSequential() {
	if len(data) == 0 {
		fmt.Println("❌ Tidak ada data untuk dicari.")
		return
	}

	var cari string
	var ketemu bool = false

	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scanln(&cari)

	for i := 0; i < len(data); i++ {
		if data[i].Kategori == cari {
			fmt.Println("✅ Ditemukan pada data nomor", i+1)
			fmt.Println("Kategori:", data[i].Kategori, "Jumlah:", data[i].Jumlah)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("❌ Data tidak ditemukan.")
	}
}

func urutkanJumlah() {
	if len(data) == 0 {
		fmt.Println("❌ Tidak ada data untuk diurutkan.")
		return
	}

	var pilihan string
	fmt.Print("Urutkan jumlah (terkecil/terbesar): ")
	fmt.Scanln(&pilihan)

	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (pilihan == "terkecil" && data[j].Jumlah > data[j+1].Jumlah) ||
				(pilihan == "terbesar" && data[j].Jumlah < data[j+1].Jumlah) {

				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}

	fmt.Println("📊 Data berhasil diurutkan berdasarkan jumlah (" + pilihan + ").")
	tampilkanData()
}

func urutkanKategori() {
	if len(data) == 0 {
		fmt.Println("❌ Tidak ada data untuk diurutkan.")
		return
	}

	var pilihan string
	fmt.Print("Urutkan kategori (terkecil/terbesar): ")
	fmt.Scanln(&pilihan)

	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (pilihan == "terkecil" && data[j].Kategori > data[j+1].Kategori) ||
				(pilihan == "terbesar" && data[j].Kategori < data[j+1].Kategori) {

				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}

	fmt.Println("📊 Data berhasil diurutkan berdasarkan kategori (" + pilihan + ").")
	tampilkanData()
}

func main() {
	var pilihMenu int

	fmt.Print("Masukkan total budget Anda: Rp")
	fmt.Scanln(&totalBudget)

	for {
		fmt.Println("\n=== Menu Budget Traveling ===")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Tampilkan Pengeluaran")
		fmt.Println("3. Ubah Pengeluaran")
		fmt.Println("4. Hapus Pengeluaran")
		fmt.Println("5. Tampilkan Anggaran")
		fmt.Println("6. Cari Pengeluaran")
		fmt.Println("7. Urutkan Berdasarkan Jumlah")
		fmt.Println("8. Urutkan Berdasarkan Kategori")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihMenu)

		switch pilihMenu {
		case 1:
			tambahData()
		case 2:
			tampilkanData()
		case 3:
			ubahData()
		case 4:
			hapusData()
		case 5:
			anggaran()
		case 6:
			cariDataSequential()
		case 7:
			urutkanJumlah()
		case 8:
			urutkanKategori()
		case 9:
			fmt.Println("👋 Terima kasih sudah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("❗ Pilihan tidak valid.")
		}
	}
}
