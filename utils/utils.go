package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadFileLineByLine ouvre le fichier passé en paramètre et renvoie toutes ses lignes dans la channel.
// La channel est fermé quand la lecture du fichier est fini
func ReadFileLineByLine(path string, c chan string) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	close(c)
	return nil
}
