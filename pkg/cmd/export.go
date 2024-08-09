package cmd

import "github.com/spf13/cobra"

func ExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Print the command to export KUBECONFIG",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
