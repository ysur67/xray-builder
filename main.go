package main

import (
	"encoding/json"
	"fmt"
	"log"
	"xraybuilder/models"

	bashexecutor "xraybuilder/domain/commands/bash"
	clientservice "xraybuilder/domain/services/clients/impl"
	linuxService "xraybuilder/domain/services/osservice/linux"
	serverservice "xraybuilder/domain/services/server/impl"

	"github.com/alexflint/go-arg"
)

const InitialUserComment = "Initial user"

func sudoRequired() {
	log.Fatalln("must be run as superuser")
}

func main() {
	var args models.Args
	argParser := arg.MustParse(&args)

	cmdExecutor := bashexecutor.New(args.Verbose)

	if args.InstallMisc != nil {
		cmdExecutor.Shell("chmod +x shell/iptables.sh; shell/iptables.sh")
		cmdExecutor.Shell("chmod +x shell/enable-tcp-bbr.sh; shell/enable-tcp-bbr.sh")
		return
	}

	osService := linuxService.New(args.XrayConfigPath, args.XrayKeypairPath, cmdExecutor)
	isSuperUser, err := osService.IsSuperUser()
	if err != nil {
		panic(err)
	}

	if args.Setup != nil {
		if !isSuperUser {
			sudoRequired()
			return
		}

		Setup(osService, args.Setup)
		return
	}

	if args.User != nil {
		if args.User.Add != nil {
			if !isSuperUser {
				sudoRequired()
				return
			}

			AddClient(osService, args.User.Add)
			return
		}

		if args.User.List != nil {
			ListUsers(osService)
			return
		}
	}

	argParser.WriteHelp(log.Writer())
}

func Setup(osService *linuxService.LinuxOsService, args *models.SetupArgs) {
	clientService := clientservice.New(osService)
	serverService := serverservice.New()

	if args.InstallXray != "" {
		err := osService.DownloadAndInstallXray(args.InstallXray)
		if err != nil {
			panic(err)
		}
	}

	cfg, err := serverService.ReadConfig("")
	if err != nil {
		panic(err)
	}

	keyPair, err := osService.GenerateKeyPair()
	if err != nil {
		panic(err)
	}

	err = osService.SaveKeyPair(keyPair)
	if err != nil {
		panic(err)
	}

	client, err := clientService.CreateClient(InitialUserComment)
	if err != nil {
		panic(err)
	}

	serverService.InflateServerConfig(cfg, client, keyPair, args.Destination)
	clientConfig, err := clientService.CreateClientConfig(cfg.ServerName(), client, keyPair)
	if err != nil {
		panic(err)
	}
	osService.WriteConfigs(cfg, clientConfig, 0)
	if err = osService.RestartXray(); err != nil {
		panic(err)
	}
}

func AddClient(osService *linuxService.LinuxOsService, args *models.AddArgs) {
	clientService := clientservice.New(osService)
	serverService := serverservice.New()
	serverConfig, err := serverService.ReadConfig(osService.XrayConfigPath)
	if err != nil {
		panic(err)
	}
	client, err := clientService.CreateClient(args.Comment)
	if err != nil {
		panic(err)
	}
	serverService.AppendClient(serverConfig, client)
	keyPair, err := serverService.ReadKeyPair(osService.XrayKeypairPath)
	if err != nil {
		panic(err)
	}

	clientConfig, err := clientService.CreateClientConfig(serverConfig.ServerName(), client, keyPair)
	if err != nil {
		panic(err)
	}

	err = osService.WriteConfigs(
		serverConfig,
		clientConfig,
		len(*serverService.GetUsers(serverConfig)),
	)
	if err != nil {
		panic(err)
	}
}

func ListUsers(osService *linuxService.LinuxOsService) {
	serverService := serverservice.New()
	cfg, err := serverService.ReadConfig(osService.XrayConfigPath)
	if err != nil {
		panic(err)
	}

	users := serverService.GetUsers(cfg)
	result, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
