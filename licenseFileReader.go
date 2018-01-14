package main

import (
	"fmt"
	"getCopyrights"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	timemeasure := time.Now()
	dir := "/home/drayasha/Schreibtisch/bosh"
	findLicenseFiles(dir)
	searchDirectory(dir)
	fmt.Println(time.Now().Sub(timemeasure))
}

func searchDirectory(dir string) {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		fmt.Println(path)
		return nil
	})
	for _, file := range fileList {
		content, _ := ioutil.ReadFile(file)
		copyrights := getCopyrights.GetClaims(string(content))

		if len(copyrights) != 0 {
			fmt.Println(file)
			filename := "./copyrights/CopyrightsIn" + strings.Replace(file, "/", "-", -1) + ".txt"
			writeStringArray(copyrights, filename)
		}
	}
}

func findLicenseFiles(dir string) {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, "LICENSE") || strings.Contains(path, "license") {
			fileList = append(fileList, path)
			fmt.Println(strings.Replace(path, dir, "~", -1))
		}
		return nil
	})
	filename := "licensesIn" + "dir" + ".txt"
	writeStringArray(fileList, filename)
}

func writeStringArray(arrayToWrite []string, filename string) {
	fileString := ""
	for _, value := range arrayToWrite {
		fileString += value + "\n"
	}
	ioutil.WriteFile(filename, []byte(fileString), 0777)
}
