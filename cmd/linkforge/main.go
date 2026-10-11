package main

import (
	"fmt"
	"linkforge/internal/base62"
)

func main() {

	fmt.Println(base62.Encode(123456))
	fmt.Println(base62.EncodeWidth(123456, 3))

}
