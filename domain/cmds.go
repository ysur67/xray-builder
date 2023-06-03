package domain

type Command int

const (
	DownloadCmd Command = iota
	GenerateKeyPair
	GetServerAddr
	GenerateShortId
)
