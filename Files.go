package wiz

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed functions:
//			Executable() string
//				Returns full path to currently running executable
//			ProgramName() string
//				Returns just the name of the running executable (no path)
//			Dir() string
//				Returns the absolute path of the directory containing the executable
//			FileExists(file string) bool
//				Returns true if a file exists at a given relative path
//			FolderExists(file string) bool
//				Returns true if a folder exists at a given relative path
//			DeleteFile(file string) error
//				Deletes a file at a given relative path
//			MkDir(dir string) error
//				Makes a folder at a given relative path (uses permissions 0644)
//			WriteFile(file string, data []byte) error
//				Writes data to a file at given relative path (overwrites existing file if any)
//		  ReadFile(file string) ([]byte, error)
//				Reads a file at a given relative path

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

func Executable() string {
	p, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return p
}

func ProgramName() string {
	return filepath.Base(Executable())
}

func Dir() string {
	p := Executable()
	return filepath.FromSlash(filepath.Dir(p) + "/")
}

func FileExists(file string) bool {
	path := filepath.FromSlash(Dir() + file)
	info, err := os.Stat(path)
	if err != nil {
		//Formerly if os.IsNotExist(err)
		return false
	}
	return !info.IsDir()
}

func FolderExists(file string) bool {
	path := filepath.FromSlash(Dir() + file)
	info, err := os.Stat(path)
	if err != nil {
		//Formerly if os.IsNotExist(err)
		return false
	}
	return info.IsDir()
}

// DeleteFile deletes a file (or empty directory) at a location
func DeleteFile(file string) error {
	path := filepath.FromSlash(Dir() + file)
	err := os.Remove(path)
	return errors.Wrap(err, "Deletefile")
}

// MkDir Creates a folder with specified relative path + name ('dir')
func MkDir(dir string) error {
	path := filepath.FromSlash(Dir() + dir)
	err := os.Mkdir(path, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return errors.Wrap(err, "MkDir")
	}
}

// WriteFile writes to (or overwrites) a file
func WriteFile(file string, data []byte) error {
	path := filepath.FromSlash(Dir() + file)
	err := ioutil.WriteFile(path, data, 0644)
	return errors.Wrap(err, "WriteFile")
}

// ReadFile read a whole file and returns it
func ReadFile(file string) ([]byte, error) {
	path := filepath.FromSlash(Dir() + file)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, errors.Wrap(err, "ReadFile")
	}
	return contents, nil
}
