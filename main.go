package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func renameAndMoveFiles(OUTPUTNAME string, fileType string) {

	_, err := os.Stat(OUTPUTNAME)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(OUTPUTNAME, 0755)
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
		if strings.Contains(file.Name(), fileType) {
			fileNameWithoutExtension := strings.Split(file.Name(), fileType)[0]
			fmt.Println(fileNameWithoutExtension)

			os.Rename(file.Name(), strconv.Itoa(counter)+"_new"+fileType)
			oldDir := dir + "/" + strconv.Itoa(counter) + "_new" + fileType
			finalDir := dir + "/" + OUTPUTNAME + "/" + strconv.Itoa(counter) + "_new" + fileType
			os.Rename(oldDir, finalDir)

			counter++

		}
	}

}

func main() {
	OUTPUTNAME := "test"
	var typeFlag string
	flag.StringVar(&typeFlag, "filetype", "00000", "Enter type of file do you want to rename")
	flag.Parse()

	renameAndMoveFiles(OUTPUTNAME, typeFlag)
}
