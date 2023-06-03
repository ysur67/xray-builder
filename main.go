package main

import (
	"fmt"
	"os"
	"strings"
	bashexecutor "xraybuilder/domain/commands/bash"
	bashservice "xraybuilder/domain/services/bash"
	"xraybuilder/internal"
	"xraybuilder/models"
	"xraybuilder/service/serverclients"
	"xraybuilder/service/serverconfig"

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
	return
}

func RunInstall() {
	var args models.InstallArgs
	arg.MustParse(&args)

	osService := bashservice.NewBashOsService(bashexecutor.NewBashExecutor())

	if args.InstallXray != "" {
		err := osService.DownloadAndInstallXray(args.InstallXray)
		if err != nil {
			panic(err)
		}
	}
	_, err := serverclients.CreateClients(args.UsersCount)
	if err != nil {
		panic(err)
	}
	cfg, err := ReadServerConfig("")
	if err != nil {
		panic(err)
	}
	keyPair, err := osService.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	clients, err := serverclients.CreateClients(args.UsersCount)
	if err != nil {
		panic(err)
	}
	InflateServerConfig(cfg, clients, keyPair, args.Destination)
	clientConfigs := CreateClientConfigs(cfg, clients, keyPair)
	internal.WriteToFile("config.json", &cfg)
	for ind, elem := range *clientConfigs {
		internal.WriteToFile(fmt.Sprintf("client%v.json", ind), &elem)
	}
}

func InflateServerConfig(
	cfg *models.ServerConfig,
	clients *[]models.ClientDto,
	keyPair *models.KeyPair,
	destination string,
) {
	serverconfig.AppendClients(
		cfg,
		clients,
		&cfg.FirstInbound().StreamSettings,
	)
	serverconfig.SetPrivateKey(cfg, keyPair)
	serverconfig.SetDestinationAddress(cfg, destination)
}

func CreateClientConfigs(
	cfg *models.ServerConfig,
	clients *[]models.ClientDto,
	keyPair *models.KeyPair) *[]models.ClientConfig {
	result := make([]models.ClientConfig, len(*clients))
	for ind, elem := range *clients {
		clientConfig := models.ClientConfig{}
		internal.ReadJson("client.template.json", &clientConfig)
		result[ind] = *serverclients.CreateClientConfig(cfg, &elem, keyPair)
	}
	return &result
}

func AddClients() {

}

func ReadServerConfig(path string) (*models.ServerConfig, error) {
	if path == "" {
		path = "server.template.json"
	}
	config := models.ServerConfig{}
	err := internal.ReadJson(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
