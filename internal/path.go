package internal

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func ResolveConfigPath() string {
	executablePath, _ := os.Executable()

	// go run
	if strings.HasPrefix(executablePath, "/tmp/go-build") {
		return "./configs"
	}

	return "/usr/local/etc/xray-builder"
}
