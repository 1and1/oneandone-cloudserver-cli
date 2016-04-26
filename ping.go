package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var pingOps []cli.Command

func init() {
	pingOps = []cli.Command{
		{
			Name:        "ping",
			Description: "1&1 ping operations",
			Usage:       "Ping operations.",
			Subcommands: []cli.Command{
				{
					Name:   "api",
					Usage:  "Checks if API is running.",
					Action: ping,
				},
				{
					Name:   "auth",
					Usage:  "Validates API token key being used.",
					Action: pingAuth,
				},
			},
		},
	}
}

func ping(ctx *cli.Context) {
	pong, err := api.Ping()
	exitOnError(err)
	if len(pong) > 0 {
		fmt.Println("Response: " + pong[0])
		if pong[0] == "PONG" {
			fmt.Println("The API is running.")
		}
	}
}

func pingAuth(ctx *cli.Context) {
	pong, err := api.PingAuth()
	exitOnError(err)
	if len(pong) > 0 {
		fmt.Println("Response: " + pong[0])
		if pong[0] == "PONG" {
			fmt.Println("The token is valid.")
		}
	}
}
