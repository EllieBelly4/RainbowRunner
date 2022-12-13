package common

import "go/format"

// formatScript formats a byte array of golang using go/format
func FormatScript(data []byte) []byte {
	data, err := format.Source(data)

	if err != nil {
		panic(err)
	}

	return data
}
