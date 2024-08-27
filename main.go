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

const InitialUserComment = "initial-user"

func sudoMaybeRequired() {
	log.Println("Probably you need to run this command with sudo.")
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
			sudoMaybeRequired()
		}

		Setup(osService, args.Setup)
		return
	}

	if args.User != nil {
		if args.User.Add != nil {
			if !isSuperUser {
				sudoMaybeRequired()
			}

			AddClient(osService, args.User.Add)
			return
		}

		if args.User.List != nil {
			ListClients(osService)
			return
		}

		if args.User.Remove != nil {
			if !isSuperUser {
				sudoMaybeRequired()
			}

			RemoveClient(osService, args.User.Remove)
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

	serverService.SetupServer(cfg, keyPair, args.Destination)
	serverService.AppendClient(cfg, client)
	clientConfig, err := clientService.CreateClientConfig(cfg.ServerName(), client, keyPair)
	if err != nil {
		panic(err)
	}
	osService.WriteConfigs(cfg, clientConfig, 0)
	if err = osService.RestartXray(); err != nil {
		panic(err)
	}
}

func AddClient(osService *linuxService.LinuxOsService, args *models.UserAddArgs) {
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

func ListClients(osService *linuxService.LinuxOsService) {
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

func RemoveClient(osService *linuxService.LinuxOsService, args *models.UserRemoveArgs) {
	serverService := serverservice.New()
	cfg, err := serverService.ReadConfig(osService.XrayConfigPath)
	if err != nil {
		panic(err)
	}

	user := serverService.RemoveUser(cfg, args.IdOrComment)
	if user == nil {
		log.Fatalln("user not found")
		return
	}

	err = osService.WriteServerConfig(cfg)
	if err != nil {
		panic(err)
	}

	userJson, _ := json.MarshalIndent(user, "", "    ")
	fmt.Println(string(userJson))
}
