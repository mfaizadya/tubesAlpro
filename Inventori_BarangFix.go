package main

import (
	"fmt"
)

const NMAX int = 1000

type db struct {
	kodeBarang string
	namaBarang string
	stok       int
	harga      int
}

type transaction struct {
	tanggal    string
	kodeBarang string
	jenis      string // "masuk" atau "keluar"
	jumlah     int
}

type tabdatbar [NMAX]db
type tabTrans [NMAX]transaction

func main() {
	var tdatbar tabdatbar
	var ttrans tabTrans
	var n, m int
	var pilihan string

	for {
		welcome()
		fmt.Scan(&pilihan)
		fmt.Println(" ")
		if pilihan == "x" {
			menu()
			var choice int
			fmt.Scan(&choice)
			if choice == 1 {
				dataBarang(&tdatbar, &n)
			} else if choice == 2 {
				dataTransaksi(&ttrans, &m)
			} else if choice == 3 {
				pencarianDataBarang(&tdatbar, n)
			} else if choice == 4 {
				fmt.Println("Selesai")
				return
			} else {
				fmt.Println("Pilihan tidak valid, coba lagi.")
			}
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func welcome() {

	fmt.Println("=======================================================================================================================================================")
	fmt.Println(" 			   #####                                              ######                                     ")
	fmt.Println(" 			  #     # ###### #        ##   #    #   ##   #####    #     #   ##   #####   ##   #    #  ####   ")
	fmt.Println(" 			  #       #      #       #  #  ##  ##  #  #    #      #     #  #  #    #    #  #  ##   # #    #  ")
	fmt.Println(" 			   #####  #####  #      #    # # ## # #    #   #      #     # #    #   #   #    # # #  # #       ")
	fmt.Println(" 			        # #      #      ###### #    # ######   #      #     # ######   #   ###### #  # # #  ###  ")
	fmt.Println(" 			  #     # #      #      #    # #    # #    #   #      #     # #    #   #   #    # #   ## #    #  ")
	fmt.Println(" 			   #####  ###### ###### #    # #    # #    #   #      ######  #    #   #   #    # #    #  ####   ")

	fmt.Println("    #                                              ###                                                      ######                                     ")
	fmt.Println("   # #   #####  #      # #    #   ##    ####  #     #  #    # #    # ###### #    # #####  ####  #####  #    #     #   ##   #####    ##   #    #  ####  ")
	fmt.Println("  #   #  #    # #      # #   #   #  #  #      #     #  ##   # #    # #      ##   #   #   #    # #    # #    #     #  #  #  #    #  #  #  ##   # #    # ")
	fmt.Println(" #     # #    # #      # ####   #    #  ####  #     #  # #  # #    # #####  # #  #   #   #    # #    # #    ######  #    # #    # #    # # #  # #      ")
	fmt.Println(" ####### #####  #      # #  #   ######      # #     #  #  # # #    # #      #  # #   #   #    # #####  #    #     # ###### #####  ###### #  # # #  ### ")
	fmt.Println(" #     # #      #      # #   #  #    # #    # #     #  #   ##  #  #  #      #   ##   #   #    # #   #  #    #     # #    # #   #  #    # #   ## #    # ")
	fmt.Println(" #     # #      ###### # #    # #    #  ####  #    ### #    #   ##   ###### #    #   #    ####  #    # #    ######  #    # #    # #    # #    #  ####  ")
	fmt.Println("=======================================================================================================================================================")
	fmt.Println("Tekan 'x' untuk lanjut: ")
}
func menu() {
	fmt.Println("==============")
	fmt.Println("MENU INVENTORI")
	fmt.Println("==============")
	fmt.Println("1. DATA BARANG")
	fmt.Println("2. DATA TRANSAKSI")
	fmt.Println("3. PENCARIAN DATA BARANG")
	fmt.Println("4. SELESAI")
	fmt.Print("PILIH MENU DI ATAS: ")
}

func menu_dataBarang() {
	fmt.Println("==============")
	fmt.Println("MENU DATA BARANG")
	fmt.Println("==============")
	fmt.Println("1. TAMBAH DATA BARANG")
	fmt.Println("2. HAPUS DATA BARANG")
	fmt.Println("3. UBAH DATA BARANG")
	fmt.Println("4. TAMPILKAN DATA TERURUT")
	fmt.Println("5. KEMBALI")
	fmt.Print("PILIH MENU DI ATAS: ")
}

func menu_dataTransaksi() {
	fmt.Println("==============")
	fmt.Println("MENU DATA TRANSAKSI")
	fmt.Println("==============")
	fmt.Println("1. TAMBAH DATA TRANSAKSI")
	fmt.Println("2. TAMPILKAN SEMUA TRANSAKSI")
	fmt.Println("3. CARI TRANSAKSI BERDASARKAN TANGGAL")
	fmt.Println("4. KEMBALI")
	fmt.Print("PILIH MENU DI ATAS: ")
}

func menu_sort() {
	fmt.Println("==============")
	fmt.Println("MENU SORTING")
	fmt.Println("==============")
	fmt.Println("1. SORT BERDASARKAN KODE BARANG")
	fmt.Println("2. SORT BERDASARKAN STOK")
	fmt.Println("3. SORT BERDASARKAN HARGA")
	fmt.Println("4. KEMBALI")
	fmt.Print("PILIH MENU DI ATAS: ")
}

func menu_search() {
	fmt.Println("==============")
	fmt.Println("MENU PENCARIAN DATA BARANG")
	fmt.Println("==============")
	fmt.Println("1. PENCARIAN DATA BARANG")
	fmt.Println("2. KEMBALI")
	fmt.Print("PILIH MENU DI ATAS: ")
}

func menu_searchCat() {
	fmt.Println("==============")
	fmt.Println("MENU PENCARIAN DATA BARANG")
	fmt.Println("==============")
	fmt.Println("1. PENCARIAN BERDASARKAN NAMA BARANG")
	fmt.Println("2. PENCARIAN BERDASARKAN KODE BARANG")
	fmt.Println("3. KEMBALI")
	fmt.Print("PILIH MENU DI ATAS: ")
}

func menuHarga() {
	fmt.Println("==============")
	fmt.Println("MENU SORTING")
	fmt.Println("==============")
	fmt.Println("1. URUTKAN HARGA TERBESAR")
	fmt.Println("2. URUTKAN HARGA TERKECIL")
	fmt.Println("3. KEMBALI")
	fmt.Print("PILIH MENU DI ATAS: ")
}

// MENU DATA BARANG
func dataBarang(A *tabdatbar, n *int) {
	for {
		menu_dataBarang()
		var input int
		fmt.Scan(&input)
		if input == 1 {
			tambahDataBarang(A, n)
		} else if input == 2 {
			hapusDataBarang(A, n)
		} else if input == 3 {
			ubahDataBarang(A, n)
		} else if input == 4 {
			sortDataBarang(A, n)
		} else if input == 5 {
			return
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func tambahDataBarang(A *tabdatbar, n *int) {
	//I.S. : `A` adalah array dari struktur data `db` yang berisi data barang. `n` adalah jumlah barang dalam array `A`, bertipe integer.
	//F.S. : Data barang baru (kode barang, nama barang, stok, harga) ditambahkan ke dalam array `A` pada indeks ke-`n`.
	//Nilai `n` meningkat satu. Pesan "Data barang berhasil ditambahkan." ditampilkan.
	var kodeBarang, namaBarang string
	var stok, harga int

	fmt.Println("Masukkan kode barang:")
	fmt.Scan(&kodeBarang)
	fmt.Println("Masukkan nama barang:")
	fmt.Scan(&namaBarang)
	fmt.Println("Masukkan jumlah stok:")
	fmt.Scan(&stok)
	fmt.Println("Masukkan harga barang:")
	fmt.Scan(&harga)

	A[*n] = db{kodeBarang, namaBarang, stok, harga}
	*n++
	fmt.Println("Data barang berhasil ditambahkan.")
}

func hapusDataBarang(A *tabdatbar, n *int) {
	//I.S. : `A` adalah array dari struktur data `db` yang berisi data barang. `n` adalah jumlah barang dalam array `A`, bertipe integer. Data barang yang akan dihapus belum ditentukan.
	//F.S. : Data barang yang sesuai dengan kode barang yang diberikan telah dihapus dari array `A`. Nilai `n` berkurang satu jika data ditemukan dan dihapus.
	//Pesan "Data telah dihapus." atau "DATA TIDAK DITEMUKAN" ditampilkan.
	var kodeBarang string

	cetakBarang(*A, *n)
	fmt.Println("Masukkan Kode Barang yang ingin dihapus:")
	fmt.Scan(&kodeBarang)
	hapus(A, n, kodeBarang)
}

func hapus(A *tabdatbar, n *int, x string) {
	//I.S. : `A` adalah array dari struktur data `db` yang berisi data barang. `n` adalah jumlah barang dalam array `A`, bertipe integer. `x` adalah string yang berisi kode barang yang akan dihapus. Data barang yang sesuai dengan kode barang `x` belum dihapus.
	//F.S. : Jika data barang yang sesuai dengan kode barang `x` ditemukan di dalam array `A`, maka data tersebut telah dihapus dari array `A`, dan elemen-elemen berikutnya digeser untuk mengisi posisi yang kosong. Nilai `n` berkurang satu. Pesan "Data telah dihapus." ditampilkan.
	//Jika data barang yang sesuai dengan kode barang `x` tidak ditemukan, pesan "DATA TIDAK DITEMUKAN" ditampilkan dan tidak ada perubahan pada array `A` atau nilai `n`
	idx := findBarang(*A, *n, x)
	if idx == -1 {
		fmt.Println("DATA TIDAK DITEMUKAN")
	} else {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data telah dihapus.")
	}
}

func ubahDataBarang(A *tabdatbar, n *int) {
	//I.S. : 'A' adalah array dari struktur data db yang berisi data barang. n adalah jumlah barang dalam array A, bertipe integer. Kode barang yang akan diubah belum ditentukan.
	//F.S. : Kode barang yang akan diubah telah dimasukkan oleh pengguna. Fungsi edit telah dipanggil dengan A, n, dan kode barang yang diberikan sebagai argumen.
	var kodeBarang string

	cetakBarang(*A, *n)
	fmt.Println("Masukkan Kode Barang yang ingin diubah:")
	fmt.Scan(&kodeBarang)
	edit(A, *n, kodeBarang)
}

func findBarang(A tabdatbar, n int, x string) int {
	//I.S. : A adalah array dari struktur data db yang berisi data barang. n adalah jumlah barang dalam array A, bertipe integer. x adalah string yang berisi kode barang yang akan dicari.
	//F.S. : Jika data barang dengan kode x ditemukan, fungsi mengembalikan indeks barang dalam array A. Jika tidak ditemukan, fungsi mengembalikan nilai -1.
	for i := 0; i < n; i++ {
		if A[i].kodeBarang == x {
			return i
		}
	}
	return -1
}

func edit(A *tabdatbar, n int, x string) {
	//I.S. : A adalah array dari struktur data db yang berisi data barang. n adalah jumlah barang dalam array A, bertipe integer. x adalah string yang berisi kode barang yang akan diubah. Data barang yang sesuai dengan kode barang x belum diubah.
	//F.S. : Jika data barang dengan kode x ditemukan, data barang tersebut telah diubah sesuai input baru (nama barang, stok, harga). Pesan "Data sudah diedit." ditampilkan. Jika data barang dengan kode x tidak ditemukan, pesan "DATA TIDAK DITEMUKAN" ditampilkan.
	idx := findBarang(*A, n, x)
	if idx == -1 {
		fmt.Println("DATA TIDAK DITEMUKAN")
	} else {
		var namaBarang string
		var stok, harga int

		fmt.Println("Masukkan nama barang baru:")
		fmt.Scan(&namaBarang)
		fmt.Println("Masukkan jumlah stok baru:")
		fmt.Scan(&stok)
		fmt.Println("Masukkan harga barang baru:")
		fmt.Scan(&harga)

		A[idx] = db{x, namaBarang, stok, harga}
		fmt.Println("Data sudah diedit.")
	}
}

func cetakBarang(A tabdatbar, n int) {
	//I.S. : A` adalah array dari struktur data `db` yang berisi data barang. `n` adalah jumlah barang dalam array `A`, bertipe integer.
	//F.S. : Seluruh data barang dalam array `A` dicetak ke layar.
	fmt.Println("Data Barang:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", A[i].kodeBarang, A[i].namaBarang, A[i].stok, A[i].harga)
	}
}

func sortDataBarang(A *tabdatbar, n *int) {
	for {
		menu_sort()
		var input, pilih int
		fmt.Scan(&input)
		if input == 1 {
			selectionSort(A, *n, "kodeBarang")
			cetakBarang(*A, *n)
		} else if input == 2 {
			selectionSort(A, *n, "stok")
			cetakBarang(*A, *n)
		} else if input == 3 {
			cetakBarang(*A, *n)
			menuHarga()
			fmt.Scan(&pilih)
			if pilih == 1 {
				insortDesc(A, *n)
			} else if pilih == 2 {
				insortAsc(A, *n)
			} else if pilih == 3 {
				return
			} else {
				fmt.Println("Tidak Cocok")
			}
			cetakBarang(*A, *n)
		} else if input == 4 {
			return
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func insortAsc(A *tabdatbar, n int) {
	//I.S. : `A` adalah array dari struktur data `db` yang berisi data barang tidak terurut atau sebagian terurut. `n` adalah jumlah barang dalam array `A`, bertipe integer.
	//F.S. : Array `A` terurut secara ascending berdasarkan harga barang.
	for i := 1; i < n; i++ {
		temp := (*A)[i]
		j := i - 1
		for j >= 0 && (*A)[j].harga > temp.harga {
			(*A)[j+1] = (*A)[j]
			j = j - 1
		}
		(*A)[j+1] = temp
	}
}

func insortDesc(A *tabdatbar, n int) {
	//I.S. : `A` adalah array dari struktur data `db` yang berisi data barang tidak terurut atau sebagian terurut. `n` adalah jumlah barang dalam array `A`, bertipe integer.
	//F.S. : Array `A` terurut secara descending berdasarkan harga barang.
	for i := 1; i < n; i++ {
		temp := (*A)[i]
		j := i - 1
		for j >= 0 && (*A)[j].harga < temp.harga {
			(*A)[j+1] = (*A)[j]
			j = j - 1
		}
		(*A)[j+1] = temp
	}
}

func selectionSort(A *tabdatbar, n int, key string) {
	//I.S. : `A` adalah array dari struktur data `db` yang berisi data barang tidak terurut atau sebagian terurut. `n` adalah jumlah barang dalam array `A`, bertipe integer.
	//F.S. : Array `A` terurut secara descending berdasarkan harga barang.
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if key == "kodeBarang" && A[j].kodeBarang < A[minIdx].kodeBarang {
				minIdx = j
			} else if key == "stok" && A[j].stok < A[minIdx].stok {
				minIdx = j
			} else if key == "harga" && A[j].harga < A[minIdx].harga {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
}

// MENU DATA TRANSAKSI
func dataTransaksi(T *tabTrans, m *int) {
	//I.S. : T adalah array dari struktur data transaction yang berisi data transaksi. m adalah jumlah transaksi dalam array T, bertipe integer. Pengguna belum memilih operasi pada menu data transaksi.
	//F.S. : Pengguna telah memilih dan menjalankan operasi dari menu data transaksi. Data transaksi dapat bertambah, seluruh data transaksi dapat ditampilkan, atau data transaksi tertentu dapat dicari. Fungsi kembali ke menu utama jika pengguna memilih keluar.
	for {
		menu_dataTransaksi()
		var input int
		fmt.Scan(&input)
		if input == 1 {
			tambahDataTransaksi(T, m)
		} else if input == 2 {
			tampilkanSemuaTransaksi(T, *m)
		} else if input == 3 {
			cariTransaksiTanggal(T, *m)
		} else if input == 4 {
			return
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func tambahDataTransaksi(T *tabTrans, m *int) {
	//I.S : `T` adalah array dari struktur data `transaction` yang berisi data transaksi. `m` adalah jumlah transaksi dalam array `T`, bertipe integer
	//F.S : Data transaksi baru (tanggal, kode barang, jenis transaksi, jumlah) ditambahkan ke dalam array `T` pada indeks ke-`m`. Nilai `m` meningkat satu. Pesan "Data transaksi berhasil ditambahkan." ditampilkan.
	var tanggal, kodeBarang, jenis string
	var jumlah int

	fmt.Println("Masukkan tanggal transaksi (DD-MM-YYYY):")
	fmt.Scan(&tanggal)
	fmt.Println("Masukkan kode barang:")
	fmt.Scan(&kodeBarang)
	fmt.Println("Masukkan jenis transaksi (masuk/keluar):")
	fmt.Scan(&jenis)
	fmt.Println("Masukkan jumlah:")
	fmt.Scan(&jumlah)

	T[*m] = transaction{tanggal, kodeBarang, jenis, jumlah}
	*m++
	fmt.Println("Data transaksi berhasil ditambahkan.")
}

func tampilkanSemuaTransaksi(T *tabTrans, m int) {
	//I.S. : `T` adalah array dari struktur data `transaction` yang berisi data transaksi. `m` adalah jumlah transaksi dalam array `T`, bertipe integer.
	//F.S. : Seluruh data transaksi dalam array `T` dicetak ke layar.
	fmt.Println("Data Transaksi:")
	for i := 0; i < m; i++ {
		fmt.Printf("%s %s %s %d\n", T[i].tanggal, T[i].kodeBarang, T[i].jenis, T[i].jumlah)
	}
}

func cariTransaksiTanggal(T *tabTrans, m int) {
	//I.S. : `T` adalah array dari struktur data `transaction` yang berisi data transaksi. `m` adalah jumlah transaksi dalam array `T`, bertipe integer. Tanggal transaksi yang akan dicari belum ditentukan.
	//F.S. : Data transaksi yang sesuai dengan tanggal yang diberikan dicetak ke layar. Pesan "Transaksi tidak ditemukan pada tanggal tersebut." jika data tidak ditemukan.
	var tanggal string
	fmt.Println("Masukkan tanggal transaksi yang dicari (DD-MM-YYYY):")
	fmt.Scan(&tanggal)

	found := false
	for i := 0; i < m; i++ {
		if T[i].tanggal == tanggal {
			fmt.Printf("%s %s %s %d\n", T[i].tanggal, T[i].kodeBarang, T[i].jenis, T[i].jumlah)
			found = true
		}
	}
	if !found {
		fmt.Println("Transaksi tidak ditemukan pada tanggal tersebut.")
	}
}

// MENU PENCARIAN DATA BARANG
func pencarianDataBarang(A *tabdatbar, n int) {
	//I.S. : A adalah array dari struktur data db yang berisi data barang. n adalah jumlah barang dalam array A, bertipe integer. Pengguna belum memilih operasi pada menu pencarian data barang.
	//F.S. :Pengguna telah memilih dan menjalankan operasi dari menu pencarian data barang.
	//Pengguna dapat melakukan pencarian secara sequential atau binary pada data barang, atau kembali ke menu utama. Fungsi kembali ke menu utama jika pengguna memilih keluar. Pesan "Pilihan tidak valid, coba lagi." ditampilkan jika pengguna memasukkan pilihan yang tidak valid.
	for {
		menu_search()
		var input int
		fmt.Scan(&input)
		if input == 1 {
			menu_searchCat()
			var pilih int
			fmt.Scan(&pilih)
			if pilih == 1 {
				sequentialSearch(A, n)
			} else if pilih == 2 {
				binarySearch(A, n)
			} else if pilih == 3 {
				return
			} else {
				fmt.Println("Pilihan tidak valid, coba lagi.")
			}
		} else if input == 2 {
			return
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func sequentialSearch(A *tabdatbar, n int) {
	//I.S. :`A` adalah array dari struktur data `db` yang berisi data barang. `n` adalah jumlah barang dalam array `A`, bertipe integer. Nama barang yang akan dicari belum ditentukan.
	//F.S. : Data barang yang sesuai dengan nama barang yang diberikan dicetak ke layar. Pesan "Data ditemukan" atau "Data tidak ditemukan." ditampilkan.
	var namaBarang string
	fmt.Println("Masukkan nama barang yang dicari:")
	fmt.Scan(&namaBarang)

	found := false
	for i := 0; i < n; i++ {
		if A[i].namaBarang == namaBarang {
			fmt.Printf("Data ditemukan: %s %s %d %d\n", A[i].kodeBarang, A[i].namaBarang, A[i].stok, A[i].harga)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func binarySearch(A *tabdatbar, n int) {
	//I.S. : `A` adalah array dari struktur data `db` yang berisi data barang, terurut berdasarkan `kodeBarang`. `n` adalah jumlah barang dalam array `A`, bertipe integer. Kode barang yang akan dicari belum ditentukan.
	//F.S. : Data barang yang sesuai dengan kode barang yang diberikan dicetak ke layar. Pesan "Data ditemukan" atau "Data tidak ditemukan." ditampilkan.
	var kodeBarang string
	fmt.Println("Masukkan kode barang yang dicari:")
	fmt.Scan(&kodeBarang)

	left, right := 0, n-1
	found := false

	for left <= right && !found {
		mid := (left + right) / 2
		if A[mid].kodeBarang < kodeBarang {
			left = mid + 1
		} else if A[mid].kodeBarang > kodeBarang {
			right = mid - 1
		} else {
			fmt.Printf("Data ditemukan: %s %s %d %d\n", A[mid].kodeBarang, A[mid].namaBarang, A[mid].stok, A[mid].harga)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}
