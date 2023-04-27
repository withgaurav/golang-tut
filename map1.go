package main

import (
	"fmt"
)

func main() {
	codes := map[string]int{"en": 1, "fr": 2, "hi": 3}
	value, found := codes["en"]
	fmt.Println(found, value)
	value, found = codes["hh"]
	fmt.Println(found, value)
}
