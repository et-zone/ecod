package template

import (
	// "errors"
	"go/format"
	// "net/url"
	"os"
	"path/filepath"
	"strings"
	// "unicode"
)

var UPword = []string{"id"}

func isWordInUPword(s string) bool {
	for _, sw := range UPword {
		if strings.ToLower(s) == sw {
			return true
		}
	}
	return false
}

func SaveFile(dirPath, fileName string, text []byte) error {
	file, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	p, err := format.Source(text)
	if err != nil {
		return err
	}
	_, err = file.Write(p)
	return err
}

func MkdirPathIfNotExist(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, os.ModePerm)
	}
	return nil
}

func CleanUpGenFiles(dir string) error {
	exist, err := FileExists(dir)
	if err != nil {
		return err
	}
	if exist {
		return os.RemoveAll(dir)
	}
	return nil
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) (bool, error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

func replaceInvalidChars(str string) string {
	str = strings.ReplaceAll(str, "-", "_")
	str = strings.ReplaceAll(str, " ", "_")
	return strings.ReplaceAll(str, ".", "_")
}
func UpFirstChar(s string) string {
	f := strings.ToUpper(string(s[0]))
	return string(f) + s[1:]
}

func LowFirstChar(s string) string {
	f := strings.ToLower(string(s[0]))
	return string(f) + s[1:]
}
