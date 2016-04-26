package main

import (
	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var privateNetOps []cli.Command

func init() {
	pnIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the private network.",
	}
	pnNameFlag := cli.StringFlag{
		Name:  "desc, d",
		Usage: "Description of the private network.",
	}
	pnDescFlag := cli.StringFlag{
		Name:  "name, n",
		Usage: "Name of the private network.",
	}
	pnAddressFlag := cli.StringFlag{
		Name:  "netip",
		Usage: "Private network IP address.",
	}
	pnMaskFlag := cli.StringFlag{
		Name:  "netmask",
		Usage: "Subnet mask.",
	}
	pnServerIdFlag := cli.StringFlag{
		Name:  "serverid",
		Usage: "ID of the server.",
	}
	dcIdFlag := cli.StringFlag{
		Name:  "datacenterid",
		Usage: "Data center ID of the private network.",
	}

	pnCreateFlags := []cli.Flag{dcIdFlag, pnDescFlag, pnNameFlag, pnAddressFlag, pnMaskFlag}

	privateNetOps = []cli.Command{
		{
			Name:        "privatenet",
			Description: "1&1 private network operations",
			Usage:       "Private network operations.",
			Subcommands: []cli.Command{
				{
					Name:  "assign",
					Usage: "Assigns servers to private network.",
					Flags: []cli.Flag{
						pnIdFlag,
						cli.StringSliceFlag{
							Name:  "serverid",
							Usage: "List of server IDs.",
						},
					},
					Action: assignPrivateNetServers,
				},
				{
					Name:   "create",
					Usage:  "Creates new private network.",
					Flags:  pnCreateFlags,
					Action: createPrivateNet,
				},
				{
					Name:   "info",
					Usage:  "Shows information about private network.",
					Flags:  []cli.Flag{pnIdFlag},
					Action: showPrivateNet,
				},
				{
					Name:   "list",
					Usage:  "Lists available private networks.",
					Flags:  queryFlags,
					Action: listPrivateNets,
				},
				{
					Name:   "server",
					Usage:  "Shows information about server attached to private network.",
					Flags:  []cli.Flag{pnIdFlag, pnServerIdFlag},
					Action: showPrivateNetServer,
				},
				{
					Name:   "servers",
					Usage:  "Lists servers attached to private network.",
					Flags:  []cli.Flag{pnIdFlag},
					Action: listPrivateNetServers,
				},
				{
					Name:   "rm",
					Usage:  "Removes private network.",
					Flags:  []cli.Flag{pnIdFlag},
					Action: deletePrivateNet,
				},
				{
					Name:   "unassign",
					Usage:  "Unassigns server from private network.",
					Flags:  []cli.Flag{pnIdFlag, pnServerIdFlag},
					Action: removePrivateNetServer,
				},
				{
					Name:   "update",
					Usage:  "Updates private network.",
					Flags:  []cli.Flag{pnIdFlag, pnDescFlag, pnNameFlag, pnAddressFlag, pnMaskFlag},
					Action: updatePrivateNet,
				},
			},
		},
	}
}

func listPrivateNets(ctx *cli.Context) {
	pNets, err := api.ListPrivateNetworks(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(pNets))
	for i, pn := range pNets {
		data[i] = []string{
			pn.Id,
			pn.Name,
			pn.NetworkAddress,
			pn.SubnetMask,
			pn.State,
			getDatacenter(pn.Datacenter),
		}
	}
	header := []string{"ID", "Name", "Network Address", "Subnet Mask", "State", "Data Center"}
	output(ctx, pNets, "", false, &header, &data)
}

func showPrivateNet(ctx *cli.Context) {
	pnId := getRequiredOption(ctx, "id")
	pNet, err := api.GetPrivateNetwork(pnId)
	exitOnError(err)
	output(ctx, pNet, "", true, nil, nil)
}

func createPrivateNet(ctx *cli.Context) {
	req := oneandone.PrivateNetworkRequest{
		Name:           getRequiredOption(ctx, "name"),
		DatacenterId:   ctx.String("datacenterid"),
		Description:    ctx.String("desc"),
		NetworkAddress: ctx.String("netip"),
		SubnetMask:     ctx.String("netmask"),
	}
	_, privateNet, err := api.CreatePrivateNetwork(&req)
	exitOnError(err)
	output(ctx, privateNet, okWaitMessage, false, nil, nil)
}

func updatePrivateNet(ctx *cli.Context) {
	pnId := getRequiredOption(ctx, "id")
	req := oneandone.PrivateNetworkRequest{
		Name:           ctx.String("name"),
		Description:    ctx.String("desc"),
		NetworkAddress: ctx.String("netip"),
		SubnetMask:     ctx.String("netmask"),
	}
	privateNet, err := api.UpdatePrivateNetwork(pnId, &req)
	exitOnError(err)
	output(ctx, privateNet, okWaitMessage, false, nil, nil)
}

func deletePrivateNet(ctx *cli.Context) {
	pnId := getRequiredOption(ctx, "id")
	privateNet, err := api.DeletePrivateNetwork(pnId)
	exitOnError(err)
	output(ctx, privateNet, okWaitMessage, false, nil, nil)
}

func listPrivateNetServers(ctx *cli.Context) {
	pnId := getRequiredOption(ctx, "id")
	servers, err := api.ListPrivateNetworkServers(pnId)
	exitOnError(err)
	data := make([][]string, len(servers))
	for i, server := range servers {
		data[i] = []string{server.Id, server.Name}
	}
	header := []string{"ID", "Name"}
	output(ctx, servers, "", false, &header, &data)
}

func assignPrivateNetServers(ctx *cli.Context) {
	pnId := getRequiredOption(ctx, "id")
	serverIds := getStringSliceOption(ctx, "serverid", true)
	privateNet, err := api.AttachPrivateNetworkServers(pnId, serverIds)
	exitOnError(err)
	output(ctx, privateNet, okWaitMessage, false, nil, nil)
}

func showPrivateNetServer(ctx *cli.Context) {
	pnId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")
	server, err := api.GetPrivateNetworkServer(pnId, serverId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func removePrivateNetServer(ctx *cli.Context) {
	pnId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")
	privateNet, err := api.DetachPrivateNetworkServer(pnId, serverId)
	exitOnError(err)
	output(ctx, privateNet, okWaitMessage, false, nil, nil)
}
