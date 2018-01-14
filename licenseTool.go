package main

import (
	"fmt"
	"getCopyrights"
	"getLicense"
	"io/ioutil"
)

func main() {
	fullTest("cloudfoundry", "cli")
	fullTest("cloudfoundry", "bosh")
	fullTest("rabbitmq", "rabbitmq-server")
}

func fullTest(repo string, name string) {
	fmt.Println("Scanning ", repo, " ", name)
	licenseObject := getLicense.GetLicense(repo, name)
	noticeFile := getLicense.GetFile(repo, name, "NOTICE")
	fileString := ""
	copyrightClaims := getCopyrights.GetClaims(licenseObject.Content)
	for _, value := range copyrightClaims {
		fileString += value + "\n"
	}
	copyrightClaims = getCopyrights.GetClaims(noticeFile)
	for _, value := range copyrightClaims {
		fileString += value + "\n"
	}

	ioutil.WriteFile("Copyright_Report_"+repo+"_"+name+".txt", []byte(fileString), 0777)
}
