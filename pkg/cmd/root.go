package cmd

import "github.com/spf13/cobra"

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vctx",
		Short: "kubevctx is a tool for managing virtual contexts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(HookCommand())
	cmd.AddCommand(ExportCommand())

	return cmd
}
