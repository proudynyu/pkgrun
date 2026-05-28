package cwd

import (
	"fmt"
	"os"
	"os/exec"
)

const PACKAGE = "package.json"

type PackageError struct {}
func (e *PackageError) Error() string {
	return "No package.json was found in the current directory"
}

func ReadCurrentDir() []os.DirEntry {
	dir, err := os.ReadDir(".")
	if err != nil {
		panic("Was not possible to read the current directory")
	}
	return dir
}

func FindPackageJson() (os.DirEntry, error) {
	dir := ReadCurrentDir()
	for _, f := range dir {
		if f.Name() == PACKAGE && !f.IsDir() {
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
