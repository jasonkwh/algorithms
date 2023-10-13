package others

import (
	"strconv"
)

func FindLowestPrice(products [][]string, discounts [][]string) int32 {
	if len(products) == 0 {
		return 0
	}

	// use a map for easier discount get
	discountsMap, type0sales := generateDiscountMap(discounts)

	// get lowest product price
	lowestPrice := 0.0
	if type0sales {
		lowestPrice = getLowestProductPrice(products)
	}

	return getTotalProductPrice(lowestPrice, products, discountsMap)
}

func getTotalProductPrice(lowestPrice float64, products [][]string, discountsMap map[string][]string) int32 {
	output := 0.0

	for _, product := range products {
		price, _ := strconv.ParseFloat(product[0], 32)
		if len(product) > 1 {
			// limit to 1000
			if len(product) > 1000 {
				product = product[:1000]
			}
			newLowestPrice := price
			for i := 1; i < len(product); i++ {
				tempPrice := price

				// if it can find a discount from the map...
				if discountsMap[product[i]] != nil {
					tempPrice = calculateDiscount(price, lowestPrice, discountsMap[product[i]][0], discountsMap[product[i]][1])
				}
				if tempPrice < newLowestPrice {
					newLowestPrice = tempPrice
				}
			}
			price = newLowestPrice
		}
		output += price
	}

	return int32(output)
}

func generateDiscountMap(discounts [][]string) (map[string][]string, bool) {
	discountsMap := make(map[string][]string, len(discounts))
	type0sales := false

	// limit to 1000 discounts
	if len(discounts) > 1000 {
		discounts = discounts[:1000]
	}
	for i := 0; i < len(discounts); i++ {
		if len(discounts[i]) != 3 {
			continue
		}
		if discounts[i][1] == "0" {
			type0sales = true
		}
		discountsMap[discounts[i][0]] = []string{discounts[i][1], discounts[i][2]}
	}

	return discountsMap, type0sales
}

func getLowestProductPrice(products [][]string) float64 {
	lowestPrice := 0.0

	for _, product := range products {
		price, _ := strconv.ParseFloat(product[0], 32)

		if lowestPrice == 0 || price < lowestPrice {
			lowestPrice = price
		}
	}

	return lowestPrice
}

func calculateDiscount(price, lowestPrice float64, discountType, discountPc string) float64 {
	switch discountType {
	case "0":
		return lowestPrice
	case "1":
		discount, _ := strconv.ParseFloat(discountPc, 32)
		return price - (price * discount / 100)
	case "2":
		discount, _ := strconv.ParseFloat(discountPc, 32)
		return price - discount
	default:
		return price
	}
}
