package main

import (
	"fmt"

	"struktur-data-dan-fungsi/input"
	"struktur-data-dan-fungsi/operator"
	"struktur-data-dan-fungsi/queue"
	"struktur-data-dan-fungsi/struk"
)

func main() {
	fmt.Println("============== 01. Struk Pembelanjaan ==============")
	jumlahBarang, err := input.InputInt("input jumlah barang : ")
	if err != nil {
		fmt.Println("Error: jumlah barang harus berupa angka")
		return
	}
	if jumlahBarang < 0 {
		fmt.Println("Error: jumlah barang tidak boleh negatif")
		return
	}

	fmt.Println("--------------------------")

	keranjang := make([]struk.Barang, 0, jumlahBarang) // slice of struct
	for i := 0; i < jumlahBarang; i++ {
		nama, err := input.InputString("input nama barang : ")
		if err != nil {
			fmt.Println("Error: nama barang tidak valid")
			return
		}

		harga, err := input.InputInt("input harga barang : ")
		if err != nil {
			fmt.Println("Error: harga barang harus berupa angka")
			return
		}
		if harga < 0 {
			fmt.Println("Error: harga barang tidak boleh negatif")
			return
		}

		jumlah, err := input.InputInt("input jumlah barang : ")
		if err != nil {
			fmt.Println("Error: jumlah barang harus berupa angka")
			return
		}
		if jumlah < 0 {
			fmt.Println("Error: jumlah barang tidak boleh negatif")
			return
		}

		keranjang = append(keranjang, struk.Barang{
			Nama:        nama,
			HargaSatuan: harga,
			Jumlah:      jumlah,
		})

		fmt.Println("-------------------------")
	}

	fmt.Println("======== struk anda ========")
	totalBelanja := struk.HitungStruk(keranjang)

	for _, item := range keranjang {
		fmt.Printf("%-15s %dx %d = %d\n", item.Nama, item.Jumlah, item.HargaSatuan, struk.SubTotal(item))
	}

	fmt.Println("---------------------------")
	fmt.Println("total belanja : ", totalBelanja)

	fmt.Println("============== 02. check duplicates ==============")
	items := []string{"b", "a", "e", "b", "c", "a"}
	fmt.Println(operator.CheckDuplicates(items))
	items = []string{"kiwi", "kiwi"}
	fmt.Println(operator.CheckDuplicates(items))
	items = []string{"x", "y", "z", "x", "y", "z", "x", "y", "z"}
	fmt.Println(operator.CheckDuplicates(items))

	fmt.Println("============== 03. check intersect ==============")
	items1 := []string{"apple", "mango", "banana", "guava"}
	items2 := []string{"apple", "biscuit", "cereal", "guava"}
	fmt.Println(operator.CheckIntersect(items1, items2)) // [apple guava]

	fmt.Println(operator.CheckIntersect([]string{"apple", "banana"}, []string{}))                  // []
	fmt.Println(operator.CheckIntersect([]string{"manggo", "orang"}, []string{"apple", "banana"})) // []

	items1 = []string{"apple", "manggo", "banana"}
	items2 = []string{"banana", "manggo"}
	fmt.Println(operator.CheckIntersect(items1, items2)) // [banana manggo]

	items1 = []string{"apple", "apple", "banana"}
	items2 = []string{"apple", "orange"}

	fmt.Println(operator.CheckIntersect(items1, items2)) // [apple]

	fmt.Println("============== 04. frequency calculator ==============")
	fmt.Println(operator.CalculateFreqs("i love vibe coding"))
	fmt.Println(operator.CalculateFreqs("go go go"))
	fmt.Println(operator.CalculateFreqs(""))

	fmt.Println("============== 05. queue ==============")
	var q1, q2 queue.Queue

	q1.Push(1)
	q1.Push(2)
	q1.Push(3)
	q1.Pop()
	q1.Push(4)

	q1.LenQueue()
	q1.DisplayQueue()

	fmt.Println("-------------------------------------")

	q2.Push(1)
	q2.Push(2)
	q2.Push(3)
	q2.Push(4)
	q2.Push(4)
	q2.Push(1)
	q2.Pop()
	q2.Pop()
	q2.Pop()
	q2.Push(10)

	q2.LenQueue()
	q2.DisplayQueue()
}
