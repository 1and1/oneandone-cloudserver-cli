package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var monitorPolicyOps []cli.Command

func init() {
	mpNotifFlags := cli.BoolFlag{
		Name:  "notify",
		Usage: "Set true for sending e-mail notifications.",
	}
	mpPortFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "alertif",
			Usage: "Alert if port is 'RESPONDING' or 'NOT_RESPONDING' (aliases 'R' and 'NR').",
		},
		mpNotifFlags,
		cli.StringFlag{
			Name:  "port",
			Usage: "Port number.",
		},
		cli.StringFlag{
			Name:  "protocol",
			Usage: "Internet protocol: TCP or UDP.",
		},
	}
	mpProcessFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "alertif",
			Usage: "Alert if process is 'RUNNING' or 'NOT_RUNNING' (aliases 'R' and 'NR').",
		},
		mpNotifFlags,
		cli.StringFlag{
			Name:  "process",
			Usage: "Name of the process.",
		},
	}

	mpPortSliceFlags := []cli.Flag{
		cli.IntSliceFlag{
			Name:  "port",
			Usage: "Port number.",
		},
		cli.StringSliceFlag{
			Name:  "protocol",
			Usage: "Internet protocol: TCP or UDP.",
		},
		cli.StringSliceFlag{
			Name:  "ptalert",
			Usage: "Alert if port is 'RESPONDING' or 'NOT_RESPONDING' (aliases 'R' and 'NR').",
		},
		cli.StringSliceFlag{
			Name:  "ptnotify",
			Usage: "Set true for sending e-mail notifications.",
		},
	}

	mpProcessSliceFlags := []cli.Flag{
		cli.StringSliceFlag{
			Name:  "pcalert",
			Usage: "Alert if process is 'RUNNING' or 'NOT_RUNNING' (aliases 'R' and 'NR').",
		},
		cli.StringSliceFlag{
			Name:  "pcnotify",
			Usage: "Set true for sending e-mail notifications.",
		},
		cli.StringSliceFlag{
			Name:  "process",
			Usage: "Name of the process.",
		},
	}

	mpIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the monitoring policy.",
	}
	mpAgentFlag := cli.BoolFlag{
		Name:  "agent",
		Usage: "Set true for using agent.",
	}
	mpEmailFlag := cli.StringFlag{
		Name:  "email",
		Usage: "User's email.",
	}
	mpDescFlag := cli.StringFlag{
		Name:  "desc, d",
		Usage: "Description of the monitoring policy.",
	}
	mpNameFlag := cli.StringFlag{
		Name:  "name, n",
		Usage: "Name of the monitoring policy.",
	}
	mpServerIdFlag := cli.StringFlag{
		Name:  "serverid",
		Usage: "ID of the server.",
	}
	mpPortIdFlag := cli.StringFlag{
		Name:  "portid",
		Usage: "ID of the port.",
	}
	mpProcessIdFlag := cli.StringFlag{
		Name:  "processid",
		Usage: "ID of the process.",
	}

	mpCPUWarnAlertFlag := cli.BoolFlag{
		Name:  "cpuwa",
		Usage: "Enable CPU warning alert.",
	}
	mpCPUWarnThresholdFlag := cli.StringFlag{
		Name:  "cpuwv",
		Usage: "Warning value (%) for CPU threshold [1 95].",
	}
	mpCPUCritAlertFlag := cli.BoolFlag{
		Name:  "cpuca",
		Usage: "Enable CPU critical alert.",
	}
	mpCPUCritThresholdFlag := cli.StringFlag{
		Name:  "cpucv",
		Usage: "Critical value (%) for CPU threshold.",
	}

	mpRAMWarnAlertFlag := cli.BoolFlag{
		Name:  "ramwa",
		Usage: "Enable RAM warning alert.",
	}
	mpRAMWarnThresholdFlag := cli.StringFlag{
		Name:  "ramwv",
		Usage: "Warning value (%) for RAM threshold [1 95].",
	}
	mpRAMCritAlertFlag := cli.BoolFlag{
		Name:  "ramca",
		Usage: "Enable RAM critical alert.",
	}
	mpRAMCritThresholdFlag := cli.StringFlag{
		Name:  "ramcv",
		Usage: "Critical value (%) for RAM threshold.",
	}

	mpDiskWarnAlertFlag := cli.BoolFlag{
		Name:  "diskwa",
		Usage: "Enable hard disk warning alert.",
	}
	mpDiskWarnThresholdFlag := cli.StringFlag{
		Name:  "diskwv",
		Usage: "Warning value (%) for hard disk threshold [1 95].",
	}
	mpDiskCritAlertFlag := cli.BoolFlag{
		Name:  "diskca",
		Usage: "Enable hard disk critical alert.",
	}
	mpDiskCritThresholdFlag := cli.StringFlag{
		Name:  "diskcv",
		Usage: "Critical value (%) for hard disk threshold.",
	}

	mpPingWarnAlertFlag := cli.BoolFlag{
		Name:  "pingwa",
		Usage: "Enable internal ping warning alert.",
	}
	mpPingWarnThresholdFlag := cli.StringFlag{
		Name:  "pingwv",
		Usage: "Warning value (ms) for internal ping threshold (min. 1).",
	}
	mpPingCritAlertFlag := cli.BoolFlag{
		Name:  "pingca",
		Usage: "Enable internal ping critical alert.",
	}
	mpPingCritThresholdFlag := cli.StringFlag{
		Name:  "pingcv",
		Usage: "Critical value (ms) for internal ping threshold (max. 100).",
	}

	mpTransferWarnAlertFlag := cli.BoolFlag{
		Name:  "transferwa",
		Usage: "Enable data transfer warning alert.",
	}
	mpTransferWarnThresholdFlag := cli.StringFlag{
		Name:  "transferwv",
		Usage: "Warning value (kbps) for data transfer threshold (min. 1).",
	}
	mpTransferCritAlertFlag := cli.BoolFlag{
		Name:  "transferca",
		Usage: "Enable data transfer critical alert.",
	}
	mpTransferCritThresholdFlag := cli.StringFlag{
		Name:  "transfercv",
		Usage: "Critical value (kbps) for data transfer threshold (max. 2000).",
	}

	mpUpdateFlags := []cli.Flag{
		mpAgentFlag,
		mpNameFlag,
		mpDescFlag,
		mpEmailFlag,
		mpCPUCritAlertFlag,
		mpCPUCritThresholdFlag,
		mpCPUWarnAlertFlag,
		mpCPUWarnThresholdFlag,
		mpDiskCritAlertFlag,
		mpDiskCritThresholdFlag,
		mpDiskWarnAlertFlag,
		mpDiskWarnThresholdFlag,
		mpPingCritAlertFlag,
		mpPingCritThresholdFlag,
		mpPingWarnAlertFlag,
		mpPingWarnThresholdFlag,
		mpRAMCritAlertFlag,
		mpRAMCritThresholdFlag,
		mpRAMWarnAlertFlag,
		mpRAMWarnThresholdFlag,
		mpTransferCritAlertFlag,
		mpTransferCritThresholdFlag,
		mpTransferWarnAlertFlag,
		mpTransferWarnThresholdFlag,
	}
	mpCreateFlags := append(mpUpdateFlags, mpPortSliceFlags...)
	mpCreateFlags = append(mpCreateFlags, mpProcessSliceFlags...)

	monitorPolicyOps = []cli.Command{
		{
			Name:        "monitorpolicy",
			Description: "1&1 monitoring policy operations",
			Usage:       "Monitoring policy operations.",
			Subcommands: []cli.Command{
				{
					Name:  "assign",
					Usage: "Assigns servers to monitoring policy.",
					Flags: []cli.Flag{
						mpIdFlag,
						cli.StringSliceFlag{
							Name:  "serverid",
							Usage: "List of server IDs.",
						},
					},
					Action: assignMonitorPolicyServers,
				},
				{
					Name:   "create",
					Usage:  "Creates new monitoring policy.",
					Flags:  mpCreateFlags,
					Action: createMonitorPolicy,
				},
				{
					Name:   "info",
					Usage:  "Shows information about monitoring policy.",
					Flags:  []cli.Flag{mpIdFlag},
					Action: showMonitorPolicy,
				},
				{
					Name:   "list",
					Usage:  "Lists available monitoring policies.",
					Flags:  queryFlags,
					Action: listMonitorPolicies,
				},
				{
					Name:   "port",
					Usage:  "Shows information about monitoring policy port.",
					Flags:  []cli.Flag{mpIdFlag, mpPortIdFlag},
					Action: showMonitorPolicyPort,
				},
				{
					Name:   "portadd",
					Usage:  "Adds new ports to monitoring policy.",
					Flags:  append([]cli.Flag{mpIdFlag}, mpPortSliceFlags...),
					Action: addMonitorPolicyPorts,
				},
				{
					Name:   "portmod",
					Usage:  "Modifies port of monitoring policy.",
					Flags:  append([]cli.Flag{mpIdFlag, mpPortIdFlag}, mpPortFlags...),
					Action: modifyMonitorPolicyPort,
				},
				{
					Name:   "portrm",
					Usage:  "Removes port from monitoring policy.",
					Flags:  []cli.Flag{mpIdFlag, mpPortIdFlag},
					Action: removeMonitorPolicyPort,
				},
				{
					Name:   "ports",
					Usage:  "Lists monitoring policy ports",
					Flags:  []cli.Flag{mpIdFlag},
					Action: listMonitorPolicyPorts,
				},
				{
					Name:   "process",
					Usage:  "Shows information about monitoring policy process.",
					Flags:  []cli.Flag{mpIdFlag, mpProcessIdFlag},
					Action: showMonitorPolicyProcess,
				},
				{
					Name:   "processadd",
					Usage:  "Adds new processes to monitoring policy.",
					Flags:  append([]cli.Flag{mpIdFlag}, mpProcessSliceFlags...),
					Action: addMonitorPolicyProcesses,
				},
				{
					Name:   "processmod",
					Usage:  "Modifies process of monitoring policy.",
					Flags:  append([]cli.Flag{mpIdFlag, mpProcessIdFlag}, mpProcessFlags...),
					Action: modifyMonitorPolicyProcess,
				},
				{
					Name:   "processrm",
					Usage:  "Removes process from monitoring policy.",
					Flags:  []cli.Flag{mpIdFlag, mpProcessIdFlag},
					Action: removeMonitorPolicyProcess,
				},
				{
					Name:   "processes",
					Usage:  "Lists monitoring policy processes.",
					Flags:  []cli.Flag{mpIdFlag},
					Action: listMonitorPolicyProcesses,
				},
				{
					Name:   "server",
					Usage:  "Shows information about server attached to monitoring policy.",
					Flags:  []cli.Flag{mpIdFlag, mpServerIdFlag},
					Action: showMonitorPolicyServer,
				},
				{
					Name:   "servers",
					Usage:  "Lists servers attached to monitoring policies.",
					Flags:  []cli.Flag{mpIdFlag},
					Action: listMonitorPolicyServers,
				},
				{
					Name:   "rm",
					Usage:  "Removes monitoring policy.",
					Flags:  []cli.Flag{mpIdFlag},
					Action: deleteMonitorPolicy,
				},
				{
					Name:   "unassign",
					Usage:  "Unassigns servers from monitoring policy.",
					Flags:  []cli.Flag{mpIdFlag, mpServerIdFlag},
					Action: removeMonitorPolicyServer,
				},
				{
					Name:   "update",
					Usage:  "Updates monitoring policy.",
					Flags:  append([]cli.Flag{mpIdFlag}, mpUpdateFlags...),
					Action: updateMonitorPolicy,
				},
			},
		},
	}
}

// Helper functions
////////////////////////////////////////////////////////////////////////////
func parseMonitorPolicyPorts(ctx *cli.Context) []oneandone.MonitoringPort {
	ports := getIntSliceOption(ctx, "port", true)
	protocols := getStringSliceOption(ctx, "protocol", true)
	alertState := getStringSliceOption(ctx, "ptalert", true)
	notifications := getStringSliceOption(ctx, "ptnotify", true)

	if len(ports) != len(alertState) || len(ports) != len(protocols) || len(ports) != len(notifications) {
		exitOnError(fmt.Errorf("equal number of --port, --protocol, --ptalert and --ptnotify arguments must be specified"))
	}

	var mpPorts []oneandone.MonitoringPort
	for i := 0; i < len(protocols); i++ {
		protocols[i] = verifyPortProtocol(protocols[i])

		alertState[i] = verifyPortAlert(alertState[i])

		notify, _ := strconv.ParseBool(notifications[i])

		port := oneandone.MonitoringPort{
			Port:              validateIntRange("port", ports[i], 1, 65535),
			Protocol:          protocols[i],
			AlertIf:           alertState[i],
			EmailNotification: notify,
		}
		mpPorts = append(mpPorts, port)
	}
	return mpPorts
}

func verifyPortProtocol(protocol string) string {
	protocol = strings.ToUpper(protocol)

	if protocol != "TCP" && protocol != "UDP" {
		exitOnError(fmt.Errorf("Invalid value for --protocol flag. Valid values are TCP and UDP"))
	}
	return protocol
}

func verifyPortAlert(alert string) string {
	alert = strings.ToUpper(alert)

	if alert == "R" {
		alert = "RESPONDING"
	} else if alert == "NR" {
		alert = "NOT_RESPONDING"
	} else if alert != "RESPONDING" && alert != "NOT_RESPONDING" {
		exitOnError(fmt.Errorf("Invalid value for port --alertif flag. Valid values are RESPONDING and NOT_RESPONDING"))
	}
	return alert
}

func parseMonitorPolicyProcs(ctx *cli.Context) []oneandone.MonitoringProcess {
	processes := getStringSliceOption(ctx, "process", true)
	alertState := getStringSliceOption(ctx, "pcalert", true)
	notifications := getStringSliceOption(ctx, "pcnotify", true)

	if len(processes) != len(alertState) || len(processes) != len(notifications) {
		exitOnError(fmt.Errorf("equal number of --process, --pcalert and --pcnotify arguments must be specified"))
	}

	var mpProcesses []oneandone.MonitoringProcess
	for i := 0; i < len(processes); i++ {
		alertState[i] = verifyProcessAlert(alertState[i])

		notify, _ := strconv.ParseBool(notifications[i])

		process := oneandone.MonitoringProcess{
			Process:           processes[i],
			AlertIf:           alertState[i],
			EmailNotification: notify,
		}
		mpProcesses = append(mpProcesses, process)
	}
	return mpProcesses
}

func verifyProcessAlert(alert string) string {
	alert = strings.ToUpper(alert)

	if alert == "R" {
		alert = "RUNNING"
	} else if alert == "NR" {
		alert = "NOT_RUNNING"
	} else if alert != "RUNNING" && alert != "NOT_RUNNING" {
		exitOnError(fmt.Errorf("Invalid value for process --alertif flag. Valid values are RUNNING and NOT_RUNNING"))
	}
	return alert
}

func getRequest(ctx *cli.Context, isCreate bool) *oneandone.MonitoringPolicy {
	mp := new(oneandone.MonitoringPolicy)
	if isCreate {
		mp.Name = getRequiredOption(ctx, "name")
		mp.Email = getRequiredOption(ctx, "email")
		mp.Ports = parseMonitorPolicyPorts(ctx)
		mp.Processes = parseMonitorPolicyProcs(ctx)
	} else {
		mp.Name = ctx.String("name")
		mp.Email = ctx.String("email")
	}
	mp.Agent = ctx.Bool("agent")
	mp.Description = ctx.String("desc")

	var cpu, ram, disk, transfer, ping *oneandone.MonitoringLevel

	if isCreate || ctx.IsSet("cpuwa") || ctx.IsSet("cpuwv") || ctx.IsSet("cpuca") || ctx.IsSet("cpucv") {
		cpu = &oneandone.MonitoringLevel{
			Warning: &oneandone.MonitoringValue{
				Alert: ctx.Bool("cpuwa"),
				Value: getIntOption(ctx, "cpuwv", isCreate),
			},
			Critical: &oneandone.MonitoringValue{
				Alert: ctx.Bool("cpuca"),
				Value: getIntOption(ctx, "cpucv", isCreate),
			},
		}
	}
	if isCreate || ctx.IsSet("ramwa") || ctx.IsSet("ramwv") || ctx.IsSet("ramca") || ctx.IsSet("ramcv") {
		ram = &oneandone.MonitoringLevel{
			Warning: &oneandone.MonitoringValue{
				Alert: ctx.Bool("ramwa"),
				Value: getIntOption(ctx, "ramwv", isCreate),
			},
			Critical: &oneandone.MonitoringValue{
				Alert: ctx.Bool("ramca"),
				Value: getIntOption(ctx, "ramcv", isCreate),
			},
		}
	}
	if isCreate || ctx.IsSet("diskwa") || ctx.IsSet("diskwv") || ctx.IsSet("diskca") || ctx.IsSet("diskcv") {
		disk = &oneandone.MonitoringLevel{
			Warning: &oneandone.MonitoringValue{
				Alert: ctx.Bool("diskwa"),
				Value: getIntOption(ctx, "diskwv", isCreate),
			},
			Critical: &oneandone.MonitoringValue{
				Alert: ctx.Bool("diskca"),
				Value: getIntOption(ctx, "diskcv", isCreate),
			},
		}
	}
	if isCreate || ctx.IsSet("transferwa") || ctx.IsSet("transferwv") || ctx.IsSet("transferca") || ctx.IsSet("transfercv") {
		transfer = &oneandone.MonitoringLevel{
			Warning: &oneandone.MonitoringValue{
				Alert: ctx.Bool("transferwa"),
				Value: getIntOption(ctx, "transferwv", isCreate),
			},
			Critical: &oneandone.MonitoringValue{
				Alert: ctx.Bool("transferca"),
				Value: getIntOption(ctx, "transfercv", isCreate),
			},
		}
	}
	if isCreate || ctx.IsSet("diskwa") || ctx.IsSet("diskwv") || ctx.IsSet("diskca") || ctx.IsSet("diskcv") {
		ping = &oneandone.MonitoringLevel{
			Warning: &oneandone.MonitoringValue{
				Alert: ctx.Bool("pingwa"),
				Value: getIntOption(ctx, "pingwv", isCreate),
			},
			Critical: &oneandone.MonitoringValue{
				Alert: ctx.Bool("pingca"),
				Value: getIntOption(ctx, "pingcv", isCreate),
			},
		}
	}
	mp.Thresholds = &oneandone.MonitoringThreshold{
		Cpu:          cpu,
		Ram:          ram,
		Disk:         disk,
		Transfer:     transfer,
		InternalPing: ping,
	}

	return mp
}

////////////////////////////////////////////////////////////////////////////

func listMonitorPolicies(ctx *cli.Context) {
	policies, err := api.ListMonitoringPolicies(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(policies))
	for i, policy := range policies {
		data[i] = []string{
			policy.Id,
			policy.Name,
			policy.Email,
			policy.CreationDate,
			strconv.FormatBool(policy.Agent),
		}
	}
	header := []string{"ID", "Name", "E-mail", "Creation Date", "Agent"}
	output(ctx, policies, "", false, &header, &data)
}

func showMonitorPolicy(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	monPolicy, err := api.GetMonitoringPolicy(mpId)
	exitOnError(err)
	output(ctx, monPolicy, "", true, nil, nil)
}

func createMonitorPolicy(ctx *cli.Context) {
	_, monPolicy, err := api.CreateMonitoringPolicy(getRequest(ctx, true))
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func updateMonitorPolicy(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	monPolicy, err := api.UpdateMonitoringPolicy(mpId, getRequest(ctx, false))
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func deleteMonitorPolicy(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	monPolicy, err := api.DeleteMonitoringPolicy(mpId)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func listMonitorPolicyServers(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	servers, err := api.ListMonitoringPolicyServers(mpId)
	exitOnError(err)
	data := make([][]string, len(servers))
	for i, server := range servers {
		data[i] = []string{server.Id, server.Name}
	}
	header := []string{"ID", "Name"}
	output(ctx, servers, "", false, &header, &data)
}

func assignMonitorPolicyServers(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	serverIds := getStringSliceOption(ctx, "serverid", true)
	monPolicy, err := api.AttachMonitoringPolicyServers(mpId, serverIds)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func showMonitorPolicyServer(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")
	server, err := api.GetMonitoringPolicyServer(mpId, serverId)
	exitOnError(err)
	output(ctx, server, "", true, nil, nil)
}

func removeMonitorPolicyServer(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	serverId := getRequiredOption(ctx, "serverid")
	monPolicy, err := api.RemoveMonitoringPolicyServer(mpId, serverId)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func listMonitorPolicyPorts(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	ports, err := api.ListMonitoringPolicyPorts(mpId)
	exitOnError(err)
	data := make([][]string, len(ports))
	for i, port := range ports {
		data[i] = []string{
			port.Id,
			strconv.Itoa(port.Port),
			port.Protocol,
			strconv.FormatBool(port.EmailNotification),
			port.AlertIf,
		}
	}
	header := []string{"ID", "Port", "Protocol", "Send E-Mail", "Alerting State"}
	output(ctx, ports, "", false, &header, &data)
}

func showMonitorPolicyPort(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	portId := getRequiredOption(ctx, "portid")
	port, err := api.GetMonitoringPolicyPort(mpId, portId)
	exitOnError(err)
	output(ctx, port, "", true, nil, nil)
}

func addMonitorPolicyPorts(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	ports := parseMonitorPolicyPorts(ctx)
	monPolicy, err := api.AddMonitoringPolicyPorts(mpId, ports)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func modifyMonitorPolicyPort(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	portId := getRequiredOption(ctx, "portid")
	protocol := verifyPortProtocol(getRequiredOption(ctx, "protocol"))
	portValue := stringFlag2Int(ctx, "port")
	alertIf := verifyPortAlert(getRequiredOption(ctx, "alertif"))

	port := &oneandone.MonitoringPort{
		Port:              portValue,
		Protocol:          protocol,
		AlertIf:           alertIf,
		EmailNotification: ctx.Bool("notify"),
	}
	monPolicy, err := api.ModifyMonitoringPolicyPort(mpId, portId, port)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func removeMonitorPolicyPort(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	portId := getRequiredOption(ctx, "portid")
	monPolicy, err := api.DeleteMonitoringPolicyPort(mpId, portId)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func listMonitorPolicyProcesses(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	processes, err := api.ListMonitoringPolicyProcesses(mpId)
	exitOnError(err)
	data := make([][]string, len(processes))
	for i, process := range processes {
		data[i] = []string{
			process.Id,
			process.Process,
			strconv.FormatBool(process.EmailNotification),
			process.AlertIf,
		}
	}
	header := []string{"ID", "Process", "Send E-Mail", "Alerting State"}
	output(ctx, processes, "", false, &header, &data)
}

func showMonitorPolicyProcess(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	processId := getRequiredOption(ctx, "processid")
	process, err := api.GetMonitoringPolicyProcess(mpId, processId)
	exitOnError(err)
	output(ctx, process, "", true, nil, nil)
}

func addMonitorPolicyProcesses(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	processes := parseMonitorPolicyProcs(ctx)
	monPolicy, err := api.AddMonitoringPolicyProcesses(mpId, processes)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func modifyMonitorPolicyProcess(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	processId := getRequiredOption(ctx, "processid")
	processName := getRequiredOption(ctx, "process")
	alertIf := verifyProcessAlert(getRequiredOption(ctx, "alertif"))

	process := &oneandone.MonitoringProcess{
		Process:           processName,
		AlertIf:           alertIf,
		EmailNotification: ctx.Bool("notify"),
	}
	monPolicy, err := api.ModifyMonitoringPolicyProcess(mpId, processId, process)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}

func removeMonitorPolicyProcess(ctx *cli.Context) {
	mpId := getRequiredOption(ctx, "id")
	processId := getRequiredOption(ctx, "processid")
	monPolicy, err := api.DeleteMonitoringPolicyProcess(mpId, processId)
	exitOnError(err)
	output(ctx, monPolicy, okWaitMessage, false, nil, nil)
}
