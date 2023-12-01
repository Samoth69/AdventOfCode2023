package day_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Run() {
	file, err := os.OpenFile("data/01.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCounter := 0
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		first, err := getFirstDigit(text, false)
		if err != nil {
			log.Fatalf("error on line %d : %s", lineCounter, err)
		}
		last, err := getFirstDigit(text, true)
		if err != nil {
			log.Fatalf("error on line %d : %s", lineCounter, err)
		}

		combined := first + last

		decodedNumber, err := strconv.Atoi(combined)
		if err != nil {
			log.Fatal(err)
		}
		sum += decodedNumber
		lineCounter++
	}
	log.Printf("%d", sum)
}

func getFirstDigit(text string, reverse bool) (string, error) {
	if reverse {
		for index := len(text) - 1; index >= 0; index-- {
			elem := int32(text[index])
			if unicode.IsNumber(elem) {
				return string(elem), nil
			}
		}
	} else {
		for _, elem := range text {
			if unicode.IsNumber(elem) {
				return string(elem), nil
			}
		}
	}
	return "", fmt.Errorf("no digit found")
}

var numbersInText = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"height",
	"nine",
	"ten",
}

func isDigitInText(text string, startIndex int) string {
	substring := text[startIndex:]
	for i := 0; i < len(numbersInText); i++ {
		currentTextNumber := numbersInText[i]
		if len(substring) < len(currentTextNumber) {
			continue
		}
		substringTrimmed := substring[:len(currentTextNumber)]
		if currentTextNumber == substringTrimmed {
			return strconv.Itoa(i + 1)
		}
	}
	return ""
}
