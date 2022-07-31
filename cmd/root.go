package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "built@home",
	Short: "Homelab as a bundle",
	Long: `An community supported homelab implementation for
	for increasing the ability to learn and develop with reducing the cost of time`,
	Run: func(cmd *cobra.Command, args []string) {
		flagVal, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return
		}
		if flagVal {
			fmt.Println("I like these words better!")
			return
		}
		fmt.Println("Welcome")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", false, "Toggle a different message")
}
