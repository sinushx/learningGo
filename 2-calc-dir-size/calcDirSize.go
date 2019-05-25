package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var size int64

func getDirectory(inputArgs []string) string {
	if len(inputArgs) > 0 {
		return inputArgs[0]
	} else {
		directoryAddress, _ := os.Getwd()
		return directoryAddress
	}
}

func getListOfFilesInDir(dirAdd string) []string {
	dirFile, _ := os.Open(dirAdd)

	dirFileNames, _ := dirFile.Readdirnames(-1)
	return dirFileNames
}

/** calcSize (add) => if add is File then size += add.Size
  else if add is a Directory => for each file in add calcSize
*/
func calcSize(currentPath string) {
	fileInfo, _ := os.Lstat(currentPath)
	fileModeStr := string(fileInfo.Mode().String()[0])
	if strings.Contains(fileModeStr, "l") {
		return
	}
	if fileInfo.IsDir() {

		dirFilesList := getListOfFilesInDir(currentPath)
		for _, value := range dirFilesList {
			calcSize(currentPath + "/" + value)
		}
	} else { // currentPath is a File so just add its size
		size += fileInfo.Size()
	}
}

func main() {

	flag.Parse()
	root := getDirectory(flag.Args())

	calcSize(root)

	fmt.Println("Total Size:", size, "Bytes")
}

//TODO: add proper behavior for links. hard => add its size - soft => add link size
