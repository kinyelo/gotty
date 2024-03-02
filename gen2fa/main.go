package main

import (
	"os"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		println("error! Username expected")
		println("Usage: gen2fa <username>")
		os.Exit(1)
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "gotty.surancebay.com",
		AccountName: args[1],
		SecretSize:  15,
	})
	if err != nil {
		panic("can't generate key " + err.Error())
	}
	println("Generated key: " + key.URL())

	qrCode, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		panic("could not generate a QR code " + err.Error())
	}
	err = os.WriteFile(args[1]+".png", qrCode, 0600)
	if err != nil {
		panic("Can't write qa code image file " + err.Error())
	}
}
