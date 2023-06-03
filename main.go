package main

import (
	"os"
	"strings"
	"xraybuilder/models"

	bashexecutor "xraybuilder/domain/commands/bash"
	clientservice "xraybuilder/domain/services/clients/impl"
	"xraybuilder/domain/services/osservice/linux"
	serverservice "xraybuilder/domain/services/server/impl"

	"github.com/alexflint/go-arg"
)

func main() {
	mode := os.Args[0]

	if mode == "install" {
		RunInstall()
		return
	}

	if mode == "add" {
		AddClients()
		return
	}

	os.Args = strings.Split("--help", " ")
	var args models.InstallArgs
	arg.MustParse(&args)
}

func RunInstall() {
	var args models.InstallArgs
	arg.MustParse(&args)

	osService := linux.NewLinuxOsService(bashexecutor.NewBashExecutor())
	clientService := clientservice.NewClientCfgServiceImpl(osService)
	serverService := serverservice.NewServerServiceImpl()

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
	clients, err := clientService.CreateClients(args.UsersCount)
	if err != nil {
		panic(err)
	}
	serverService.InflateServerConfig(cfg, clients, keyPair, args.Destination)
	clientConfigs, err := clientService.CreateMultipleConfigs(cfg.ServerName(), clients, keyPair)
	if err != nil {
		panic(err)
	}
	osService.WriteConfigs(cfg, clientConfigs)
	if err = osService.RestartXray(); err != nil {
		panic(err)
	}
}

func AddClients() {

}
