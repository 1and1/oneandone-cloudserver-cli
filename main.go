package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

var AppVersion = "0.0.1"

const (
	appName       = "oneandone"
	appHelpName   = "1&1 Cloud Server CLI"
	appCopyright  = "Copyright (c) 2016 1&1 Internet SE"
	okWaitMessage = "OK, wait for the action to complete.\n"
)

var (
	api        *oneandone.API
	queryFlags = []cli.Flag{
		cli.IntFlag{
			Name:  "page",
			Usage: "Current page to show.",
		},
		cli.IntFlag{
			Name:  "perpage",
			Usage: "Number of servers that will be shown in each page.",
		},
		cli.StringFlag{
			Name:  "sort, s",
			Usage: "Property to sort the list by priority. E.g., 'name' or '-creation_date'",
		},
		cli.StringFlag{
			Name:  "query, q",
			Usage: "Search for a string in the response and return the elements that contain it.",
		},
		cli.StringFlag{
			Name:  "fields",
			Usage: "Return only the requested fields. E.g., 'id,name'",
		},
	}
	periodFlag = cli.StringFlag{
		Name:  "period",
		Usage: "Time range: LAST_HOUR, LAST_24H, LAST_7D, LAST_30D, LAST_365D or CUSTOM.",
	}
	startDateFlag = cli.StringFlag{
		Name:  "startdate",
		Usage: "The first date in a custom range. Required only if selected period is 'CUSTOM'.",
	}
	endDateFlag = cli.StringFlag{
		Name:  "enddate",
		Usage: "The second date in a custom range. Required only if selected period is 'CUSTOM'.",
	}
)

func main() {
	setHelpTemplates()

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appHelpName
	app.Version = AppVersion
	app.Copyright = appCopyright
	app.EnableBashCompletion = true

	cli.HelpFlag.Usage = "Show help."
	cli.VersionFlag.Usage = "Print the version."

	// global flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "about",
			Usage: "Show info about the application.",
		},
		cli.StringFlag{
			EnvVar: "ONEANDONE_API_KEY",
			Name:   "apikey",
			Usage:  "The API token key.",
		},
		cli.StringFlag{
			EnvVar: "ONEANDONE_BASE_URL",
			Name:   "baseurl",
			Usage:  "The API base endpoint. Default: https://cloudpanel-api.1and1.com/v1",
		},
		cli.BoolFlag{
			EnvVar: "ONEANDONE_JSON_OUTPUT",
			Name:   "json",
			Usage:  "Print output as JSON string.",
		},
		cli.BoolFlag{
			EnvVar: "ONEANDONE_DISPLAY_WRAP",
			Name:   "wrap",
			Usage:  "Try to fit the screen display by wrapping long table cells' content.",
		},
	}

	app.Before = beforeCommandRun

	// Operations
	app.Commands = append(app.Commands, applianceOps...)
	app.Commands = append(app.Commands, datacenterOps...)
	app.Commands = append(app.Commands, dvdIsoOps...)
	app.Commands = append(app.Commands, firewallOps...)
	app.Commands = append(app.Commands, imageOps...)
	app.Commands = append(app.Commands, ipOps...)
	app.Commands = append(app.Commands, loadbalancerOps...)
	app.Commands = append(app.Commands, logOps...)
	app.Commands = append(app.Commands, monitorCenterOps...)
	app.Commands = append(app.Commands, monitorPolicyOps...)
	app.Commands = append(app.Commands, pingOps...)
	app.Commands = append(app.Commands, pricingOps...)
	app.Commands = append(app.Commands, privateNetOps...)
	app.Commands = append(app.Commands, roleOps...)
	app.Commands = append(app.Commands, serverOps...)
	app.Commands = append(app.Commands, sharedStorageOps...)
	app.Commands = append(app.Commands, usageOps...)
	app.Commands = append(app.Commands, userOps...)
	app.Commands = append(app.Commands, vpnOps...)

	app.Run(os.Args)
}

func setHelpTemplates() {
	cli.AppHelpTemplate = `{{.HelpName}}
	{{if .Version}}
Version: {{.Version}}
   {{end}}
Usage: {{.Name}} {{if .Flags}}[OPTIONS]{{end}}{{if .Commands}} OPERATION COMMAND{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if .Flags}}
Options:
   {{range .Flags}}{{.}}
   {{end}}{{end}}{{if .Commands}}
Operations:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}
Run '{{.Name}} OPERATION --help' for more information on an operation's commands.
`

	cli.CommandHelpTemplate = `{{.Usage}}

Usage: {{.HelpName}}{{if .Flags}} [COMMAND OPTIONS]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{if .Description}}

Description: {{.Description}}{{end}}{{if .Flags}}

Options:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`

	cli.SubcommandHelpTemplate = `{{.Usage}}

Usage: {{.HelpName}} COMMAND{{if .Flags}} [COMMAND OPTIONS]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}

Commands:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
Options:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`
}

func beforeCommandRun(ctx *cli.Context) error {
	if ctx.GlobalIsSet("about") {
		fmt.Fprintf(os.Stdout, ctx.App.HelpName+"\n\n")
		fmt.Fprintf(os.Stdout, appCopyright)
		fmt.Fprintf(os.Stdout, "\n\nThis software is using the following open source components:\n\n")
		fmt.Fprintf(os.Stdout, "- codegangsta CLI framework\n")
		fmt.Fprintf(os.Stdout, "\tCopyright (C) 2013 Jeremy Saenz\n")
		fmt.Fprintf(os.Stdout, "\tAll Rights Reserved.\n")
		fmt.Fprintf(os.Stdout, "\tMIT license: https://github.com/codegangsta/cli/blob/master/LICENSE\n\n")
		fmt.Fprintf(os.Stdout, "- ASCII Table Writer\n")
		fmt.Fprintf(os.Stdout, "\tCopyright (C) 2014 by Oleku Konko\n")
		fmt.Fprintf(os.Stdout, "\tLicense terms: https://github.com/olekukonko/tablewriter/blob/master/LICENCE.md\n")
		os.Exit(0)
	}
	var err error

	if ctx.NArg() > 1 {
		last := ctx.Args()[ctx.NArg()-1]
		if last != "--help" && last != "-help" && last != "-h" && last != "--h" {
			api, err = newClient(ctx.GlobalString("apikey"), ctx.GlobalString("baseurl"))
		}
	}
	return err
}

func newClient(token, url string) (*oneandone.API, error) {
	if token == "" {
		return nil, fmt.Errorf("No API key specified, use either --apikey global option or environment variable ONEANDONE_API_KEY")
	}
	if url == "" {
		url = oneandone.BaseUrl
	}
	return oneandone.New(token, url), nil
}

func getRequiredOption(ctx *cli.Context, flag string) string {
	option := ctx.String(flag)
	if !ctx.IsSet(flag) || strings.TrimSpace(option) == "" {
		exitOnError(fmt.Errorf("--%s option is required", flag))
	}
	return option
}

func getIntSliceOption(ctx *cli.Context, flag string, required bool) []int {
	slice := ctx.IntSlice(flag)
	if required && (!ctx.IsSet(flag) || len(slice) == 0) {
		exitOnError(fmt.Errorf("--%s must specify at least one integer value", flag))
	}
	return slice
}

func getStringSliceOption(ctx *cli.Context, flag string, required bool) []string {
	slice := ctx.StringSlice(flag)
	if required && (!ctx.IsSet(flag) || len(slice) == 0) {
		exitOnError(fmt.Errorf("--%s must specify at least one string value", flag))
	}
	return slice
}

func getIntOptionInRange(ctx *cli.Context, flag string, min, max int) int {
	value := ctx.Int(flag)
	return validateIntRange(flag, value, min, max)
}

func getIntOption(ctx *cli.Context, flag string, required bool) int {
	if required && !ctx.IsSet(flag) {
		exitOnError(fmt.Errorf("--%s option is required", flag))
	}
	return ctx.Int(flag)
}

func getDateOption(ctx *cli.Context, flag string, required bool) time.Time {
	dateStr := ctx.String(flag)
	if required && (!ctx.IsSet(flag) || strings.TrimSpace(dateStr) == "") {
		exitOnError(fmt.Errorf("--%s option is required", flag))
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		exitOnError(fmt.Errorf("--%s should be in RFC3339 date format, e.g. 2016-01-29T23:00:00Z ", flag))
	}
	return date
}

func validateIntRange(name string, value, min, max int) int {
	if value < min || value > max {
		exitOnError(fmt.Errorf("--%s must be an integer in range [%d %d]", name, min, max))
	}
	return value
}

func validatePeriod(period string) string {
	switch period {
	case "LAST_HOUR":
		break
	case "LAST_24H":
		break
	case "LAST_7D":
		break
	case "LAST_30D":
		break
	case "LAST_365D":
		break
	case "CUSTOM":
		break
	default:
		exitOnError(fmt.Errorf("--period must be either LAST_HOUR, LAST_24H, LAST_7D, LAST_30D, LAST_365D or CUSTOM"))
	}
	return period
}

func stringFlag2Int(ctx *cli.Context, flag string) int {
	if ctx.IsSet(flag) {
		option := ctx.String(flag)
		if strings.TrimSpace(option) == "" {
			exitOnError(fmt.Errorf("--%s must be an integer", flag))
		}

		n, err := strconv.Atoi(option)

		if err != nil {
			exitOnError(fmt.Errorf("--%s must be an integer", flag))
		}
		return n
	}
	return 0
}

func stringFlag2Float32(ctx *cli.Context, flag string) float32 {
	if ctx.IsSet(flag) {
		option := ctx.String(flag)
		if strings.TrimSpace(option) == "" {
			exitOnError(fmt.Errorf("--%s must be a number", flag))
		}

		n, err := strconv.ParseFloat(option, 32)

		if err != nil {
			exitOnError(fmt.Errorf("--%s must be a number", flag))
		}
		return float32(n)
	}
	return 0
}

func getQueryParams(ctx *cli.Context) (int, int, string, string, string) {
	var page, perPage int
	var sort, query, fields string

	if ctx.IsSet("page") && ctx.Int("page") > 0 {
		page = ctx.Int("page")
	}
	if ctx.IsSet("perpage") && ctx.Int("perpage") > 0 {
		perPage = ctx.Int("perpage")
	}
	if ctx.IsSet("sort") {
		sort = ctx.String("sort")
	}
	if ctx.IsSet("query") {
		query = ctx.String("query")
	}
	if ctx.IsSet("fields") {
		fields = ctx.String("fields")
	}
	return page, perPage, sort, query, fields
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(0)
	}
}

func output(ctx *cli.Context, in interface{}, m string, forceJson bool, header *[]string, data *[][]string) {
	if forceJson || ctx.GlobalBool("json") {
		bytes, _ := json.MarshalIndent(in, "", "    ")
		fmt.Printf("%v\n", string(bytes))
	} else if header != nil && data != nil {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(ctx.GlobalBool("wrap"))
		table.SetAlignment(3)
		table.SetHeader(*header)
		table.AppendBulk(*data)
		table.Render()
	}
	fmt.Print(m)
}

func getDatacenter(dc *oneandone.Datacenter) string {
	if dc != nil {
		return dc.CountryCode
	}
	return ""
}

func formatDateTime(layout string, value string) string {
	dt, err := time.Parse(layout, value)
	if err != nil {
		return value
	}
	return dt.Format(layout)
}
