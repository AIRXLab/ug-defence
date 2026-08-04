package main

import (
  "os"
  "github.com/airxlab/ug-defence/server/cmd/defence/commands"
)

func main() {
  if err := commands.Run(os.Args[1:]); err != nil {
    os.Exit(1)
  }
}
