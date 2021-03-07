package cli

import (
	"github.com/jpillora/opts"
	"os"
)

type Config struct {
	LatestVersion LatestVersionCommandConfig  `opts:"mode=cmd,help=Get latest version from git tags."`
	NextVersion NextVersionCommandConfig  `opts:"mode=cmd,help=Get next for latest version from git tags."`
	Version VersionCommandConfig  `opts:"mode=cmd,help=Get version of versioner."`
}

func Run() {
	args := os.Args
	if len(args) < 2 {
		args = []string{os.Args[0], "--help"}
	}

	config := Config{}
	opts.New(&config).ParseArgs(args).Run()
}
