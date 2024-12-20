package main

import (
	"encoding/json"
	"fmt"
	"log"
	"xray-builder/models"
	"xray-builder/qr"

	bashexecutor "xray-builder/domain/commands/bash"
	clientservice "xray-builder/domain/services/clients"
	linuxService "xray-builder/domain/services/osservice/linux"
	serverservice "xray-builder/domain/services/server"

	"github.com/alexflint/go-arg"
)

func sudoMaybeRequired() {
	log.Println("You may need to run this command with sudo.")
}

func main() {
	var args models.Args
	argParser := arg.MustParse(&args)

	cmdExecutor := bashexecutor.New(args.Verbose)
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

		if args.User.Disable != nil {
			if !isSuperUser {
				sudoMaybeRequired()
			}

			ToggleClientEnabled(osService, args.User.Disable, false)
		}

		if args.User.Enable != nil {
			if !isSuperUser {
				sudoMaybeRequired()
			}

			ToggleClientEnabled(osService, args.User.Enable, true)
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
		log.Fatalln("Xray restart failed. Please restart it manually.")
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

	if err = osService.RestartXray(); err != nil {
		log.Fatalln("Xray restart failed. Please restart it manually.")
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

func ToggleClientEnabled(osService *linuxService.LinuxOsService, args *models.UserIdentificationArgs, isEnabled bool) {
	serverService := serverservice.New()
	cfg, err := serverService.ReadConfig(osService.XrayConfigPath)
	if err != nil {
		panic(err)
	}

	user := serverService.ToggleUserEnabled(cfg, args.IdOrComment, isEnabled)
	if user == nil {
		fmt.Println("user not found")
		return
	}

	err = osService.WriteServerConfig(cfg)
	if err != nil {
		panic(err)
	}

	err = osService.RestartXray()
	if err != nil {
		log.Fatalln("Xray restart failed. Please restart it manually.")
	}
}

func RemoveClient(osService *linuxService.LinuxOsService, args *models.UserIdentificationArgs) {
	serverService := serverservice.New()
	cfg, err := serverService.ReadConfig(osService.XrayConfigPath)
	if err != nil {
		panic(err)
	}

	user := serverService.RemoveUser(cfg, args.IdOrComment)
	if user == nil {
		fmt.Println("user not found")
		return
	}

	err = osService.WriteServerConfig(cfg)
	if err != nil {
		panic(err)
	}

	userJson, _ := json.MarshalIndent(user, "", "    ")
	fmt.Println(string(userJson))

	err = osService.RestartXray()
	if err != nil {
		log.Fatalln("Xray restart failed. Please restart it manually.")
	}
}
