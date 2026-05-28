package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error has ocurred loading the file: [%s]", path)
		fmt.Printf("[Error]: %e", err)
		return nil, err
	}
	return buffer, nil
}

type PackageFormat struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Scripts  map[string]string `json:"scripts"`
}
func ReadPackageJson(f os.DirEntry) *PackageFormat {
	json_string, err := ReadFile(f.Name())
	if err != nil {
		// TODO: fix the error treatment
		panic("Not possible to read the json file")
	}
	data := PackageFormat{}
	err = json.Unmarshal(json_string, &data)
	if err != nil {
		// TODO: fix the error treatment
		panic("Was not possible to decode in json format")
	}
	return &data
}
