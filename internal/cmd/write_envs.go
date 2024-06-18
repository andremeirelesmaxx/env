package cmd

import (
	"fmt"

	"github.com/maxxcard/env-inject/internal/action"
	"github.com/spf13/cobra"
)

func WriteEnvs() *cobra.Command {
	secretId := new(string)
	projectDir := new(string)
	command := &cobra.Command{
		Use:   "inject [secrect-id]",
		Short: "Get env from aws secret manager and add as env",
		Run: func(cmd *cobra.Command, args []string) {
			if *secretId == "" {
				fmt.Println("Secret id must not be null")

				return
			}

			if *projectDir == "" {
				fmt.Println("Project dir must not be null")
				return
			}

			secrets, err := action.GetSecrets(*secretId)
			if err != nil {
				fmt.Println("Error: ", err)

				return
			}

			action.Write(secrets, *projectDir)
		},
	}
	command.PersistentFlags().StringVarP(secretId, "secret-id", "s", "", "")
	command.PersistentFlags().StringVarP(projectDir, "project-dir", "p", "/var/www/html", "directory where project is located")

	return command
}
