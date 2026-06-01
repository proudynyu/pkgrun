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
		fmt.Printf("[Error]: %s", err.Error())
		return nil, err
	}
	return buffer, nil
}

type PackageFormat struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Scripts  map[string]string `json:"scripts"`
}

type PackageReadError struct { msg string }
func (p *PackageReadError) Error() string {
	return fmt.Sprintf("An error has occured when reading package.json: \n\n[ERROR]: %s", p.msg)
}
func ReadPackageJson(f os.DirEntry) (*PackageFormat, error) {
	json_string, err := ReadFile(f.Name())
	if err != nil {
		return nil, &PackageReadError{"Not possible to read the json file"}
	}
	data := PackageFormat{}
	err = json.Unmarshal(json_string, &data)
	if err != nil {
		return nil, &PackageReadError{"Was not possible to decode in json format"}
	}
	return &data, nil
}
