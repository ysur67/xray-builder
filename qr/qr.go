package qr

import (
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func RenderString(s string, inverseColor bool) {
	q, err := qrcode.New(s, qrcode.Low)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q.ToSmallString(inverseColor))
}
