package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type QA struct {
	q string `json:"q"`
	a string `json:"a"`
}
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

}

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

	words := strings.Split(string(sDec), "\n")
	// fmt.Println(words[10])
	// fmt.Println(words[11])
	// fmt.Println(words[0])
	// fmt.Println(len(words))

	// sellout random
	s1 := rand.NewSource(time.Now().UnixNano())
	randomNum := rand.New(s1).Intn(len(words))
	fmt.Println(words[randomNum])

	// // loop over the array
	// for _, word := range words {
	// 	questionAnswers := strings.Split(word, ";")
	// 	question := questionAnswers[0]
	// 	answer := questionAnswers[1]

	// 	if question == "angry birds" {
	// 		fmt.Println("answer", answer)
	// 		q = question
	// 		a = answer

	// 		fmt.Println(qa)
	// 		// return
	// 	}
	// }

}
