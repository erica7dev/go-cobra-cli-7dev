package cobrao

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/erica7dev/cobrao/pkg/cobrao"
)

var onlyDigits bool
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Aliases: []string{"insp"},
	Short: "Inspect a string",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		res, kind := cobrao.Inspect(i, onlyDigits)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("%s has %d %s\n", i, res, kind+pluralS)
},
}

func init(){
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "only count digits")
	rootCmd.AddCommand(inspectCmd)
}