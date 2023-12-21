package image

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/google/subcommands"

	"github.com/yunomu/blog/lib/image"
)

type Command struct {
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "image" }
func (c *Command) Synopsis() string { return "image resize" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	buf, err := image.Resize(os.Stdin)
	if err != nil {
		slog.Error("Reize", "err", err)
		return subcommands.ExitFailure
	}

	buf.WriteTo(os.Stdout)

	return subcommands.ExitSuccess
}
