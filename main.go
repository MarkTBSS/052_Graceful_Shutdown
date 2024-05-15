package main

import (
	"os"

	_pkgConfig "github.com/MarkTBSS/052_Graceful_Shutdown/config"
	_pkgModulesServers "github.com/MarkTBSS/052_Graceful_Shutdown/modules/servers"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := _pkgConfig.LoadConfig(envPath())
	_pkgModulesServers.NewServer(cfg).Start()
}
