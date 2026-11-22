package lumna

import (
	"fmt"
	"runtime/debug"
)

var (
	Version   string
	Commit    string
	BuildDate string
)

func BuildVersion() (version, commit, date string) {
	version, commit, date = Version, Commit, BuildDate

	if info, ok := debug.ReadBuildInfo(); ok {
		if version == "" && info.Main.Version != "" && info.Main.Version != "(devel)" {
			version = info.Main.Version
		}
		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision":
				if commit == "" {
					commit = s.Value
					if len(commit) > 8 {
						commit = commit[:8]
					}
				}
			case "vcs.time":
				if date == "" {
					date = s.Value
				}
			}
		}
	}

	if version == "" {
		version = "dev"
	}
	if commit == "" {
		commit = "unknown"
	}
	if date == "" {
		date = "unknown"
	}
	return version, commit, date
}

func VersionString() string {
	version, commit, date := BuildVersion()
	return fmt.Sprintf("%s (commit %s, built %s)", version, commit, date)
}
