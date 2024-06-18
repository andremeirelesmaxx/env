package action

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Write(envs map[string]string, projectDir string) {
	envExists := false
	_, err := os.Stat(projectDir + "/.env")

	if err == nil {
		envExists = true
	}

	if !envExists {
		createEnv(envs, projectDir)

		return
	}

	updateEnvs(envs, projectDir)
}

func createEnv(envs map[string]string, projectDir string) {
	content := ""

	_, ok := envs["APP_KEY"]

	if !ok {
		envs["APP_KEY"] = ""
	}

	for k, v := range envs {
		if k == "DB_PASSWORD" {
			content += fmt.Sprintf("%s='%s'\n", k, v)
			continue
		}
		content += fmt.Sprintf("%s=%s\n", k, v)
	}

	file, err := os.Create(projectDir + "/.env")
	if err != nil {
		fmt.Println("Cannot create file, Error: ", err)

		return
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Cannot write file, Error: ", err)

		return
	}

	fmt.Println("Env file created.")
}

func updateEnvs(envs map[string]string, projectDir string) {
	filePath := projectDir + "/.env"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Cannot read file, Error: ", err)

		return
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	appKey := ""

	for scanner.Scan() {
		line := scanner.Text()
		value := strings.SplitN(line, "=", 2)

		if value[0] == "APP_KEY" {
			appKey = value[1]
			break
		}
	}

	if appKey == "" {
		fmt.Println("APP_KEY not found")

		return
	}

	envs["APP_KEY"] = appKey
	err = os.Remove(filePath)

	if err != nil {
		fmt.Println("Cannot remove file, Error: ", err)

		return
	}

	createEnv(envs, projectDir)
}
