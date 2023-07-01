package models

type ClientConfig struct {
	Log       Log              `json:"log"`
	Routing   Routing          `json:"routing"`
	Inbounds  []ClientInbound  `json:"inbounds"`
	Outbounds []ClientOutbound `json:"outbounds"`
}

type Log struct {
	Loglevel string `json:"loglevel"`
}

type Routing struct {
	DomainStrategy string `json:"domainStrategy"`
	Rules          []struct {
		Type        string   `json:"type"`
		Domain      []string `json:"domain,omitempty"`
		OutboundTag string   `json:"outboundTag"`
		IP          []string `json:"ip,omitempty"`
	} `json:"rules"`
}

type ClientInbound struct {
	Listen   string `json:"listen"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Settings *struct {
		UDP bool `json:"udp"`
	} `json:"settings,omitempty"`
	Sniffing struct {
		Enabled      bool     `json:"enabled"`
		DestOverride []string `json:"destOverride"`
	} `json:"sniffing"`
}

type ClientOutbound struct {
	Protocol string `json:"protocol"`
	Settings *struct {
		Vnext []ClientVnext `json:"vnext"`
	} `json:"settings,omitempty"`
	StreamSettings *struct {
		Network         string `json:"network"`
		Security        string `json:"security"`
		RealitySettings struct {
			Show        bool   `json:"show"`
			Fingerprint string `json:"fingerprint"`
			ServerName  string `json:"serverName"`
			PublicKey   string `json:"publicKey"`
			ShortID     string `json:"shortId"`
			SpiderX     string `json:"spiderX"`
		} `json:"realitySettings"`
	} `json:"streamSettings,omitempty"`
	Tag string `json:"tag"`
}

type ClientVnext struct {
	Address string       `json:"address"`
	Port    int          `json:"port"`
	Users   []ClientUser `json:"users"`
}

type ClientUser struct {
	ID         string `json:"id"`
	Flow       string `json:"flow"`
	Encryption string `json:"encryption"`
}

func (c *ClientConfig) FirstOutbound() *ClientOutbound {
	return &c.Outbounds[0]
}
