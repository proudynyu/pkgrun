package cwd

import (
	"fmt"
	"os"
	"os/exec"
)

const PackageJsonFile = "package.json"

type PackageError struct { msg string }
func (e *PackageError) Error() string {
	if e.msg != "" {
		return e.msg
	}
	return "No package.json was found in the current directory"
}

func ReadCurrentDir() ([]os.DirEntry, error) {
	dir, err := os.ReadDir(".")
	if err != nil {
		return nil, &PackageError{"Was not possible to read the current directory"}
	}
	return dir, nil
}

func FindPackageJson() (os.DirEntry, error) {
	dir, err := ReadCurrentDir()
	if err != nil {
		return nil, err
	}
	for _, f := range dir {
		if f.Name() == PackageJsonFile && !f.IsDir() {
			return f, nil
		}
	}
	return nil, &PackageError{}
}

func PackageIsInstalled(pkg string) bool {
	_, err := exec.LookPath(pkg)
	if err != nil {
		fmt.Printf("The package manager %s is not installed, proceeding to next one compatible with the current repository\n", pkg)
		return false
	}
	return true
}
