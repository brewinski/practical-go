package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/brewinski")
	if err != nil {
		log.Fatalf("Error fetching github user: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from github: %s", resp.Status)
	}

	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("Error copying response body: %s", err)
	// }

	var r Reply

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&r); err != nil {
		log.Fatalf("Error decoding json: %s", err)
	}

	fmt.Println(r)
}

func githubInfo(login string) (string, int, error) {
	// TODO: your code goes here
	return "", 0, nil
}

type Reply struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}
