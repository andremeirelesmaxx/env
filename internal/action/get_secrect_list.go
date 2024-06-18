package action

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/maxxcard/env-inject/internal"
)

func GetSecrets(secretId string) (map[string]string, error) {
	command := exec.Command("aws", "secretsmanager", "get-secret-value", "--secret-id", secretId)
	output, err := command.CombinedOutput()

	if err != nil {
		return nil, err
	}

	aws := new(internal.AWSResponse)
	err = json.Unmarshal(output, aws)

	if err != nil {
		return nil, err
	}

	secretsMap := new(map[string]any)
	json.Unmarshal([]byte(aws.SecretString), secretsMap)
	castMap := make(map[string]string)

	for k, v := range *secretsMap {
		castMap[k] = fmt.Sprintf("%v", v)
	}

	return castMap, nil
}
