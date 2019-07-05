package goin

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
)

// ReadFromFile returns the contents of a file
// input from an editor in the form of a
// string representing the full file contents
func ReadFromFile(tempBufferFileName string) (string, error) {
	byteValue, err := ReadBytesFromFile(tempBufferFileName)
	return string(byteValue), err
}

// ReadLinesFromFile returns the contents of a file
// input from an editor in the form of an array of strings
// representing each line of the file.
func ReadLinesFromFile(tempBufferFileName string) ([]string, error) {
	err := openEditor(tempBufferFileName)
	file, err := os.Open(tempBufferFileName)
	defer file.Close()

	if fileEmpty(file) {
		return []string{}, nil
	}
	fileScanner := bufio.NewScanner(file)
	var a []string
	for fileScanner.Scan() {
		a = append(a, fileScanner.Text())
	}

	os.Remove(tempBufferFileName)
	return a, err
}

// ReadBytesFromFile returns the contents of a file
// input from an editor in the form of an
// array of bytes representing the full file contents
func ReadBytesFromFile(tempBufferFileName string) ([]byte, error) {
	err := openEditor(tempBufferFileName)
	file, err := os.Open(tempBufferFileName)
	defer file.Close()

	if fileEmpty(file) {
		return []byte{}, nil
	}

	byteValue, err := ioutil.ReadAll(file)
	os.Remove(tempBufferFileName)
	return byteValue, err
}

func openEditor(fileName string) error {
	ed := os.Getenv("EDITOR")
	if ed == "" {
		ed = "vim"
	}

	cmd := exec.Command(ed, fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func fileEmpty(file *os.File) bool {
	fi, err := file.Stat()
	if err != nil || fi.Size() <= 0 {
		return true
	}
	return false
}
