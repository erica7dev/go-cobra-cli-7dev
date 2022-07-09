package cobrao

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/erica7dev/cobrao/pkg/cobrao"
)

var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Aliases: []string{"rev"},
	Short: "Reverse a string",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := cobrao.Reverse(args[0])
		fmt.Println(res)
	},
}

func init(){
	rootCmd.AddCommand(reverseCmd)
}