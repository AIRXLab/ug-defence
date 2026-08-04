package commands

import (
  "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
  Use:   "help",
  Short: "defence system",
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}


func Run(args []string) error {
  RootCmd.SetArgs(args)
  return RootCmd.Execute()
}