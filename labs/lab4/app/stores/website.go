package stores

import (
	"encoding/json"
	"os"
)

const datafile = "db.json"

type Website struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Load() *Website {
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
