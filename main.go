package main

import (
	"encoding/json"
	"fmt"
	"log"
	"xraybuilder/models"
	"xraybuilder/qr"

	bashexecutor "xraybuilder/domain/commands/bash"
	clientservice "xraybuilder/domain/services/clients"
	linuxService "xraybuilder/domain/services/osservice/linux"
	serverservice "xraybuilder/domain/services/server"

	"github.com/alexflint/go-arg"
)

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

		if args.User.Remove != nil {
			if !isSuperUser {
				sudoMaybeRequired()
			}

			RemoveClient(osService, args.User.Remove)
			return
		}

		if args.User.List != nil {
			ListClients(osService)
			return
		}

		if args.User.Share != nil {
			Share(osService, args.User.Share)
			return
		}
	}

	argParser.WriteHelp(log.Writer())
}

func Share(osService *linuxService.LinuxOsService, args *models.ShareArgs) {
	serverService := serverservice.New()
	serverConfig, err := serverService.ReadConfig(osService.XrayConfigPath)

	if err != nil {
		panic(err)
	}

	client := serverService.GetUser(serverConfig, args.IdOrComment)
	if client == nil {
		log.Fatalln("user not found")
		return
	}

	keyPair, err := serverService.ReadKeyPair(osService.XrayKeypairPath)
	if err != nil {
		panic(err)
	}

	clientService := clientservice.New(osService)
	clientConfig, err := clientService.CreateClientConfig(serverConfig.ServerName(), client, keyPair)
	if err != nil {
		panic(err)
	}

	switch {
	case args.Format == "link":
		link := clientConfig.FirstOutbound().ShareLink(client.ShortId)
		fmt.Println(link.String())

	case args.Format == "qr":
		link := clientConfig.FirstOutbound().ShareLink(client.ShortId)
		qr.RenderString(link.String(), false)

	case args.Format == "json":
		result, _ := json.MarshalIndent(clientConfig, "", "    ")
		fmt.Println(string(result))
	}
}

func Setup(osService *linuxService.LinuxOsService, args *models.SetupArgs) {
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

	serverService.SetupServer(cfg, keyPair, args.Destination)
	osService.WriteServerConfig(cfg)
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

	err = osService.WriteServerConfig(serverConfig)
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
	result, _ := json.MarshalIndent(users, "", "    ")
	fmt.Println(string(result))
}

func RemoveClient(osService *linuxService.LinuxOsService, args *models.UserIdentificationArgs) {
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
