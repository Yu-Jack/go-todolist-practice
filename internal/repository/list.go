package repository

import (
	"io/ioutil"
)

var fileName string = "data.json"

func GetList() ([]byte, error) {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func SaveList(content []byte) error {
	err := ioutil.WriteFile(fileName, content, 0644)

	if err != nil {
		return err
	}

	return nil
}
