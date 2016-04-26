package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var firewallOps []cli.Command

func init() {
	fRuleFlags := []cli.Flag{
		cli.IntSliceFlag{
			Name:  "portfrom",
			Usage: "First port in range.",
		},
		cli.IntSliceFlag{
			Name:  "portto",
			Usage: "Second port in range.",
		},
		cli.StringSliceFlag{
			Name:  "protocol",
			Usage: "Internet protocol: TCP, UDP, TCP/UDP, ICMP, IPSEC or GRE",
		},
		cli.StringSliceFlag{
			Name:  "source",
			Usage: "IPs from which access is available. Default is 0.0.0.0 and all IPs are allowed.",
		},
	}

	fIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the firewall policy.",
	}
	fNameFlag := cli.StringFlag{
		Name:  "desc, d",
		Usage: "Description of the firewall policy.",
	}
	fDescFlag := cli.StringFlag{
		Name:  "name, n",
		Usage: "Name of the firewall policy.",
	}
	fIpIdFlag := cli.StringFlag{
		Name:  "ipid",
		Usage: "ID of the server's IP.",
	}
	fRuleIdFlag := cli.StringFlag{
		Name:  "ruleid",
		Usage: "ID of the rule.",
	}

	fCreateFlags := []cli.Flag{fDescFlag, fNameFlag}
	fCreateFlags = append(fCreateFlags, fRuleFlags...)

	firewallOps = []cli.Command{
		{
			Name:        "firewall",
			Description: "1&1 firewall policy operations",
			Usage:       "Firewall policy operations.",
			Subcommands: []cli.Command{
				{
					Name:  "assign",
					Usage: "Assigns servers/IPs to firewall policy.",
					Flags: []cli.Flag{
						fIdFlag,
						cli.StringSliceFlag{
							Name:  "ipid",
							Usage: "List of server/IP IDs.",
						},
					},
					Action: assignFirewallServers,
				},
				{
					Name:   "create",
					Usage:  "Creates new firewall policy.",
					Flags:  fCreateFlags,
					Action: createFirewall,
				},
				{
					Name:   "info",
					Usage:  "Shows information about firewall policy.",
					Flags:  []cli.Flag{fIdFlag},
					Action: showFirewall,
				},
				{
					Name:   "list",
					Usage:  "Lists available firewall policies.",
					Flags:  queryFlags,
					Action: listFirewalls,
				},
				{
					Name:   "rule",
					Usage:  "Shows information about firewall policy rule.",
					Flags:  []cli.Flag{fIdFlag, fRuleIdFlag},
					Action: showFirewallRule,
				},
				{
					Name:   "ruleadd",
					Usage:  "Adds new rules to firewall policy.",
					Flags:  append([]cli.Flag{fIdFlag}, fRuleFlags...),
					Action: addFirewallRules,
				},
				{
					Name:   "rulerm",
					Usage:  "Removes rule from firewall policy.",
					Flags:  []cli.Flag{fIdFlag, fRuleIdFlag},
					Action: removeFirewallRule,
				},
				{
					Name:   "rules",
					Usage:  "Lists firewall policy rules.",
					Flags:  []cli.Flag{fIdFlag},
					Action: listFirewallRules,
				},
				{
					Name:   "server",
					Usage:  "Shows information about server attached to firewall policy.",
					Flags:  []cli.Flag{fIdFlag, fIpIdFlag},
					Action: showFirewallServer,
				},
				{
					Name:   "servers",
					Usage:  "Lists servers/IPs attached to firewall policies.",
					Flags:  []cli.Flag{fIdFlag},
					Action: listFirewallServers,
				},
				{
					Name:   "rm",
					Usage:  "Removes firewall policy.",
					Flags:  []cli.Flag{fIdFlag},
					Action: deleteFirewall,
				},
				{
					Name:   "unassign",
					Usage:  "Unassigns servers/IPs to firewall policy.",
					Flags:  []cli.Flag{fIdFlag, fIpIdFlag},
					Action: removeFirewallServer,
				},
				{
					Name:   "update",
					Usage:  "Updates name and description of firewall policy.",
					Flags:  []cli.Flag{fIdFlag, fDescFlag, fNameFlag},
					Action: updateFirewall,
				},
			},
		},
	}
}

// Helper function
func parseFirewallRules(ctx *cli.Context) []oneandone.FirewallPolicyRule {
	portsFrom := getIntSliceOption(ctx, "portfrom", true)
	portsTo := getIntSliceOption(ctx, "portto", true)
	protocols := getStringSliceOption(ctx, "protocol", true)

	if len(portsFrom) != len(portsTo) {
		exitOnError(fmt.Errorf("equal number of --portfrom and --portto arguments must be specified"))
	}

	sources := getStringSliceOption(ctx, "source", false)

	var rules []oneandone.FirewallPolicyRule
	for i := 0; i < len(protocols); i++ {
		protocols[i] = strings.ToUpper(protocols[i])

		switch protocols[i] {
		case "TCP":
			break
		case "UDP":
			break
		case "TCP/UDP":
			break
		case "ICMP":
			break
		case "IPSEC":
			break
		case "GRE":
			break
		default:
			exitOnError(fmt.Errorf("Invalid value for --protocol flag. Valid values are TCP, UDP, TCP/UDP, ICMP, IPSEC or GRE."))
		}

		var source string
		var fromPort *int
		var toPort *int

		if len(sources) > i {
			source = sources[i]
		}
		if protocols[i] != "ICMP" && protocols[i] != "IPSEC" && protocols[i] != "GRE" && len(portsFrom) > i {
			fromPort = oneandone.Int2Pointer(validateIntRange("portfrom", portsFrom[i], 1, 65535))
			toPort = oneandone.Int2Pointer(validateIntRange("portto", portsTo[i], 1, 65535))
		}

		rule := oneandone.FirewallPolicyRule{
			PortFrom: fromPort,
			PortTo:   toPort,
			Protocol: protocols[i],
			SourceIp: source,
		}
		rules = append(rules, rule)
	}
	return rules
}

func listFirewalls(ctx *cli.Context) {
	policies, err := api.ListFirewallPolicies(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(policies))
	for i, policy := range policies {
		data[i] = []string{policy.Id, policy.Name, policy.State}
	}
	header := []string{"ID", "Name", "State"}
	output(ctx, policies, "", false, &header, &data)
}

func showFirewall(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	firewall, err := api.GetFirewallPolicy(fwId)
	exitOnError(err)
	output(ctx, firewall, "", true, nil, nil)
}

func createFirewall(ctx *cli.Context) {
	name := getRequiredOption(ctx, "name")
	rules := parseFirewallRules(ctx)

	req := oneandone.FirewallPolicyRequest{
		Name:        name,
		Description: ctx.String("desc"),
		Rules:       rules,
	}
	_, firewall, err := api.CreateFirewallPolicy(&req)
	exitOnError(err)
	output(ctx, firewall, okWaitMessage, false, nil, nil)
}

func updateFirewall(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	firewall, err := api.UpdateFirewallPolicy(fwId, ctx.String("name"), ctx.String("desc"))
	exitOnError(err)
	output(ctx, firewall, "", false, nil, nil)
}

func deleteFirewall(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	firewall, err := api.DeleteFirewallPolicy(fwId)
	exitOnError(err)
	output(ctx, firewall, okWaitMessage, false, nil, nil)
}

func listFirewallServers(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	servers, err := api.ListFirewallPolicyServerIps(fwId)
	exitOnError(err)
	data := make([][]string, len(servers))
	for i, server := range servers {
		data[i] = []string{server.Id, server.ServerName, server.Ip}
	}
	header := []string{"ID", "Name", "IP Address"}
	output(ctx, servers, "", false, &header, &data)
}

func assignFirewallServers(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	ipIds := getStringSliceOption(ctx, "ipid", true)
	firewall, err := api.AddFirewallPolicyServerIps(fwId, ipIds)
	exitOnError(err)
	output(ctx, firewall, okWaitMessage, false, nil, nil)
}

func showFirewallServer(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	server, err := api.GetFirewallPolicyServerIp(fwId, ipId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func removeFirewallServer(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	firewall, err := api.DeleteFirewallPolicyServerIp(fwId, ipId)
	exitOnError(err)
	output(ctx, firewall, okWaitMessage, false, nil, nil)
}

func listFirewallRules(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	rules, err := api.ListFirewallPolicyRules(fwId)
	exitOnError(err)
	data := make([][]string, len(rules))
	for i, rule := range rules {
		var portFrom, portTo string
		if rule.PortFrom != nil {
			portFrom = strconv.Itoa(*rule.PortFrom)
		}
		if rule.PortTo != nil {
			portTo = strconv.Itoa(*rule.PortTo)
		}
		data[i] = []string{
			rule.Id,
			portFrom,
			portTo,
			rule.Protocol,
			rule.SourceIp,
		}
	}
	header := []string{"ID", "Port From", "Port To", "Protocol", "Source IP"}
	output(ctx, rules, "", false, &header, &data)
}

func showFirewallRule(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	ruleId := getRequiredOption(ctx, "ruleid")
	rule, err := api.GetFirewallPolicyRule(fwId, ruleId)
	exitOnError(err)
	output(ctx, rule, "", true, nil, nil)
}

func addFirewallRules(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	rules := parseFirewallRules(ctx)
	firewall, err := api.AddFirewallPolicyRules(fwId, rules)
	exitOnError(err)
	output(ctx, firewall, okWaitMessage, false, nil, nil)
}

func removeFirewallRule(ctx *cli.Context) {
	fwId := getRequiredOption(ctx, "id")
	ruleId := getRequiredOption(ctx, "ruleid")
	firewall, err := api.DeleteFirewallPolicyRule(fwId, ruleId)
	exitOnError(err)
	output(ctx, firewall, okWaitMessage, false, nil, nil)
}
