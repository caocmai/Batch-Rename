package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func renameAndMoveFiles(fileType string, outputFileName string, newfileName string, inputFolder string) {

	// Checking if output file name already exsits
	_, err := os.Stat(outputFileName)

	// If not create a file
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(outputFileName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	// Get the working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Read the current directory
	workingDir := dir + "/" + inputFolder
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0

	// Loop through all files in current working directory
	// TODO: Maybe change this to a specfied folder within the working dir to keep it organzied
	for _, file := range files {
		// fmt.Println("in for loop")
		if strings.Contains(file.Name(), fileType) {
			// fileNameWithoutExtension := strings.Split(file.Name(), fileType)[0]
			// fmt.Println(fileNameWithoutExtension)

			// Rename and move file
			finalFileDir := dir + "/" + outputFileName + "/" + strconv.Itoa(counter) + "_" + newfileName + fileType

			err := os.Rename(filepath.Join(workingDir, file.Name()), finalFileDir)

			if err != nil {
				log.Fatal(err)
			}
			counter++
		}
	}

	fmt.Println(counter, fileType, "was renamed")

}

func main() {

	var config struct {
		inputFolderName  string
		filetypeName     string
		outputFolderName string
		outputFileName   string
	}

	flag.StringVar(&config.filetypeName, "filetype", "0000000", "Enter filetype you want to rename")
	flag.StringVar(&config.outputFolderName, "outputFolderName", "output_files", "Enter file name to store renamed files")
	flag.StringVar(&config.outputFileName, "newFileName", "renamed_file", "What to call the renamed files")
	flag.StringVar(&config.inputFolderName, "inputFolder", "", "Enter folder of files to rename")
	flag.Parse()

	renameAndMoveFiles(config.filetypeName, config.outputFolderName, config.outputFileName, config.inputFolderName)
}
