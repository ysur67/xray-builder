package models

import "github.com/google/uuid"

type LogObject struct {
	LogLevel string `json:"logLevel"`
}

type Rule struct {
	Type        string   `json:"type"`
	Domain      []string `json:"domain"`
	OutBoundTag string   `json:"outboundTag"`
	Ip          []string `json:"ip"`
}

type InboudSettings struct {
	Clients    []Client `json:"clients"`
	Decryption string   `json:"decryption"`
}

type Client struct {
	ID   string `json:"id"`
	Flow string `json:"flow"`
}

type RealitySettingsObject struct {
	Show         bool     `json:"show"`
	Dest         string   `json:"dest"`
	Xver         int      `json:"xver"`
	ServerNames  []string `json:"serverNames"`
	PrivateKey   string   `json:"privateKey"`
	MinClientVer string   `json:"minClientVer"`
	MaxClientVer string   `json:"maxClientVer"`
	MaxTimeDiff  int      `json:"maxTimeDiff"`
	ShortIds     []string `json:"shortIds"`
}

type StreamSettingsObject struct {
	Network         string                `json:"network"`
	Security        string                `json:"security"`
	RealitySettings RealitySettingsObject `json:"realitySettings"`
}

type SniffingObject struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
}

type Inbound struct {
	Listen         string               `json:"listen"`
	Port           int                  `json:"port"`
	Protocol       string               `json:"protocol"`
	Settings       InboudSettings       `json:"settings"`
	StreamSettings StreamSettingsObject `json:"streamSettings"`
	Sniffing       SniffingObject       `json:"sniffing"`
}

type RoutingObject struct {
	DomainStrategy string `json:"domainStrategy"`
	Rules          []Rule `json:"rules"`
}

type Outbound struct {
	Protocol string `json:"protocol"`
	Tag      string `json:"tag"`
}

type PolicyObject struct {
	Levels struct {
		Num0 struct {
			Handshake int `json:"handshake"`
			ConnIdle  int `json:"connIdle"`
		} `json:"0"`
	} `json:"levels"`
}

type ServerConfig struct {
	Log       LogObject     `json:"log"`
	Routing   RoutingObject `json:"routing"`
	OutBounds []Outbound    `json:"outbounds"`
	Policy    PolicyObject  `json:"policy"`
	Inbounds  []Inbound     `json:"inbounds"`
}

func (s *ServerConfig) FirstInbound() *Inbound {
	return &s.Inbounds[0]
}

type ClientDto struct {
	Client  Client
	ShortId string
}

func NewClient(shortId string) *ClientDto {
	return &ClientDto{
		Client: Client{
			ID:   uuid.New().String(),
			Flow: "xtls-rprx-vision",
		},
		ShortId: shortId,
	}
}
