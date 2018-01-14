package getCopyrights

import (
	"fmt"
	"regexp"
)

//TODO: Read from Configfile
var reClaims = regexp.MustCompile("[^\\n]*Copyright\\s*(\\(c\\)|\\(C\\))[^\\n]*")
var reLicense = regexp.MustCompile("[^\\n]*Licensed[^\\n]*")

// GetClaims finds all sentences Containing Copyright
func GetClaims(text string) []string {
	reArray := reClaims.FindAllString(text, -1)
	if len(reArray) == 0 {
		//fmt.Println("No Copyright Claims Found")
	} else {
		//	for index, value := range reArray {
		//fmt.Println(index+1, " ", value)
		//	}
	}
	return reArray
}

func GetLicense(text string) []string {
	reArray := reLicense.FindAllString(text, -1)
	if len(reArray) == 0 {
		//fmt.Println("No Copyright Claims Found")
	} else {
		for index, value := range reArray {
			fmt.Println(index+1, " ", value)
		}
	}
	return reArray
}
