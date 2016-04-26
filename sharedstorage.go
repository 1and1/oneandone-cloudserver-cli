package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var sharedStorageOps []cli.Command

func init() {
	ssDriveIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the shared storage drive.",
	}
	sharedStorageOps = []cli.Command{
		{
			Name:        "sharedstorage",
			Description: "1&1 shared storage operations",
			Usage:       "Shared storage operations.",
			Subcommands: []cli.Command{
				{
					Name:  "access",
					Usage: "Shows access credentials or changes password for shared storages.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "newpass",
							Usage: "New password for accessing the shared storages.",
						},
					},
					Action: accessShDrives,
				},
				{
					Name:  "attach",
					Usage: "Attaches servers to shared storage.",
					Flags: []cli.Flag{
						ssDriveIdFlag,
						cli.StringSliceFlag{
							Name:  "serverid",
							Usage: "ID of the servers to attach the shared storage to.",
						},
						cli.StringSliceFlag{
							Name:  "perm",
							Usage: "Permissions for accessing from servers: R or RW.",
						},
					},
					Action: attachShDrive,
				},
				{
					Name:  "create",
					Usage: "Creates new shared storage.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "datacenterid",
							Usage: "Data center ID of the shared storage.",
						},
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the shared storage.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "Description of the shared storage.",
						},
						cli.StringFlag{
							Name:  "size",
							Usage: "Size of the shared storage: 50 - 2000 GB, multiple of 50.",
						},
					},
					Action: createShDrive,
				},
				{
					Name:  "detach",
					Usage: "Detaches server from shared storage.",
					Flags: []cli.Flag{
						ssDriveIdFlag,
						cli.StringFlag{
							Name:  "serverid",
							Usage: "ID of the server.",
						},
					},
					Action: detachShDrive,
				},
				{
					Name:   "info",
					Usage:  "Shows information about shared storage drive.",
					Flags:  []cli.Flag{ssDriveIdFlag},
					Action: showShDrive,
				},
				{
					Name:   "list",
					Usage:  "Lists available shared storage drives.",
					Flags:  queryFlags,
					Action: listShDrives,
				},
				{
					Name:   "rm",
					Usage:  "Deletes shared storage drive.",
					Flags:  []cli.Flag{ssDriveIdFlag},
					Action: deleteShDrive,
				},
				{
					Name:  "serverinfo",
					Usage: "Shows information about shared storage server.",
					Flags: []cli.Flag{
						ssDriveIdFlag,
						cli.StringFlag{
							Name:  "serverid",
							Usage: "ID of the server.",
						},
					},
					Action: showShDriveServer,
				},
				{
					Name:   "serverlist",
					Usage:  "Lists shared storage servers.",
					Flags:  []cli.Flag{ssDriveIdFlag},
					Action: listShDriveServers,
				},
				{
					Name:  "update",
					Usage: "Updates shared storage.",
					Flags: []cli.Flag{
						ssDriveIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "New name of the shared storage.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "New description of the shared storage.",
						},
						cli.StringFlag{
							Name:  "size",
							Usage: "New size of the shared storage: 50 - 2000 GB, multiple of 50.",
						},
					},
					Action: updateShDrive,
				},
			},
		},
	}
}

func createShDrive(ctx *cli.Context) {
	name := getRequiredOption(ctx, "name")
	size := getIntOptionInRange(ctx, "size", 50, 2000)
	if size%50 != 0 {
		exitOnError(fmt.Errorf("--size must be multiple of 50"))
	}
	req := oneandone.SharedStorageRequest{
		DatacenterId: ctx.String("datacenterid"),
		Name:         name,
		Description:  ctx.String("desc"),
		Size:         &size,
	}
	_, storage, err := api.CreateSharedStorage(&req)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}

func listShDrives(ctx *cli.Context) {
	sharedstores, err := api.ListSharedStorages(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(sharedstores))
	for i, drive := range sharedstores {
		var size string
		if drive.Size > 0 {
			size = strconv.Itoa(drive.Size)
		}
		data[i] = []string{
			drive.Id,
			drive.Name,
			size,
			drive.SizeUsed,
			drive.State,
			getDatacenter(drive.Datacenter),
		}
	}
	header := []string{"ID", "Name", "Total Size (GB)", "Used (%)", "State", "Data Center"}
	output(ctx, sharedstores, "", false, &header, &data)
}

func showShDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	storage, err := api.GetSharedStorage(driveId)
	exitOnError(err)
	output(ctx, storage, "", true, nil, nil)
}

func updateShDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	var size *int
	if ctx.IsSet("size") {
		s := getIntOptionInRange(ctx, "size", 50, 2000)
		if s%50 != 0 {
			exitOnError(fmt.Errorf("--size must be multiple of 50"))
		}
		size = &s
	}
	req := oneandone.SharedStorageRequest{
		Name:        ctx.String("name"),
		Description: ctx.String("desc"),
		Size:        size,
	}
	storage, err := api.UpdateSharedStorage(driveId, &req)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}

func attachShDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	servers := getStringSliceOption(ctx, "serverid", true)
	rights := getStringSliceOption(ctx, "perm", true)
	if len(servers) != len(rights) {
		exitOnError(fmt.Errorf("equal number of --serverid and --perm arguments must be specified"))
	}
	var ssServers []oneandone.SharedStorageServer
	for i := 0; i < len(servers); i++ {
		sss := oneandone.SharedStorageServer{
			Id:     servers[i],
			Rights: strings.ToUpper(rights[i]),
		}
		ssServers = append(ssServers, sss)
	}
	storage, err := api.AddSharedStorageServers(driveId, ssServers)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}

func listShDriveServers(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	servers, err := api.ListSharedStorageServers(driveId)
	exitOnError(err)
	data := make([][]string, len(servers))
	for i, server := range servers {
		data[i] = []string{server.Id, server.Name, server.Rights}
	}
	header := []string{"ID", "Name", "Permissions"}
	output(ctx, servers, "", false, &header, &data)
}

func showShDriveServer(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")
	server, err := api.GetSharedStorageServer(driveId, serverId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func detachShDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")
	storage, err := api.DeleteSharedStorageServer(driveId, serverId)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}

func accessShDrives(ctx *cli.Context) {
	if ctx.IsSet("newpass") {
		storage, err := api.UpdateSharedStorageCredentials(ctx.String("newpass"))
		exitOnError(err)
		output(ctx, storage, okWaitMessage, false, nil, nil)
		return
	}
	access, err := api.GetSharedStorageCredentials()
	exitOnError(err)
	output(ctx, access, "", true, nil, nil)
}

func deleteShDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	storage, err := api.DeleteSharedStorage(driveId)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}
