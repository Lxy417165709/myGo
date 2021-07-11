package main

import qrcode "github.com/skip2/go-qrcode"

func main() {
	q, err := qrcode.New("https://example.org", qrcode.Medium)
	if err != nil {
		panic(err)
	}

	q.DisableBorder = true

	err = q.WriteFile(256, "example-noborder.png")
	if err != nil {
		panic(err)
	}
}
