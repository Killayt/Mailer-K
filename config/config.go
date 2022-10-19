package config

import (
	"encoding/json"
	"os"
)

type WorkList struct {
	Subscribers []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"subscribers"`
}

func LoadSubscribers(filename string) (*WorkList, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return &WorkList{[]struct {
			Email string "json:\"email\""
			Name  string "json:\"name\""
			Phone string "json:\"phone\""
		}{}}, err
	}

	var c *WorkList
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return &WorkList{[]struct {
			Email string "json:\"email\""
			Name  string "json:\"name\""
			Phone string "json:\"phone\""
		}{}}, err
	}
	return c, nil
}
