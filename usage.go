package main

import (
	"strings"
	"time"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var usageOps []cli.Command

func init() {
	usageOps = []cli.Command{
		{
			Name:        "usage",
			Description: "1&1 usage operations",
			Usage:       "Usage operations.",
			Subcommands: []cli.Command{
				{
					Name:   "images",
					Usage:  "Lists all image usages in the specified time period.",
					Flags:  []cli.Flag{periodFlag, startDateFlag, endDateFlag},
					Action: listUsages,
				},
				{
					Name:   "loadbalancers",
					Usage:  "Lists all load balancer usages in the specified time period.",
					Flags:  []cli.Flag{periodFlag, startDateFlag, endDateFlag},
					Action: listUsages,
				},
				{
					Name:   "ips",
					Usage:  "Lists all public IP usages in the specified time period.",
					Flags:  []cli.Flag{periodFlag, startDateFlag, endDateFlag},
					Action: listUsages,
				},
				{
					Name:   "servers",
					Usage:  "Lists all server usages in the specified time period.",
					Flags:  []cli.Flag{periodFlag, startDateFlag, endDateFlag},
					Action: listUsages,
				},
				{
					Name:   "sharedstorages",
					Usage:  "Lists all shared storage usages in the specified time period.",
					Flags:  []cli.Flag{periodFlag, startDateFlag, endDateFlag},
					Action: listUsages,
				},
			},
		},
	}
}

// Helper function
func getUsageData(cmd string, us *oneandone.Usages) [][]string {
	var data [][]string
	switch cmd {
	case "images":
		data = make([][]string, len(us.Images))
		for i, im := range us.Images {
			data[i] = []string{im.Id, im.Name}
		}
		break
	case "loadbalancers":
		data = make([][]string, len(us.LoadBalancers))
		for i, lb := range us.LoadBalancers {
			data[i] = []string{lb.Id, lb.Name}
		}
		break
	case "ips":
		data = make([][]string, len(us.PublicIPs))
		for i, ip := range us.PublicIPs {
			data[i] = []string{ip.Id, ip.Name}
		}
		break
	case "servers":
		data = make([][]string, len(us.Servers))
		for i, s := range us.Servers {
			data[i] = []string{s.Id, s.Name}
		}
		break
	case "sharedstorages":
		data = make([][]string, len(us.SharedStorages))
		for i, ss := range us.SharedStorages {
			data[i] = []string{ss.Id, ss.Name}
		}
		break
	}
	return data
}

func listUsages(ctx *cli.Context) {
	period := validatePeriod(strings.ToUpper(getRequiredOption(ctx, "period")))

	var startDate, endDate *time.Time
	if period == "CUSTOM" {
		startDate = new(time.Time)
		*startDate = getDateOption(ctx, "startdate", true)
		endDate = new(time.Time)
		*endDate = getDateOption(ctx, "enddate", true)
	}

	page, perPage, sort, query, fields := getQueryParams(ctx)
	usages, err := api.ListUsages(period, startDate, endDate, page, perPage, sort, query, fields)
	exitOnError(err)
	header := []string{"ID", "Name"}
	data := getUsageData(ctx.Command.Name, usages)
	output(ctx, usages, "", false, &header, &data)
}
