package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("String Manipulation Program")
	fmt.Print("Enter a string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error: %s", err)
	}
	input = strings.TrimSpace(input)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Reverse the string")
		fmt.Println("2. Convert to uppercase")
		fmt.Println("3. Convert to lowercase")
		fmt.Println("4. Count vowels")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Reversed string:", reverseString(input))
		case 2:
			fmt.Println("Uppercase string:", strings.ToUpper(input))
		case 3:
			fmt.Println("Lowercase string:", strings.ToLower(input))
		case 4:
			fmt.Printf("Number of vowels: %d\n", countVowels(input))
		case 5:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}
