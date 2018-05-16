package main

import (
	"fmt"
	"strconv"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var blockStorageOps []cli.Command

func init() {
	bsDriveIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the block storage drive.",
	}
	blockStorageOps = []cli.Command{
		{
			Name:        "blockstorage",
			Description: "1&1 block storage operations",
			Usage:       "Block storage operations.",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "Creates new block storage.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the block storage.",
						},
						cli.StringFlag{
							Name:  "size",
							Usage: "Size of the block storage: 20 - 500 GB, multiple of 10.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "Description of the block storage.",
						},
						cli.StringFlag{
							Name:  "datacenterid",
							Usage: "ID of the datacenter where the shared storage will be created.",
						},
						cli.StringFlag{
							Name:  "serverid",
							Usage: "ID of the server that will be attached to the block storage.",
						},
					},
					Action: createBsDrive,
				},
				{
					Name:   "info",
					Usage:  "Shows information about block storage drive.",
					Flags:  []cli.Flag{bsDriveIdFlag},
					Action: showBsDrive,
				},
				{
					Name:   "list",
					Usage:  "Lists available block storage drives.",
					Flags:  queryFlags,
					Action: listBsDrives,
				},
				{
					Name:   "rm",
					Usage:  "Deletes block storage drive.",
					Flags:  []cli.Flag{bsDriveIdFlag},
					Action: deleteBsDrive,
				},
				{
					Name:  "attach",
					Usage: "Attaches block storage to a server.",
					Flags: []cli.Flag{
						bsDriveIdFlag,
						cli.StringFlag{
							Name:  "serverid",
							Usage: "ID of the server to which to attach the block storage.",
						},
					},
					Action: attachBsDrive,
				},
				{
					Name:   "serverinfo",
					Usage:  "Shows information about block storage server.",
					Flags:  []cli.Flag{bsDriveIdFlag},
					Action: showBsDriveServer,
				},
				{
					Name:   "detach",
					Usage:  "Detaches a block storage from a server.",
					Flags:  []cli.Flag{bsDriveIdFlag},
					Action: detachBsDrive,
				},
			},
		},
	}
}

func createBsDrive(ctx *cli.Context) {
	name := getRequiredOption(ctx, "name")
	size := getIntOptionInRange(ctx, "size", 20, 500)
	if size%10 != 0 {
		exitOnError(fmt.Errorf("--size must be multiple of 50"))
	}
	req := oneandone.BlockStorageRequest{
		Name:         name,
		Description:  ctx.String("desc"),
		Size:         &size,
		ServerId:     ctx.String("serverid"),
		DatacenterId: ctx.String("datacenterid"),
	}
	_, storage, err := api.CreateBlockStorage(&req)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}

func showBsDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	storage, err := api.GetBlockStorage(driveId)
	exitOnError(err)
	output(ctx, storage, "", true, nil, nil)
}

func listBsDrives(ctx *cli.Context) {
	blockstores, err := api.ListBlockStorages(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(blockstores))
	for i, drive := range blockstores {
		var size string
		if drive.Size > 0 {
			size = strconv.Itoa(drive.Size)
		}
		data[i] = []string{
			drive.Id,
			drive.Name,
			size,
			drive.State,
			getDatacenter(drive.Datacenter),
			getBsServer(drive.Server),
		}
	}
	header := []string{"ID", "Name", "Total Size (GB)", "State", "Data Center", "Server"}
	output(ctx, blockstores, "", false, &header, &data)
}

func deleteBsDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	storage, err := api.DeleteBlockStorage(driveId)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}

func attachBsDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")

	storage, err := api.AddBlockStorageServer(driveId, serverId)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}

func showBsDriveServer(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	server, err := api.GetBlockStorageServer(driveId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func detachBsDrive(ctx *cli.Context) {
	driveId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")
	storage, err := api.RemoveBlockStorageServer(driveId, serverId)
	exitOnError(err)
	output(ctx, storage, okWaitMessage, false, nil, nil)
}
