package logging

import (
	"RainbowRunner/internal/message"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

type LoggingOptions struct {
	LogMoves             bool
	LogGenericSent       bool
	LogSmallAs           bool
	LogHashes            bool
	LogGCObjectSerialise bool
	LogRandomEquipment   bool
	LogSentMessages      bool
	LogSentMessageTypes  map[message.OpType]bool
	LogFileName          string
	LogTruncate          bool
	LogEMessages         bool
}

var LoggingOpts = LoggingOptions{
	LogSentMessages:      true,
	LogMoves:             false,
	LogGenericSent:       false,
	LogSmallAs:           false,
	LogEMessages:         false,
	LogHashes:            false,
	LogGCObjectSerialise: false,
	LogRandomEquipment:   false,
	LogFileName:          "inventory_logs",
	LogTruncate:          true,
	LogSentMessageTypes: map[message.OpType]bool{
		message.OpTypeAvatarMovement:            false,
		message.OpTypeCreateNPC:                 true,
		message.OpTypeEquippedItemClickResponse: true,
		message.OpTypeOther:                     false,
	},
}

func Init() {
	timestamp := time.Now().Format(time.RFC3339)
	safeName := strings.Replace(timestamp, ":", "_", -1)

	if len(LoggingOpts.LogFileName) > 0 {
		safeName = LoggingOpts.LogFileName
	}

	flag := os.O_APPEND | os.O_CREATE

	if LoggingOpts.LogTruncate {
		flag |= os.O_TRUNC
	}

	logFile, err := os.OpenFile("resources\\Logs\\"+safeName+".txt", flag, 0755)

	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})

	mw := io.MultiWriter(logFile)
	log.SetOutput(mw)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
