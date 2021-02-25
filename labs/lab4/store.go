package main

import (
	"encoding/json"
	"os"
	"strings"
)

const datafile = "db.json"

type Website struct {
	Title  string            `json:"title"`
	Body   string            `json:"body"`
	Errors map[string]string `json:"-"`
}

func LoadWebsite() *Website {
	file, err := os.Open(datafile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var website Website
	json.NewDecoder(file).Decode(&website)
	return &website
}

func (website *Website) Save() error {
	file, err := os.Create(datafile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(&website)
}

func (website *Website) Validate() bool {
	website.Errors = make(map[string]string)

	if strings.TrimSpace(website.Title) == "" {
		website.Errors["Title"] = "Please enter title"
	}

	if strings.TrimSpace(website.Body) == "" {
		website.Errors["Body"] = "Please enter body"
	}

	return len(website.Errors) == 0
}
