package logger

import (
	"os"

	"github.com/sehun-cha/zommon/lib/log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func SetLogRoot(useTerminal bool, useFile bool, verbosityTerminal int, verbosityFile int, filePath string) {
	// set the log handler according to the log setting in config
	hs := func() []log.Handler {
		h := []log.Handler{}
		if useTerminal { // terminal mode
			handler := log.LvlFilterHandler(log.Lvl(verbosityTerminal),
				log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
			h = append(h, handler)
		}
		if useFile { // file mode
			handler := log.LvlFilterHandler(log.Lvl(verbosityFile), log.StreamHandler(&lumberjack.Logger{
				Filename:   filePath,
				MaxSize:    1024, // megabytes
				MaxBackups: 5,
				MaxAge:     28,   // days
				Compress:   true, // disabled by default
			}, log.JSONFormat()))
			h = append(h, handler)
		}
		return h
	}()
	log.Root().SetHandler(log.MultiHandler(hs...))
	log.Debug("Set log root", "terminal", useTerminal, "file", useFile, "terminal verbosity", verbosityTerminal, "file verbosity", verbosityFile, "file", filePath)
}
