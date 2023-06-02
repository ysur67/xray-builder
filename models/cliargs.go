package models

type AppendClientsArgs struct {
	ClientCount  int
	RedirectAddr string
}

type InstallArgs struct {
	ClientCount  int
	RedirectAddr string
	DownloadXray bool
	XrayVersion  string
}
