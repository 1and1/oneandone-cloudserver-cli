package main

import (
	"strconv"

	"github.com/codegangsta/cli"
)

var applianceOps []cli.Command

func init() {
	applianceOps = []cli.Command{
		{
			Name:        "appliance",
			Description: "1&1 server appliance operations",
			Usage:       "Server appliance operations.",
			Subcommands: []cli.Command{
				{
					Name:  "info",
					Usage: "Shows information about server appliance.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Usage: "ID of the appliance.",
						},
					},
					Action: showAppliance,
				},
				{
					Name:   "list",
					Usage:  "Lists available server appliances.",
					Flags:  queryFlags,
					Action: listAppliances,
				},
			},
		},
	}
}

func showAppliance(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	appliance, err := api.GetServerAppliance(id)
	exitOnError(err)
	output(ctx, appliance, "", true, nil, nil)
}

func listAppliances(ctx *cli.Context) {
	saps, err := api.ListServerAppliances(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(saps))
	for i, a := range saps {
		var arch string
		ar, isNum := a.Architecture.(float64)
		if isNum {
			arch = strconv.FormatFloat(ar, 'f', -1, 64)
		} else {
			arch, _ = a.Architecture.(string)
		}
		data[i] = []string{a.Id, a.Name, a.Type, a.OsVersion, arch}
	}
	header := []string{"ID", "Name", "Type", "OS", "Architecture"}
	output(ctx, saps, "", false, &header, &data)
}
