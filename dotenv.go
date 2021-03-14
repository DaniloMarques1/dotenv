package dotenv

import (
	"io/ioutil"
	"os"
	"strings"
)

// loads the key value pair from the files
// if no files is passed tries to read from
// .env file if exists
func Load(files ...string) error {
	var fileName string
	if len(files) == 0 {
		fileName = ".env"
		if err := readAndSetEnv(fileName); err != nil {
			return err
		}
	} else {
		for _, fileName = range files {
			if err := readAndSetEnv(fileName); err != nil {
				return err
			}
		}
	}

	return nil
}

// reads the file with the filename name
// and calls the setEnv function
func readAndSetEnv(fileName string) error {
	lines, err := readFromFile(fileName)
	if err != nil {
		return err
	}
	setEnv(lines)

	return nil
}

// receives a slice of strings where each index is a line
// and separe based on key=value to set to the env
func setEnv(lines []string) {
	for _, line := range lines {
		keyValue := strings.Split(line, "=")
		os.Setenv(strings.Trim(keyValue[0], " "), strings.Trim(keyValue[1], " "))
	}
}

// read from file and return a []string with its content
// separated by line
func readFromFile(fileName string) ([]string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	str := string(bytes)
	lines := strings.Split(strings.TrimSuffix(str, "\n"), "\n")
	return lines, nil
}
