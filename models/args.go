package models

type Args struct {
	XrayConfigPath  string `arg:"-c,--config" default:"/usr/local/etc/xray/config.json"`
	XrayKeypairPath string `arg:"-k,--keypair" default:"/usr/local/etc/xray/keypair.json"`

	Verbose     bool       `arg:"-v,--verbose"`
	User        *UserArgs  `arg:"subcommand:user" help:"user management"`
	Setup       *SetupArgs `arg:"subcommand:setup" help:"install xray and generate configs"`
	InstallMisc *struct{}  `arg:"subcommand:install-misc" help:"Install additional iptables and TCP BBR configuration, suppress ssh MOTD"`
}

type UserArgs struct {
	Add  *AddArgs  `arg:"subcommand:add" help:"add user to xray config"`
	List *struct{} `arg:"subcommand:list"`
}

type AddArgs struct {
	Comment string `arg:"positional,required" help:"Amount of users in generated config"`
}

type SetupArgs struct {
	Destination string `arg:"-d,--destination required" placeholder:"rkn.gov.ru/" help:"destination and serverName in xray server config"`
	InstallXray string `arg:"-i,--install-xray" placeholder:"1.8.1" help:"Install xray-core by version"`
}

func (SetupArgs) Description() string {
	return `This script is used for installing, configuring
			an xray-core and generating user configs for it`
}
func (SetupArgs) Epilogue() string {
	return "For more information visit https://github.com/ysur67/xray-builder"
}

func (AddArgs) Description() string {
	return `This script is used for installing, configuring
			an xray-core and generating user configs for it`
}
func (AddArgs) Epilogue() string {
	return "For more information visit https://github.com/ysur67/xray-builder"
}
