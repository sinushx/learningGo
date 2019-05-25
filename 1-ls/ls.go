package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDirectory (otherInputArgs [] string) string {
	if len(otherInputArgs) > 0 {
		return otherInputArgs[0]
	} else {
		directoryAddress, _ := os.Getwd()
		return directoryAddress
	}
}


func getListOfFilesInDir (dirAdd string) [] string {
	dirFile, err := os.Open(dirAdd)

	dirFileNames, err := dirFile.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}
	return dirFileNames
}


func showLongList (dirFileNames [] string, dirAddress string, isHumanReadable bool) {

	color.Set(color.FgYellow)
	fmt.Print("Mode")
	fmt.Print("           ")
	fmt.Print("Size")
	fmt.Print("       ")
	fmt.Println("Name")
	fmt.Println("--------------------------------")
	color.Unset()

	for _, value := range dirFileNames {
		fullPath := dirAddress + "/" + value

		valueFileInfo, _ := os.Stat(fullPath)

		fmt.Print(valueFileInfo.Mode())
		fmt.Print("     ")

		if isHumanReadable {
			fmt.Print(makeHumanReadable(valueFileInfo.Size()))
		} else {
			fmt.Print(valueFileInfo.Size(), "B")
		}
		fmt.Print("     ")
		if valueFileInfo.IsDir() {
			color.Cyan(valueFileInfo.Name())
		} else {
			fmt.Println(valueFileInfo.Name())
		}

	}

}


func makeHumanReadable(size int64) string {
	if size < 1024 { // less than 1 KB
		return strconv.FormatInt(int64(size), 10) + "B"
	} else if 1024 <= size && size < 1048576 { // between 1KB and 1MB
		return strconv.FormatInt(int64(size/1024), 10) + "KB"
	} else if 1048576 <= size && size < 1073741824 { // between 1MB and 1GB
		return strconv.FormatInt(int64(size/1048576), 10) + "MB"
	} else { // Bigger than 1GB
		return strconv.FormatInt(int64(size/1073741824), 10) + "GB"
	}
}


func removeHiddenFiles (dirFileNames [] string) [] string {
	for i, value := range dirFileNames {
		if strings.Index(value, ".") == 0 {
			dirFileNames = append(dirFileNames[:i], dirFileNames[i+1:]...)
		}
	}
	return dirFileNames
}


func main() {

	longListPtr := flag.Bool("l", false, "Shows more detailed list.")
	showHiddensPtr := flag.Bool("a", false, "Shows all files including hidden ones.")
	humanReadablePtr := flag.Bool("h", false, "With -l shows size in human readable way.")

	flag.Parse()

	directoryAddress:= getDirectory(flag.Args())
	directoryFileNames := getListOfFilesInDir(directoryAddress)


	if !*showHiddensPtr { directoryFileNames = removeHiddenFiles(directoryFileNames) }

	if *longListPtr { showLongList (directoryFileNames, directoryAddress,*humanReadablePtr)} else {
		for _, value := range directoryFileNames {
			fullPath := directoryAddress + "/" + value
			valueFileInfo, _ := os.Stat(fullPath)
			if valueFileInfo.IsDir() {color.Cyan(valueFileInfo.Name())} else {
				fmt.Println(value)
			}
		}
	}
}