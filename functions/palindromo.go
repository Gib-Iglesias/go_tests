package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)             // Ignorar mayúsculas/minúsculas
	s = strings.ReplaceAll(s, " ", "") // Ignorar espacios
	//TODO: Ignorar caracteres especiales
	s = strings.ReplaceAll(s, "?", "") // Ignora solo signos de interrogación
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("radar"))
	fmt.Println(isPalindrome("A man a plan a canal Panama"))
	fmt.Println(isPalindrome("hello"))
	fmt.Println(isPalindrome("Anita lava la tina"))
	fmt.Println(isPalindrome("Was it a car or a cat I saw?"))
	fmt.Println(isPalindrome("No lemon no melon"))
}
