package getLicense

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Download_URL string
	Html_URL     string
	Content      string
	License      License
}

type License struct {
	Name string
	URL  string
}

type GitFile struct {
	Name         string
	Download_URL string
}

func GetLicense(user string, repo string) Response {
	gitURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/license", user, repo)
	resp, err := http.Get(gitURL)
	if err != nil {
		fmt.Println(err)
	}
	responseJSON := Response{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		panic(err)
	}
	Content, err := base64.StdEncoding.DecodeString(responseJSON.Content)
	if err != nil {
		fmt.Println("decode error:", err)
	}
	responseJSON.Content = string(Content)
	fmt.Println("License Object : ", responseJSON.License)
	return responseJSON
}

func GetFile(user string, repo string, filename string) string {
	gitURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents", user, repo)
	resp, err := http.Get(gitURL)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var returnArray []GitFile
	err = json.Unmarshal(body, &returnArray)
	if err != nil {
		panic(err)
	}
	var notice string
	for _, value := range returnArray {
		if value.Name == filename {
			resp, err := http.Get(value.Download_URL)
			if err != nil {
				fmt.Println(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			notice = string(body)
		}
	}
	return notice
}
