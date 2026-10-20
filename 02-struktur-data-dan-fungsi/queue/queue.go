package queue

import (
	"fmt"
	"slices"
)

type Queue struct {
	items []int
}

// memasukkan nilai kedalam queue
func (q *Queue) Push(n int) {
	if slices.Contains(q.items, n) == false {
		q.items = append(q.items, n)
	}
}

// mengeluarkan nilai dari antrian
func (q *Queue) Pop() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	item := q.items[0]
	q.items = q.items[1:] // dari 1 sampai akhir

	return item, true
}

// panjang queue
func (q *Queue) LenQueue() {
	fmt.Println("length ", len(q.items))
}

// menampilkan queue
func (q *Queue) DisplayQueue() {
	fmt.Println("items : ", q.items)
}
