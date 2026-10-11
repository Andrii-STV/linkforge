package base62

import "slices"

var code string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func encodeLogic(v uint64, width *int) string {
	division := v
	var remainder uint64
	var encodedString string

	for {
		remainder = division % 62
		division = division / 62
		encodedString = encodedString + code[remainder:remainder+1]
		if division == 0 {
			break
		}
	}

	encodedStringReverted := []byte(encodedString)
	slices.Reverse(encodedStringReverted)
	return string(encodedStringReverted)
}

func Encode(v uint64) string {

	return encodeLogic(v, nil)

}

func EncodeWidth(v uint64, width int) string {
	encoded := []byte(encodeLogic(v, &width))

	if length := len(encoded); width > length {
		for i := 0; i < (width - length); i++ {
			encoded = slices.Insert(encoded, 0, '0')
		}
	}

	return string(encoded)

}
