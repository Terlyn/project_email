package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func processSubfolder(subfolderPath string) ([]map[string]string, error) {
	var folderList []map[string]string
	jsonDict := make(map[string]string)
	fileCount := 0

	err := filepath.Walk(subfolderPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			relativePath, err := filepath.Rel(subfolderPath, filePath)
			if err != nil {
				return err
			}

			clave := strings.ReplaceAll(relativePath, string(os.PathSeparator), "_")

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			contentStr := string(content)
			jsonDict[clave] = contentStr
			fileCount++

			if fileCount%50 == 0 {
				folderList = append(folderList, jsonDict)
				jsonDict = make(map[string]string)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Add remaining files if any
	if len(jsonDict) > 0 {
		folderList = append(folderList, jsonDict)
	}

	return folderList, nil
}

func processFolder(folderPath, outputFolder string) error {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subfolderPath := filepath.Join(folderPath, entry.Name())
			folderList, err := processSubfolder(subfolderPath)
			if err != nil {
				return err
			}

			jsonData, err := json.MarshalIndent(folderList, "", "    ")
			if err != nil {
				return err
			}

			jsonFileName := filepath.Join(outputFolder, entry.Name()+".json")
			err = ioutil.WriteFile(jsonFileName, jsonData, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	rootFolder := "/Users/macbook/Documents/proyecto_v2_mail/Data/enron_mail_20110402/maildir"
	outputFile := "/Users/macbook/Documents/proyecto_v2_mail/Data/bachs"

	err := processFolder(rootFolder, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Archivos procesados con éxito")
	}
}
