package daemonstatus

import (
	"github.com/Wil3on/nordvik_gameap_gameapctl/pkg/daemon"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func Handle(cliCtx *cli.Context) error {
	err := daemon.Status(cliCtx.Context)
	if err != nil {
		return errors.WithMessage(err, "failed to get daemon status")
	}

	return nil
}
