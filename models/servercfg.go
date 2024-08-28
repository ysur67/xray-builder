package models

import "github.com/google/uuid"

type LogObject struct {
	LogLevel string `json:"loglevel"`
}

type Rule struct {
	Type        string    `json:"type"`
	Domain      []string  `json:"domain,omitempty"`
	OutBoundTag string    `json:"outboundTag"`
	Ip          *[]string `json:"ip,omitempty"`
}

type InboundSettings struct {
	Clients    []Client `json:"clients"`
	Decryption string   `json:"decryption"`
}

type Client struct {
	Id   string `json:"id"`
	Flow string `json:"flow"`

	// Used to identify client in the `realitySettings.shortIds`.
	ShortId string `json:"shortId"`

	// Comment for the user.
	// Field name `email` is for compatibility with 3x-ui.
	Comment string `json:"email"`
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
	Settings       InboundSettings      `json:"settings"`
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
	System struct {
		StatsInboundUplink    bool `json:"statsInboundUplink,omitempty"`
		StatsInboundDownlink  bool `json:"statsInboundDownlink,omitempty"`
		StatsOutboundUplink   bool `json:"statsOutboundUplink,omitempty"`
		StatsOutboundDownlink bool `json:"statsOutboundDownlink,omitempty"`
	} `json:"system,omitempty"`
	Levels struct {
		Num0 struct {
			StatsUserUplink   bool `json:"statsUserUplink,omitempty"`
			StatsUserDownlink bool `json:"statsUserDownlink,omitempty"`
			Handshake         int  `json:"handshake"`
			ConnIdle          int  `json:"connIdle"`
		} `json:"0"`
	} `json:"levels"`
}

type ServerConfig struct {
	Log       LogObject     `json:"log"`
	Routing   RoutingObject `json:"routing"`
	Outbounds []Outbound    `json:"outbounds"`
	Policy    PolicyObject  `json:"policy"`
	Inbounds  []Inbound     `json:"inbounds"`
	Stats     struct{}      `json:"stats"`
	Api       ApiConfig     `json:"api,omitempty"`
}

type ApiConfig struct {
	Tag      string   `json:"tag,omitempty"`
	Listen   string   `json:"listen,omitempty"`
	Services []string `json:"services,omitempty"`
}

func (s *ServerConfig) FirstInbound() *Inbound {
	return &s.Inbounds[0]
}

func (s *ServerConfig) ServerName() string {
	return s.FirstInbound().StreamSettings.RealitySettings.ServerNames[0]
}

func NewClient(shortId string, comment string) *Client {
	return &Client{
		Comment: comment,
		ShortId: shortId,
		Id:      uuid.New().String(),
		Flow:    "xtls-rprx-vision",
	}
}
