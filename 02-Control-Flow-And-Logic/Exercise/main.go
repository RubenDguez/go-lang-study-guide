package main

import (
	"fmt"
	"strings"
)

const SALE = 0.25
const SALE_TAG = "**SALE**"
const SALE_TAX = 0.07

var products = map[string]float32{
	"T-SHIRT": 9.99,
	"MUG":     12.99,
	"HAT":     18.99,
	"BOOK":    25.99,
}

var shoppingCart = map[string]int{
	"MUG":              1,
	"**SALE** MUG":     3,
	"HAT":              2,
	"**SALE** HAT":     4,
	"T-SHIRT":          0,
	"**SALE** T-SHIRT": 3,
	"BOOK":             1,
	"**SALE** BOOK":    0,
	"SHORTS":           10,
	"**SALE** SHORTS":  10,
}

func calculateItemPrice(itemCode string) (price float32, ok bool) {
	if price, ok = products[itemCode]; ok {
		return
	}

	after, found := strings.CutPrefix(itemCode, SALE_TAG)
	if !found {
		return
	}

	price, ok = products[strings.Trim(after, " ")]
	if !ok {
		return
	}

	price = (price - (price * SALE))
	return
}

func printShoppingCart(sc map[string]int) {
	fmt.Println()
	for item, quantity := range sc {
		if quantity <= 0 {
			continue
		}

		if price, exists := calculateItemPrice(item); exists {
			var total float32 = price * float32(quantity)
			fmt.Printf("|%4.d | %-30s | %10.2f | %10.2f |\n", quantity, item, price, total)
		}
	}
}

func calculateShoppingCartTotal(sc map[string]int) (total float32) {
	for item, quantity := range sc {
		if quantity <= 0 {
			continue
		}

		if price, exists := calculateItemPrice(item); exists {
			total += price * float32(quantity)
		}
	}
	
	return
}

func calculateSaleTax(sale float32) (tax, total float32) {
	tax = sale * SALE_TAX
	total = sale + tax
	return
}

func main() {
	fmt.Println("\nSales Order Processor")
	
	printShoppingCart(shoppingCart)
	
	cartTotal := calculateShoppingCartTotal(shoppingCart)
	tax, total := calculateSaleTax(cartTotal)

	fmt.Println()
	fmt.Printf("Sub total......: $%10.2f\n", cartTotal)
	fmt.Printf("Sale tax.......: $%10.2f\n", tax)
	fmt.Printf("Gran total.....: $%10.2f\n", total)
	fmt.Println()
}
