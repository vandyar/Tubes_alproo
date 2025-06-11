package main

import "fmt"

const NMAX int = 100

type SparePart struct {
	ID    int
	Nama  string
	Harga int
	Freq  int
}

type Service struct {
	Pelanggan string
	SparePart string
}

type Transaksi struct {
	Nama      string
	SparePart string
	HargaJasa float64
}

type pelanggan struct {
	ID        int
	Nama      string
	NoHp      int
	HargaJasa float64
}

var dataSparePart [NMAX]SparePart
var nSparepart int

var dataPelanggan [NMAX]pelanggan
var nPelanggan int

var dataService [NMAX]Service
var nService int

var dataTransaksi [NMAX]Transaksi
var nTransaksi int

func main() {

	var klik int

	for klik != 5 {

		fmt.Println("=====menu utama=====")
		fmt.Println("1. menu sparepat ")
		fmt.Println("2. menu Pelanggan")
		fmt.Println("3. menu Transaksi ")
		fmt.Println("4. keluar ")
		fmt.Println("====================")
		fmt.Print("Silahkan masukkan pilihan :")
		fmt.Scan(&klik)

		if klik == 1 {
			menuSparePart()
		} else if klik == 2 {
			menuPelanggan()
		} else if klik == 3 {
			menuTransaksi()
			return
		} else {

			fmt.Print("menu tidak tersedia")
		}
	}
}

func menuSparePart() {
	var pilih int

	for pilih != 6 {
		fmt.Println("===== Menu Service Motor =====")
		fmt.Println("1. Tambah Data Sparepart      ")
		fmt.Println("2. Edit Sparepart             ")
		fmt.Println("3. Hapus Sparepart            ")
		fmt.Println("4. Lihat semua Sparepart      ")
		fmt.Println("5. Kembali                    ")
		fmt.Println("6. Tambah Pelanggan           ")
		fmt.Println("==============================")
		fmt.Print("Masukan pilihan anda: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahData()
		} else if pilih == 2 {
			editData()
		} else if pilih == 3 {
			hapusData()
		} else if pilih == 4 {
			hasilData()
		} else if pilih == 5 {
			fmt.Println("Terima kasih telah menggunakan layanan kami!")
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}

	}
}

func menuPelanggan() {
	var klik int 

	 for !=  5 (
		fmt.Println("=============")
		fmt.Println("1.Tambah pelanggan ")
	 )




}

func menuTransaksi() {

}

func tambahData() {
	var pelanggan Service
	var kategori, subKategori int
	if nService >= NMAX {
		fmt.Println("Data service sudah penuh!")
		return

	}

	fmt.Print("silahkan masukan pelanggan : ")
	fmt.Scan(&pelanggan.Pelanggan)

	for kategori != 6 {
		fmt.Println("pilih kategori sparepart yang ingin diservice:")
		fmt.Println("1.mesin                              ")
		fmt.Println("2.Kelistrikan                        ")
		fmt.Println("3.kaki - kaki                        ")
		fmt.Println("4.Rem                               m ")
		fmt.Println("5. kembali kemenu utama              ")
		fmt.Print("masukan pilihan (1-5) : ")
		fmt.Scan(&kategori)

		if kategori == 5 {
			return 

		}

		if kategori == 1 {
			fmt.Println("Pilih sparePart untuk Mesin :")
			fmt.Println("1.Oli mesin")
			fmt.Println("2.Busi")
			fmt.Println("3.filter udara")
			fmt.Println("4.kampas kopling")
			fmt.Print("Mauskan pilihan (1-4) :")
			fmt.Scan(&subKategori)

			if subKategori >= 1 && subKategori <= 4 {
				if subKategori == 1 {
					pelanggan.SparePart = "Oli mesin"
				} else if subKategori == 2 {
					pelanggan.SparePart = "Busi"
				} else if subKategori == 3 {
					pelanggan.SparePart = "filter udara"
				} else if subKategori == 4 {
					pelanggan.SparePart = "kampas kopling"
				}
			} else {
				fmt.Println("Pilhan tidak ada!")
				return
			}

		} else if kategori == 2 {
			fmt.Println("Pilih SparePart untuk kelistrikan :")
			fmt.Println("1.aki")
			fmt.Println("2.lampu")
			fmt.Println("3.Saklar")
			fmt.Println("4.CDI unit")
			fmt.Print("Masukan pilihan (1-4) :")
			fmt.Scan(&subKategori)

			if subKategori >= 1 && subKategori <= 4 {
				if subKategori == 1 {
					pelanggan.SparePart = "Aki"
				} else if subKategori == 2 {
					pelanggan.SparePart = "Lampu"
				} else if subKategori == 3 {
					pelanggan.SparePart = "Saklar"
				} else if subKategori == 4 {
					pelanggan.SparePart = "CDI unit "
				}
			} else {
				fmt.Println("Pilhan tidak ada!")
				return
			}

		} else if kategori == 3 {
			fmt.Println("Pilih sparePart untuk kaki - kaki :")
			fmt.Println("1.Ban")
			fmt.Println("2.Shockbreaker")
			fmt.Println("3.Velg")
			fmt.Println("4.Rantai dan Gear")
			fmt.Print("Masukan pilihan (1-4) :")
			fmt.Scan(&subKategori)

			if subKategori >= 1 && subKategori <= 4 {
				if subKategori == 1 {
					pelanggan.SparePart = "Ban"
				} else if subKategori == 2 {
					pelanggan.SparePart = "Shockbreaker"
				} else if subKategori == 3 {
					pelanggan.SparePart = "Velg"
				} else if subKategori == 4 {
					pelanggan.SparePart = "Rantai dan Gear"
				}
			} else {
				fmt.Println("Pilhan tidak ada!")
				return
			}

		} else if kategori == 4 {
			fmt.Println("Pilih sparePart untuk rem : ")
			fmt.Println("1.kampas rem")
			fmt.Println("2.Master rem")
			fmt.Println("3.selang rem")
			fmt.Print(" masukan pilihan (1-3) :")
			fmt.Scan(&subKategori)

			if subKategori >= 1 && subKategori <= 3 {
				if subKategori == 1 {
					pelanggan.SparePart = "kampas rem"
				} else if subKategori == 2 {
					pelanggan.SparePart = "master rem"
				} else if subKategori == 3 {
					pelanggan.SparePart = "selang rem"
				}
			} else {
				fmt.Println("Pilhan tidak ada!")
				return

			}

		} else {
			fmt.Println("Pilihan tidak valid")
			return

		}

		dataService[nService] = pelanggan
		nService++
		fmt.Println("data berhasil ditambah !")

	}

}

func hasilData() {
	if nService == 0 {
		fmt.Println("belum ada data service")
		return
	}

	fmt.Println("daftar data service :")
	for i := 0; i < nService; i++ {
		fmt.Printf("%d. Pelanggan : %s, SparePart: %s\n", i+1, dataService[i].Pelanggan, dataService[i].SparePart)
	}
}

func editData() {
	if nService == 0 {
		fmt.Println("Belum ada data service")
		return

	}

	var namaCari string

	fmt.Println("=======Edit data=======")
	hasilData()
	fmt.Print("masukan data yang mau di edit: ")
	fmt.Scan(&namaCari)

	for i := 0; i < nService; i++ {
		if dataService[i].Pelanggan == namaCari {
			fmt.Println("Masukan nama pelanggan baru :")
			fmt.Scan(&dataService[i].Pelanggan)
			fmt.Println("masukan sparepart baru :")
			fmt.Scan(&dataService[i].SparePart)
			fmt.Print("Data berhasil diubah")
			return
		}
	}

}

func hapusData() {
	if nService == 0 {
		fmt.Println("Masukan data pelanggan :")
		return
	}

	hasilData()
	var idx int
	fmt.Print("Masukkan nomom data yang ingin dihapus: ")
	fmt.Scan(&idx)
	if idx < 1 || idx > nService {
		fmt.Println("Nomor tidak ada")
		return

	}

	for i := idx - 1; i < nService-1; i++ {
		dataService[i] = dataService[i+1]
	}
	nService--
	fmt.Println("Data berhasil di Hapuas ")
}

func tambahPelanggan() {
	var p pelanggan
	if nPelanggan >= NMAX {
		fmt.Println("Data pelanggan sudah penuh!")
		return
	}

	fmt.Println("MAsukan Id Pelanggan :")
	fmt.Scan(&p.ID)
	fmt.Println("Masukan Nama pelanggan :")
	fmt.Scan(&p.Nama)
	fmt.Println("Masukan No Hp pelanggan :")
	fmt.Scan(&p.NoHp)
	fmt.Println(" Masukan harga jasa: ")
	fmt.Scan(&p.HargaJasa)

	dataPelanggan[nPelanggan] = p
	nPelanggan++
	fmt.Println("Data berhasil tambah pelanggan")

}
