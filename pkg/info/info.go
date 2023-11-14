package info

import (
	"encoding/json"
	"net/http"
)

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

type AppInfo struct {
	Name        string `json:"name"`        // Name - Application name
	Description string `json:"description"` // Description - A brief description of what the application does
	Version     string `json:"version"`     // AppVersion - The version of the application i.e. v1.0.0
}

type GitInfo struct {
	CommitAuthor  string `json:"commit_author"` // CommitAuthor - The username/email of the person who authored the commit
	CommitID      string `json:"commit_id"`     // CommitID - The SHA1 checksum of the commit
	CommitTime    string `json:"commit_time"`   // CommitTime - The time that the commit occurred
	BuildTime     string `json:"build_time"`    // BuildTime - Timestamp when the build occurred
	RepositoryUrl string `json:"url"`           // RepositoryUrl - The URL where the repository is located at
	Branch        string `json:"name"`          // Branch - The branch the commit exists in
}

type RuntimeInfo struct {
	Arch           string `json:"arch"`    // Arch - The operating system architecture i.e. x86_64
	OS             string `json:"os"`      // OS - The OS which built the application i.e. darwin, linux
	RuntimeVersion string `json:"version"` // RuntimeVersion - The version of Go which was used to build the application i.e. 1.17
}

type Info struct {
	Application AppInfo     `json:"app"`     // Application - A nested field which contains all of the application information
	Git         GitInfo     `json:"git"`     // Git - A nested struct which contains all information related to the Git repository
	Runtime     RuntimeInfo `json:"runtime"` // Runtime - Information about the runtime environment
}

// HandleInfo - handles compiling the info provided by flags into a JSON string and
// writes the string out to the response
func HandleInfo(writer http.ResponseWriter, _ *http.Request) {
	// Prepare the response payload
	payload := Info{
		Application: AppInfo{
			Name:        AppName,
			Description: AppDescription,
			Version:     AppVersion,
		},
		Git: GitInfo{
			CommitAuthor:  CommitAuthor,
			CommitID:      CommitID,
			Branch:        Branch,
			RepositoryUrl: RepositoryUrl,
			CommitTime:    CommitTime,
			BuildTime:     BuildTime,
		},
		Runtime: RuntimeInfo{
			Arch:           Arch,
			OS:             OS,
			RuntimeVersion: RuntimeVersion,
		},
	}

	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(payload)
}
