package main
/*
Author: Evan Gardner
Version: 1.0
*/

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	)
var palindrome string

func checkPalindrome() bool {
	s := strings.Split(palindrome, "")
	for x := 0; x < len(s)/2; x++ {
		if !(s[x] == s[len(s) - x - 1]) {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("Welcome to Palindrome Tester! Enter a potential palindrome to see if it truly is one. \n")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("----------------\nEnter your potential palindrome: ")
		scanner.Scan()
		palindrome = scanner.Text()
		palindrome = strings.Replace(palindrome, " ", "", -1)
		if checkPalindrome() {
			fmt.Println("That is a palindrome\n----------------\n")
		} else {
			fmt.Println("That's not a palindrome\n-------------\n")
		}
	}

}
