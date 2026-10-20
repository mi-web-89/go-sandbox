package struk

// Barang merepresentasikan item pembelian dalam keranjang.
type Barang struct {
	Nama        string
	HargaSatuan int
	Jumlah      int
}

// SubTotal menghitung total untuk satu item.
func SubTotal(item Barang) int {
	if item.HargaSatuan < 0 || item.Jumlah < 0 {
		return 0
	}
	return item.HargaSatuan * item.Jumlah
}

// HitungStruk menghitung total belanja dari semua item.
func HitungStruk(items []Barang) (total int) {
	for _, item := range items {
		total += SubTotal(item)
	}
	return total
}
