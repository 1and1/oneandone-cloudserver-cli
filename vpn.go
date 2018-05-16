package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	fp "path/filepath"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

var vpnOps []cli.Command

func init() {
	vpnIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the VPN.",
	}
	vpnOps = []cli.Command{
		{
			Name:        "vpn",
			Description: "1&1 vpn operations",
			Usage:       "VPN operations.",
			Subcommands: []cli.Command{
				{
					Name:  "configfile",
					Usage: "Downloads VPN configuration files as a zip arhive.",
					Flags: []cli.Flag{
						vpnIdFlag,
						cli.StringFlag{
							Name:  "dir",
							Usage: "Directory where to store the VPN configuration.",
						},
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the confiration file arhive.",
						},
					},
					Action: downloadVPNConfig,
				},
				{
					Name:  "create",
					Usage: "Creates new VPN.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "datacenterid",
							Usage: "Data center ID of the VPN.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "Description of the VPN.",
						},
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the VPN.",
						},
					},
					Action: createVPN,
				},
				{
					Name:   "info",
					Usage:  "Shows information about VPN.",
					Flags:  []cli.Flag{vpnIdFlag},
					Action: showVPN,
				},
				{
					Name:   "list",
					Usage:  "Lists all available VPNs.",
					Flags:  queryFlags,
					Action: listVPNs,
				},
				{
					Name:  "modify",
					Usage: "Modifies VPN configuration.",
					Flags: []cli.Flag{
						vpnIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "New name of the VPN.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "New description of the VPN.",
						},
					},
					Action: modifyVPN,
				},
				{
					Name:   "rm",
					Usage:  "Deletes VPN.",
					Flags:  []cli.Flag{vpnIdFlag},
					Action: deleteVPN,
				},
			},
		},
	}
}

func createVPN(ctx *cli.Context) {
	vpnName := getRequiredOption(ctx, "name")
	vpnDesc := ctx.String("desc")
	datacenterId := ctx.String("datacenterid")
	_, vpn, err := api.CreateVPN(vpnName, vpnDesc, datacenterId)
	exitOnError(err)
	output(ctx, vpn, okWaitMessage, false, nil, nil)
}

func listVPNs(ctx *cli.Context) {
	vpns, err := api.ListVPNs(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(vpns))
	for i, vpn := range vpns {
		data[i] = []string{
			vpn.Id,
			vpn.Name,
			vpn.Type,
			formatDateTime(time.RFC3339, vpn.CreationDate),
			vpn.State,
			getDatacenter(vpn.Datacenter),
		}
	}
	header := []string{"ID", "Name", "Type", "Creation Date", "State", "Data Center"}
	output(ctx, vpns, "", false, &header, &data)
}

func showVPN(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	vpn, err := api.GetVPN(id)
	exitOnError(err)
	output(ctx, vpn, "", true, nil, nil)
}

func deleteVPN(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	vpn, err := api.DeleteVPN(id)
	exitOnError(err)
	output(ctx, vpn, okWaitMessage, false, nil, nil)
}

func modifyVPN(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	vpn, err := api.ModifyVPN(id, ctx.String("name"), ctx.String("desc"))
	exitOnError(err)
	output(ctx, vpn, okWaitMessage, false, nil, nil)
}

func downloadVPNConfig(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")

	var fileName, directory string
	var err error

	if ctx.IsSet("dir") {
		directory = ctx.String("dir")
	} else {
		directory, err = os.Getwd()
		exitOnError(err)
	}

	// Make it absolute
	if !fp.IsAbs(directory) {
		directory, err = fp.Abs(directory)
		exitOnError(err)
	}

	// Check if the directory exists
	_, err = os.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			// make all dirs
			exitOnError(os.MkdirAll(directory, 0666))
		} else {
			exitOnError(err)
		}
	}

	content, err := api.GetVPNConfigFile(id, directory)
	exitOnError(err)
	var data []byte
	data, err = base64.StdEncoding.DecodeString(content)
	exitOnError(err)

	if ctx.IsSet("name") {
		fileName = ctx.String("name")
		if !strings.HasSuffix(fileName, ".zip") {
			fileName += ".zip"
		}
	} else {
		fileName = "vpn_" + fmt.Sprintf("%x", md5.Sum(data)) + ".zip"
	}

	fpath := fp.Join(directory, fileName)

	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	exitOnError(err)

	var n int
	n, err = f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	exitOnError(err)
	fmt.Printf("VPN configuration written to: \"%s\"\n", fpath)
}
