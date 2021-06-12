package main

import (
	"fmt"
	"strings"
)

func main() {

	var original string = "Madam"
        var reverse string = ""
	var length = len(original)

	for i := length - 1; i >= 0; i -- {
		reverse = reverse + string(original[i])
	}

	
	if strings.ToLower(original) == strings.ToLower(reverse) {
		fmt.Println("The given string",original,"is Palindrome");
	} else {
		fmt.Println("The given string", original, "is NOT a Palindrome");
	}
}