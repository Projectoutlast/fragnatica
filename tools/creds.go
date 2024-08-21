package creds

import (
	"math/rand"
	"os"
	"strings"
)

func getListOfUserAgents(sourceFile string) ([]string, error) {
	file, err := os.ReadFile(sourceFile)
	if err != nil {
		return nil, err
	}

	splittedFile := strings.Split(string(file), "\n")

	return splittedFile, nil
}

func GetRandomUserAgent(sourceFile string) (string, error) {
	userAgents, err := getListOfUserAgents(sourceFile)
	if err != nil {
		return "", err
	}

	randomIndex := rand.Intn(len(userAgents))

	return userAgents[randomIndex], nil
}
