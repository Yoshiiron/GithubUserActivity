package cmd

import (
	"GithubUserActivity/internal/activity"
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github-activity",
		Short: "Github User Activity tool",
		Long:  `Apllication for viewing last user activity.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDisplayActivityCmd(args)
		},
	}
	return cmd
}

func RunDisplayActivityCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a username")
	}

	username := args[0]
	activities, err := activity.FetchGithubActvity(username)
	if err != nil {
		return err
	}

	return activity.DisplayActivity(username, activities)
}
