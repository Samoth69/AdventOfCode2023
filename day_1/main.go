package day_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	file, err := os.OpenFile("data/01.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lineCounter := 0
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

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
