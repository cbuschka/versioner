package cli

import (
	"github.com/jpillora/opts"
	"os"
)

type Config struct {
	LatestVersion LatestVersionConfig  `opts:"mode=cmd,help=Get latest version from git tags."`
}

func Run() {
	args := os.Args
	if len(args) < 2 {
		args = []string{os.Args[0], "--help"}
	}

	config := Config{}
	opts.New(&config).ParseArgs(args).Run()
}
