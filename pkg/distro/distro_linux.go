// +build linux freebsd openbsd netbsd
// +build !windows
// +build !darwin

package distro

import (
	"gitlab.com/mjwhitta/sysinfo"
)

var Info = sysinfo.New()

func Vendor() string {
	return "linux"
}

func Name() string {
	return Info.OS
}
