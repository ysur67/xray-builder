package main

import (
	"fmt"
	"os"
	"xraybuilder/internal"
	"xraybuilder/models"

	bashexecutor "xraybuilder/domain/commands/bash"
	clientservice "xraybuilder/domain/services/clients/impl"
	"xraybuilder/domain/services/osservice/linux"
	serverservice "xraybuilder/domain/services/server/impl"

	"github.com/alexflint/go-arg"
)

func main() {
	mode := os.Args[1]
	os.Args = internal.RemoveByIndex(os.Args, 1)
	if mode == "create" {
		RunInstall()
		return
	}

	if mode == "add" {
		AddClients()
		return
	}

	fmt.Println(`Select one of the commands "add" or "create".`)
}

func RunInstall() {
	var args models.InstallArgs
	arg.MustParse(&args)

	osService := linux.NewLinuxOsService(bashexecutor.NewBashExecutor())
	clientService := clientservice.NewClientCfgServiceImpl(osService)
	serverService := serverservice.NewServerServiceImpl()

	isSuperUser, err := osService.IsSuperUser()
	if err != nil {
		panic(err)
	}

	if !isSuperUser {
		fmt.Println("Must be run as superuser")
		return
	}

	if args.InstallMisc {
		err := osService.SuppressLoginMessage()
		if err != nil {
			panic(err)
		}
		err = osService.ApplyIptablesRules()
		if err != nil {
			panic(err)
		}
		err = osService.EnableTcpBBR()
		if err != nil {
			panic(err)
		}
	}

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
	if err = osService.SaveKeyPair(keyPair); err != nil {
		panic(err)
	}
	clients, err := clientService.CreateClients(args.UsersCount)
	if err != nil {
		panic(err)
	}
	serverService.InflateServerConfig(cfg, clients, keyPair, args.Destination)
	clientConfigs, err := clientService.CreateMultipleConfigs(cfg.ServerName(), clients, keyPair)
	if err != nil {
		panic(err)
	}
	osService.WriteConfigs(cfg, clientConfigs, 0)
	if err = osService.RestartXray(); err != nil {
		panic(err)
	}
}

func AddClients() {
	var args models.AddArgs
	arg.MustParse(&args)

	if args.Add < 1 {
		fmt.Println("The number of users must be greater than 0")
		return
	}

	osService := linux.NewLinuxOsService(bashexecutor.NewBashExecutor())

	isSuperUser, err := osService.IsSuperUser()
	if err != nil {
		panic(err)
	}

	if !isSuperUser {
		fmt.Println("Must be run as superuser")
		return
	}

	clientService := clientservice.NewClientCfgServiceImpl(osService)
	serverService := serverservice.NewServerServiceImpl()
	serverConfig, err := serverService.ReadConfig(internal.LinuxConfigPath)
	if err != nil {
		panic(err)
	}
	usersCount := serverService.CurrentUsers(serverConfig)
	clients, err := clientService.CreateClients(args.Add)
	if err != nil {
		panic(err)
	}
	serverService.AppendClients(serverConfig, clients, &serverConfig.FirstInbound().StreamSettings)
	keyPair, err := serverService.ReadKeyPair(internal.LinuxKeyPairPath)
	if err != nil {
		panic(err)
	}
	clientConfigs, err := clientService.CreateMultipleConfigs(serverConfig.ServerName(), clients, keyPair)
	if err != nil {
		panic(err)
	}
	osService.WriteConfigs(serverConfig, clientConfigs, usersCount)
}
