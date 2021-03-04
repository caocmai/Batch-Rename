package main

import "testing"

func TestGetFiles(t *testing.T) {
	dir := getWorkingDir()
	files := getFilesFromDir(dir)
	if len(files) < 0 {
		t.Error("Expected there to be more files in current directory")
	}

}

func TestTableFolderSizes(t *testing.T) {
	var foldersTest = []struct {
		name           string
		expectedLength int
	}{
		{"output_files", 2},
		{"folder3", 3},
	}

	for _, eachFolderTest := range foldersTest {
		workingDir := getWorkingDir()
		specifiedWorkingDir := workingDir + "/" + eachFolderTest.name

		if output := getFilesFromDir(specifiedWorkingDir); len(output) != eachFolderTest.expectedLength {
			t.Error("Test Failed: {} folder inputted, {} expected, but recieved: {}", eachFolderTest.name, eachFolderTest.expectedLength, output)
		}

	}
}
