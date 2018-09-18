package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "",
	Long:  `示例命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("example called")
	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}
