package logging

import (
	"RainbowRunner/internal/config"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

func Init() {
	timestamp := time.Now().Format(time.RFC3339)
	safeName := strings.Replace(timestamp, ":", "_", -1)

	if len(config.Config.Logging.LogFileName) > 0 {
		safeName = config.Config.Logging.LogFileName
	}

	flag := os.O_APPEND | os.O_CREATE

	if config.Config.Logging.LogTruncate {
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
