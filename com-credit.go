package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func validateCreditCard(cardNumber string) string {
	// Define the regex pattern for valid credit card numbers
	pattern := `^[456]([\d]{15}|[\d]{3}(-[\d]{4}){3})$`
	re := regexp.MustCompile(pattern)

	// Check if the card number matches the pattern
	if re.MatchString(cardNumber) {
		// Remove hyphens and check for consecutive repeated digits
		cardNumber = strings.ReplaceAll(cardNumber, "-", "")
		if match, _ := regexp.MatchString(`(\d)\1{3,}`, cardNumber); match {
			return "Invalid"
		} else {
			return "Valid"
		}
	} else {
		return "Invalid"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the number of credit card numbers
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Read and validate each credit card number
	for i := 0; i < n; i++ {
		scanner.Scan()
		cardNumber := scanner.Text()
		fmt.Println(validateCreditCard(cardNumber))
	}
}
