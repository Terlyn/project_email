package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func indexJSONToServer(jsonFolderPath, apiURL, username, password string) {
	files, err := ioutil.ReadDir(jsonFolderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			jsonFilePath := filepath.Join(jsonFolderPath, file.Name())
			jsonContent, err := ioutil.ReadFile(jsonFilePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
				continue
			}

			req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonContent))
			if err != nil {
				fmt.Printf("Error creating request for file %s: %v\n", file.Name(), err)
				continue
			}
			req.SetBasicAuth(username, password)
			req.Header.Set("Content-Type", "application/json")

			// Disable SSL verification
			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error sending request for file %s: %v\n", file.Name(), err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				fmt.Printf("Successfully indexed %s\n", file.Name())
			} else {
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Printf("Failed to index %s: %d, %s\n", file.Name(), resp.StatusCode, string(body))
			}
		}
	}
}

func main() {
	// apiURL := "https://api.openobserve.ai/api/terlyn_organization_418/enron_mail/_json"
	// username := "galeanoterlyn@gmail.com"
	// password := "348n2GsAw5IU96EK170o"
	// jsonFolderPath := "/Users/macbook/Documents/proyecto_v2_mail/Data/bachs"
	apiURL := "https://api.openobserve.ai/api/anthony_Q2ralmQFAgWI8oo/enron_mail/_json"
	username := "galeanoterlyn@gmail.com"
	password := "9h17Gt86j0vRT352F4aq"
	jsonFolderPath := "/Users/macbook/Documents/proyecto_v2_mail/Data/test"

	indexJSONToServer(jsonFolderPath, apiURL, username, password)
}
