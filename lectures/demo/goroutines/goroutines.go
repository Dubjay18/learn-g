package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	for _, r := range data {
		go capIt(r)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("capitalized %s\n", string(capitalized))
}
