package api

import (
	"github.com/ipfs/go-ipfs-api"
)

func NewShell(address string) *shell.Shell {
	return shell.NewShell(address)
}
