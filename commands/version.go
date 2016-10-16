package commands

import (
	"fmt"
	"runtime"

	"github.com/centurylinkcloud/clc-go-cli/base"
)

var banner = `
-------------------------------------------------------------

   _____           __                    __    _        __
  / ___/___  ___  / /_ __ __ ____ __ __ / /   (_)___   / /__
 / /__ / -_)/ _ \/ __// // // __// // // /__ / // _ \ /  '_/
 \___/ \__//_//_/\__/ \_,_//_/   \_, //____//_//_//_//_/\_\
                                /___/

-------------------------------------------------------------
`

type Version struct {
	CommandBase
}

func NewVersion(info CommandExcInfo) *Version {
	v := Version{}
	v.ExcInfo = info
	return &v
}

func (v *Version) IsOffline() bool {
	return true
}

func (v *Version) ExecuteOffline() (string, error) {
	fmt.Printf("%s", banner)
	fmt.Printf("CenturyLink Cloud CLI (Version %s)\n", base.BuildVersion)
	fmt.Printf("%s\n", base.PROJ_URL)
	fmt.Printf("\n")
	fmt.Printf("Go Version: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Built on: %s\n", base.BuildDate)
	fmt.Printf("Git Commit: %s\n", base.BuildGitCommit)
	fmt.Printf("\n")
	fmt.Printf("For more information on CenturyLink Cloud visit: %s\n", base.CTL_URL)

	return "", nil
}
