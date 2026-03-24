package models

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed fixtures/client-outbound.json
var clientTest []byte

//go:embed fixtures/client-outbound-xhttp.json
var clientXhttpTest []byte

const ExpectedLink = "vless://4bdb184f-263d-47ce-8a68-c3267278a078@127.0.0.1:443?encryption=none&flow=xtls-rprx-vision&fp=chrome&pbk=some-publicKey&security=reality&sid=server-short-id-for-this-user&sni=www.gggg.com&spx=%2F&type=tcp#reality-tcp"

const ExpectedXhttpLink = "vless://4bdb184f-263d-47ce-8a68-c3267278a078@127.0.0.1:443?encryption=none&fp=chrome&path=%2Fxh&pbk=some-publicKey&security=reality&sid=some-short-id&sni=www.microsoft.com&type=xhttp#reality-xhttp"

func TestShareLink(t *testing.T) {
	var outbound ClientOutbound
	json.Unmarshal(clientTest, &outbound)

	link := outbound.ShareLink("server-short-id-for-this-user")
	t.Log(link.String())

	if link.String() != ExpectedLink {
		t.Fatalf("\nExpected %s \ngot %s", ExpectedLink, link.String())
	}
}

func TestShareLinkXhttp(t *testing.T) {
	var outbound ClientOutbound
	json.Unmarshal(clientXhttpTest, &outbound)

	link := outbound.ShareLink("some-short-id")
	t.Log(link.String())

	if link.String() != ExpectedXhttpLink {
		t.Fatalf("\nExpected %s \ngot %s", ExpectedXhttpLink, link.String())
	}
}
