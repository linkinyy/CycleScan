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

type TargetOption struct {
	Ip     string
	Ports  cli.StringSlice
	OsScan bool
}

type Options struct {
	LogOption
	TargetOption
}

var Option = &Options{}

var (
	LogCategory    = "Log Config"
	TargetCategory = "Target Config"
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
		&cli.StringFlag{
			Name:        "ip",
			Usage:       "Target Ip",
			Required:    true,
			Category:    TargetCategory,
			Destination: &Option.Ip,
		},
		&cli.StringSliceFlag{
			Name:        "ports",
			Usage:       "Target Ports",
			Category:    TargetCategory,
			Destination: &Option.Ports,
		},
		&cli.BoolFlag{
			Name:        "os",
			Usage:       "Os Scan Guess",
			Value:       false,
			Category:    TargetCategory,
			Destination: &Option.OsScan,
		},
	}

	app.Action = func(c *cli.Context) error {
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("App Error: ", err)
		os.Exit(-1)
	}
}
