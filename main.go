package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"os"
	"path/filepath"
	"strconv"
	"strings"

	// Using this 3rd party package to log out errors and save it to ErrorLog.log file
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func getWorkingDir() (dir string) {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		logError()
	}
	return workingDir
}

func createFolder(folderName string) {
	errDir := os.MkdirAll(folderName, 0755)
	if errDir != nil {
		log.Fatal(errDir)
		logError()
	}
}

func numFilesInFolder(outputFolderName string) (fileCount int) {
	// Checks if output file name already exsits
	_, err := os.Stat(outputFolderName)

	// If not create a file
	if os.IsNotExist(err) {
		createFolder(outputFolderName)
	} else {
		workingDir := getWorkingDir()
		specifiedWorkingDir := workingDir + "/" + outputFolderName
		files := getFilesFromDir(specifiedWorkingDir)
		return len(files)
	}
	// Or return 0, meaning folder is empty
	return 0
}

func deleteEmptyFolder(dir string, files []os.FileInfo) {
	// Delete folder if is empty
	if len(files) == 0 {
		os.Remove(dir)
	}
}

func getFilesFromDir(specifiedWorkingDir string) (files []os.FileInfo) {
	// Get files from dir
	files, err := ioutil.ReadDir(specifiedWorkingDir)

	if err != nil {
		// log.Fatal(err)
		log.Println(err)
		logError()
	}
	return files
}

func renameAndMoveFiles(fileType string, outputFolderName string, newfileName string, inputFolder string) {

	// Get the current working directory
	workingDir := getWorkingDir()

	// Read the current working directory + inputFolder name for files
	specifiedWorkingDir := workingDir + "/" + inputFolder
	files := getFilesFromDir(specifiedWorkingDir)

	// Get the number of files in the ouput folder
	outputFolderCount := numFilesInFolder(outputFolderName)

	counter := 0

	// Loop through all files in specified folder
	for _, file := range files {
		if strings.Contains(file.Name(), fileType) {
			// fileNameWithoutExtension := strings.Split(file.Name(), fileType)[0]
			// fmt.Println(fileNameWithoutExtension)

			// Rename and move file
			finalFileDir := workingDir + "/" + outputFolderName + "/" + newfileName + "_" + strconv.Itoa(counter+outputFolderCount) + fileType

			err := os.Rename(filepath.Join(specifiedWorkingDir, file.Name()), finalFileDir)

			if err != nil {
				log.Fatal(err)
				logError()
			}
			counter++
		}
	}

	// Log out how many files were renamed, if any
	if counter+outputFolderCount > outputFolderCount {
		fmt.Println("Renamed:", counter, fileType, "files to ", outputFolderName, "as", newfileName)
	}

	deleteEmptyFolder(specifiedWorkingDir, getFilesFromDir(inputFolder))

}

// This saves error to a file called "ErrorLog.log", so user can dedug what went wrong
func logError() {
	log.Out = os.Stdout

	// Opens ErrorLog if exists if not create it
	file, err := os.OpenFile("ErrorLog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	// Set error log as JSON format
	log.SetFormatter(&logrus.JSONFormatter{})

	// config stuct to store flags
	var config struct {
		inputFolderName  string
		filetypeName     string
		outputFolderName string
		outputFileName   string
	}

	flag.StringVar(&config.filetypeName, "filetype", "0000000", "Enter filetype you want to rename")
	flag.StringVar(&config.outputFolderName, "outputFolder", "output_files", "Enter folder name to store renamed files in")
	flag.StringVar(&config.outputFileName, "renameFileAs", "renamed_file", "What to call the renamed files")
	flag.StringVar(&config.inputFolderName, "inputFolder", "", "Enter the folder of files to rename")
	flag.Parse()

	renameAndMoveFiles(config.filetypeName, config.outputFolderName, config.outputFileName, config.inputFolderName)

}
