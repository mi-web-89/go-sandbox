package main

import "math"

func hitungJumlahTanaman(ukuranTanah float64) float64 {
	jumlahTanaman := math.Ceil(math.Pow(ukuranTanah, 2) / 2)
	return jumlahTanaman
}
