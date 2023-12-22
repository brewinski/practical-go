package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	name, repos, err := githubInfo("brewinski")
	if err != nil {
		log.Fatalf("error getting github info: %s", err)
	}

	fmt.Println(name, repos)
}

func githubInfo(username string) (string, int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(username))
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, fmt.Errorf("unable to make request, %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("received non-200 status code: %s", resp.Status)
	}

	// var r Reply
	var r struct {
		Name        string `json:"name"`
		PublicRepos int    `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&r); err != nil {
		return "", 0, fmt.Errorf("unable to decode github response body, %w", err)
	}
	return r.Name, r.PublicRepos, nil
}

// type Reply struct {
// 	Name        string `json:"name"`
// 	PublicRepos int    `json:"public_repos"`
// }
