package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	directory := "."
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0

	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			fileNameWithoutExtension := strings.Split(file.Name(), ".txt")[0]
			fmt.Println(fileNameWithoutExtension)

			os.Rename(file.Name(), strconv.Itoa(counter)+"_new.txt")

			oldDir := dir + "/" + strconv.Itoa(counter) + "_new.txt"
			finalDir := dir + "/test/" + strconv.Itoa(counter) + "_new.txt"
			os.Rename(oldDir, finalDir)

			counter++

		}
	}
}
