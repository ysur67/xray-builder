package models

type AddArgs struct {
	Add int `arg:"-u, --users required" help:"Amount of users in generated config"`
}

type InstallArgs struct {
	Destination string `arg:"-d,--destination required" placeholder:"https://rkn.gov.ru/" help:"destination and serverName in xray server config"`
	InstallXray string `arg:"-i,--install-xray" placeholder:"1.8.1" help:"Is installation of xray-core reguired, also you need to specify the version"`
	UsersCount  int    `arg:"-u, --users" help:"Amount of users in generated config"`
	InstallMisc bool   `arg:"-m,--misc" help:"Additional iptables and TCP BBR configuration, see README.md for additional information"`
}

func (InstallArgs) Description() string {
	return `This script is used for installing, configuring
			an xray-core and generating user configs for it`
}
func (InstallArgs) Epilogue() string {
	return "For more information visit https://github.com/ysur67/xray-builder"
}

func (AddArgs) Description() string {
	return `This script is used for installing, configuring
			an xray-core and generating user configs for it`
}
func (AddArgs) Epilogue() string {
	return "For more information visit https://github.com/ysur67/xray-builder"
}
