package main

import (
	"fmt"
	"strconv"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var serverOps []cli.Command

func init() {
	serverIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the server.",
	}
	ipIdFlag := cli.StringFlag{
		Name:  "ipid",
		Usage: "ID of the IP.",
	}
	fpIdFlag := cli.StringFlag{
		Name:  "firewallid, f",
		Usage: "ID of the firewall policy.",
	}
	lbIdFlag := cli.StringFlag{
		Name:  "loadbalancerid, l",
		Usage: "ID of the load balancer.",
	}
	hddIdFlag := cli.StringFlag{
		Name:  "hddid",
		Usage: "ID of the hard disk.",
	}
	passwordFlag := cli.StringFlag{
		Name:  "password, p",
		Usage: "Password of the server.",
	}
	cpuFlag := cli.StringFlag{
		Name:  "cpu",
		Usage: "Number of processors.",
	}
	coresFlag := cli.StringFlag{
		Name:  "cores",
		Usage: "Number of cores per processor.",
	}
	flavorFlag := cli.StringFlag{
		Name:  "fixsizeid, s",
		Usage: "Fixed size ID desired for the server.",
	}
	hdSizeFlag := cli.StringFlag{
		Name:  "hdsize",
		Usage: "Size of the hard disk in GB.",
	}
	ramFlag := cli.StringFlag{
		Name:  "ram",
		Usage: "Size of RAM memory in GB.",
	}
	datacenterIDFlag := cli.StringFlag{
		Name:  "datacenterid",
		Usage: "Datacenter ID.",
	}
	hwFlags := []cli.Flag{cpuFlag, coresFlag, flavorFlag, hdSizeFlag, ramFlag}

	tcsFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Usage: "Name of the server.",
		},
		cli.StringFlag{
			Name:  "desc, d",
			Usage: "Description of the server.",
		},
		passwordFlag,
		cli.BoolFlag{
			Name:  "poweron",
			Usage: "Power on the server after creating.",
		},
		cli.StringFlag{
			Name:  "osid, a",
			Usage: "Server appliance ID.",
		},
		ipIdFlag,
		datacenterIDFlag,
		fpIdFlag,
		lbIdFlag,
		cli.StringFlag{
			Name:  "monitorpolicyid, m",
			Usage: "Monitoring policy ID to use with the server.",
		},
	}

	serverOps = []cli.Command{
		{
			Name:        "server",
			Description: "1&1 cloud server operations",
			Usage:       "Server operations.",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "Creates new server.",
					Flags:  append(hwFlags, tcsFlags...),
					Action: createServer,
				},
				{
					Name:  "clone",
					Usage: "Clones server.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the new server.",
						},
						datacenterIDFlag,
					},
					Action: cloneServer,
				},
				{
					Name:   "info",
					Usage:  "Shows information about server.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: showServer,
				},
				{
					Name:   "list",
					Usage:  "Lists available servers.",
					Flags:  queryFlags,
					Action: listServers,
				},
				{
					Name:  "fixedsize",
					Usage: "Shows information about fixed size.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Usage: "ID of the fixed size.",
						},
					},
					Action: flavorInfo,
				},
				{
					Name:   "fixedsizes",
					Usage:  "Lists available fixed-size flavors.",
					Action: listServerFlavors,
				},
				{
					Name:   "start",
					Usage:  "Turns server on.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: startServer,
				},
				{
					Name:   "status",
					Usage:  "Shows server's status.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: showServerStatus,
				},
				{
					Name:  "stop",
					Usage: "Turns server off.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.BoolFlag{
							Name:  "force, f",
							Usage: "Force hardware shutdown.",
						},
					},
					Action: shutdownServer,
				},
				{
					Name:  "reboot",
					Usage: "Reboots server.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.BoolFlag{
							Name:  "force, f",
							Usage: "Force hardware reboot.",
						},
					},
					Action: rebootServer,
				},
				{
					Name:  "rm",
					Usage: "Removes server.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.BoolFlag{
							Name:  "keepips",
							Usage: "Keep server IPs after deleting the server.",
						},
					},
					Action: deleteServer,
				},
				{
					Name:  "update",
					Usage: "Updates server's name and description.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "New name.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "New description.",
						},
					},
					Action: renameServer,
				},
				{
					Name:   "dvdinfo",
					Usage:  "Shows server's DVD info.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: serverDvdInfo,
				},
				{
					Name:  "dvdload",
					Usage: "Loads server's DVD.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "dvdid",
							Usage: "ID of the DVD.",
						},
					},
					Action: loadServerDvd,
				},
				{
					Name:   "dvdrm",
					Usage:  "Ejects server's DVD.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: ejectServerDvd,
				},
				{
					Name:   "fwadd",
					Usage:  "Assigns firewall policy to server's IP.",
					Flags:  []cli.Flag{serverIdFlag, ipIdFlag, fpIdFlag},
					Action: addServerFirewall,
				},
				{
					Name:   "fwinfo",
					Usage:  "Shows information about firewall policy assigned to server's IP.",
					Flags:  []cli.Flag{serverIdFlag, ipIdFlag},
					Action: showServerFirewall,
				},
				{
					Name:   "fwrm",
					Usage:  "Removes firewall policy from server's IP.",
					Flags:  []cli.Flag{serverIdFlag, ipIdFlag},
					Action: deleteServerFirewall,
				},
				{
					Name:  "hddadd",
					Usage: "Adds one or more hard disks to server.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.IntSliceFlag{
							Name:  "size",
							Usage: "List of HDD sizes in GB.",
						},
					},
					Action: addServerHdds,
				},
				{
					Name:   "hddinfo",
					Usage:  "Shows information about server's hard disk.",
					Flags:  []cli.Flag{serverIdFlag, hddIdFlag},
					Action: infoServerHdd,
				},
				{
					Name:   "hddlist",
					Usage:  "Lists server's hard disks.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: listServerHdds,
				},
				{
					Name:   "hddrm",
					Usage:  "Removes server's hard disk.",
					Flags:  []cli.Flag{serverIdFlag, hddIdFlag},
					Action: deleteServerHdd,
				},
				{
					Name:  "hddupdate",
					Usage: "Resizes server's hard disk.",
					Flags: []cli.Flag{
						serverIdFlag,
						hddIdFlag,
						cli.IntFlag{
							Name:  "newsize",
							Usage: "New size of the hard disk.",
						},
					},
					Action: resizeServerHdd,
				},
				{
					Name:   "hwinfo",
					Usage:  "Shows information about server's hardware.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: showServerHardware,
				},
				{
					Name:   "hwupdate",
					Usage:  "Modifies server's hardware.",
					Flags:  append([]cli.Flag{cpuFlag, coresFlag, flavorFlag, ramFlag}, serverIdFlag),
					Action: modifyServerHardware,
				},
				{
					Name:   "imginfo",
					Usage:  "Shows information about server's image.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: showServerImage,
				},
				{
					Name:  "imgupdate",
					Usage: "Reinstalls new image into server.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "imgid",
							Usage: "ID of the image.",
						},
						passwordFlag,
						fpIdFlag,
					},
					Action: reInsServerImage,
				},
				{
					Name:  "ipadd",
					Usage: "Assigns new public IP address to server.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "type, t",
							Usage: "IP address type, IPV4 or IPV6. Currently, only IPV4 is allowed.",
						},
					},
					Action: addServerIp,
				},
				{
					Name:   "ipinfo",
					Usage:  "Shows server's public IP.",
					Flags:  []cli.Flag{serverIdFlag, ipIdFlag},
					Action: infoServerIp,
				},
				{
					Name:   "iplist",
					Usage:  "Lists server IP addresses.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: listServerIps,
				},
				{
					Name:  "iprm",
					Usage: "Removes public IP address from server.",
					Flags: []cli.Flag{
						serverIdFlag,
						ipIdFlag,
						cli.BoolFlag{
							Name:  "keepip",
							Usage: "Releases the IP without removing it.",
						},
					},
					Action: deleteServerIp,
				},
				{
					Name:   "lbadd",
					Usage:  "Assigns new load balancer to server's IP.",
					Flags:  []cli.Flag{serverIdFlag, ipIdFlag, lbIdFlag},
					Action: addServerLoadbalancer,
				},
				{
					Name:   "lblist",
					Usage:  "Lists load balancers assigned to server's IP.",
					Flags:  []cli.Flag{serverIdFlag, ipIdFlag},
					Action: listServerLoadbalancers,
				},
				{
					Name:   "lbrm",
					Usage:  "Removes load balancer from server's IP.",
					Flags:  []cli.Flag{serverIdFlag, ipIdFlag, lbIdFlag},
					Action: deleteServerLoadbalancer,
				},
				{
					Name:  "pnadd",
					Usage: "Adds server to private network.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "pnetid",
							Usage: "ID of the private network.",
						},
					},
					Action: addServerPrivateNet,
				},
				{
					Name:  "pninfo",
					Usage: "Shows server's private network.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "pnetid",
							Usage: "ID of the private network.",
						},
					},
					Action: showServerPrivateNet,
				},
				{
					Name:   "pnlist",
					Usage:  "Lists server's private networks.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: listServerPrivateNets,
				},
				{
					Name:  "pnrm",
					Usage: "Removes server from private network.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "pnetid",
							Usage: "ID of the private network.",
						},
					},
					Action: deleteServerPrivateNet,
				},
				{
					Name:   "snapshotmake",
					Usage:  "Creates server's snapshot.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: snapshotServer,
				},
				{
					Name:   "snapshotinfo",
					Usage:  "Shows server's snapshot.",
					Flags:  []cli.Flag{serverIdFlag},
					Action: showServerSnapshot,
				},
				{
					Name:  "snapshotrest",
					Usage: "Restores server's snapshot.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "snapshotid",
							Usage: "ID of the snapshot.",
						},
					},
					Action: restoreServerSnapshot,
				},
				{
					Name:  "snapshotrm",
					Usage: "Removes server's snapshot.",
					Flags: []cli.Flag{
						serverIdFlag,
						cli.StringFlag{
							Name:  "snapshotid",
							Usage: "ID of the snapshot.",
						},
					},
					Action: deleteServerSnapshot,
				},
			},
		},
	}
}

// Helper method
func getHardwareConfig(ctx *cli.Context) oneandone.Hardware {
	var hardware oneandone.Hardware

	fixedSizeID := ctx.String("fixsizeid")
	if fixedSizeID != "" {
		hardware = oneandone.Hardware{
			FixedInsSizeId: fixedSizeID,
		}
	} else {
		hardware = oneandone.Hardware{
			FixedInsSizeId:    ctx.String("fixsizeid"),
			Vcores:            stringFlag2Int(ctx, "cpu"),
			CoresPerProcessor: stringFlag2Int(ctx, "cores"),
			Ram:               stringFlag2Float32(ctx, "ram"),
			Hdds: []oneandone.Hdd{
				oneandone.Hdd{
					Size:   stringFlag2Int(ctx, "hdsize"),
					IsMain: true,
				},
			},
		}
	}

	return hardware
}

func createServer(ctx *cli.Context) {
	req := oneandone.ServerRequest{
		Name:               getRequiredOption(ctx, "name"),
		Description:        ctx.String("desc"),
		ApplianceId:        getRequiredOption(ctx, "osid"),
		Password:           ctx.String("password"),
		PowerOn:            ctx.Bool("poweron"),
		FirewallPolicyId:   ctx.String("firewallid"),
		IpId:               ctx.String("ipid"),
		LoadBalancerId:     ctx.String("loadbalancerid"),
		MonitoringPolicyId: ctx.String("monitorpolicyid"),
		DatacenterId:       ctx.String("datacenterid"),
		Hardware:           getHardwareConfig(ctx),
	}
	_, server, err := api.CreateServer(&req)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func cloneServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	name := getRequiredOption(ctx, "name")
	datacenterID := ctx.String("datacenterid")
	server, err := api.CloneServer(id, name, datacenterID)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func renameServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.RenameServer(id, ctx.String("name"), ctx.String("desc"))
	exitOnError(err)
	output(ctx, server, "", false, nil, nil)
}

func listServers(ctx *cli.Context) {
	servers, err := api.ListServers(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(servers))
	for i, server := range servers {
		state := ""
		if server.Status != nil {
			state = server.Status.State
		}
		data[i] = []string{server.Id, server.Name, state, getDatacenter(server.Datacenter)}
	}
	header := []string{"ID", "Name", "State", "Data Center"}
	output(ctx, servers, "", false, &header, &data)
}

func showServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.GetServer(id)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func listServerFlavors(ctx *cli.Context) {
	flavors, err := api.ListFixedInstanceSizes()
	exitOnError(err)
	data := make([][]string, len(flavors))
	for i, f := range flavors {
		data[i] = []string{
			f.Id, f.Name,
			strconv.FormatFloat(float64(f.Hardware.Ram), 'f', -1, 32),
			strconv.Itoa(f.Hardware.Vcores),
			strconv.Itoa(f.Hardware.CoresPerProcessor),
			strconv.Itoa(f.Hardware.Hdds[0].Size),
		}
	}
	header := []string{"ID", "Name", "RAM (GB)", "Processor No.", "Cores per Processor", "Disk Size (GB)"}
	output(ctx, flavors, "", false, &header, &data)
}

func flavorInfo(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	flavor, err := api.GetFixedInstanceSize(id)
	exitOnError(err)
	output(ctx, flavor, "", true, nil, nil)
}

func deleteServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.DeleteServer(id, ctx.Bool("keepips"))
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func listServerIps(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	ips, err := api.ListServerIps(id)
	exitOnError(err)
	data := make([][]string, len(ips))
	for i, ip := range ips {
		data[i] = []string{ip.Id, ip.Ip, ip.ReverseDns}
	}
	header := []string{"ID", "IP Address", "Reverse DNS"}
	output(ctx, ips, "", false, &header, &data)
}

func addServerIp(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.AssignServerIp(id, ctx.String("iptype"))
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func infoServerIp(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	server, err := api.GetServerIp(serverId, ipId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func deleteServerIp(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	server, err := api.DeleteServerIp(serverId, ipId, ctx.Bool("keepip"))
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func startServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.StartServer(id)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func rebootServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.RebootServer(id, ctx.Bool("force"))
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func shutdownServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.ShutdownServer(id, ctx.Bool("force"))
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func serverDvdInfo(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	dvd, err := api.GetServerDvd(id)
	exitOnError(err)
	output(ctx, dvd, "", true, nil, nil)
}

func loadServerDvd(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	dvdId := getRequiredOption(ctx, "dvdid")
	server, err := api.LoadServerDvd(serverId, dvdId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func ejectServerDvd(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.EjectServerDvd(id)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func showServerHardware(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.GetServerHardware(id)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func modifyServerHardware(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	flavor := ctx.String("fixsizeid")
	processors := stringFlag2Int(ctx, "cpu")
	cores := stringFlag2Int(ctx, "cores")
	ram := stringFlag2Float32(ctx, "ram")
	hardware := oneandone.Hardware{
		FixedInsSizeId:    flavor,
		Vcores:            processors,
		CoresPerProcessor: cores,
		Ram:               ram,
	}
	server, err := api.UpdateServerHardware(id, &hardware)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func listServerHdds(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	hdds, err := api.ListServerHdds(id)
	exitOnError(err)
	data := make([][]string, len(hdds))
	for i, hdd := range hdds {
		data[i] = []string{hdd.Id, strconv.Itoa(hdd.Size), strconv.FormatBool(hdd.IsMain)}
	}
	header := []string{"ID", "Size (GB)", "Main"}
	output(ctx, hdds, "", false, &header, &data)
}

func infoServerHdd(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	hddId := getRequiredOption(ctx, "hddid")
	server, err := api.GetServerHdd(id, hddId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func addServerHdds(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	sizes := getIntSliceOption(ctx, "size", true)
	hdds := new(oneandone.ServerHdds)
	for _, s := range sizes {
		hdds.Hdds = append(hdds.Hdds, oneandone.Hdd{Size: s})
	}
	server, err := api.AddServerHdds(id, hdds)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func deleteServerHdd(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	hddId := getRequiredOption(ctx, "hddid")
	server, err := api.DeleteServerHdd(serverId, hddId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func resizeServerHdd(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	hddId := getRequiredOption(ctx, "hddid")
	newSize := ctx.Int("newsize")

	server, err := api.GetServer(serverId)
	exitOnError(err)

	for _, hdd := range server.Hardware.Hdds {
		if hdd.Id == hddId {
			if hdd.Size >= newSize {
				exitOnError(fmt.Errorf("--newsize must be greater than %d, the current size.", hdd.Size))
			}
			break
		}
	}
	if newSize < 20 || newSize > 2000 || newSize%10 != 0 {
		exitOnError(fmt.Errorf("Invalid value for hard disk size. The size must be at least 20, at most 2000 and multiple of 0.5."))
	}

	server, err = api.ResizeServerHdd(serverId, hddId, newSize)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func showServerImage(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	im, err := api.GetServerImage(serverId)
	exitOnError(err)
	output(ctx, im, "", true, nil, nil)
}

func reInsServerImage(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	imageId := getRequiredOption(ctx, "imgid")
	pass := ctx.String("password")
	fpId := ctx.String("firewallid")
	server, err := api.ReinstallServerImage(serverId, imageId, pass, fpId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func showServerFirewall(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	firewall, err := api.GetServerIpFirewallPolicy(serverId, ipId)
	exitOnError(err)
	output(ctx, firewall, "", true, nil, nil)
}

func addServerFirewall(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	firewallId := getRequiredOption(ctx, "firewallid")
	server, err := api.AssignServerIpFirewallPolicy(serverId, ipId, firewallId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func deleteServerFirewall(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	server, err := api.UnassignServerIpFirewallPolicy(serverId, ipId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func addServerLoadbalancer(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	lbId := getRequiredOption(ctx, "loadbalancerid")
	server, err := api.AssignServerIpLoadBalancer(serverId, ipId, lbId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func listServerLoadbalancers(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	lbs, err := api.ListServerIpLoadBalancers(serverId, ipId)
	exitOnError(err)
	data := make([][]string, len(lbs))
	for i, lb := range lbs {
		data[i] = []string{lb.Id, lb.Name}
	}
	header := []string{"ID", "Name"}
	output(ctx, lbs, "", false, &header, &data)
}

func deleteServerLoadbalancer(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	lbId := getRequiredOption(ctx, "loadbalancerid")
	server, err := api.UnassignServerIpLoadBalancer(serverId, ipId, lbId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func showServerStatus(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	status, err := api.GetServerStatus(id)
	exitOnError(err)
	output(ctx, status, "", true, nil, nil)
}

func listServerPrivateNets(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	privateNets, err := api.ListServerPrivateNetworks(id)
	exitOnError(err)
	data := make([][]string, len(privateNets))
	for i, pn := range privateNets {
		data[i] = []string{pn.Id, pn.Name}
	}
	header := []string{"ID", "Name"}
	output(ctx, privateNets, "", false, &header, &data)
}

func showServerPrivateNet(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	pNetId := getRequiredOption(ctx, "pnetid")
	pn, err := api.GetServerPrivateNetwork(serverId, pNetId)
	exitOnError(err)
	output(ctx, pn, "", true, nil, nil)
}

func addServerPrivateNet(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	pNetId := getRequiredOption(ctx, "pnetid")
	server, err := api.AssignServerPrivateNetwork(serverId, pNetId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func deleteServerPrivateNet(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	pNetId := getRequiredOption(ctx, "pnetid")
	server, err := api.RemoveServerPrivateNetwork(serverId, pNetId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func snapshotServer(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	server, err := api.CreateServerSnapshot(id)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func showServerSnapshot(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	snapshot, err := api.GetServerSnapshot(id)
	exitOnError(err)
	output(ctx, snapshot, "", true, nil, nil)
}

func restoreServerSnapshot(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	snapshotId := getRequiredOption(ctx, "snapshotid")
	server, err := api.RestoreServerSnapshot(serverId, snapshotId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}

func deleteServerSnapshot(ctx *cli.Context) {
	serverId := getRequiredOption(ctx, "id")
	snapshotId := getRequiredOption(ctx, "snapshotid")
	server, err := api.DeleteServerSnapshot(serverId, snapshotId)
	exitOnError(err)
	output(ctx, server, okWaitMessage, false, nil, nil)
}
