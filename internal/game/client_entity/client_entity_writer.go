package client_entity

import (
	byter "RainbowRunner/pkg/byter"
)

type ClientEntityWriter struct {
	byter *byter.Byter
}

func (w ClientEntityWriter) Create() {

}

func NewClientEntityWriter(b *byter.Byter) *ClientEntityWriter {
	return &ClientEntityWriter{
		byter: b,
	}
}
