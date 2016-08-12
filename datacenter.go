package main

import "github.com/codegangsta/cli"

var datacenterOps []cli.Command

func init() {
	datacenterOps = []cli.Command{
		{
			Name:        "datacenter",
			Description: "1&1 data center operations",
			Usage:       "Data center operations.",
			Subcommands: []cli.Command{
				{
					Name:  "info",
					Usage: "Shows information about data center.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Usage: "ID of the data center.",
						},
					},
					Action: showDatacenter,
				},
				{
					Name:   "list",
					Usage:  "Lists all available data centers.",
					Flags:  queryFlags,
					Action: listDatacenters,
				},
			},
		},
	}
}

func listDatacenters(ctx *cli.Context) {
	datacenters, err := api.ListDatacenters(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(datacenters))
	for i, dc := range datacenters {
		data[i] = []string{dc.Id, dc.Location, dc.CountryCode}
	}
	header := []string{"ID", "Location", "Country Code"}
	output(ctx, datacenters, "", false, &header, &data)
}

func showDatacenter(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	datacenter, err := api.GetDatacenter(id)
	exitOnError(err)
	output(ctx, datacenter, "", true, nil, nil)
}
