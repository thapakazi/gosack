package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type GithubContentResponse struct {
	_links struct {
		Git  string `json:"git"`
		HTML string `json:"html"`
		Self string `json:"self"`
	} `json:"_links"`
	Content     string `json:"content"`
	DownloadURL string `json:"download_url"`
	Encoding    string `json:"encoding"`
	GitURL      string `json:"git_url"`
	HTMLURL     string `json:"html_url"`
	Name        string `json:"name"`
}

func main() {

	url := "https://api.github.com/repos/foss-np/np-l10n-glossary/contents/en2ne/fun.tra"
	downloadFromUrl(url)

	// parse the json

}

// func (resp GithubContentResponse) getContent() {

// }

func downloadFromUrl(url string) {

	tmp_path := "/tmp/"
	tokens := strings.Split(url, "/")
	fileName := tmp_path + tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// decode and print
	githubcontentresponse := GithubContentResponse{}
	json.Unmarshal(raw, &githubcontentresponse)
	contentsInside := githubcontentresponse.Content

	sDec, _ := b64.StdEncoding.DecodeString(contentsInside)
	fmt.Println(string(sDec))

}
