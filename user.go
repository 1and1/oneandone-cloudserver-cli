package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var userOps []cli.Command

func init() {
	userIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the user.",
	}
	userPassFlag := cli.StringFlag{
		Name:  "password, p",
		Usage: "User's password. Pass must contain at least 8 characters using uppercase letters, numbers and other special symbols.",
	}
	userDescFlag := cli.StringFlag{
		Name:  "desc, d",
		Usage: "Description of the user.",
	}
	userEmailFlag := cli.StringFlag{
		Name:  "email",
		Usage: "User's e-mail",
	}
	userOps = []cli.Command{
		{
			Name:        "user",
			Description: "1&1 user operations",
			Usage:       "User operations.",
			Subcommands: []cli.Command{
				{
					Name:   "api",
					Usage:  "Shows information about user's API access.",
					Flags:  []cli.Flag{userIdFlag},
					Action: showUserApiAccess,
				},
				{
					Name:   "apitoken",
					Usage:  "Shows user's API key.",
					Flags:  []cli.Flag{userIdFlag},
					Action: showUserApiToken,
				},
				{
					Name:   "disableapi",
					Usage:  "Disables API access.",
					Flags:  []cli.Flag{userIdFlag},
					Action: modifyUserApiAccess,
				},
				{
					Name:   "enableapi",
					Usage:  "Enables API access.",
					Flags:  []cli.Flag{userIdFlag},
					Action: modifyUserApiAccess,
				},
				{
					Name:  "create",
					Usage: "Creates new user.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Username.",
						},
						userPassFlag,
						userDescFlag,
						userEmailFlag,
					},
					Action: createUser,
				},
				{
					Name:   "info",
					Usage:  "Shows information about user.",
					Flags:  []cli.Flag{userIdFlag},
					Action: showUser,
				},
				{
					Name:  "ipadd",
					Usage: "Adds one or more IPs from which access to API is allowed.",
					Flags: []cli.Flag{
						userIdFlag,
						cli.StringSliceFlag{
							Name:  "ip",
							Usage: "List of IP addresses.",
						},
					},
					Action: addUserIps,
				},
				{
					Name:   "ips",
					Usage:  "Lists IPs from which access to API is allowed.",
					Flags:  []cli.Flag{userIdFlag},
					Action: listUserIps,
				},
				{
					Name:  "iprm",
					Usage: "Removes IP from which access to API is allowed.",
					Flags: []cli.Flag{
						userIdFlag,
						cli.StringFlag{
							Name:  "ip",
							Usage: "IP address to forbid API access for.",
						},
					},
					Action: deleteUserIp,
				},
				{
					Name:   "list",
					Usage:  "Lists available users.",
					Flags:  queryFlags,
					Action: listUsers,
				},
				{
					Name:  "modify",
					Usage: "Modifies user.",
					Flags: []cli.Flag{
						userIdFlag,
						userPassFlag,
						userDescFlag,
						userEmailFlag,
						cli.StringFlag{
							Name:  "status",
							Usage: "Enable or disable user: ACTIVE or DISABLED.",
						},
					},
					Action: modifyUser,
				},
				{
					Name:   "newtoken",
					Usage:  "Renews user's API key.",
					Flags:  []cli.Flag{userIdFlag},
					Action: renewUserApiToken,
				},
				{
					Name:   "rm",
					Usage:  "Deletes user.",
					Flags:  []cli.Flag{userIdFlag},
					Action: deleteUser,
				},
				{
					Name:   "permissions",
					Usage:  "Shows current user permissions.",
					Action: showPermissions,
				},
			},
		},
	}
}

func createUser(ctx *cli.Context) {
	name := getRequiredOption(ctx, "name")
	password := getRequiredOption(ctx, "password")

	req := oneandone.UserRequest{
		Name:        name,
		Password:    password,
		Description: ctx.String("desc"),
		Email:       ctx.String("email"),
	}
	_, user, err := api.CreateUser(&req)
	exitOnError(err)
	output(ctx, user, okWaitMessage, false, nil, nil)
}

func listUsers(ctx *cli.Context) {
	users, err := api.ListUsers(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(users))
	for i, u := range users {
		var role, apiState, apiKey string
		if u.Role != nil {
			role = u.Role.Name
		}
		if u.Api != nil {
			apiState = strconv.FormatBool(u.Api.Active)
			apiKey = u.Api.UserApiKey.Key
		}
		data[i] = []string{
			u.Id,
			u.Name,
			u.Email,
			role,
			u.State,
			apiState,
			apiKey,
		}
	}
	header := []string{"ID", "Name", "E-Mail", "Role", "State", "API Enabled", "API KEY"}
	output(ctx, users, "", false, &header, &data)
}

func showUser(ctx *cli.Context) {
	userId := getRequiredOption(ctx, "id")
	user, err := api.GetUser(userId)
	exitOnError(err)
	output(ctx, user, "", true, nil, nil)
}

func modifyUser(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	status := strings.ToUpper(ctx.String("status"))

	if status != "ACTIVE" && status != "DISABLE" {
		exitOnError(fmt.Errorf("Invalid value for --status flag. Valid values are ACTIVE and DISABLE."))
	}

	req := oneandone.UserRequest{
		Password:    ctx.String("password"),
		Description: ctx.String("desc"),
		Email:       ctx.String("email"),
		State:       status,
	}
	user, err := api.ModifyUser(id, &req)
	exitOnError(err)
	output(ctx, user, okWaitMessage, false, nil, nil)
}

func modifyUserApiAccess(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	var active bool
	if ctx.Command.Name == "enableapi" {
		active = true
	}
	user, err := api.ModifyUserApi(id, active)
	exitOnError(err)
	output(ctx, user, okWaitMessage, false, nil, nil)
}

func addUserIps(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	ips := getStringSliceOption(ctx, "ip", true)
	user, err := api.AddUserApiAlowedIps(id, ips)
	exitOnError(err)
	output(ctx, user, okWaitMessage, false, nil, nil)
}

func listUserIps(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	ips, err := api.ListUserApiAllowedIps(id)
	exitOnError(err)
	data := make([][]string, len(ips))
	for i, ip := range ips {
		data[i] = append(data[i], ip)
	}
	output(ctx, ips, "", false, &[]string{"IP Address"}, &data)
}

func deleteUserIp(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	ip := getRequiredOption(ctx, "ip")
	user, err := api.RemoveUserApiAllowedIp(id, ip)
	exitOnError(err)
	output(ctx, user, okWaitMessage, false, nil, nil)
}

func showUserApiAccess(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	access, err := api.GetUserApi(id)
	exitOnError(err)
	output(ctx, access, "", true, nil, nil)
}

func showUserApiToken(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	apiKey, err := api.GetUserApiKey(id)
	exitOnError(err)
	output(ctx, apiKey, "", true, nil, nil)
}

func renewUserApiToken(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	user, err := api.RenewUserApiKey(id)
	exitOnError(err)
	output(ctx, user, "", true, nil, nil)
}

func deleteUser(ctx *cli.Context) {
	userId := getRequiredOption(ctx, "id")
	user, err := api.DeleteUser(userId)
	exitOnError(err)
	output(ctx, user, okWaitMessage, false, nil, nil)
}

func showPermissions(ctx *cli.Context) {
	perm, err := api.GetCurrentUserPermissions()
	exitOnError(err)
	output(ctx, perm, "", true, nil, nil)
}
