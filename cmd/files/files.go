package files

import "bytes"

func GetExtensionForFile(buf []byte, fileLength uint32) (string, string) {
	fileType := "unk"
	ext := ""

	if string(buf[:3]) == "Ogg" {
		fileType = "ogg"
		ext = ".ogg"
	} else if string(buf[:3]) == "DDS" {
		fileType = "dds"
		ext = ".dds"
	} else if string(buf[:8]) == "Material" {
		fileType = "mat"
		ext = ".mat"
	} else if string(buf[1:10]) == "DFC3DNode" {
		fileType = "DFC3DNode"
		ext = ".3dnode"
	} else if string(buf[1:11]) == "DFCControl" {
		fileType = "DFCControl"
		ext = ".control"
	} else if string(buf[1:11]) == "DFCDRENode" {
		fileType = "DFCDRENode"
		ext = ".drenode"
	} else if string(buf[1:10]) == "DFCButton" {
		fileType = "DFCButton"
		ext = ".button"
	} else if string(buf[1:18]) == "DFC3DSkinMeshNode" {
		fileType = "DFC3DSkinMeshNode"
		ext = ".3dskinmeshnode"
	} else if string(buf[1:15]) == "ParticleSystem" {
		fileType = "ParticleSystem"
		ext = ".particlesystem"
	} else if string(buf[1:18]) == "AdvParticleSystem" {
		fileType = "AdvParticleSystem"
		ext = ".advparticlesystem"
	} else if string(buf[1:18]) == "ShortNameLabel" {
		fileType = "ShortNameLabel"
		ext = ".shortnamelabel"
	} else if string(buf[1:17]) == "DescriptionLabel" {
		fileType = "DescriptionLabel"
		ext = ".descriptionlabel"
	} else if string(buf[1:9]) == "DFCLabel" {
		fileType = "DFCLabel"
		ext = ".label"
	} else if string(buf[1:13]) == "DFCScrollBar" {
		fileType = "DFCScrollBar"
		ext = ".scrollbar"
	} else if string(buf[1:22]) == "HybridCollisionObject" {
		fileType = "HybridCollisionObject"
		ext = ".hybridcollisionobj"
	} else if bytes.Compare(buf[1:5], []byte{0xEF, 0xE0, 0xFF, 0xD7}) == 0 {
		fileType = "Effect"
		ext = ".fx"
	} else if string(buf[:4]) == "<xml" {
		fileType = "XML"
		ext = ".xml"
	} else {
		fileType = string(buf[:12])
	}

	if ext == "" {
		ascii := true
		for i := 0; i < int(fileLength); i++ {
			if buf[i] > 0x92 {
				ascii = false
				break
			}
		}

		if ascii {
			fileType = "txt"
			ext = ".txt"
		}
	}
	return fileType, ext
}
