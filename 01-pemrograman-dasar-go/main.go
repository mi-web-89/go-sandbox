package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("======= No.1 Phytagoras =======")
	var a, b, c float64

	fmt.Print("input nilai a : ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Error: nilai a harus berupa angka")
		return
	}

	fmt.Print("input nilai b : ")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Error: nilai b harus berupa angka")
		return
	}

	fmt.Print("input nilai c : ")
	if _, err := fmt.Scanln(&c); err != nil {
		fmt.Println("Error: nilai c harus berupa angka")
		return
	}

	fmt.Println(hitungPythagoras(a, b, c))

	fmt.Println("========= No. 2 Nilai Evaluasi Siswa ========")
	var evaluasiHarian, evaluasiPekanan, evaluasiAkhir int

	fmt.Print("input nilai evaluasi harian : ")
	if _, err := fmt.Scanln(&evaluasiHarian); err != nil {
		fmt.Println("Error: nilai evaluasi harian harus berupa angka")
		return
	}
	if evaluasiHarian < 0 {
		fmt.Println("Error: nilai evaluasi harian harus positif")
		return
	}

	fmt.Print("input nilai evaluasi pekananan : ")
	if _, err := fmt.Scanln(&evaluasiPekanan); err != nil {
		fmt.Println("Error: nilai evaluasi pekanan harus berupa angka")
		return
	}
	if evaluasiPekanan < 0 {
		fmt.Println("Error: nilai evaluasi pekanan harus positif")
		return
	}

	fmt.Print("input nilai evaluasi akhir : ")
	if _, err := fmt.Scanln(&evaluasiAkhir); err != nil {
		fmt.Println("Error: nilai evaluasi akhir harus berupa angka")
		return
	}
	if evaluasiAkhir < 0 {
		fmt.Println("Error: nilai evaluasi akhir harus positif")
		return
	}

	hitungNilaiEvaluasi(evaluasiHarian, evaluasiPekanan, evaluasiAkhir)

	fmt.Println("======== No.3 Faktor Bilangan ========")
	fmt.Print("input nilai faktor bilangan : ")
	var n int
	fmt.Scanln(&n)

	hasilFaktorBilangan := hitungFaktorBilangan(n)
	fmt.Println("nilai faktor bilangan ", hasilFaktorBilangan)

	fmt.Println("======= No.4 Hitung Jumlah Tanaman =======")
	var tanah float64
	fmt.Print("input ukuran tanah : ")
	fmt.Scanln(&tanah)

	totalTanaman := hitungJumlahTanaman(tanah)
	fmt.Printf("Jumlah Tanaman : %.0f\n", totalTanaman)

	fmt.Println("======= No.5 Pencarian string =======")
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("input kalimat : ")
	kalimat, _ := reader.ReadString('\n')
	kalimat = strings.TrimSpace(kalimat)

	fmt.Printf("input keyword : ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)

	if kalimat == "" || keyword == "" {
		fmt.Println("Input invalid, harap isi kalimat dan kata kunci")
		return
	}

	searchString(kalimat, keyword)

}
