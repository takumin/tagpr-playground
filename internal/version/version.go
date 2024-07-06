package version

var (
	version    string = "0.1.0"
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
