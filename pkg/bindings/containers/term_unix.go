//go:build !windows
// +build !windows

package containers

import (
	"context"
	"os"
	"os/signal"

	sig "github.com/hanks177/podman/v4/pkg/signal"
	"golang.org/x/term"
)

func makeRawTerm(stdin *os.File) (*term.State, error) {
	return term.MakeRaw(int(stdin.Fd()))
}

func notifyWinChange(ctx context.Context, winChange chan os.Signal, stdin *os.File, stdout *os.File) {
	signal.Notify(winChange, sig.SIGWINCH)
}

func getTermSize(stdin *os.File, stdout *os.File) (width, height int, err error) {
	return term.GetSize(int(stdin.Fd()))
}
