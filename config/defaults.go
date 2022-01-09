package config

import (
	"fmt"
	"os"
	"runtime"
)

var (
	configSections = []string{"logger", "http", "performance", "ftp_public", "ftp_internal"}
	deflogdir      = home + "/.config/" + Title + "/logs/"
	defNoColor     = false
)

var defOpts = map[string]map[string]interface{}{
	"logger": {
		"debug":             true,
		"directory":         deflogdir,
		"nocolor":           defNoColor,
		"use_date_filename": true,
	},
	"habbo": {
		"bind":     "127.0.0.1",
		"port":     11235,
		"maxconns": 2,
	},
	"http": {
		"bind": "127.0.0.1",
		"port": 8080,
	},
	"data": {
		"dbdir": "./.data",
	},
}

func setDefaults() {
	if runtime.GOOS == "windows" {
		deflogdir = "logs/"
		defNoColor = true
	}

	for _, def := range configSections {
		snek.SetDefault(def, defOpts[def])
	}

	if GenConfig {
		if err = snek.SafeWriteConfigAs("./config.toml"); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

}
