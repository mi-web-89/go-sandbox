package main

import (
	"fmt"
	"pemrograman-go-lanjutan/product"
	"strings"
)

func main() {

	fmt.Println("==============3. Purchasable Products ===============")
	fmt.Println("Result: ", product.GetPurchasableProducts([]product.Product{
		{Name: "milk", Price: 9000},
		{Name: "apple", Price: 5000},
		{Name: "cake", Price: 15000},
		{Name: "candybar", Price: 2000},
		{Name: "ice cream", Price: 10000},
	}, 20000)) // [{candybar 2000},{apple 5000},{milk 9000}]
	fmt.Println(strings.Repeat("-", 20))

	fmt.Println("Result: ", product.GetPurchasableProducts([]product.Product{
		{Name: "candy", Price: 100},
		{Name: "snack", Price: 200},
		{Name: "coffee", Price: 300},
	}, 500)) // [{candy 100} {snack 200}]
	fmt.Println(strings.Repeat("-", 20))

	fmt.Println("Result: ", product.GetPurchasableProducts([]product.Product{
		{Name: "milk", Price: 10000},
		{Name: "mango", Price: 20000},
	}, 5000)) // []
	fmt.Println(strings.Repeat("-", 20))

	fmt.Println("Result: ", product.GetPurchasableProducts([]product.Product{
		{Name: "rice bowl", Price: 1000},
		{Name: "cupcake", Price: 2000},
	}, 1000)) // [{rice bowl 1000}]
	fmt.Println(strings.Repeat("-", 20))

	fmt.Println("Result: ", product.GetPurchasableProducts([]product.Product{
		{Name: "tea", Price: 500},
		{Name: "chai", Price: 500},
		{Name: "milk", Price: 1000},
	}, 1000)) // [{tea 500},{chai 500}]
	fmt.Println(strings.Repeat("-", 20))
}
