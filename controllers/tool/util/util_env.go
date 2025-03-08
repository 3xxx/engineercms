package util

import (
	"logs"
	"os"
	"os/user"
	"strings"
)

// whether windows develop environment
func EnvWinDevelopment() bool {
	ex, err := os.Executable()
	if err != nil {
		// panic(err)
		logs.Error(err)
	}

	//if exPath contains \\AppData\\Local\\Temp we regard as dev.
	systemUser, err := user.Current()
	if systemUser != nil {
		return strings.HasPrefix(ex, systemUser.HomeDir+"\\AppData\\Local\\Temp")
	}
	return false
}

// whether mac develop environment
func EnvMacDevelopment() bool {
	ex, err := os.Executable()
	if err != nil {
		// panic(err)
		logs.Error(err)
	}
	return strings.HasPrefix(ex, "/private/var/folders")
}

// whether develop environment (whether run in IDE)
func EnvDevelopment() bool {
	return EnvWinDevelopment() || EnvMacDevelopment()
}
