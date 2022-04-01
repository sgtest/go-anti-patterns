package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world!")
}

func outdatedReplace(s string) string {
	return strings.ReplaceAll(s, "\\", "/")
}

func outdatedReplace2(s string) string {
	return strings.ReplaceAll(s, "\\", "/")
}

func outdatedReplace3(s string) string {
	return strings.ReplaceAll(s, "\\", "/")
}

func stringsContains(s, sub string) bool {
	return strings.Index(s, sub) != -1
}

func unusedRangeValue(s []string) {
	for i, _ := range s {
		fmt.Println(s[i])
	}
}

func redundantTypeDeclaration(a string, b string) string {
	return a + b
}

func expensiveStringComparasion(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}
