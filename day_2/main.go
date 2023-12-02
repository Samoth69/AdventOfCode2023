package day_2

import (
	"AdventOfCode2023/utils"
	"log"
	"strconv"
	"strings"
)

func Run() {
	c := make(chan string, 10)
	go func() {
		err := utils.ReadFileLineByLine("data/02.txt", c)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	sum := 0
	for text := range c {
		ret, err := decodeLinePart2(text)
		if err != nil {
			log.Fatalln(err)
		}
		if ret > 0 {
			sum += ret
		}
	}
	log.Printf("%d", sum)
}

// Lit la ligne et renvoie le game id si celle-ci est valide
// renvoie game id si ok
// renvoie -1 si game impossible
func processLinePart1(line string) int {
	gameId, err := strconv.Atoi(line[5:strings.Index(line, ":")])
	if err != nil {
		panic(err)
	}
	//log.Printf("--- %d ---", gameId)

	ret, err := decodeLinePart1(line)
	if err != nil {
		panic(err)
	}
	if ret == 0 {
		return gameId
	}

	return -1
}

// gère les différentes parties présente sur une ligne
// return 0 si la game est possible
// return 1 si la game est impossible
func decodeLinePart1(line string) (int, error) {
	arr := strings.Split(line[strings.Index(line, ":")+1:], ";")
	for _, party := range arr {
		ret, err := decodePartyPart1(party)
		if err != nil {
			return 1, nil
		}
		if ret == 1 {
			return 1, nil
		}
	}
	return 0, nil
}

// Renvoie vrai si la partie est possible avec 12 cubes rouges, 13 cubes verts et 14 cubes bleus
// return 0 si la partie est possible
// return 1 si elle n'est PAS possible
// return non défini en cas d'erreur (l'objet error contiendra des détails)
func decodePartyPart1(text string) (int, error) {
	for _, color := range strings.Split(text, ",") {
		color = strings.TrimSpace(color)
		if color == "" {
			continue
		}
		splitted := strings.Split(color, " ")
		num, err := strconv.Atoi(splitted[0])
		if err != nil {
			return 2, err
		}
		color = splitted[1]
		switch color {
		case "red":
			if num > 12 {
				return 1, nil
			}
		case "green":
			if num > 13 {
				return 1, nil
			}
		case "blue":
			if num > 14 {
				return 1, nil
			}
		}
	}
	return 0, nil
}

func decodeLinePart2(line string) (int, error) {
	cutted := strings.ReplaceAll(line[strings.Index(line, ":")+1:], ";", ",")

	ret, err := decodePartyPart2(cutted)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func decodePartyPart2(text string) (int, error) {
	minRed, minGreen, minBlue := 0, 0, 0
	for _, color := range strings.Split(text, ",") {
		color = strings.TrimSpace(color)
		if color == "" {
			continue
		}
		splitted := strings.Split(color, " ")
		num, err := strconv.Atoi(splitted[0])
		if err != nil {
			return 2, err
		}
		color = splitted[1]
		switch color {
		case "red":
			if num > minRed {
				minRed = num
			}
		case "green":
			if num > minGreen {
				minGreen = num
			}
		case "blue":
			if num > minBlue {
				minBlue = num
			}
		}
	}
	return minRed * minGreen * minBlue, nil
}
