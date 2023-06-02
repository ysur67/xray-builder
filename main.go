package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
		fmt.Println("Select mode: install, add")
		return
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
	_, err = internal.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	clients, err := serverclients.CreateClients(args.ClientCount)
	if err != nil {
		panic(err)
	}
	serverconfig.AppendClients(
		cfg,
		clients,
		&cfg.Inbounds[0].StreamSettings,
	)
}

func AddClients() {

}

func ReadServerConfig(path string) (*models.ServerConfig, error) {
	if path == "" {
		path = "server.template.json"
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := models.ServerConfig{}
	err = json.Unmarshal([]byte(file), &config)
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
