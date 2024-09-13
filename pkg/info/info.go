package info

import (
	"encoding/json"
	"net/http"
)

type AppInfo struct {
	Name        string `json:"name,omitempty"`        // Name - Application name
	Description string `json:"description,omitempty"` // Description - A brief description of what the application does
	Version     string `json:"version,omitempty"`     // AppVersion - The version of the application i.e. v1.0.0
}

type GitInfo struct {
	CommitAuthor  string `json:"commitAuthor,omitempty"` // CommitAuthor - The username/email of the person who authored the commit
	CommitID      string `json:"commitId,omitempty"`     // CommitID - The SHA1 checksum of the commit
	CommitTime    string `json:"commitTime,omitempty"`   // CommitTime - The time that the commit occurred
	BuildTime     string `json:"buildTime,omitempty"`    // BuildTime - Timestamp when the build occurred
	RepositoryUrl string `json:"url,omitempty"`          // RepositoryUrl - The URL where the repository is located at
	Branch        string `json:"name,omitempty"`         // Branch - The branch the commit exists in
}

type RuntimeInfo struct {
	Arch           string `json:"arch,omitempty"`    // Arch - The operating system architecture i.e. x86_64
	OS             string `json:"os,omitempty"`      // OS - The OS which built the application i.e. darwin, linux
	RuntimeVersion string `json:"version,omitempty"` // RuntimeVersion - The version of Go which was used to build the application i.e. 1.17
}

type Info struct {
	Application AppInfo     `json:"app,omitempty"`     // Application - A nested field which contains all of the application information
	Git         GitInfo     `json:"git,omitempty"`     // Git - A nested struct which contains all information related to the Git repository
	Runtime     RuntimeInfo `json:"runtime,omitempty"` // Runtime - Information about the runtime environment
}

type Payload[T interface{}] func(info Info) T

func HandleInfo(writer http.ResponseWriter, req *http.Request) {
	HandleCustomInfo[Info](writer, req)
}

// HandleInfo - handles compiling the info provided by flags into a JSON string and
// writes the string out to the response
func HandleCustomInfo[T interface{}](writer http.ResponseWriter, _ *http.Request, payloadFuncs ...Payload[T]) {
	// Prepare the response payload (default implementation)
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

	var output interface{} = payload
	for _, payloadFunc := range payloadFuncs {
		output = payloadFunc(payload)
	}

	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(output)
}
