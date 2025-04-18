package main

import (
	"fmt"

	"github.com/JHONYKK22/morse-code-mod/morse"
)

func main() {

	text := "ABC abc aa"

	// Fail test
	// text += ";"

	encodeText, err := morse.Encode(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encodeText)

	// Fail test
	// encodeText += " ......"

	decodeText, err := morse.Decode(encodeText)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(decodeText)

}
