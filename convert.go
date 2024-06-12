package humannumbers

import "fmt"

// NumToWords takes a number between 0 and 100,
// and returns the same in words.
func NumToWords(number int) (string, error) {
	if number < 0 || number > 100 {
		return "",
			fmt.Errorf("the number %v does not fall in the range 0-100", number)
	}

	if number == 100 {
		return "hundred", nil
	}

	unitsWord := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
		"eighteen", "nineteen"}

	if number < 20 {
		return unitsWord[number], nil
	}

	tensWord := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if number%10 == 0 {
		return tensWord[number/10], nil
	}

	return tensWord[number/10] + " " + unitsWord[number%10], nil

}
