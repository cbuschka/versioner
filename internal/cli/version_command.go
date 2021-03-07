package cli

import (
	"github.com/cbuschka/versioner/internal/build"
)

type VersionCommandConfig struct {}

func (config *VersionCommandConfig) Run() error {
	buildInfo := build.GetBuildInfo()
	Print("%s", buildInfo.Version)
	return nil
}
