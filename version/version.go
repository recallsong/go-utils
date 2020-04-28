package version

import (
	"fmt"
	"os"
)

var (
	// 版本号
	Version string
	// 编译时间
	BuildTime string
	// 编译时间
	GoVersion string
	// 版本库中的提交版本
	CommitID string
)

// String 返回版本信息
func String() string {
	return fmt.Sprintf("Version: %s\nBuildTime: %s\nGoVersion: %s\nCommitID: %s\n",
		Version, BuildTime, GoVersion, CommitID)
}

// Print 打印版本信息
func Print() {
	fmt.Print(String())
}

// PrintIfCommand 如果命令行有version参数，则打印并退出
func PrintIfCommand() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		Print()
		os.Exit(0)
	}
}
