package version

import (
	"fmt"
	"os"
)

var (
	BuildVersion string
	BuildTime    string
	BuildName    string
)

func init() {
	args := os.Args
	if nil == args || len(args) < 2 {
		return
	}
	if "-v" == args[1] {
		fmt.Printf("%s: v%s (%s)\n", BuildName, BuildVersion, BuildTime)
	} else if "-h" == args[1] {
		fmt.Println("Usage:")
		fmt.Printf("./%s\n", BuildName)
		fmt.Printf("./%s -v\n", BuildName)
		fmt.Printf("./%s -h\n", BuildName)
	}
	os.Exit(0)
}
