package logging

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

type LoggingOptions struct {
	LogMoves             bool
	LogSent              bool
	LogSmallAs           bool
	LogHashes            bool
	LogGCObjectSerialise bool
}

var LoggingOpts = LoggingOptions{
	LogMoves:             true,
	LogSent:              true,
	LogSmallAs:           false,
	LogHashes:            false,
	LogGCObjectSerialise: true,
}

func Init() {
	timestamp := time.Now().Format(time.RFC3339)
	safeTime := strings.Replace(timestamp, ":", "_", -1)

	logFile, err := os.OpenFile("resources\\Logs\\"+safeTime+".txt", os.O_APPEND|os.O_CREATE, 0755)

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
