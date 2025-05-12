package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var name string
var shout bool

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet the user",
	Long: "This command is used to greet the user",

	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			name = "stranger"
		}

		greeting := fmt.Sprintf("Hello, %s", name)

		if shout {
			greeting = strings.ToUpper(greeting)
		}

		fmt.Println(greeting)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	greetCmd.Flags().StringVarP(&name, "name", "n", "", "Name to greet")
	greetCmd.Flags().BoolVarP(&shout, "shout", "s", false, "Shout the greeting in uppercase")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
