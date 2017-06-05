package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var loadbalancerOps []cli.Command

func init() {
	lbRuleFlags := []cli.Flag{
		cli.IntSliceFlag{
			Name:  "portbalancer",
			Usage: "First port in range.",
		},
		cli.IntSliceFlag{
			Name:  "portserver",
			Usage: "Second port in range.",
		},
		cli.StringSliceFlag{
			Name:  "protocol",
			Usage: "Internet protocol: TCP or UDP",
		},
		cli.StringSliceFlag{
			Name:  "source",
			Usage: "IPs from which access is available. Default is 0.0.0.0 and all IPs are allowed.",
		},
	}

	lbIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the load balancer.",
	}
	dcIdFlag := cli.StringFlag{
		Name:  "datacenterid",
		Usage: "Data center ID of the load balancer.",
	}
	lbNameFlag := cli.StringFlag{
		Name:  "desc, d",
		Usage: "Description of the load balancer.",
	}
	lbDescFlag := cli.StringFlag{
		Name:  "name, n",
		Usage: "Name of the load balancer.",
	}
	lbIpIdFlag := cli.StringFlag{
		Name:  "ipid",
		Usage: "ID of the server's IP.",
	}
	lbRuleIdFlag := cli.StringFlag{
		Name:  "ruleid",
		Usage: "ID of the rule.",
	}
	hctFlag := cli.StringFlag{
		Name:  "hctest",
		Usage: "Health check test: NONE, TCP or ICMP.",
	}
	hciFlag := cli.StringFlag{
		Name:  "hctime",
		Usage: "Health check period in seconds.",
	}
	persistenceFlag := cli.BoolFlag{
		Name:  "persistence",
		Usage: "Is persistence enabled.",
	}
	persistTimeFlag := cli.StringFlag{
		Name:  "persint",
		Usage: "Persistence time in seconds: 30 - 1200.",
	}
	lbMethodFlag := cli.StringFlag{
		Name:  "method",
		Usage: "Balancing procedure: ROUND_ROBIN or LEAST_CONNECTIONS (aliases RR and LC).",
	}
	hcPathFlag := cli.StringFlag{
		Name:  "hcpath",
		Usage: "URL to call for checking. Required for HTTP health check.",
	}
	hcParseFlag := cli.StringFlag{
		Name:  "hcregex",
		Usage: "Regular expression to check. Required for HTTP health check.",
	}

	lbCreateFlags := []cli.Flag{
		dcIdFlag,
		lbDescFlag,
		hcPathFlag,
		hcParseFlag,
		hctFlag,
		hciFlag,
		lbMethodFlag,
		lbNameFlag,
		persistenceFlag,
		persistTimeFlag,
	}
	lbCreateFlags = append(lbCreateFlags, lbRuleFlags...)

	loadbalancerOps = []cli.Command{
		{
			Name:        "loadbalancer",
			Description: "1&1 load balancer operations",
			Usage:       "Load balancer operations.",
			Subcommands: []cli.Command{
				{
					Name:  "assign",
					Usage: "Assigns servers/IPs to load balancer.",
					Flags: []cli.Flag{
						lbIdFlag,
						cli.StringSliceFlag{
							Name:  "ipid",
							Usage: "List of server/IP IDs.",
						},
					},
					Action: assignLoadBalancerServers,
				},
				{
					Name:   "create",
					Usage:  "Creates new load balancer.",
					Flags:  lbCreateFlags,
					Action: createLoadBalancer,
				},
				{
					Name:   "info",
					Usage:  "Shows information about load balancer.",
					Flags:  []cli.Flag{lbIdFlag},
					Action: showLoadBalancer,
				},
				{
					Name:   "list",
					Usage:  "Lists available load balancers.",
					Flags:  queryFlags,
					Action: listLoadBalancers,
				},
				{
					Name:   "rule",
					Usage:  "Shows information about load balancer rule.",
					Flags:  []cli.Flag{lbIdFlag, lbRuleIdFlag},
					Action: showLoadBalancerRule,
				},
				{
					Name:   "ruleadd",
					Usage:  "Adds new rules to load balancer.",
					Flags:  append([]cli.Flag{lbIdFlag}, lbRuleFlags...),
					Action: addLoadBalancerRules,
				},
				{
					Name:   "rulerm",
					Usage:  "Removes rule from load balancer.",
					Flags:  []cli.Flag{lbIdFlag, lbRuleIdFlag},
					Action: removeLoadBalancerRule,
				},
				{
					Name:   "rules",
					Usage:  "Lists load balancer rules.",
					Flags:  []cli.Flag{lbIdFlag},
					Action: listLoadBalancerRules,
				},
				{
					Name:   "server",
					Usage:  "Shows information about server attached to load balancer.",
					Flags:  []cli.Flag{lbIdFlag, lbIpIdFlag},
					Action: showLoadBalancerServer,
				},
				{
					Name:   "servers",
					Usage:  "Lists servers/IPs attached to load balancers.",
					Flags:  []cli.Flag{lbIdFlag},
					Action: listLoadBalancerServers,
				},
				{
					Name:   "rm",
					Usage:  "Removes load balancer.",
					Flags:  []cli.Flag{lbIdFlag},
					Action: deleteLoadBalancer,
				},
				{
					Name:   "unassign",
					Usage:  "Unassigns servers/IPs to load balancer.",
					Flags:  []cli.Flag{lbIdFlag, lbIpIdFlag},
					Action: removeLoadBalancerServer,
				},
				{
					Name:  "update",
					Usage: "Updates load balancer.",
					Flags: []cli.Flag{
						lbIdFlag,
						lbDescFlag,
						hcPathFlag,
						hcParseFlag,
						hctFlag,
						hciFlag,
						lbMethodFlag,
						lbNameFlag,
						persistenceFlag,
						persistTimeFlag,
					},
					Action: updateLoadBalancer,
				},
			},
		},
	}
}

// Helper functions
////////////////////////////////////////////////////////////////////////////
func parseLoadBalancerRules(ctx *cli.Context) []oneandone.LoadBalancerRule {
	lbPorts := getIntSliceOption(ctx, "portbalancer", true)
	serverPorts := getIntSliceOption(ctx, "portserver", true)
	protocols := getStringSliceOption(ctx, "protocol", true)

	if len(lbPorts) != len(serverPorts) || len(lbPorts) != len(protocols) {
		exitOnError(fmt.Errorf("equal number of --portbalancer, --portserver and --protocol arguments must be specified"))
	}

	sources := getStringSliceOption(ctx, "source", false)

	var rules []oneandone.LoadBalancerRule
	for i := 0; i < len(protocols); i++ {
		protocols[i] = strings.ToUpper(protocols[i])

		if protocols[i] != "TCP" && protocols[i] != "UDP" {
			exitOnError(fmt.Errorf("Invalid value for --protocol flag. Valid values are TCP and UDP."))
		}

		var source string

		if len(sources) > i {
			source = sources[i]
		}

		rule := oneandone.LoadBalancerRule{
			PortBalancer: uint16(validateIntRange("portbalancer", lbPorts[i], 0, 65535)),
			PortServer:   uint16(validateIntRange("portserver", serverPorts[i], 0, 65535)),
			Protocol:     protocols[i],
			Source:       source,
		}
		rules = append(rules, rule)
	}
	return rules
}

func parseLBMethod(ctx *cli.Context) string {
	method := strings.ToUpper(getRequiredOption(ctx, "method"))
	switch method {
	case "RR":
		method = "ROUND_ROBIN"
		break
	case "LC":
		method = "LEAST_CONNECTIONS"
		break
	case "ROUND_ROBIN":
		break
	case "LEAST_CONNECTIONS":
		break
	default:
		exitOnError(fmt.Errorf("Invalid value for --method flag. Valid values are ROUND_ROBIN and LEAST_CONNECTIONS."))
	}
	return method
}

func parseHCTest(ctx *cli.Context) string {
	hcTest := strings.ToUpper(getRequiredOption(ctx, "hctest"))
	if hcTest != "NONE" && hcTest != "TCP" && hcTest != "ICMP" {
		exitOnError(fmt.Errorf("Invalid value for --hctest flag. Valid values are NONE, TCP and ICMP."))
	}
	return hcTest
}

////////////////////////////////////////////////////////////////////////////

func listLoadBalancers(ctx *cli.Context) {
	loadbalancers, err := api.ListLoadBalancers(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(loadbalancers))
	for i, lb := range loadbalancers {
		data[i] = []string{
			lb.Id,
			lb.Name,
			lb.Ip,
			lb.Method,
			lb.State,
			getDatacenter(lb.Datacenter),
		}
	}
	header := []string{"ID", "Name", "IP Address", "Method", "State", "Data Center"}
	output(ctx, loadbalancers, "", false, &header, &data)
}

func showLoadBalancer(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	loadbalancer, err := api.GetLoadBalancer(lbId)
	exitOnError(err)
	output(ctx, loadbalancer, "", true, nil, nil)
}

func createLoadBalancer(ctx *cli.Context) {
	req := oneandone.LoadBalancerRequest{
		DatacenterId:          ctx.String("datacenterid"),
		HealthCheckTest:       parseHCTest(ctx),
		HealthCheckInterval:   oneandone.Int2Pointer(getIntOptionInRange(ctx, "hctime", 5, 300)),
		Method:                parseLBMethod(ctx),
		Name:                  getRequiredOption(ctx, "name"),
		Rules:                 parseLoadBalancerRules(ctx),
		Description:           ctx.String("desc"),
		HealthCheckPath:       ctx.String("hcpath"),
		HealthCheckPathParser: ctx.String("hcregex"),
	}

	if ctx.IsSet("persistence") {
		req.Persistence = oneandone.Bool2Pointer(ctx.Bool("persistence"))
		req.PersistenceTime = oneandone.Int2Pointer(getIntOptionInRange(ctx, "persint", 30, 1200))
	}
	_, loadbalancer, err := api.CreateLoadBalancer(&req)
	exitOnError(err)
	output(ctx, loadbalancer, okWaitMessage, false, nil, nil)
}

func updateLoadBalancer(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	var hcTest, method string
	var hcTime *int

	if ctx.IsSet("hctest") {
		hcTest = parseHCTest(ctx)
	}
	if ctx.IsSet("hctime") {
		hcTime = oneandone.Int2Pointer(getIntOptionInRange(ctx, "hctime", 5, 300))
	}
	if ctx.IsSet("method") {
		method = parseLBMethod(ctx)
	}

	req := oneandone.LoadBalancerRequest{
		HealthCheckTest:       hcTest,
		HealthCheckInterval:   hcTime,
		Method:                method,
		Name:                  ctx.String("name"),
		Description:           ctx.String("desc"),
		HealthCheckPath:       ctx.String("hcpath"),
		HealthCheckPathParser: ctx.String("hcregex"),
	}

	if ctx.IsSet("persistence") {
		req.Persistence = oneandone.Bool2Pointer(ctx.Bool("persistence"))
		req.PersistenceTime = oneandone.Int2Pointer(getIntOptionInRange(ctx, "persint", 30, 1200))
	}

	loadbalancer, err := api.UpdateLoadBalancer(lbId, &req)
	exitOnError(err)
	output(ctx, loadbalancer, okWaitMessage, false, nil, nil)
}

func deleteLoadBalancer(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	loadbalancer, err := api.DeleteLoadBalancer(lbId)
	exitOnError(err)
	output(ctx, loadbalancer, okWaitMessage, false, nil, nil)
}

func listLoadBalancerServers(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	servers, err := api.ListLoadBalancerServerIps(lbId)
	exitOnError(err)
	data := make([][]string, len(servers))
	for i, server := range servers {
		data[i] = []string{server.Id, server.ServerName, server.Ip}
	}
	header := []string{"ID", "Name", "IP Address"}
	output(ctx, servers, "", false, &header, &data)
}

func assignLoadBalancerServers(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	ipIds := getStringSliceOption(ctx, "ipid", true)
	loadbalancer, err := api.AddLoadBalancerServerIps(lbId, ipIds)
	exitOnError(err)
	output(ctx, loadbalancer, okWaitMessage, false, nil, nil)
}

func showLoadBalancerServer(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	server, err := api.GetLoadBalancerServerIp(lbId, ipId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func removeLoadBalancerServer(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	ipId := getRequiredOption(ctx, "ipid")
	loadbalancer, err := api.DeleteLoadBalancerServerIp(lbId, ipId)
	exitOnError(err)
	output(ctx, loadbalancer, okWaitMessage, false, nil, nil)
}

func listLoadBalancerRules(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	rules, err := api.ListLoadBalancerRules(lbId)
	exitOnError(err)
	data := make([][]string, len(rules))
	for i, rule := range rules {
		data[i] = []string{
			rule.Id,
			strconv.Itoa(int(rule.PortBalancer)),
			strconv.Itoa(int(rule.PortServer)),
			rule.Protocol,
			rule.Source,
		}
	}
	header := []string{"ID", "Balancer Port", "Server Port", "Protocol", "Source IP"}
	output(ctx, rules, "", false, &header, &data)
}

func showLoadBalancerRule(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	ruleId := getRequiredOption(ctx, "ruleid")
	rule, err := api.GetLoadBalancerRule(lbId, ruleId)
	exitOnError(err)
	output(ctx, rule, "", true, nil, nil)
}

func addLoadBalancerRules(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	rules := parseLoadBalancerRules(ctx)
	loadbalancer, err := api.AddLoadBalancerRules(lbId, rules)
	exitOnError(err)
	output(ctx, loadbalancer, okWaitMessage, false, nil, nil)
}

func removeLoadBalancerRule(ctx *cli.Context) {
	lbId := getRequiredOption(ctx, "id")
	ruleId := getRequiredOption(ctx, "ruleid")
	loadbalancer, err := api.DeleteLoadBalancerRule(lbId, ruleId)
	exitOnError(err)
	output(ctx, loadbalancer, okWaitMessage, false, nil, nil)
}
