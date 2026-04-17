package main

import (
	"bufio"
	"fmt"
	"masking/hw4mask_test/cmd/internal/product"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите номер карты: ")
	scanner.Scan()
	input := scanner.Text()

	clean := strings.ReplaceAll(input, " ", "")
	clean = strings.ReplaceAll(clean, "-", "")
	
	masked := product.Mask(clean)

	fmt.Println("Карта:", masked)
}
