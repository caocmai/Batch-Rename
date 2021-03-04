[![Go Report Card](https://goreportcard.com/badge/github.com/caocmai/Make-Utility)](https://goreportcard.com/report/github.com/caocmai/Make-Utility)

# Make Utility
This Go program renames all files of specified file type and stores them into a new folder.

## Prerequisites
* [Go - v1.16](https://golang.org/doc/install)

## Usage 
To use this program, provide the information to the flag items. The filetype flag is the only required flag, and when an input folder is not specified it looks at the root of the project directory for that file type. All other flags are optional. 

### Example Usage
Case 1: If you want to rename all `.jpg` files in a folder called `pictures`.

In terminal run:

`$ go run main.go -filetype=.jpg -inputFolder=pictures`

You can use other flags to customize for renaming of files, else they will take the default values, shown in the Flags Table below.

Case 2: If you want to rename all `.jpg` files in a folder called `pictures` to something that starts with `Vacation_2021_XX` and saved all the renamed files into a folder called `Vacation2021`.

In terminal run:

`$ go run main.go -filetype=.jpg -inputFolder=pictures -ouputFolder=Vacation2021 -renameFileAs=Vacation_2021`


#### Flags
A list of all the flags and their usage for this project

| Flag | Default Value | Helper Text |
| :--- | :---: | :--- |
| `filetype` | `nil` | [Required] Enter filetype you want to rename, ie `.txt` |
| `inputFolder` | ""  | Enter the folder of files to rename |
| `outputFolder` | "output_files" | Enter folder name to store renamed files in |
| `renameFileAs` | "renamed_file" | What to call the renamed files |

