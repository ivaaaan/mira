package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ivaaaan/mira/parser/markdown"
	"github.com/ivaaaan/mira/provider"
	"github.com/ivaaaan/mira/provider/jira"
	"github.com/ivaaaan/mira/provider/plain"
	"github.com/spf13/cobra"
)

var (
	Plain bool
	File  string
)

func init() {
	pushCmd.Flags().BoolVarP(&Plain, "plain", "p", false, "Print parsed tasks to stdout")
	pushCmd.Flags().StringVarP(&File, "file", "f", "", "Path to a file with tasks")
	pushCmd.MarkFlagFilename("file")
	pushCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push <file with tasks>",
	Short: "Push tasks from a file to provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.ReadFile(File)
		if err != nil {
			return fmt.Errorf("failed to read tasks file %q: %v", File, err)
		}
		parser := markdown.NewParser()
		t, err := parser.Parse(f)
		if err != nil {
			return fmt.Errorf("failed to parse file %q: %v", File, err)
		}

		jiraProvider, err := jira.NewJiraProvider(
			Config.Jira.URL,
			Config.Jira.Username,
			Config.Jira.Api_Token,
			Config.Jira.Project_Key,
		)
		plainProvider := plain.NewProvider(os.Stdout)

		providers := map[string]provider.Provider{
			"jira":  jiraProvider,
			"plain": plainProvider,
		}

		provider := "jira"
		if Plain {
			provider = "plain"
		}

		if err := providers[provider].Push(context.TODO(), t); err != nil {
			return fmt.Errorf("failed to push tasks from file %q: %v", File, err)
		}

		return nil
	},
}
