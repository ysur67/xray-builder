package models

import (
	"fmt"
	"net/url"
)

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
	Id         string `json:"id"`
	Flow       string `json:"flow"`
	Encryption string `json:"encryption"`
	Comment    string `json:"comment"`
}

// https://github.com/XTLS/libXray/blob/ebd5526e3ad531ab184fffdaf281512084e8de26/nodep/generate_share.go#L40
// https://github.com/2dust/v2rayNG/blob/b8939763d4f67d5c28950d18da3ced9da88bbb6d/V2rayNG/app/src/main/kotlin/com/v2ray/ang/util/fmt/VlessFmt.kt#L22

func (outbound *ClientOutbound) ShareLink(shortId string) url.URL {
	if outbound.Protocol != "vless" {
		panic("only vless is supported")
	}

	shareUrl := outbound.vlessLink()
	outbound.addStreamSettingsQueryParamsTo(&shareUrl)

	query := shareUrl.Query()
	addQueryIfNonEmpty(&query, "sid", shortId)
	shareUrl.RawQuery = query.Encode()

	return shareUrl
}

func (outbound *ClientOutbound) vlessLink() url.URL {
	// link.Fragment = proxy.Name
	var link url.URL
	link.Scheme = "vless"
	query := link.Query()

	for _, vnext := range outbound.Settings.Vnext {
		link.Host = fmt.Sprintf("%s:%d", vnext.Address, vnext.Port)
		for _, user := range vnext.Users {
			link.User = url.User(user.Id)
			addQueryIfNonEmpty(&query, "flow", user.Flow)
		}
	}

	link.RawQuery = query.Encode()
	return link
}

func (outbound *ClientOutbound) addStreamSettingsQueryParamsTo(link *url.URL) {
	streamSettings := outbound.StreamSettings
	if streamSettings == nil {
		return
	}

	if streamSettings.Security != "reality" {
		panic("only reality is supported")
	}

	if len(streamSettings.Network) == 0 {
		streamSettings.Network = "tcp"
	}

	query := link.Query()

	addQueryIfNonEmpty(&query, "type", streamSettings.Network)
	addQueryIfNonEmpty(&query, "security", streamSettings.Security)
	addQueryIfNonEmpty(&query, "fp", streamSettings.RealitySettings.Fingerprint)
	addQueryIfNonEmpty(&query, "sni", streamSettings.RealitySettings.ServerName)
	addQueryIfNonEmpty(&query, "pbk", streamSettings.RealitySettings.PublicKey)
	addQueryIfNonEmpty(&query, "spx", streamSettings.RealitySettings.SpiderX)

	link.RawQuery = query.Encode()
}

func addQueryIfNonEmpty(query *url.Values, key string, value string) {
	if len(value) == 0 {
		return
	}

	query.Add(key, value)
}

func (c *ClientConfig) FirstOutbound() *ClientOutbound {
	return &c.Outbounds[0]
}
