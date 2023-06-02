package models

type ClientInboundSettings struct {
	Udp bool
}

type ClientInbound struct {
	Listen   string         `json:"listen"`
	Port     int            `json:"port"`
	Protocol string         `json:"protocol,omitempty"`
	Settings InboudSettings `json:"settings"`
	Sniffing SniffingObject `json:"sniffing"`
}

type ClientUser struct {
	ID         string `json:"id"`
	Flow       string `json:"flow"`
	Encryption string `json:"encryption"`
}

type ClientVnext struct {
	Address string       `json:"address"`
	Port    int          `json:"port"`
	Users   []ClientUser `json:"users"`
}

type ClientOutboundSettings struct {
	Vnext ClientVnext
}

type ClientRealitySettings struct {
	Show        bool   `json:"show"`
	Fingerprint string `json:"fingerprint"`
	ServerName  string `json:"serverName"`
	PublicKey   string `json:"publicKey"`
	ShortID     string `json:"shortId"`
	SpiderX     string `json:"spiderX"`
}

type ClientStreamSettings struct {
	Network         string                `json:"network"`
	Security        string                `json:"security"`
	RealitySettings ClientRealitySettings `json:"realitySettings"`
}

type ClientOutbound struct {
	Protocol       string                 `json:"protocol"`
	Settings       ClientOutboundSettings `json:"settings,omitempty"`
	StreamSettings ClientStreamSettings   `json:"streamSettings,omitempty"`
	Tag            string                 `json:"tag"`
}

type ClientConfig struct {
	Log       LogObject        `json:"log"`
	Routing   RoutingObject    `json:"routing"`
	Inbounds  []Inbound        `json:"inbounds"`
	Outbounds []ClientOutbound `json:"outbounds"`
}

func (c *ClientConfig) FirstOutbound() *ClientOutbound {
	return &c.Outbounds[0]
}
