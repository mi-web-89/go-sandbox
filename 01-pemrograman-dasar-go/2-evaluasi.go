package main

import "fmt"

func hitungNilaiEvaluasi(nilaiHarian, nilaiPekanan, nilaiAkhir int) {

	hasilAkhir := (nilaiHarian * 20 / 100) + (nilaiPekanan * 25 / 100) + (nilaiAkhir * 55 / 100)

	switch {
	case hasilAkhir < 90 && hasilAkhir >= 100:
		fmt.Println("nilai anda sangat baik ")
	case hasilAkhir > 79 && hasilAkhir < 91:
		fmt.Println("nilai anda baik ")
	case hasilAkhir > 65 && hasilAkhir < 80:
		fmt.Println("nilai anda cukup ")
	case hasilAkhir > 50 && hasilAkhir < 66:
		fmt.Println("nilai anda kurang")
	default:
		fmt.Println("nilai anda sangat kurang")
	}

	fmt.Println("nilai anda : ", hasilAkhir)
}
