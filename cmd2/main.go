package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	cmd = &cobra.Command{
		Use:   "secret",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: RunWrapper,
	}
	envValues []string
	cmdName   string
)

func RunWrapper(cmd *cobra.Command, args []string) {
	for _, s := range envValues {
		strs := strings.SplitN(s, "=", 2)
		se := strings.SplitN(s, "|", 2)

		_, token, err := GetSecret(se[0], se[1])
		if err != nil {
			panic(err)
		}
		os.Setenv(strs[0], token)
	}
	ExecCommand(cmdName, args...)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cmd.Flags().StringArrayVarP(&envValues, "env-var", "v", []string{}, "vars")
	cmd.MarkFlagRequired("env-var")
	cmd.Flags().StringVarP(&cmdName, "command", "c", "", "command")
	cmd.MarkFlagRequired("command")
}

func main() {
	cmd.Execute()
}

func ExecCommand(command string, args ...string) error {
	// Create the command with our context
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
	}
	return nil
}
