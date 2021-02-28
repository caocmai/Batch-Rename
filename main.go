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

func renameAndMoveFiles(fileType string, outputName string, newfileName string) {

	_, err := os.Stat(outputName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(outputName, 0755)
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

			os.Rename(file.Name(), strconv.Itoa(counter)+"_"+newfileName+fileType)
			oldFileDir := dir + "/" + strconv.Itoa(counter) + "_" + newfileName + fileType
			finalFileDir := dir + "/" + outputName + "/" + strconv.Itoa(counter) + "_" + newfileName + fileType
			os.Rename(oldFileDir, finalFileDir)

			counter++

		}
	}

}

func main() {

	var config struct { // [1]
		filetypeName     string
		outputFolderName string
		outputFileName   string
	}

	flag.StringVar(&config.filetypeName, "filetype", "0000000", "Enter filetype you want to rename")
	flag.StringVar(&config.outputFolderName, "outputFolderName", "test", "Enter file name to store renamed files")
	flag.StringVar(&config.outputFileName, "newFileName", "renamed_default", "Rename files as this")
	flag.Parse()

	fmt.Println(config.filetypeName)
	fmt.Println(config.outputFolderName)

	renameAndMoveFiles(config.filetypeName, config.outputFolderName, config.outputFileName)
}
