package main

import (
	"strings"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var monitorCenterOps []cli.Command

func init() {
	statusFlags := []cli.Flag{
		cli.BoolFlag{
			Name:  "cpu",
			Usage: "Set true to show CPU status.",
		},
		cli.BoolFlag{
			Name:  "disk",
			Usage: "Set true to  show disk status.",
		},
		cli.BoolFlag{
			Name:  "ram",
			Usage: "Set true to show RAM status.",
		},
		cli.BoolFlag{
			Name:  "ping",
			Usage: "Set true to show internal ping status.",
		},
		cli.BoolFlag{
			Name:  "transfer",
			Usage: "Set true to show transfer status.",
		},
	}
	monitorCenterOps = []cli.Command{
		{
			Name:        "monitor",
			Description: "1&1 monitoring center operations",
			Usage:       "Monitoring center operations.",
			Subcommands: []cli.Command{
				{
					Name:  "info",
					Usage: "Shows monitoring information about server.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Usage: "ID of the server.",
						},
						periodFlag,
						startDateFlag,
						endDateFlag,
					},
					Action: showMonitor,
				},
				{
					Name:   "list",
					Usage:  "Lists usages and alerts for all monitored servers.",
					Flags:  append(statusFlags, queryFlags...),
					Action: listMonitors,
				},
			},
		},
	}
}

func listMonitors(ctx *cli.Context) {
	ms, err := api.ListMonitoringServersUsages(getQueryParams(ctx))
	exitOnError(err)
	count := len(ms)
	data := make([][]string, count)
	header := []string{"ID", "Name"}
	cpu := ctx.Bool("cpu") && count > 0 && ms[0].Status != nil && ms[0].Status.Cpu != nil
	disk := ctx.Bool("disk") && count > 0 && ms[0].Status != nil && ms[0].Status.Disk != nil
	ram := ctx.Bool("ram") && count > 0 && ms[0].Status != nil && ms[0].Status.Ram != nil
	ping := ctx.Bool("ping") && count > 0 && ms[0].Status != nil && ms[0].Status.InternalPing != nil
	transfer := ctx.Bool("transfer") && count > 0 && ms[0].Status != nil && ms[0].Status.Transfer != nil

	if cpu {
		header = append(header, "CPU State")
	}
	if disk {
		header = append(header, "Disk State")
	}
	if ram {
		header = append(header, "RAM State")
	}
	if ping {
		header = append(header, "Ping State")
	}
	if transfer {
		header = append(header, "Transfer State")
	}

	for i, m := range ms {
		data[i] = []string{
			m.Id,
			m.Name,
		}
		if cpu {
			data[i] = append(data[i], m.Status.Cpu.State)
		}
		if disk {
			data[i] = append(data[i], m.Status.Disk.State)
		}
		if ram {
			data[i] = append(data[i], m.Status.Ram.State)
		}
		if ping {
			data[i] = append(data[i], m.Status.InternalPing.State)
		}
		if transfer {
			data[i] = append(data[i], m.Status.Transfer.State)
		}
	}
	output(ctx, ms, "", false, &header, &data)
}

func showMonitor(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	period := validatePeriod(strings.ToUpper(getRequiredOption(ctx, "period")))

	var ms *oneandone.MonServerUsageDetails
	var err error

	if period == "CUSTOM" {
		startDate := getDateOption(ctx, "startdate", true)
		endDate := getDateOption(ctx, "enddate", true)

		ms, err = api.GetMonitoringServerUsage(id, period, startDate, endDate)
	} else {
		ms, err = api.GetMonitoringServerUsage(id, period)
	}

	exitOnError(err)
	output(ctx, ms, "", true, nil, nil)
}
