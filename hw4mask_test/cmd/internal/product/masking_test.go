package product_test

import (
	"fmt"
	"masking/hw4mask_test/cmd/internal/product"
)

func ExampleMask() {
	card := "8600123456789012"
	result := product.Mask(card)
	fmt.Println(result)

	// Output:
	// 8600 **** **** 9012
}

func ExampleMask_withSpaces() {
	card := "9860-1234-5678-4321"
	result := product.Mask(card)
	fmt.Println(result)

	// Output:
	// 9860 **** **** 4321
}

func ExampleMask_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	invalidCard := "12345"
	product.Mask(invalidCard)

	// Output:
	// card number must be 16 digits
}
