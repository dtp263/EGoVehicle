package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// cobra.OnInitialize(utilities.ConfigureEnvironment)
}

var RootCmd = &cobra.Command{
	Use:   "bessie",
	Short: "A tool for running and managaing an electric vehicle.",
	Long: `bessie will get you there... This application provides
a flexible and extensible tool set for driving a dual motor
electric "vehicle".`,
}

func Execute() error {
	return RootCmd.Execute()
}
