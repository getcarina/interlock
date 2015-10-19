package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/rackerlabs/libcarina"
	"github.com/rgbkrk/interlocarina/plugins"
	"github.com/rgbkrk/interlocarina/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "interlock"
	app.Version = version.FullVersion()
	app.Author = "@rgbkrk"
	app.Email = ""
	app.Usage = "event driven docker plugins"
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username",
			Usage: "carina username",
		},
		cli.StringFlag{
			Name:  "apikey",
			Usage: "carina API key",
		},
		cli.StringFlag{
			Name:  "clustername",
			Usage: "name of the swarm cluster",
		},
		cli.StringFlag{
			Name:  "endpoint",
			Value: libcarina.BetaEndpoint,
			Usage: "endpoint for carina",
		},
		cli.StringSliceFlag{
			Name:  "plugin, p",
			Usage: "enable plugin",
			Value: &cli.StringSlice{},
		},
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
	}
	// base commands
	baseCommands := []cli.Command{
		{
			Name:   "start",
			Action: cmdStart,
		},
		{
			Name:   "list-plugins",
			Action: cmdListPlugins,
		},
		{
			Name:   "info",
			Action: cmdInfo,
		},
	}
	// plugin supplied commands
	baseCommands = append(baseCommands, plugins.GetCommands()...)

	app.Commands = baseCommands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
