package image

import (
	"context"
	"flag"
	"image"
	"log/slog"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/subcommands"
)

type Command struct {
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "image" }
func (c *Command) Synopsis() string { return "image utils" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	img, s, err := image.Decode(os.Stdin)
	if err != nil {
		slog.Error("image.Decode", "err", err)
		return subcommands.ExitFailure
	}

	slog.Info("decoded", "s", s, "rect", img.Bounds())

	return subcommands.ExitSuccess
}
