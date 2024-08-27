package models

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed fixtures/client-outbound.json
var clientTest []byte

const ExpectedLink = "vless://4bdb184f-263d-47ce-8a68-c3267278a078@127.0.0.1:443?flow=xtls-rprx-vision&fp=chrome&pbk=some-publicKey&security=reality&sid=server-short-id-for-this-user&sni=www.gggg.com&spx=%2F&type=tcp"

func TestShareLink(t *testing.T) {
	var outbound ClientOutbound
	json.Unmarshal(clientTest, &outbound)

	link := outbound.ShareLink("server-short-id-for-this-user")
	t.Log(link.String())

	if link.String() != ExpectedLink {
		t.Fatalf("\nExpected %s \ngot %s", ExpectedLink, link.String())
	}
}
