package product

import "strings"

func Mask(cardNum string) string {

	clean := strings.ReplaceAll(cardNum, " ", "")
	clean = strings.ReplaceAll(clean, "-", "")

	if len(clean) != 16 {
		panic("card number must be 16 digits")
	}

	var masked string
	for i := 0; i < len(clean); i++ {
		if i > 0 && i%4 == 0 {
			masked += " "
		}
		if i < 4 || i >= 12 {
			masked += string(clean[i])
		} else {
			masked += "*"
		}
	}
	return masked
}
