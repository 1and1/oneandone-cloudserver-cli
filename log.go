package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

var logOps []cli.Command

func init() {
	logOps = []cli.Command{
		{
			Name:        "log",
			Description: "1&1 log operations",
			Usage:       "Log operations.",
			Subcommands: []cli.Command{
				{
					Name:  "info",
					Usage: "Shows log information.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Usage: "ID of the log.",
						},
					},
					Action: showLog,
				},
				{
					Name:   "list",
					Usage:  "Lists all logs in time period.",
					Flags:  []cli.Flag{periodFlag, startDateFlag, endDateFlag},
					Action: listLogs,
				},
			},
		},
	}
}

func listLogs(ctx *cli.Context) {
	period := validatePeriod(strings.ToUpper(getRequiredOption(ctx, "period")))

	var startDate *time.Time
	var endDate *time.Time
	if period == "CUSTOM" {
		startDate = new(time.Time)
		*startDate = getDateOption(ctx, "startdate", true)
		endDate = new(time.Time)
		*endDate = getDateOption(ctx, "enddate", true)
	}

	page, perPage, sort, query, fields := getQueryParams(ctx)
	logs, err := api.ListLogs(period, startDate, endDate, page, perPage, sort, query, fields)
	exitOnError(err)
	data := make([][]string, len(logs))
	for i, log := range logs {
		data[i] = []string{
			log.Id,
			log.Type,
			log.Action,
			log.StartDate,
			strconv.Itoa(log.Duration),
			log.Status.State,
		}
	}
	header := []string{"ID", "Type", "Action", "Start Date", "Duration (S)", "Status"}
	output(ctx, logs, "", false, &header, &data)
}

func showLog(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	log, err := api.GetLog(id)
	exitOnError(err)
	output(ctx, log, "", true, nil, nil)
}
