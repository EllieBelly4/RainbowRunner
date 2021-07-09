package game

import (
	byter "RainbowRunner/pkg/byter"
	"encoding/hex"
	"regexp"
)

func NewLEByterFromCommandString(str string) *byter.Byter {
	cmd, err := CommandStringToBytes(str)

	if err != nil {
		panic(err)
	}

	return byter.NewLEByter(cmd)
}

func CommandStringToBytes(str string) ([]byte, error) {
	reg := regexp.MustCompile("#.*(?:\n|$)")
	reg2 := regexp.MustCompile("[ \t\n]")
	cleanString := reg.ReplaceAllString(str, "")
	cleanString = reg2.ReplaceAllString(cleanString, "")
	hexData, err := hex.DecodeString(cleanString)
	return hexData, err
}
