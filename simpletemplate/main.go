package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Chdir(os.Getenv("HOME"))
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("letter.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(createLetter("Jan")))
	if err != nil {
		log.Fatalln()
	}

}

func createLetter(str string) string {
	return fmt.Sprintf(
		`Dear %s,
How are you doing?
	
Kind Regards`+"\n", str)
}
