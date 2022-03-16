package conv

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Color struct {
	Col []string `json:"colors"`
}

func GetColor() (*Color, error) {
	file, err := ioutil.ReadFile("json/color.json")
	if err != nil {
		log.Fatalf("error in reading file >> %v", err)
	}

	c := &Color{}

	if err := json.Unmarshal([]byte(file), &c); err != nil {
		return nil, err
	}
	return c, nil
}
