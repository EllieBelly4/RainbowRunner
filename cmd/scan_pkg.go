package main

import (
	"RainbowRunner/cmd/files"
	byter "RainbowRunner/pkg/byter"
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	dest := "D:\\Work\\dungeon-runners\\666 dumps new"

	pki, err := ioutil.ReadFile("D:\\Work\\dungeon-runners\\game_decompressed.pki")

	if err != nil {
		panic(err)
	}

	pkg, err := os.Open("E:\\Games\\DungeonRunners v666\\game.pkg")

	if err != nil {
		panic(err)
	}

	b := byter.NewLEByter(pki)
	b.Bytes(0x6C) // Skip header

	// This came from the game client, not sure where it comes from
	//local uint entry_count = 0x78AD;
	//
	//typedef struct MetadataEntry {
	//    uint32 str <bgcolor=cLtRed,read=ReadOffsetString>;
	//    uint16 some_flag <bgcolor=cLtBlue>; //0x40 or
	//    uint16 unk_1 <bgcolor=cLtYellow>;
	//    uint32 unk_2;
	//    uint32 file_offset;
	//    uint32 uncompressed_file_length;
	//    uint16 unk_5;
	//    uint16 unk_6;
	//    uint32 unk_7;
	//    uint32 unk_8;
	//    uint32 unk_9;
	//    uint32 unk_10;
	//    uint32 unk_11;
	//};

	buf := make([]byte, 1000000000)

	for i := 0; i < 0x78AD; i++ {
		strOffset := b.UInt32()
		b.UInt16()
		b.UInt16()
		b.UInt32()
		fileOffset := b.UInt32()
		fileLength := b.UInt32()
		isCompressed := b.Bool() // isCompressed
		b.UInt8()
		b.UInt16()
		b.UInt32()
		b.UInt32()
		b.UInt32()
		b.UInt32()
		b.UInt32()

		strByter := byter.NewLEByter(pki[0x6C+0x2C*0x78AD+4+strOffset:])
		str := strByter.CString()

		//if str != "00-shadow_EL_Effect_buff" {
		//	continue
		//}

		pkg.Seek(int64(fileOffset), 0)
		pkg.Read(buf[:fileLength])

		firstBytes := binary.BigEndian.Uint16(buf)

		if firstBytes == 0x78DA {
			decompressed := make([]byte, 1000000000)

			r := bytes.NewReader(buf)
			zReader, err := zlib.NewReader(r)

			if err != nil {
				panic(err)
			}

			var bytesRead = 0
			var totalBytes = 0

			for {
				bytesRead, err = zReader.Read(decompressed[totalBytes:])

				if err != nil && !errors.Is(err, io.EOF) {
					panic(err)
				}

				if bytesRead == 0 {
					break
				}

				totalBytes += bytesRead
			}

			copy(buf, decompressed[:totalBytes])
		}

		fileType, ext := files.GetExtensionForFile(buf, fileLength)

		if isCompressed {
			fileType = "Z " + fileType
		}

		str = strings.ReplaceAll(str, "\\", "_")

		ioutil.WriteFile(path.Join(dest, str+ext), buf[:fileLength], 644)
		fmt.Printf("%d [%s] %s %d %d\n", i, fileType, str, fileOffset, fileLength)
	}
}
