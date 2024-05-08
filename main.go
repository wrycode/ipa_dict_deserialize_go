package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type IPADict map[string]string


func LoadJson(lang string) (IPADict, error) {
	var jsonDict struct {
		Lang []IPADict `json:"en_US"`
	}

	file := fmt.Sprintf("./ipa_dicts/%s.json", lang)

	jsonFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonFile, &jsonDict)
	if err != nil {
		return nil, err
	}

	return jsonDict.Lang[0], nil
}

func main() {
	ipa, err := LoadJson("en_US")
	if err != nil {
		log.Fatal(err)
	}
	for word, pronunciation := range ipa {
		fmt.Printf("Word: %s, Pronunciation: %s\n", word, pronunciation)
	}
}
