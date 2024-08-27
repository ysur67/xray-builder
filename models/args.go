package models

type Args struct {
	XrayConfigPath  string `arg:"-c,--config" default:"/usr/local/etc/xray/config.json"`
	XrayKeypairPath string `arg:"-k,--keypair" default:"/usr/local/etc/xray/keypair.json"`

	Verbose     bool       `arg:"-v,--verbose"`
	User        *UserArgs  `arg:"subcommand:user" help:"user management"`
	Setup       *SetupArgs `arg:"subcommand:setup" help:"install xray and generate configs"`
	InstallMisc *struct{}  `arg:"subcommand:install-misc" help:"Install additional iptables and TCP BBR configuration"`
}

type UserArgs struct {
	Add    *UserAddArgs    `arg:"subcommand:add" help:"add user to xray config"`
	Remove *UserRemoveArgs `arg:"subcommand:remove"`
	List   *struct{}       `arg:"subcommand:list"`
}

type UserAddArgs struct {
	Comment string `arg:"positional,required" help:"Amount of users in generated config"`
}

type UserRemoveArgs struct {
	IdOrComment string `arg:"positional,required" help:"Id or comment of removed user."`
}

type SetupArgs struct {
	Destination string `arg:"-d,--destination,required" placeholder:"rkn.gov.ru/" help:"destination and serverName in xray server config"`
	InstallXray string `arg:"-i,--install-xray" placeholder:"1.8.1" help:"Install xray-core by version"`
}

func (Args) Description() string {
	return "Configure xray-core and manage users.\nhttps://github.com/ysur67/xray-builder\n"
}
