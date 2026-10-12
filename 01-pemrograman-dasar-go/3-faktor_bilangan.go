package main

func hitungFaktorBilangan(n int) int {
	var total int = 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			total += i
		}
	}
	return total
}
