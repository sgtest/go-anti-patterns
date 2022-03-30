package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world!")
}

func outdatedReplace(s string) string {
	return strings.Replace(s, "\\", "/", -1)
}

func unusedRangeValue(s []string) {
	for i, _ := range s {
		fmt.Println(s[i])
	}
}
