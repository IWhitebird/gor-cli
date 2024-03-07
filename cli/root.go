package cli

import (
	"github.com/iwhitebird/Gor"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gor",
	Short: "Gor CLI",
	Long:  `Gor Language CLI application written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {

		if (len(args) == 0) || args[0] == "repl" {
			Gor.Repl()
		} else if args[0] == "run" {
			Gor.RunFromFile(args[1])
		}
	},
}

/*
Execute adds all child commands to the root command and sets flags appropriately.
This is called by main.main(). It only needs to happen once to the rootCmd.
*/
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
