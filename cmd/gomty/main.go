package main

import (
	_ "embed"
	"io"
	"os"
	"strings"

	"github.com/demetrius-oak/gomty"
	"github.com/spf13/cobra"
)

var opts *gomty.Options

func Transform() *cobra.Command {

	opts = new(gomty.Options)
	cmd := &cobra.Command{
		Use:   "gomty [file]",
		Short: "Transform html file to gomponents",
		Args:  cobra.MinimumNArgs(0),
		Run:   transform,
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "index", "Gomponent name")
	cmd.Flags().StringVarP(&opts.Package, "package", "p", "components", "Package name")
	cmd.Flags().StringVarP(&opts.Suffix, "suffix", "s", "Component", "Suffix name")

	return cmd
}

// Transform html
func transform(cmd *cobra.Command, args []string) {

	var reader io.Reader
	var err error
	reader = os.Stdin
	if len(args) > 0 {
		reader, err = os.Open(args[0])
		if err != nil {
			cmd.PrintErr(err)
		}
	}

	err = gomty.Transform(reader, os.Stdout, &gomty.Options{
		Suffix:  strings.Title(strings.ToLower(opts.Suffix)),
		Package: strings.ToLower(opts.Package),
		Name:    strings.Title(strings.ToLower(opts.Name)),
	})
	if err != nil {
		cmd.PrintErr(err)
	}
}

func main() {
	Transform().Execute()
}
