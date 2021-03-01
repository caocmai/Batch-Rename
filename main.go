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

func createOutputFolder(outputFolderName string) {
	// Checks if output file name already exsits
	_, err := os.Stat(outputFolderName)

	// If not create a file
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(outputFolderName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

}

func getFilesFromDir(inputFolder string) (files []os.FileInfo) {
	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Read the current working directory + inputFolder name for files
	specifiedWorkingDir := workingDir + "/" + inputFolder
	files, err = ioutil.ReadDir(specifiedWorkingDir)
	if err != nil {
		log.Fatal(err)
	}

	isEmpty := false

	if len(files) == 0 {
		isEmpty = true
	}

	// If there are no files in folder then delete that folder
	if isEmpty {
		os.Remove(specifiedWorkingDir)
	}

	return files

}

func renameAndMoveFiles(fileType string, outputFolderName string, newfileName string, inputFolder string) {

	createOutputFolder(outputFolderName)

	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Read the current working directory + inputFolder name for files
	specifiedWorkingDir := workingDir + "/" + inputFolder
	files := getFilesFromDir(inputFolder)

	counter := 0

	// Loop through all files in specified folder
	for _, file := range files {
		if strings.Contains(file.Name(), fileType) {
			// fileNameWithoutExtension := strings.Split(file.Name(), fileType)[0]
			// fmt.Println(fileNameWithoutExtension)

			// Rename and move file
			finalFileDir := workingDir + "/" + outputFolderName + "/" + strconv.Itoa(counter) + "_" + newfileName + fileType

			err := os.Rename(filepath.Join(specifiedWorkingDir, file.Name()), finalFileDir)

			if err != nil {
				log.Fatal(err)
			}
			counter++
		}
	}

	// Logs out how many files were renamed
	if counter > 0 {
		fmt.Println(counter, fileType, "was renamed")
	}

	getFilesFromDir(inputFolder)

}

func main() {

	// config stuct to store flags
	var config struct {
		inputFolderName  string
		filetypeName     string
		outputFolderName string
		outputFileName   string
	}

	flag.StringVar(&config.filetypeName, "filetype", "0000000", "Enter filetype you want to rename")
	flag.StringVar(&config.outputFolderName, "output", "output_files", "Enter file name to store renamed files")
	flag.StringVar(&config.outputFileName, "renameAs", "renamed_file", "What to call the renamed files")
	flag.StringVar(&config.inputFolderName, "input", "", "Enter the folder of files to rename")
	flag.Parse()

	renameAndMoveFiles(config.filetypeName, config.outputFolderName, config.outputFileName, config.inputFolderName)
}
