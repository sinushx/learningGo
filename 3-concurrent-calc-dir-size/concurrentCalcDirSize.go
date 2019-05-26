package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

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

func calcSize(currentPath string, ch chan int64) {
	fileInfo, _ := os.Lstat(currentPath)
	fileModeStr := string(fileInfo.Mode().String()[0])
	if strings.Contains(fileModeStr, "l") {
		return
	}
	if fileInfo.IsDir() {
		dirFilesList := getListOfFilesInDir(currentPath)
		for _, value := range dirFilesList {
			calcSize(currentPath+"/"+value, ch)
		}
	} else { // currentPath is a File so just add its size
		ch <- fileInfo.Size()
	}

}

func main() {

	flag.Parse()
	root := getDirectory(flag.Args())

	fileSizesChannel := make(chan int64)

	go func() {
		calcSize(root, fileSizesChannel)
		close(fileSizesChannel)
	}()

	var totalSize int64
	for fileSize := range fileSizesChannel {
		totalSize += fileSize
	}

	fmt.Println("Total Size:", totalSize)
}
