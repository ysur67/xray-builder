package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"xraybuilder/internal"
	"xraybuilder/models"
	"xraybuilder/service/serverclients"
	"xraybuilder/service/serverconfig"
)

func main() {
	mode := os.Args[0]
	if mode == "" {
		RunInstall()
		return
		// fmt.Println("Select mode: install, add")
		// return
	}
	mode = strings.ToLower(mode)
	if mode == "install" {
		RunInstall()
		return
	}
	if mode == "add" {
		AddClients()
		return
	}
	RunInstall()
	return

}

func RunInstall() {
	args := ReadCreateArgs()
	if args.DownloadXray {
		err := internal.DownloadAndInstallXray(args)
		if err != nil {
			panic(err)
		}
	}
	_, err := serverclients.CreateClients(args.ClientCount)
	if err != nil {
		panic(err)
	}
	cfg, err := ReadServerConfig("")
	if err != nil {
		panic(err)
	}
	keyPair, err := internal.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	clients, err := serverclients.CreateClients(args.ClientCount)
	if err != nil {
		panic(err)
	}
	InflateServerConfig(cfg, clients, keyPair, args)
	clientConfigs := CreateClientConfigs(cfg, clients, keyPair, args)
	internal.WriteToFile("config.json", &cfg)
	for ind, elem := range *clientConfigs {
		internal.WriteToFile(fmt.Sprintf("client%v.json", ind), &elem)
	}
}

func InflateServerConfig(
	cfg *models.ServerConfig,
	clients *[]models.ClientDto,
	keyPair *models.KeyPair, args *models.InstallArgs,
) {
	serverconfig.AppendClients(
		cfg,
		clients,
		&cfg.FirstInbound().StreamSettings,
	)
	serverconfig.SetPrivateKey(cfg, keyPair)
	serverconfig.SetDestinationAddress(cfg, args.RedirectAddr)
}

func CreateClientConfigs(
	cfg *models.ServerConfig,
	clients *[]models.ClientDto,
	keyPair *models.KeyPair, args *models.InstallArgs,
) *[]models.ClientConfig {
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

func ReadCreateArgs() *models.InstallArgs {
	clients := flag.Int("n", 3, "Amount of clients to create")
	redirectAddress := flag.String("redir", "https://google.com", "Shadow address")
	downloadXray := flag.Bool("preload", false, "Preload Xray")
	xrayVersion := flag.String("version", "1.8.0", "Xray version, 1.8.0 default")
	return &models.InstallArgs{
		ClientCount:  *clients,
		RedirectAddr: *redirectAddress,
		DownloadXray: *downloadXray,
		XrayVersion:  *xrayVersion,
	}
}
