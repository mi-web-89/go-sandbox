package main

import "math"

func hitungPythagoras(a, b, c float64) string {
	if math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2) {
		return "Phytagoras"
	}

	return "Not Phytagoras"
}
