package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Bookmark is a link to a website with a description and descriptive tags
type Bookmark struct {
	Link        string `json:"link"`
	Description string `json:"description"`
	AddDate     string `json:"add_date"`
	Tags        string `json:"tags"`
}

// WriteBookmarksToJSONFile writes an array of `Bookmark`s to a file at a given path
func WriteBookmarksToJSONFile(bs []Bookmark, path string) error {
	bytes, err := json.MarshalIndent(bs, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// WriteBookmarksToTSVFile writes an array of `Bookmark`s to a file at a given path
func WriteBookmarksToTSVFile(bs []Bookmark, path string) error {
	data := [][]string{
		{"link", "description", "tags"},
	}

	for _, b := range bs {
		data = append(data, []string{b.Link, b.Description, b.Tags})
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Comma = '\t'

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return err
		}
	}

	return nil
}
