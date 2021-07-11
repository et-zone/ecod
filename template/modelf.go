package template

import (
	"io/ioutil"
)

func GetFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}
	}
	return data
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0666)
}
