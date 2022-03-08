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
		Use:   "my-calc",
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
)

func RunWrapper(cmd *cobra.Command, args []string) {
	// elementMap := make(map[string]string)
	for _, s := range envValues {
		strs := strings.SplitN(s, "=", 2)
		se := strings.SplitN(s, "|", 2)
		// elementMap[strs[0]] = strs[1]

		_, token, err := GetSecret(se[0], se[1])
		if err != nil {
			panic(err)
		}

		fmt.Println(token)
		os.Setenv(strs[0], token)
	}
	fmt.Println(os.Environ())
	// fmt.Print(elementMap)
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
	// cobra.OnInitialize(initConfig)

	// // Here you will define your flags and configuration settings.
	// // Cobra supports persistent flags, which, if defined here,
	// // will be global for your application.
	// // fmt.Println("flag")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-calc.yaml)")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	cmd.Flags().StringArrayVarP(&envValues, "env-var", "v", []string{}, "vars")
	cmd.MarkFlagRequired("env-var")
}

// func main() {

// 	var cmdName = flag.StringP("commasnd", "c", "", "cmd message for flag n")
// 	var nFlag = flag.StringP("tokens", "t", "", "help message for flag n")

// 	flag.Parse()

// 	fmt.Println(*nFlag)
// 	fmt.Println(*cmdName)
// 	fmt.Println("--")
// 	fmt.Println(flag.Args())

// 	_, token, err := GetSecret(serverURL, "github_token")
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	// fmt.Println(token)
// 	os.Setenv("GITHUB_TOKEN", token)
// 	ExecCommand(os.Args[1], os.Args[2:]...)
// }

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
