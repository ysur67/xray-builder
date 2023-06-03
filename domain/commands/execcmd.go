package commands

type CmdExecutor interface {
	GenerateKeyPair() (*string, error)
	DownloadAndInstallXray(version string) error
	GenerateShortId() (*string, error)
	GetServerAddr() (*string, error)
	RestartXray() error
}
