package commands

type CmdExecutor interface {
	Execute(cmd string) (*string, error)
}
