package cli

type LatestVersionConfig struct {
}

func (config *LatestVersionConfig) Run() error {

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
