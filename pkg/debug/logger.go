package debug

import (
	"flag"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	logpath := "./storage/logs/app.log"

	flag.Parse()
	file, err1 := os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
