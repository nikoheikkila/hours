package toggl

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Token     string `yaml:"api_token"`
	Workspace string `yaml:"workspace_id"`
}

func LoadConfiguration() (*Configuration, error) {
	var configuration Configuration
	path := fmt.Sprintf("%s/.togglrc", os.Getenv("HOME"))

	if !fileExists(path) {
		return &Configuration{
			Token:     env("TOGGL_API_TOKEN"),
			Workspace: env("TOGGL_WORKSPACE_ID"),
		}, nil
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func fileExists(filename string) bool {
	file, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !file.IsDir()
}

func env(key string) string {
	fmt.Printf("Attempting to read key $%s from environment\n", key)

	value := os.Getenv(key)

	if value == "" {
		fmt.Printf("Could not read environment variable $%s.\n", key)
		os.Exit(2)
	}

	return value
}
