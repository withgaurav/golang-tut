package main

import (
	"fmt"
)

func main() {
	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes)

	fmt.Println(len(codes))
	fmt.Println(codes["en"])
	fmt.Println(codes["fr"])
}
