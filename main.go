package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Website health checker",
		Usage: "A small tool that checks if the website is down or no, it's just like the ping command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Doamin name used to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "port number used to check",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			if port == "" {
				port = "80"
			}
			status := Check(ctx.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
