package core

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var (
	logger *zerolog.Logger
)

func L() *zerolog.Logger {
	if logger == nil {
		level, err := zerolog.ParseLevel(viper.GetString("loglevel"))
		if err != nil {
			level = zerolog.WarnLevel
		}
		output := viper.GetString("logfile")
		var writer io.Writer
		if output == "-" {
			writer = os.Stderr
		} else {
			writer, err = os.OpenFile(output, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
			if err != nil {
				panic(err)
			}
		}
		var formatWriter io.Writer
		switch viper.GetString("logformat") {
		case "json", "j":
			formatWriter = writer
		case "text", "t", "plain":
			formatWriter = zerolog.ConsoleWriter{Out: writer}
		}
		l := zerolog.New(formatWriter).Level(level).With().Timestamp().Logger()
		logger = &l
	}
	return logger
}
