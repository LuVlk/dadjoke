package model

import (
	"encoding/json"
)

type Joke struct {
	Joke string `json:"joke"`
}

func MakeJokeFromJson(data []byte) (*Joke, error) {
	joke := Joke{}
	err := json.Unmarshal(data, &joke)
	if err != nil {
		return nil, err
	}
	return &joke, nil
}
