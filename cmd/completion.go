package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generates autocompletion script",
	Long: `
Generates autocompletion script for sendtg.

▶ Bash:
    sendtg completion bash > /etc/bash_completion.d/sendtg
    source /etc/bash_completion.d/sendtg

▶ Zsh:
    sendtg completion zsh > "${fpath[1]}/_sendtg"
    autoload -U compinit && compinit
    source ~/.zshrc

▶ Fish:
    sendtg completion fish | source

▶ PowerShell:
    sendtg completion powershell | Out-String | Invoke-Expression
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			return rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return fmt.Errorf("unknown shell type: %s", args[0])
		}
	},
}
