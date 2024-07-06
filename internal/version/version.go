package version

var (
	version    string = "0.0.3"
	revision   string = "unknown"
	prerelease string = "dev"
)

func Version() string {
	return version
}

func Revision() string {
	return revision
}

func Prerelease() string {
	return prerelease
}
