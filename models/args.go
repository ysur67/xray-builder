package models

type Args struct {
	XrayConfigPath  string `arg:"-c,--config" default:"/usr/local/etc/xray/config.json"`
	XrayKeypairPath string `arg:"-k,--keypair" default:"/usr/local/etc/xray/keypair.json"`
	Verbose         bool   `arg:"-v,--verbose"`

	User  *UserArgs  `arg:"subcommand:user" help:"user management"`
	Setup *SetupArgs `arg:"subcommand:setup" help:"install xray and generate configs"`
}

type UserArgs struct {
	Add     *UserAddArgs            `arg:"subcommand:add" help:"add user to xray config"`
	Remove  *UserIdentificationArgs `arg:"subcommand:remove"`
	Enable  *UserIdentificationArgs `arg:"subcommand:enable"`
	Disable *UserIdentificationArgs `arg:"subcommand:disable"`
	Share   *ShareArgs              `arg:"subcommand:share"`
	List    *struct{}               `arg:"subcommand:list"`
}

type UserAddArgs struct {
	Comment string `arg:"positional,required" help:"Comment of the new user."`
}

type UserIdentificationArgs struct {
	IdOrComment string `arg:"positional,required" help:"Id or comment of removed user."`
}

type ShareArgs struct {
	UserIdentificationArgs

	Format string `arg:"-f,--format" default:"qr" help:"qr | json | link"`
}

type SetupArgs struct {
	Destination string `arg:"-d,--destination,required" placeholder:"rkn.gov.ru/" help:"destination and serverName in xray server config"`
}

func (Args) Description() string {
	return "Configure xray-core and manage users.\nhttps://github.com/ysur67/xray-builder\n"
}
