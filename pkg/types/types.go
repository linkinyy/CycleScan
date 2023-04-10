package types

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"os"
)

type LogOption struct {
	Record  bool
	NoLog   bool
	JsonLog bool
}

type Options struct {
	LogOption
}

var Option = &Options{}
var (
	LogCategory = "Log Config"
)

func init() {
	app := cli.NewApp()
	app.Name = "CycleScan"
	app.Version = "1.0"
	app.Usage = "CycleScan Tools"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "record",
			Aliases:     []string{"rd"},
			Usage:       "Record To Log File",
			Category:    LogCategory,
			Value:       false,
			Destination: &Option.Record,
		},
		&cli.BoolFlag{
			Name:        "no-log",
			Aliases:     []string{"nl"},
			Usage:       "Not Print Log To Console",
			Category:    LogCategory,
			Value:       false,
			Destination: &Option.NoLog,
		},
		&cli.BoolFlag{
			Name:        "json-log",
			Aliases:     []string{"jl"},
			Usage:       "Json Style Log",
			Category:    LogCategory,
			Value:       false,
			Destination: &Option.JsonLog,
		},
	}

	app.Action = func(c *cli.Context) error {
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		_ = fmt.Errorf("options err :%s", err)
		os.Exit(-1)
	}
}
