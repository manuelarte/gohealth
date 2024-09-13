package info

// These variables are expected to be set by the LDFLAGS arguments
// LDFLAGS would be set at compile time by the CI/CD pipeline of any code
// which leverages this library
var (
	AppName        string // AppName - Application name
	AppDescription string // AppDescription - A brief description of what the application does
	AppVersion     string // AppVersion - The version of the application i.e. v1.0.0
	CommitAuthor   string // CommitAuthor - The username/email of the person who authored the commit
	CommitID       string // CommitID - The SHA1 checksum of the commit
	CommitTime     string // CommitTime - The time that the commit occurred
	BuildTime      string // BuildTime - Timestamp that the build occurred
	RepositoryUrl  string // RepositoryUrl - The URL where the repository is located at
	Branch         string // Branch - The branch the commit exists in
	Arch           string // Arch - The operating system architecture i.e. x86_64
	OS             string // OS - The OS which built the application i.e. darwin, linux
	RuntimeVersion string // RuntimeVersion - The version of Go which was used to build the application i.e. 1.17
)
