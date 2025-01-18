package main

import (
	"fmt"
	"os"

	"github.com/dungvan/mailstation/account/cmd/idp"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "account",
	Short: "account application",
	Long:  `mailstation tool for account service.`,
}

func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.AddCommand(idp.ListIDPsCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
