package cmd

import (
	"github.com/code-ready/crc/pkg/crc/logging"
	"github.com/code-ready/crc/pkg/crc/output"

	"github.com/YourFin/binappend"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List data files embedded in the crc binary",
	Long:  `List all the data files which were embedded in the crc binary`,
	Run: func(cmd *cobra.Command, args []string) {
		runList(args)
	},
}

func runList(args []string) {
	if len(args) != 1 {
		logging.Fatalf("list takes exactly one argument")
	}
	binaryPath := args[0]
	extractor, err := binappend.MakeExtractor(binaryPath)
	if err != nil {
		logging.Fatalf("Could not access data embedded in %s: %v", binaryPath, err)
	}
	output.Outf("Data files embedded in %s:\n", binaryPath)
	for _, name := range extractor.AvalibleData() {
		output.Outln("\t", name)
	}
}
