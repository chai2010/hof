package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/cmd/context"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
)

var contextLong = `get, set, and use contexts`

var ContextCmd = &cobra.Command{

	Use: "context",

	Short: "get, set, and use contexts",

	Long: contextLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},
}

func init() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	ohelp := ContextCmd.HelpFunc()
	ousage := ContextCmd.UsageFunc()
	help := func(cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		ohelp(cmd, args)
	}
	usage := func(cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		return ousage(cmd)
	}

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	ContextCmd.SetHelpFunc(thelp)
	ContextCmd.SetUsageFunc(tusage)

	ContextCmd.AddCommand(cmdcontext.GetCmd)
	ContextCmd.AddCommand(cmdcontext.SetCmd)
	ContextCmd.AddCommand(cmdcontext.UseCmd)
	ContextCmd.AddCommand(cmdcontext.SourceCmd)
	ContextCmd.AddCommand(cmdcontext.ClearCmd)

}
