package main

import (
	"io"
	"os"

	"armor64.org"
)

func main() {
	encoded, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	b, err := armor64.DecodeString(string(encoded))
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		panic(err)
	}
}
