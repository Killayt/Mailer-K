package config

import (
	"encoding/json"
	"os"
)

type Subscriber struct {
	Mail string
	Name string
}

func LoadSubscribers(filename string) (Subscriber, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return Subscriber{}, err
	}

	var c Subscriber
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return Subscriber{}, err
	}
	return c, nil
}
