package main

import (
	"strconv"

	"github.com/codegangsta/cli"
)

var dvdIsoOps []cli.Command

func init() {
	dvdIsoOps = []cli.Command{
		{
			Name:        "dvdiso",
			Description: "1&1 DVD ISO operations",
			Usage:       "DVD ISO operations.",
			Subcommands: []cli.Command{
				{
					Name:  "info",
					Usage: "Shows information about DVD ISO.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Usage: "ID of the DVD.",
						},
					},
					Action: showDvd,
				},
				{
					Name:   "list",
					Usage:  "Lists all available DVD ISOs.",
					Flags:  queryFlags,
					Action: listDvds,
				},
			},
		},
	}
}

func listDvds(ctx *cli.Context) {
	dvds, err := api.ListDvdIsos(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(dvds))
	for i, dvd := range dvds {
		var arch string
		ar, isNum := dvd.Architecture.(float64)
		if isNum {
			arch = strconv.FormatFloat(ar, 'f', -1, 64)
		} else {
			arch, _ = dvd.Architecture.(string)
		}
		data[i] = []string{dvd.Id, dvd.Name, dvd.OsVersion, arch}
	}
	header := []string{"ID", "Name", "OS", "Architecture"}
	output(ctx, dvds, "", false, &header, &data)
}

func showDvd(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	dvd, err := api.GetDvdIso(id)
	exitOnError(err)
	output(ctx, dvd, "", true, nil, nil)
}
