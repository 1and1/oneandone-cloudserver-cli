package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

const (
	back_info = "backinfo"
	fw_info   = "fwinfo"
	img_info  = "imginfo"
	inv_info  = "invinfo"
	ip_info   = "ipinfo"
	lb_info   = "lbinfo"
	log_info  = "loginfo"
	mc_info   = "mcinfo"
	mp_info   = "mpinfo"
	pn_info   = "pninfo"
	role_info = "roleinfo"
	ser_info  = "serinfo"
	ss_info   = "ssinfo"
	usg_info  = "usginfo"
	user_info = "userinfo"
	vpn_info  = "vpninfo"

	set_all   = "setall"
	unset_all = "unsetall"

	back_mod = "backmod"
	fw_mod   = "fwmod"
	img_mod  = "imgmod"
	inv_mod  = "invmod"
	ip_mod   = "ipmod"
	lb_mod   = "lbmod"
	log_mod  = "logmod"
	mc_mod   = "mcmod"
	mp_mod   = "mpmod"
	pn_mod   = "pnmod"
	role_mod = "rolemod"
	ser_mod  = "sermod"
	ss_mod   = "ssmod"
	usg_mod  = "usgmod"
	user_mod = "usermod"
	vpn_mod  = "vpnmod"
)

var roleOps []cli.Command

func init() {
	roleIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the role.",
	}
	allFlag := cli.BoolFlag{
		Name:  "all",
		Usage: "Set or unset all permissions in the segment.",
	}
	accessFlag := cli.BoolFlag{
		Name:  "access",
		Usage: "Permission to access.",
	}
	assignipFlag := cli.BoolFlag{
		Name:  "assignip",
		Usage: "Permission to assign IP address.",
	}
	changerRoleFlag := cli.BoolFlag{
		Name:  "changerole",
		Usage: "Permission to change role.",
	}
	cloneFlag := cli.BoolFlag{
		Name:  "clone",
		Usage: "Permission to clone.",
	}
	createFlag := cli.BoolFlag{
		Name:  "create",
		Usage: "Permission to create.",
	}
	deleteFlag := cli.BoolFlag{
		Name:  "delete",
		Usage: "Permission to delete.",
	}
	disableFlag := cli.BoolFlag{
		Name:  "disable",
		Usage: "Permission to disable.",
	}
	disableAutoCreateFlag := cli.BoolFlag{
		Name:  "noautocreate",
		Usage: "Permission to disable automatic creation.",
	}
	downloadFlag := cli.BoolFlag{
		Name:  "downloadfile",
		Usage: "Permission to download configuration file.",
	}
	enableFlag := cli.BoolFlag{
		Name:  "enable",
		Usage: "Permission to enable.",
	}
	manageAPIFlag := cli.BoolFlag{
		Name:  "manageapi",
		Usage: "Permission to manage API.",
	}
	manageDVDFlag := cli.BoolFlag{
		Name:  "managedvd",
		Usage: "Permission to manage DVD.",
	}
	kvmFlag := cli.BoolFlag{
		Name:  "kvm",
		Usage: "Permission to access KVM console.",
	}
	manageSrvFlag := cli.BoolFlag{
		Name:  "manageserver",
		Usage: "Permission to manage attached servers.",
	}
	manageSrvIPFlag := cli.BoolFlag{
		Name:  "manageip",
		Usage: "Permission to manage attached server IPs.",
	}
	managePortFlag := cli.BoolFlag{
		Name:  "manageport",
		Usage: "Permission to manage ports.",
	}
	manageProcFlag := cli.BoolFlag{
		Name:  "manageprocess",
		Usage: "Permission to manage processes.",
	}
	manageRuleFlag := cli.BoolFlag{
		Name:  "managerule",
		Usage: "Permission to manage rules.",
	}
	manageSnapFlag := cli.BoolFlag{
		Name:  "managesnapshot",
		Usage: "Permission to manage snapshot.",
	}
	manageUserFlag := cli.BoolFlag{
		Name:  "manageuser",
		Usage: "Permission to manage users.",
	}
	modifyFlag := cli.BoolFlag{
		Name:  "modify",
		Usage: "Permission to modify.",
	}
	reinstallFlag := cli.BoolFlag{
		Name:  "reinstall",
		Usage: "Permission to reinstall.",
	}
	releaseFlag := cli.BoolFlag{
		Name:  "release",
		Usage: "Permission to release.",
	}
	resizeFlag := cli.BoolFlag{
		Name:  "resize",
		Usage: "Permission to resize.",
	}
	resourcesFlag := cli.BoolFlag{
		Name:  "resources",
		Usage: "Permission to modify resources.",
	}
	restartFlag := cli.BoolFlag{
		Name:  "restart",
		Usage: "Permission to restart.",
	}
	setDescFlag := cli.BoolFlag{
		Name:  "setdesc",
		Usage: "Permission to set description.",
	}
	setEmailFlag := cli.BoolFlag{
		Name:  "setemail",
		Usage: "Permission to set email.",
	}
	setNameFlag := cli.BoolFlag{
		Name:  "setname",
		Usage: "Permission to set name.",
	}
	setNetInfoFlag := cli.BoolFlag{
		Name:  "setnetinfo",
		Usage: "Permission to set network info.",
	}
	setPassFlag := cli.BoolFlag{
		Name:  "setpassword",
		Usage: "Permission to set password.",
	}
	setDNSFlag := cli.BoolFlag{
		Name:  "setdns",
		Usage: "Permission to set reverse DNS.",
	}
	showFlag := cli.BoolFlag{
		Name:  "show",
		Usage: "Permission to show.",
	}
	shutdownFlag := cli.BoolFlag{
		Name:  "shutdown",
		Usage: "Permission to shutdown.",
	}
	startFlag := cli.BoolFlag{
		Name:  "start",
		Usage: "Permission to start.",
	}
	userIDFlag := cli.StringFlag{
		Name:  "userid",
		Usage: "ID of the user.",
	}

	roleOps = []cli.Command{
		{
			Name:        "role",
			Description: "1&1 role operations",
			Usage:       "Role operations.",
			Subcommands: []cli.Command{
				{
					Name:  "clone",
					Usage: "Clones role.",
					Flags: []cli.Flag{
						roleIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the new role.",
						},
					},
					Action: cloneRole,
				},
				{
					Name:  "create",
					Usage: "Creates new role.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the role.",
						},
					},
					Action: createRole,
				},
				{
					Name:   "info",
					Usage:  "Shows information about role.",
					Flags:  []cli.Flag{roleIdFlag},
					Action: showRole,
				},
				{
					Name:   "list",
					Usage:  "Lists all available roles.",
					Flags:  queryFlags,
					Action: listRoles,
				},
				{
					Name:  "modify",
					Usage: "Modifies role configuration.",
					Flags: []cli.Flag{
						roleIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "New name of the role.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "New description of the role.",
						},
						cli.StringFlag{
							Name:  "state",
							Usage: "New state of the role (ACTIVE|DISABLE).",
						},
					},
					Action: modifyRole,
				},
				{
					Name:  "permissions",
					Usage: "Manages role's permissions.",
					Subcommands: []cli.Command{
						{
							Name:   back_info,
							Usage:  "Shows permissions for backups.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  back_mod,
							Usage: "Modifies permissions for backups.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								createFlag,
								deleteFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   fw_info,
							Usage:  "Shows permissions for firewall policies.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  fw_mod,
							Usage: "Modifies permissions for firewall policies.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								cloneFlag,
								createFlag,
								deleteFlag,
								manageSrvIPFlag,
								manageRuleFlag,
								setDescFlag,
								setNameFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   img_info,
							Usage:  "Shows permissions for images.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  img_mod,
							Usage: "Modifies permissions for images.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								createFlag,
								deleteFlag,
								disableAutoCreateFlag,
								setDescFlag,
								setNameFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   "info",
							Usage:  "Shows all permissions.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:   inv_info,
							Usage:  "Shows permissions for invoice.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:   inv_mod,
							Usage:  "Modifies permissions for invoice.",
							Flags:  []cli.Flag{roleIdFlag, allFlag, showFlag},
							Action: modifyPerm,
						},
						{
							Name:   ip_info,
							Usage:  "Shows permissions for IPs.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  ip_mod,
							Usage: "Modifies permissions for IPs.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								createFlag,
								deleteFlag,
								releaseFlag,
								setDNSFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   lb_info,
							Usage:  "Shows permissions for load balancers.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  lb_mod,
							Usage: "Modifies permissions for load balancers.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								createFlag,
								deleteFlag,
								manageSrvIPFlag,
								manageRuleFlag,
								modifyFlag,
								setDescFlag,
								setNameFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   log_info,
							Usage:  "Shows permissions for logs.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:   log_mod,
							Usage:  "Modifies permissions for logs.",
							Flags:  []cli.Flag{roleIdFlag, allFlag, showFlag},
							Action: modifyPerm,
						},
						{
							Name:   mc_info,
							Usage:  "Shows permissions for monitoring center.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:   mc_mod,
							Usage:  "Modifies permissions for monitoring center.",
							Flags:  []cli.Flag{roleIdFlag, allFlag, showFlag},
							Action: modifyPerm,
						},
						{
							Name:   mp_info,
							Usage:  "Shows permissions for monitoring policies.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  mp_mod,
							Usage: "Modifies permissions for monitoring policies.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								cloneFlag,
								createFlag,
								deleteFlag,
								manageSrvFlag,
								managePortFlag,
								manageProcFlag,
								resourcesFlag,
								setDescFlag,
								setEmailFlag,
								setNameFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   pn_info,
							Usage:  "Shows permissions for private networks.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  pn_mod,
							Usage: "Modifies permissions for private networks.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								createFlag,
								deleteFlag,
								manageSrvFlag,
								setDescFlag,
								setNameFlag,
								setNetInfoFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   role_info,
							Usage:  "Shows permissions for roles.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  role_mod,
							Usage: "Modifies permissions for roles.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								cloneFlag,
								createFlag,
								deleteFlag,
								manageUserFlag,
								modifyFlag,
								setDescFlag,
								setNameFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   ser_info,
							Usage:  "Shows permissions for servers.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  ser_mod,
							Usage: "Modifies permissions for servers.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								assignipFlag,
								cloneFlag,
								createFlag,
								deleteFlag,
								kvmFlag,
								manageDVDFlag,
								manageSnapFlag,
								reinstallFlag,
								resizeFlag,
								restartFlag,
								setDescFlag,
								setNameFlag,
								showFlag,
								shutdownFlag,
								startFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   set_all,
							Usage:  "Sets role's all permissions.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: modifyPerm,
						},
						{
							Name:   ss_info,
							Usage:  "Shows permissions for shared storages.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  ss_mod,
							Usage: "Modifies permissions for shared storages.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								accessFlag,
								createFlag,
								deleteFlag,
								manageSrvFlag,
								resizeFlag,
								setDescFlag,
								setNameFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   unset_all,
							Usage:  "Unsets role's all permissions.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: modifyPerm,
						},
						{
							Name:   usg_info,
							Usage:  "Shows permissions for usages.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:   usg_mod,
							Usage:  "Modifies permissions for usages.",
							Flags:  []cli.Flag{roleIdFlag, allFlag, showFlag},
							Action: modifyPerm,
						},
						{
							Name:   user_info,
							Usage:  "Shows permissions for users.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  user_mod,
							Usage: "Modifies permissions for users.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								changerRoleFlag,
								createFlag,
								deleteFlag,
								disableFlag,
								enableFlag,
								manageAPIFlag,
								setDescFlag,
								setEmailFlag,
								setPassFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
						{
							Name:   vpn_info,
							Usage:  "Shows permissions for VPNs.",
							Flags:  []cli.Flag{roleIdFlag},
							Action: showPerm,
						},
						{
							Name:  vpn_mod,
							Usage: "Modifies permissions for VPNs.",
							Flags: []cli.Flag{
								roleIdFlag,
								allFlag,
								createFlag,
								deleteFlag,
								downloadFlag,
								setDescFlag,
								setNameFlag,
								showFlag,
							},
							Action: modifyPerm,
						},
					},
				},
				{
					Name:   "rm",
					Usage:  "Deletes role.",
					Flags:  []cli.Flag{roleIdFlag},
					Action: deleteRole,
				},
				{
					Name:  "useradd",
					Usage: "Adds users to role.",
					Flags: []cli.Flag{
						roleIdFlag,
						cli.StringSliceFlag{
							Name:  "userid",
							Usage: "List od user IDs.",
						},
					},
					Action: addRoleUsers,
				},
				{
					Name:   "userinfo",
					Usage:  "Shows information about role's user.",
					Flags:  []cli.Flag{roleIdFlag, userIDFlag},
					Action: showRoleUser,
				},
				{
					Name:   "userlist",
					Usage:  "Lists role's users.",
					Flags:  []cli.Flag{roleIdFlag},
					Action: listRoleUsers,
				},
				{
					Name:   "userrm",
					Usage:  "Removes role's user.",
					Flags:  []cli.Flag{roleIdFlag, userIDFlag},
					Action: removeRoleUser,
				},
			},
		},
	}
}

func cloneRole(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	roleName := getRequiredOption(ctx, "name")
	role, err := api.CloneRole(id, roleName)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}

func createRole(ctx *cli.Context) {
	roleName := getRequiredOption(ctx, "name")
	_, role, err := api.CreateRole(roleName)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}

func listRoles(ctx *cli.Context) {
	roles, err := api.ListRoles(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(roles))
	for i, role := range roles {
		var isDefault = "no"
		if *role.Default == 1 {
			isDefault = "yes"
		}
		data[i] = []string{
			role.Id,
			role.Name,
			formatDateTime(time.RFC3339, role.CreationDate),
			role.State,
			isDefault,
		}
	}
	header := []string{"ID", "Name", "Creation Date", "State", "Default"}
	output(ctx, roles, "", false, &header, &data)
}

func showRole(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	role, err := api.GetRole(id)
	exitOnError(err)
	output(ctx, role, "", true, nil, nil)
}

func deleteRole(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	role, err := api.DeleteRole(id)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}

func modifyRole(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	state := strings.ToUpper(ctx.String("state"))
	if state != "ACTIVE" && state != "DISABLE" {
		exitOnError(fmt.Errorf("Invalid role state. Valid states are 'ACTIVE' and 'DISABLE'."))
	}
	role, err := api.ModifyRole(id, ctx.String("name"), ctx.String("desc"), state)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}

func addRoleUsers(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	userIDs := getStringSliceOption(ctx, "userid", true)
	role, err := api.AssignRoleUsers(id, userIDs)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}

func listRoleUsers(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	users, err := api.ListRoleUsers(id)
	exitOnError(err)
	data := make([][]string, len(users))
	for i, u := range users {
		data[i] = []string{u.Id, u.Name}
	}
	header := []string{"ID", "Name"}
	output(ctx, users, "", false, &header, &data)
}

func removeRoleUser(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	userID := getRequiredOption(ctx, "userid")
	role, err := api.RemoveRoleUser(id, userID)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}
func showRoleUser(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	userID := getRequiredOption(ctx, "userid")
	role, err := api.GetRoleUser(id, userID)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}

func showPerm(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	permissions, err := api.GetRolePermissions(id)
	exitOnError(err)

	switch ctx.Command.Name {
	case back_info:
		output(ctx, permissions.Backups, "", true, nil, nil)
		break
	case fw_info:
		output(ctx, permissions.Firewalls, "", true, nil, nil)
		break
	case img_info:
		output(ctx, permissions.Images, "", true, nil, nil)
		break
	case "info":
		output(ctx, permissions, "", true, nil, nil)
		break
	case inv_info:
		output(ctx, permissions.Invoice, "", true, nil, nil)
		break
	case ip_info:
		output(ctx, permissions.IPs, "", true, nil, nil)
		break
	case lb_info:
		output(ctx, permissions.LoadBalancers, "", true, nil, nil)
		break
	case log_info:
		output(ctx, permissions.Logs, "", true, nil, nil)
		break
	case mc_info:
		output(ctx, permissions.MonitorCenter, "", true, nil, nil)
		break
	case mp_info:
		output(ctx, permissions.MonitorPolicies, "", true, nil, nil)
		break
	case pn_info:
		output(ctx, permissions.PrivateNetworks, "", true, nil, nil)
		break
	case role_info:
		output(ctx, permissions.Roles, "", true, nil, nil)
		break
	case ser_info:
		output(ctx, permissions.Servers, "", true, nil, nil)
		break
	case ss_info:
		output(ctx, permissions.SharedStorage, "", true, nil, nil)
		break
	case usg_info:
		output(ctx, permissions.Usages, "", true, nil, nil)
		break
	case user_info:
		output(ctx, permissions.Users, "", true, nil, nil)
		break
	case vpn_info:
		output(ctx, permissions.VPNs, "", true, nil, nil)
		break
	}
}

func setPerm(ctx *cli.Context, value bool, perms *oneandone.Permissions) {
	switch ctx.Command.Name {
	case back_mod:
		perms.Backups = new(oneandone.BackupPerm)
		perms.Backups.SetAll(value)
		break
	case fw_mod:
		perms.Firewalls = new(oneandone.FirewallPerm)
		perms.Firewalls.SetAll(value)
		break
	case img_mod:
		perms.Images = new(oneandone.ImagePerm)
		perms.Images.SetAll(value)
		break
	case set_all:
		fallthrough
	case unset_all:
		perms.SetAll(value)
		break
	case inv_mod:
		perms.Invoice = new(oneandone.InvoicePerm)
		perms.Invoice.SetAll(value)
		break
	case ip_mod:
		perms.IPs = new(oneandone.IPPerm)
		perms.IPs.SetAll(value)
		break
	case lb_mod:
		perms.LoadBalancers = new(oneandone.LoadBalancerPerm)
		perms.LoadBalancers.SetAll(value)
		break
	case log_mod:
		perms.Logs = new(oneandone.LogPerm)
		perms.Logs.SetAll(value)
		break
	case mc_mod:
		perms.MonitorCenter = new(oneandone.MonitorCenterPerm)
		perms.MonitorCenter.SetAll(value)
		break
	case mp_mod:
		perms.MonitorPolicies = new(oneandone.MonitorPolicyPerm)
		perms.MonitorPolicies.SetAll(value)
		break
	case pn_mod:
		perms.PrivateNetworks = new(oneandone.PrivateNetworkPerm)
		perms.PrivateNetworks.SetAll(value)
		break
	case role_mod:
		perms.Roles = new(oneandone.RolePerm)
		perms.Roles.SetAll(value)
		break
	case ser_mod:
		perms.Servers = new(oneandone.ServerPerm)
		perms.Servers.SetAll(value)
		break
	case ss_mod:
		perms.SharedStorage = new(oneandone.SharedStoragePerm)
		perms.SharedStorage.SetAll(value)
		break
	case usg_mod:
		perms.Usages = new(oneandone.UsagePerm)
		perms.Usages.SetAll(value)
		break
	case user_mod:
		perms.Users = new(oneandone.UserPerm)
		perms.Users.SetAll(value)
		break
	case vpn_mod:
		perms.VPNs = new(oneandone.VPNPerm)
		perms.VPNs.SetAll(value)
		break
	}
}

func modifyPerm(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	ps := new(oneandone.Permissions)

	if ctx.IsSet("all") {
		setPerm(ctx, ctx.Bool("all"), ps)
	} else if ctx.Command.Name == set_all {
		setPerm(ctx, true, ps)
	} else if ctx.Command.Name == unset_all {
		setPerm(ctx, false, ps)
	} else if ctx.Command.Name == back_mod {
		setBackupPerm(ctx, ps)
	} else if ctx.Command.Name == fw_mod {
		setFirewallPerm(ctx, ps)
	} else if ctx.Command.Name == img_mod {
		setImagePerm(ctx, ps)
	} else if ctx.Command.Name == inv_mod {
		setInvoicePerm(ctx, ps)
	} else if ctx.Command.Name == ip_mod {
		setIPPerm(ctx, ps)
	} else if ctx.Command.Name == lb_mod {
		setLoadBalancerPerm(ctx, ps)
	} else if ctx.Command.Name == log_mod {
		setLogPerm(ctx, ps)
	} else if ctx.Command.Name == mc_mod {
		setMonCenterPerm(ctx, ps)
	} else if ctx.Command.Name == mp_mod {
		setMonPolicyPerm(ctx, ps)
	} else if ctx.Command.Name == pn_mod {
		setPrivateNetPerm(ctx, ps)
	} else if ctx.Command.Name == role_mod {
		setRolePerm(ctx, ps)
	} else if ctx.Command.Name == ser_mod {
		setServerPerm(ctx, ps)
	} else if ctx.Command.Name == ss_mod {
		setSharedStoragePerm(ctx, ps)
	} else if ctx.Command.Name == usg_mod {
		setUsagePerm(ctx, ps)
	} else if ctx.Command.Name == user_mod {
		setUserPerm(ctx, ps)
	} else if ctx.Command.Name == vpn_mod {
		setVPNPerm(ctx, ps)
	} else {
		return
	}
	role, err := api.ModifyRolePermissions(id, ps)
	exitOnError(err)
	output(ctx, role, "OK", false, nil, nil)
}

func setBackupPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Backups = new(oneandone.BackupPerm)
	ps.Backups.Create = ctx.Bool("create")
	ps.Backups.Delete = ctx.Bool("delete")
	ps.Backups.Show = ctx.Bool("show")
}

func setFirewallPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Firewalls = new(oneandone.FirewallPerm)
	ps.Firewalls.Clone = ctx.Bool("clone")
	ps.Firewalls.Create = ctx.Bool("create")
	ps.Firewalls.Delete = ctx.Bool("delete")
	ps.Firewalls.ManageAttachedServerIPs = ctx.Bool("manageip")
	ps.Firewalls.ManageRules = ctx.Bool("managerule")
	ps.Firewalls.SetDescription = ctx.Bool("setdesc")
	ps.Firewalls.SetName = ctx.Bool("setname")
	ps.Firewalls.Show = ctx.Bool("show")
}

func setImagePerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Images = new(oneandone.ImagePerm)
	ps.Images.Create = ctx.Bool("create")
	ps.Images.Delete = ctx.Bool("delete")
	ps.Images.DisableAutoCreate = ctx.Bool("noautocreate")
	ps.Images.SetDescription = ctx.Bool("setdesc")
	ps.Images.SetName = ctx.Bool("setname")
	ps.Images.Show = ctx.Bool("show")
}

func setInvoicePerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Invoice = new(oneandone.InvoicePerm)
	ps.Invoice.Show = ctx.Bool("show")
}

func setIPPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.IPs = new(oneandone.IPPerm)
	ps.IPs.Create = ctx.Bool("create")
	ps.IPs.Delete = ctx.Bool("delete")
	ps.IPs.Release = ctx.Bool("release")
	ps.IPs.SetReverseDNS = ctx.Bool("setdns")
	ps.IPs.Show = ctx.Bool("show")
}

func setLoadBalancerPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.LoadBalancers = new(oneandone.LoadBalancerPerm)
	ps.LoadBalancers.Create = ctx.Bool("create")
	ps.LoadBalancers.Delete = ctx.Bool("delete")
	ps.LoadBalancers.ManageAttachedServerIPs = ctx.Bool("manageip")
	ps.LoadBalancers.ManageRules = ctx.Bool("managerule")
	ps.LoadBalancers.Modify = ctx.Bool("modify")
	ps.LoadBalancers.SetDescription = ctx.Bool("setdesc")
	ps.LoadBalancers.SetName = ctx.Bool("setname")
	ps.LoadBalancers.Show = ctx.Bool("show")
}

func setLogPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Logs = new(oneandone.LogPerm)
	ps.Logs.Show = ctx.Bool("show")
}

func setMonCenterPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.MonitorCenter = new(oneandone.MonitorCenterPerm)
	ps.MonitorCenter.Show = ctx.Bool("show")
}

func setMonPolicyPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.MonitorPolicies = new(oneandone.MonitorPolicyPerm)
	ps.MonitorPolicies.Clone = ctx.Bool("clone")
	ps.MonitorPolicies.Create = ctx.Bool("create")
	ps.MonitorPolicies.Delete = ctx.Bool("delete")
	ps.MonitorPolicies.ManageAttachedServers = ctx.Bool("manageserver")
	ps.MonitorPolicies.ManagePorts = ctx.Bool("manageport")
	ps.MonitorPolicies.ManageProcesses = ctx.Bool("manageprocess")
	ps.MonitorPolicies.ModifyResources = ctx.Bool("resources")
	ps.MonitorPolicies.SetDescription = ctx.Bool("setdesc")
	ps.MonitorPolicies.SetEmail = ctx.Bool("setemail")
	ps.MonitorPolicies.SetName = ctx.Bool("setname")
	ps.MonitorPolicies.Show = ctx.Bool("show")
}

func setPrivateNetPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.PrivateNetworks = new(oneandone.PrivateNetworkPerm)
	ps.PrivateNetworks.Create = ctx.Bool("create")
	ps.PrivateNetworks.Delete = ctx.Bool("delete")
	ps.PrivateNetworks.ManageAttachedServers = ctx.Bool("manageserver")
	ps.PrivateNetworks.SetDescription = ctx.Bool("setdesc")
	ps.PrivateNetworks.SetName = ctx.Bool("setname")
	ps.PrivateNetworks.SetNetworkInfo = ctx.Bool("setnetinfo")
	ps.PrivateNetworks.Show = ctx.Bool("show")
}

func setRolePerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Roles = new(oneandone.RolePerm)
	ps.Roles.Clone = ctx.Bool("clone")
	ps.Roles.Create = ctx.Bool("create")
	ps.Roles.Delete = ctx.Bool("delete")
	ps.Roles.ManageUsers = ctx.Bool("manageuser")
	ps.Roles.Modify = ctx.Bool("modify")
	ps.Roles.SetDescription = ctx.Bool("setdesc")
	ps.Roles.SetName = ctx.Bool("setname")
	ps.Roles.Show = ctx.Bool("show")
}

func setServerPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Servers = new(oneandone.ServerPerm)
	ps.Servers.AccessKVMConsole = ctx.Bool("kvm")
	ps.Servers.AssignIP = ctx.Bool("assignip")
	ps.Servers.Clone = ctx.Bool("clone")
	ps.Servers.Create = ctx.Bool("create")
	ps.Servers.Delete = ctx.Bool("delete")
	ps.Servers.ManageDVD = ctx.Bool("managedvd")
	ps.Servers.ManageSnapshot = ctx.Bool("managesnapshot")
	ps.Servers.Reinstall = ctx.Bool("reinstall")
	ps.Servers.Resize = ctx.Bool("resize")
	ps.Servers.Restart = ctx.Bool("restart")
	ps.Servers.SetDescription = ctx.Bool("setdesc")
	ps.Servers.SetName = ctx.Bool("setname")
	ps.Servers.Show = ctx.Bool("show")
	ps.Servers.Shutdown = ctx.Bool("shutdown")
	ps.Servers.Start = ctx.Bool("start")
}

func setSharedStoragePerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.SharedStorage = new(oneandone.SharedStoragePerm)
	ps.SharedStorage.Access = ctx.Bool("access")
	ps.SharedStorage.Create = ctx.Bool("create")
	ps.SharedStorage.Delete = ctx.Bool("delete")
	ps.SharedStorage.ManageAttachedServers = ctx.Bool("manageserver")
	ps.SharedStorage.Resize = ctx.Bool("resize")
	ps.SharedStorage.SetDescription = ctx.Bool("setdesc")
	ps.SharedStorage.SetName = ctx.Bool("setname")
	ps.SharedStorage.Show = ctx.Bool("show")
}

func setUsagePerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Usages = new(oneandone.UsagePerm)
	ps.Usages.Show = ctx.Bool("show")
}

func setUserPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.Users = new(oneandone.UserPerm)
	ps.Users.ChangeRole = ctx.Bool("changerole")
	ps.Users.Create = ctx.Bool("create")
	ps.Users.Delete = ctx.Bool("delete")
	ps.Users.Enable = ctx.Bool("enable")
	ps.Users.Disable = ctx.Bool("disable")
	ps.Users.ManageAPI = ctx.Bool("manageapi")
	ps.Users.SetDescription = ctx.Bool("setdesc")
	ps.Users.SetEmail = ctx.Bool("setemail")
	ps.Users.SetPassword = ctx.Bool("setpassword")
	ps.Users.Show = ctx.Bool("show")
}

func setVPNPerm(ctx *cli.Context, ps *oneandone.Permissions) {
	ps.VPNs = new(oneandone.VPNPerm)
	ps.VPNs.Create = ctx.Bool("create")
	ps.VPNs.Delete = ctx.Bool("delete")
	ps.VPNs.DownloadFile = ctx.Bool("downloadfile")
	ps.VPNs.SetDescription = ctx.Bool("setdesc")
	ps.VPNs.SetName = ctx.Bool("setname")
	ps.VPNs.Show = ctx.Bool("show")
}
