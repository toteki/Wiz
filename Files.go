package wiz

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Returns full path + filename of currently running executable
func Executable() string {
	p, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return p
}

// Returns just the name of the running executable (no path)
func ProgramName() string {
	return filepath.Base(Executable())
}

// Returns the absolute path of the directory containing the executable
func Dir() string {
	p := Executable()
	return filepath.FromSlash(filepath.Dir(p) + "/")
}

// Returns true if a file exists at a given relative path
func FileExists(file string) bool {
	path := filepath.FromSlash(Dir() + file)
	info, err := os.Stat(path)
	if err != nil {
		//Formerly if os.IsNotExist(err)
		return false
	}
	return !info.IsDir()
}

// Returns true if a folder exists at a given relative path
func FolderExists(file string) bool {
	path := filepath.FromSlash(Dir() + file)
	info, err := os.Stat(path)
	if err != nil {
		//Formerly if os.IsNotExist(err)
		return false
	}
	return info.IsDir()
}

// DeleteFile deletes a file (or empty directory) at a given relative path
func DeleteFile(file string) error {
	path := filepath.FromSlash(Dir() + file)
	err := os.Remove(path)
	return errors.Wrap(err, "wiz.Deletefile")
}

// MkDir Creates a folder with specified relative path + name ('dir') with 0755 permissions
func MkDir(dir string) error {
	path := filepath.FromSlash(Dir() + dir)
	err := os.Mkdir(path, 0755)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return errors.Wrap(err, "wiz.MkDir")
	}
}

// WriteFile writes to (or overwrites) a file at a given relative path
func WriteFile(file string, data []byte) error {
	path := filepath.FromSlash(Dir() + file)
	err := ioutil.WriteFile(path, data, 0644)
	return errors.Wrap(err, "wiz.WriteFile")
}

// ReadFile read a whole file at a given relative path and returns it as []byte
func ReadFile(file string) ([]byte, error) {
	path := filepath.FromSlash(Dir() + file)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.ReadFile")
	}
	return contents, nil
}
