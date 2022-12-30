package logging

import (
	"RainbowRunner/internal/serverconfig"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

func Init() {
	timestamp := time.Now().Format(time.RFC3339)
	safeName := strings.Replace(timestamp, ":", "_", -1)

	if len(serverconfig.Config.Logging.LogFileName) > 0 {
		safeName = serverconfig.Config.Logging.LogFileName
	}

	flag := os.O_APPEND | os.O_CREATE

	if serverconfig.Config.Logging.LogTruncate {
		flag |= os.O_TRUNC
	}

	if _, err := os.Stat("resources/Logs"); os.IsNotExist(err) {
		fmt.Println("Logs directory missing, creating now")
		err = os.Mkdir("resources/Logs", 0755)

		if err != nil {
			panic(err)
		}
	}

	logFile, err := os.OpenFile("resources/Logs/"+safeName+".txt", flag, 0755)

	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})

	mw := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(mw)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
