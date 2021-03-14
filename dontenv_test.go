package dotenv

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	clearEnv()
	os.Remove(".test")
	os.Exit(code)
}

func TestLoad(t *testing.T) {
	clearEnv()
	fileContent := "DOTENV_USER = Danilo\nDOTENV_PASSWORD =    12345\n"
	file, err := os.Create(".test")
	if err != nil {
		t.Errorf("Failed to create file %v\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(fileContent)
	if err != nil {
		t.Errorf("Failed to write to file %v\n", err)
	}

	if err := Load(".test"); err != nil {
		t.Errorf("Error loading vars %v\n", err)
	}
	user := os.Getenv("DOTENV_USER")
	password := os.Getenv("DOTENV_PASSWORD")
	if user != "Danilo" || password != "12345" {
		t.Errorf("Failed to load env variables USER=%v PASSWORD=%v", user, password)
	}

	// testing with multiple files
	clearEnv()
	file, err = os.Create(".test_2")
	if err != nil {
		t.Errorf("Failed to create file %v\n", err)
	}
	defer file.Close()
	defer os.Remove(".test_2")
	fileContent = "DOTENV_HOST=Localhost\nDOTENV_PORT=5000"
	_, err = file.WriteString(fileContent)
	if err != nil {
		t.Errorf("Error writing to file %v\n", err)
	}
	Load(".test", ".test_2")
	user = os.Getenv("DOTENV_USER")
	password = os.Getenv("DOTENV_PASSWORD")
	host := os.Getenv("DOTENV_HOST")
	port := os.Getenv("DOTENV_PORT")

	if user != "Danilo" || password != "12345" || host != "Localhost" || port != "5000" {
		t.Errorf("Failed to load env variables USER=%v PASSWORD=%v HOST=%v PORT=%v", user, password, host, port)
	}
}

func TestReadFromFile(t *testing.T) {
	fileStr := "Danilo\nMarques"
	file, err := os.Create(".test")
	if err != nil {
		t.Errorf("Failed to create the test file")
	}
	_, err = file.WriteString(fileStr)
	if err != nil {
		t.Errorf("Failed to write to the test file")
	}
	file.Close()
	returnedFromFile, err := readFromFile(".test")
	if err != nil {
		t.Errorf("Error reading from file %v", err)
	}

	if returnedFromFile[0] != "Danilo" || returnedFromFile[1] != "Marques" {
		t.Errorf("Failed to correctly read from file")
	}
}

func TestSetEnv(t *testing.T) {
	clearEnv()
	lines := []string{"DOTENV_NAME=Danilo", "DOTENV_PASSWORD=1234"}
	setEnv(lines)
	user := os.Getenv("DOTENV_NAME")
	password := os.Getenv("DOTENV_PASSWORD")
	if user != "Danilo" || password != "1234" {
		t.Errorf("Failed to set enviroment variables")
	}
}

func TestReadAndSetEnv(t *testing.T) {
	clearEnv()
	file, err := os.Create(".test")
	if err != nil {
		t.Errorf("Problem creating file\n")
	}
	content := "DOTENV_USER=Danilo"
	file.WriteString(content)
	file.Close()
	if err := readAndSetEnv(".test"); err != nil {
		t.Errorf("Error setting up env %v\n", err)
	}

	user := os.Getenv("DOTENV_USER")
	if user != "Danilo" {
		t.Errorf("Error expecting Danilo got %s\n", user)
	}
}

func clearEnv() {
	os.Unsetenv("DOTENV_USER")
	os.Unsetenv("DOTENV_PASSWORD")
}
