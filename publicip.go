package main

import (
	"strconv"

	"github.com/codegangsta/cli"
)

var ipOps []cli.Command

func init() {
	ipdIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the IP address",
	}
	ipOps = []cli.Command{
		{
			Name:        "ip",
			Description: "1&1 public IP operations",
			Usage:       "Public IP operations.",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "Allocates new public IP.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "type, t",
							Usage: "IP address type, IPV4 or IPV6. Currently, only IPV4 is allowed.",
						},
						cli.StringFlag{
							Name:  "dns",
							Usage: "Reverse DNS name.",
						},
						cli.StringFlag{
							Name:  "datacenterid",
							Usage: "Data center ID of the IP address.",
						},
					},
					Action: createPublicIP,
				},
				{
					Name:   "info",
					Usage:  "Shows information about public IP.",
					Flags:  []cli.Flag{ipdIdFlag},
					Action: showIP,
				},
				{
					Name:   "list",
					Usage:  "Lists all available IPs.",
					Flags:  queryFlags,
					Action: listIPs,
				},
				{
					Name:   "rm",
					Usage:  "Deletes public IP.",
					Flags:  []cli.Flag{ipdIdFlag},
					Action: deleteIP,
				},
				{
					Name:  "update",
					Usage: "Updates reverse DNS of IP.",
					Flags: []cli.Flag{
						ipdIdFlag,
						cli.StringFlag{
							Name:  "dns",
							Usage: "New reverse DNS name.",
						},
					},
					Action: updateIP,
				},
			},
		},
	}
}

func createPublicIP(ctx *cli.Context) {
	ipType := ctx.String("type")
	dns := ctx.String("dns")
	datacenterId := ctx.String("datacenterid")
	_, ip, err := api.CreatePublicIp(ipType, dns, datacenterId)
	exitOnError(err)
	output(ctx, ip, okWaitMessage, false, nil, nil)
}

func listIPs(ctx *cli.Context) {
	ips, err := api.ListPublicIps(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(ips))
	for i, ip := range ips {
		var dhcp string
		if ip.IsDhcp != nil {
			dhcp = strconv.FormatBool(*ip.IsDhcp)
		}
		data[i] = []string{
			ip.Id,
			ip.IpAddress,
			dhcp,
			ip.ReverseDns,
			ip.State,
			getDatacenter(ip.Datacenter),
		}
	}
	header := []string{"ID", "IP Address", "DHCP", "Reverse DNS", "State", "Data Center"}
	output(ctx, ips, "", false, &header, &data)
}

func showIP(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	ip, err := api.GetPublicIp(id)
	exitOnError(err)
	output(ctx, ip, "", true, nil, nil)
}

func deleteIP(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	ip, err := api.DeletePublicIp(id)
	exitOnError(err)
	output(ctx, ip, okWaitMessage, false, nil, nil)
}

func updateIP(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	dns := ctx.String("dns")
	ip, err := api.UpdatePublicIp(id, dns)
	exitOnError(err)
	output(ctx, ip, "", false, nil, nil)
}
