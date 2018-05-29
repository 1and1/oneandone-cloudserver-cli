package main

import (
	"time"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var sshKeyOps []cli.Command

func init() {
	sshKeyIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the SSH Key.",
	}
	sshKeyOps = []cli.Command{
		{
			Name:        "sshkey",
			Description: "1&1 ssh key operations",
			Usage:       "SSH Key operations.",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "Creates a new SSH Key.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the SSH Key.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "Description of the SSH Key.",
						},
						cli.StringFlag{
							Name: "publickey, p",
							Usage: "Public key to import. If not given, new SSH key pair " +
								"will be created and the private key is returned in the response.",
						},
					},
					Action: createSSHKey,
				},
				{
					Name:   "info",
					Usage:  "Shows information about an SSH Key.",
					Flags:  []cli.Flag{sshKeyIdFlag},
					Action: showSSHKey,
				},
				{
					Name:   "list",
					Usage:  "Lists all available SSH Keys.",
					Flags:  queryFlags,
					Action: listSSHKeys,
				},
				{
					Name:  "modify",
					Usage: "Modifies an SSH Key.",
					Flags: []cli.Flag{
						sshKeyIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "New name of the SSH Key.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "New description of the SSH Key.",
						},
					},
					Action: modifySSHKey,
				},
				{
					Name:   "rm",
					Usage:  "Deletes an SSH Key.",
					Flags:  []cli.Flag{sshKeyIdFlag},
					Action: deleteSSHKey,
				},
			},
		},
	}
}

func createSSHKey(ctx *cli.Context) {
	sshKeyName := getRequiredOption(ctx, "name")

	req := oneandone.SSHKeyRequest{
		Name:        sshKeyName,
		Description: ctx.String("desc"),
		PublicKey:   ctx.String("publickey"),
	}
	_, sshKey, err := api.CreateSSHKey(&req)
	exitOnError(err)
	output(ctx, sshKey, okWaitMessage, false, nil, nil)
}

func listSSHKeys(ctx *cli.Context) {
	sshKeys, err := api.ListSSHKeys(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(sshKeys))
	for i, sshKey := range sshKeys {
		data[i] = []string{
			sshKey.Id,
			sshKey.Name,
			sshKey.Description,
			sshKey.State,
			getSSHServers(*sshKey.Servers),
			sshKey.Md5,
			formatDateTime(time.RFC3339, sshKey.CreationDate),
		}
	}
	header := []string{"ID", "Name", "Description", "State", "Servers", "Md5", "Creation Date"}
	output(ctx, sshKeys, "", false, &header, &data)
}

func showSSHKey(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	sshKey, err := api.GetSSHKey(id)
	exitOnError(err)
	output(ctx, sshKey, "", true, nil, nil)
}

func deleteSSHKey(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	sshKey, err := api.DeleteSSHKey(id)
	exitOnError(err)
	output(ctx, sshKey, okWaitMessage, false, nil, nil)
}

func modifySSHKey(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	sshKey, err := api.RenameSSHKey(id, ctx.String("name"), ctx.String("desc"))
	exitOnError(err)
	output(ctx, sshKey, okWaitMessage, false, nil, nil)
}
