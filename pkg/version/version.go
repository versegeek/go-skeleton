package version

import (
	"fmt"
	"runtime"

	"github.com/versegeek/go-skeleton/pkg/json"
)

var (
	// GitVersion 是语义化的版本号.
	GitVersion = "v0.0.0-main+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是 Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出.
	GitCommit = "$Format:%H$"
	// GitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean, dirty.
	GitTreeState = ""
)

type (
	Version struct {
		GitVersion   string `json:"git_version"`
		GitCommit    string `json:"git_commit"`
		BuildDate    string `json:"build_date"`
		GitTreeState string `json:"git_tree_state"`
		GoVersion    string `json:"go_version"`
		Compiler     string `json:"compiler"`
		Platform     string `json:"platform"`
	}
)

func (version *Version) ToString() string {
	return fmt.Sprintf("GitVersion: %s\nGitCommit: %s\nBuildDate: %s\nGoVersion: %s\nCompiler: %s\nPlatform: %s\n",
		version.GitVersion, version.GitCommit, version.BuildDate, version.GoVersion, version.Compiler, version.Platform)
}

func (version *Version) ToJSON() string {
	jsonBytes, _ := json.Marshal(version)
	return string(jsonBytes)
}

func Get() *Version {
	return &Version{
		GitVersion: GitVersion,
		GitCommit:  GitCommit,
		BuildDate:  BuildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
