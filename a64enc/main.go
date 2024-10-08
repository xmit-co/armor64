package main

import (
	"io"
	"os"

	"armor64.org"
)

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	b := armor64.EncodeToString(raw)
	_, err = os.Stdout.WriteString(b)
	if err != nil {
		panic(err)
	}
}
