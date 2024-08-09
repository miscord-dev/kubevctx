package cmd

import (
	"fmt"

	"github.com/miscord-dev/kubevctx/pkg/vctx/hook"
	"github.com/spf13/cobra"
)

func HookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hook bash|zsh|fish",
		Short: "Print the hook for shell",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "zsh":
				fmt.Print(hook.ZSH())
			default:
				return fmt.Errorf("unsupported shell: %s", args[0])
			}

			return nil
		},
	}

	return cmd
}
