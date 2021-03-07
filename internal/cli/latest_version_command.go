package cli

type LatestVersionCommandConfig struct {
}

func (config *LatestVersionCommandConfig) Run() error {

	latestVersion, err := getLatestVersion()
	if err != nil {
		return err
	}

	if latestVersion != nil {
		Print("%s", latestVersion.Original())
	} else {
		Debug("No versions found.")
	}

	return nil
}
