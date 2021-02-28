package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	directory := "."
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			fileNameWithoutExtension := strings.Split(file.Name(), ".txt")[0]
			fmt.Println(fileNameWithoutExtension)
			os.Rename(file.Name(), "new.txt")

		}
	}
}
