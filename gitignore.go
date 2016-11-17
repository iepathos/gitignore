package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func copyGitignoreUrl(url string) {
	// create .gitignore
	f, err := os.Create(".gitignore")
    if err != nil {
		panic(err)
	}

	// retrieve gitignore from given url
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// write to file
	_, err = f.Write(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(".gitignore created")
}

func main() {
	if (len(os.Args)) >= 2 {
		switch env := os.Args[1]; env {
		case "python":
			// https://raw.githubusercontent.com/github/gitignore/master/Python.gitignore
			copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Python.gitignore")
		case "go":
			// https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore
			copyGitignoreUrl("https://raw.githubusercontent.com/github/gitignore/master/Go.gitignore")
		default:
			fmt.Println("Please pass type of gitignore to create, like: gitignore python")
		}
	} else {
		fmt.Println("Please pass type of gitignore to create, like: gitignore python")
	}
	
}
