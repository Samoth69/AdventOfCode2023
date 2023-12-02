package day_1

import (
	"AdventOfCode2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	lineCounter := 0
	sum := 0

	c := make(chan string, 10)
	go func() {
		err := utils.ReadFileLineByLine("data/01.txt", c)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	for text := range c {
		ret, err := decodeNumber(text)
		if err != nil {
			log.Fatalln(err)
		}

		sum += ret
		lineCounter++
	}
	log.Printf("%d", sum)
}

func decodeNumber(text string) (int, error) {
	first, err := getFirstDigit(text, false)
	if err != nil {
		return 0, err
	}
	last, err := getFirstDigit(text, true)
	if err != nil {
		return 0, err
	}

	combined := first + last

	decodedNumber, err := strconv.Atoi(combined)
	if err != nil {
		return 0, err
	}
	return decodedNumber, nil
}

func getFirstDigit(text string, reverse bool) (string, error) {
	if reverse {
		for index := len(text) - 1; index >= 0; index-- {
			elem := int32(text[index])
			if unicode.IsNumber(elem) {
				return string(elem), nil
			} else {
				ret := isDigitInText(text, index)
				if ret != "" {
					return ret, nil
				}
			}
		}
	} else {
		for index, elem := range text {
			if unicode.IsNumber(elem) {
				return string(elem), nil
			} else {
				ret := isDigitInText(text, index)
				if ret != "" {
					return ret, nil
				}
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
	"eight",
	"nine",
	"ten",
}

func isDigitInText(text string, startIndex int) string {
	substring := text[startIndex:]
	if len(substring) > 2 {
		for i := 0; i < len(numbersInText); i++ {
			currentTextNumber := numbersInText[i]
			// fast skipping
			if currentTextNumber[0] != substring[0] {
				continue
			}
			if strings.HasPrefix(substring, currentTextNumber) {
				return strconv.Itoa(i + 1)
			}
		}
	}
	return ""
}
