package assets

import "fmt"

var currentVersion = "dev"

var currentCommitHash = "fooooo"

var softwareName = "Baloo"

func GenName() string {
	return fmt.Sprintf("%s %s(%s)", softwareName, currentVersion, currentCommitHash)
}
