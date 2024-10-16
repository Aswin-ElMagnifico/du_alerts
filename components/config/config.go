package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Metadata string
	WatchDir string
	Treshold float32
	Url      string
	Refresh  int
}

func GetConfiguration(args []string) Config {
	var config Config
	for i, v := range args {
		switch v {
		case "--metadata":
			config.Metadata = getValueFromIndex(i, args)
		case "--watch-dir":
			config.WatchDir = getValueFromIndex(i, args)
		case "--notify":
			treshold, _ := strconv.ParseFloat(getValueFromIndex(i, args), 32)
			config.Treshold = float32(treshold)
		case "--url":
			config.Url = getValueFromIndex(i, args)
		case "--refresh":
			minutes, _ := strconv.ParseInt(getValueFromIndex(i, args), 10, 32)
			config.Refresh = int(minutes)
		default:

		}
	}
	return config
}

func getValueFromIndex(index int, args []string) string {
	if index+1 >= len(args) {
		fmt.Printf("unspecified value for %s\n", args[index])
		os.Exit(1)
	}
	return args[index+1]
}
