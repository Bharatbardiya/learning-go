package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseConfig(envFileContent string) (map[string]string, error) {
	config := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(envFileContent))
	re := regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(?:'([^']*)'|"([^"]*)"|([^#\s]*))?(?:\s*#.*)?$`)
	lineNo := 0

	for scanner.Scan() {
		line := scanner.Text()
		trimLine := strings.TrimSpace(line)
		lineNo++

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		matches := re.FindStringSubmatch(trimLine)

		key := matches[1]
		var val string

		if matches[2] != "" {
			val = matches[2]
		} else if matches[3] != "" {
			val = matches[3]
		} else {
			val = matches[4]
		}

		config[key] = val
	}
	return config, nil
}

func main() {
	envFileContent := `
# Application Configuration
APP_NAME="My Cool App"
APP_VERSION="1.0.2-beta" # Version with quotes
PORT=8080
DEBUG_MODE="true"
# Database Settings
DB_HOST=localhost
DB_USER = admin
DB_PASSWORD = "p@s$w Ord With Sp@ces!" # Quoted password
API_ENDPOINT = https://api.example.com/v1

# An empty value
EMPTY_KEY=
ANOTHER_KEY_NO_VALUE =`

	config, err := parseConfig(envFileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for k, v := range config {
		fmt.Printf("%s=%q\n", k, v)
	}
}
