package internal

import (
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

const LinuxConfigPath = "/usr/local/etc/xray/config.json"
const LinuxKeyPairPath = "/usr/local/etc/xray/keypair.json"
