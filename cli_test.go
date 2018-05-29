package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var cliOps = []string{
	"appliance",
	"datacenter",
	"dvdiso",
	"firewall",
	"image",
	"ip",
	"loadbalancer",
	"log",
	"monitor",
	"monitorpolicy",
	"ping",
	"pricing",
	"privatenet",
	"role",
	"server",
	"sharedstorage",
	"usage",
	"user",
	"vpn",
	"blockstorage",
	"sshkey",
}

const (
	appPath          = "./" + appName
	requiredOption   = "--%s option is required\n"
	requiredIntSlice = "--%s must specify at least one integer value\n"
	requiredStrSlice = "--%s must specify at least one string value\n"
	requiredIntRange = "--%s must be an integer in range [%d %d]\n"
	pingResponse     = "Response: PONG\nThe API is running.\n"
)

func runCommand(command string, args ...string) (string, error) {
	out, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", out), nil
}

func assertEqual(t *testing.T, err error, expected, got string) {
	if err != nil {
		t.Fatal(err.Error())
	}
	if expected != got {
		t.Errorf("Command output error. Expected: '%s', Got: '%s'", expected, got)
	}
}

func assertContain(t *testing.T, err error, out string, tokens []string) {
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, token := range tokens {
		if !strings.Contains(out, token) {
			t.Errorf("command output is expected to contain '%s'", token)
		}
	}
}

func TestDefault(t *testing.T) {
	out, err := runCommand(appPath)
	assertContain(t, err, out, []string{appHelpName, AppVersion, "Usage:", "Options:", "Operations:"})
}

func TestOperationDefault(t *testing.T) {
	for _, op := range cliOps {
		out, err := runCommand(appPath, op)
		assertContain(t, err, out, []string{"Usage:", "Options:", "Commands:"})
	}
}

func TestRequiredArgs(t *testing.T) {
	ops := [][]string{
		{"id", "appliance", "info"},
		{"id", "datacenter", "info"},
		{"id", "dvdiso", "info"},
		{"id", "firewall", "info"},
		{"id", "image", "info"},
		{"id", "ip", "info"},
		{"id", "loadbalancer", "info"},
		{"id", "log", "info"},
		{"id", "monitor", "info"},
		{"id", "monitorpolicy", "info"},
		{"id", "privatenet", "info"},
		{"id", "role", "info"},
		{"id", "server", "info"},
		{"id", "sharedstorage", "info"},
		{"id", "user", "info"},
		{"id", "vpn", "info"},
		{"id", "blockstorage", "info"},
		{"id", "sshkey", "info"},
		{"period", "usage", "images"},
		{"period", "usage", "loadbalancers"},
		{"period", "usage", "ips"},
		{"period", "usage", "servers"},
		{"period", "usage", "sharedstorages"},
		{"period", "monitor", "info", "--id=dummy"},
		{"id", "firewall", "rm"},
		{"id", "image", "rm"},
		{"id", "ip", "rm"},
		{"id", "loadbalancer", "rm"},
		{"id", "monitorpolicy", "rm"},
		{"id", "privatenet", "rm"},
		{"id", "role", "rm"},
		{"id", "server", "rm"},
		{"id", "sharedstorage", "rm"},
		{"id", "user", "rm"},
		{"id", "vpn", "rm"},
		{"id", "blockstorage", "rm"},
		{"id", "sshkey", "rm"},
		{"id", "role", "clone"},
		{"id", "server", "clone"},
		{"id", "server", "update"},
		{"id", "vpn", "configfile"},
		{"id", "blockstorage", "update"},
		{"name", "role", "clone", "--id=dummy"},
		{"name", "server", "clone", "--id=dummy"},
		{"name", "server", "create"},
		{"name", "vpn", "create"},
		{"osid", "server", "create", "--name=dummy"},
	}
	for _, op := range ops {
		out, err := runCommand(appPath, op[1:len(op)]...)
		assertEqual(t, err, fmt.Sprintf(requiredOption, op[0]), out)
	}
}

func TestNoArgsCommandHelp(t *testing.T) {
	var out string
	var err error
	for _, op := range cliOps {
		switch op {
		case "ping":
			out, err = runCommand(appPath, op, "api", "--help")
			assertContain(t, err, out, []string{"Usage:"})
			out, err = runCommand(appPath, op, "auth", "--help")
			assertContain(t, err, out, []string{"Usage:"})
			break
		case "pricing":
			useOps := []string{"info", "image", "ip", "fixserver", "flexserver", "sharedstorage", "software"}
			for _, uo := range useOps {
				out, err = runCommand(appPath, op, uo, "--help")
				assertContain(t, err, out, []string{"Usage:"})
			}
			break
		case "usage":
			useOps := []string{"images", "loadbalancers", "ips", "servers", "sharedstorages"}
			for _, uo := range useOps {
				out, err = runCommand(appPath, op, uo, "--help")
				assertContain(t, err, out, []string{"Usage:", "Options:"})
			}
			break
		default:
			out, err = runCommand(appPath, op, "list", "--help")
			assertContain(t, err, out, []string{"Usage:", "Options:"})
		}
	}
}

func TestRequiredIntSlice(t *testing.T) {
	ops := [][]string{
		{"portfrom", "firewall", "create", "--name=dummy"},
		{"portto", "firewall", "create", "--name=dummy", "--portfrom=80"},
		{"portbalancer", "loadbalancer", "ruleadd", "--id=dummy"},
		{"portserver", "loadbalancer", "ruleadd", "--id=dummy", "--portbalancer=80"},
		{"port", "monitorpolicy", "create", "--name=dummy", "--email=dummy@mail.no"},
		{"size", "server", "hddadd", "--id=dummy"},
	}
	for _, op := range ops {
		out, err := runCommand(appPath, op[1:len(op)]...)
		assertEqual(t, err, fmt.Sprintf(requiredIntSlice, op[0]), out)
	}
}

func TestRequiredStringSlice(t *testing.T) {
	ops := [][]string{
		{"ipid", "firewall", "assign", "--id=dummy"},
		{"serverid", "privatenet", "assign", "--id=dummy"},
		{"serverid", "sharedstorage", "attach", "--id=dummy"},
		{"perm", "sharedstorage", "attach", "--id=dummy", "--serverid=dummy"},
		{"ip", "user", "ipadd", "--id=dummy"},
	}
	for _, op := range ops {
		out, err := runCommand(appPath, op[1:len(op)]...)
		assertEqual(t, err, fmt.Sprintf(requiredStrSlice, op[0]), out)
	}
}

func TestRequiredIntRange(t *testing.T) {
	out, err := runCommand(appPath, "image", "create", "--num=60", "--serverid=1a", "--name=dummy", "--frequency=once")
	assertEqual(t, err, fmt.Sprintf(requiredIntRange, "num", 1, 50), out)

	out, err = runCommand(appPath, "sharedstorage", "create", "--name=dummy")
	assertEqual(t, err, fmt.Sprintf(requiredIntRange, "size", 50, 2000), out)

	out, err = runCommand(appPath, "sharedstorage", "update", "--id=dummy", "--size=0")
	assertEqual(t, err, fmt.Sprintf(requiredIntRange, "size", 50, 2000), out)
}

func TestPingResponse(t *testing.T) {
	out, err := runCommand(appPath, "ping", "api")
	assertEqual(t, err, pingResponse, out)
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}
