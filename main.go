package main

//todo: fix core library that tries to use non-commands to make calls to webos, or figure out why its getting 401 in the first place

import (
	"github.com/spf13/cobra"
	"pimview.thelabshack.com/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewPlugin().Execute())
}
