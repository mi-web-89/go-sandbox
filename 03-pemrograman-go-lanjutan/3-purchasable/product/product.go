package product

import (
	"cmp"
	"slices"
)

type Product struct {
	Name  string
	Price int
}

func GetPurchasableProducts(products []Product, budget int) []Product {
	if budget < 0 {
		return []Product{}
	}

	tmpProducts := slices.Clone(products)
	slices.SortFunc(tmpProducts, func(p1, p2 Product) int {
		return cmp.Compare(p1.Price, p2.Price)
	})

	var result []Product
	var remainingBudget = budget

	for _, p := range tmpProducts {
		if remainingBudget < p.Price {
			break
		}
		result = append(result, p)
		remainingBudget -= p.Price
	}

	return result
}
