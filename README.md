# 1&amp;1 Cloud Server CLI

For more information on the 1&amp;1 Cloud Server CLI see the [1&1 Community Portal](https://www.1and1.com/cloud-community/).

# Table of Contents

- [Concepts](#concepts)
- [Getting Started](#getting-started)
- [Supported Platforms](#supported-platforms)
- [Installation](#installation)
- [Overview](#overview)
- [Configuration](#configuration)
- [How To's](#how-tos)
  - [Firewall Policy Basics](#firewall-policy-basics)
  - [Create Server](#create-server)
  - [Create Baremetal Server](#create-baremetal-server)
  - [Clone Server](#clone-server)
  - [List Servers](#list-servers)
  - [Hardware Update](#hardware-update)
  - [Restart Server](#restart-server)
  - [Create Snapshot](#create-snapshot)
  - [Delete Snapshot](#delete-snapshot)
  - [Allocate Public IP](#allocate-public-ip)
  - [Create Load Balancer](#create-load-balancer)
  - [Assign Load Balancer](#assign-load-balancer)
  - [Create Image](#create-image)
  - [Download VPN Configuration](#download-vpn-configuration)
- [Summary](#summary)
- [References](#references)
  - [Server](#server)
  - [Image](#image)
  - [Shared Storage](#shared-storage)
  - [Firewall Policy](#firewall-policy)
  - [Load Balancer](#load-balancer)
  - [Public IP](#public-ip)
  - [Private Network](#private-network)
  - [VPN](#vpn)
  - [Monitoring Center](#monitoring-center)
  - [Monitoring Policy](#monitoring-policy)
  - [Log](#log)
  - [User](#user)
  - [Role](#role)
  - [Usage](#usage)
  - [Server Appliance](#server-appliance)
  - [DVD ISO](#dvd-iso)
  - [Ping](#ping)
  - [Ping Authentication](#ping-authentication)
  - [Pricing](#pricing)
  - [Data Center](#data-center)
  - [Block Storage](#block-storage)
  - [SSH Key](#ssh-key)

## Concepts

The 1&amp;1 Cloud Server CLI wraps the [1&amp;1 Cloud Server SDK for Go](https://github.com/1and1/oneandone-cloudserver-sdk-go) allowing you to interact with [1&amp;1 Cloud Server API](https://cloudpanel-api.1and1.com/documentation/1and1/) from a command-line interface.

## Getting Started

Before you begin you will need to have signed up for a 1&amp;1 account. The credentials you create during sign-up will be used to authenticate against the API.

## Supported Platforms

The 1&amp;1 CLI binaries are available on [releases](https://github.com/1and1/oneandone-cloudserver-cli/releases) page for Linux, Windows and OS X amd64 architecture.

## Installation

### Linux

  1. Download `oneandone-linux-amd64-*.sh` shell script.
  2. Run `chmod +x oneandone-linux-amd64-*.sh`
  3. Run the installation `./oneandone-linux-amd64-*.sh`
  4. Enable bash auto-completion `source /etc/bash_completion.d/oneandone`

### Mac OS X

  1. Download `oneandone-darwin-amd64-*.sh` shell script.
  2. Run `chmod +x oneandone-darwin-amd64-*.sh`
  3. Run the installation `./oneandone-darwin-amd64-*.sh`
  4. Add the directory where you run the script from in your $PATH or copy/move `bash_autocomplete` file to your $PATH.
  5. Insert `PROG=oneandone source bash_autocomplete` to your `.bashrc` file.

### Windows

  1. Download `oneandone-windows-amd64-*.msi` setup file.
  2. Run `oneandone-windows-amd64-*.msi` installer.
  3. Accept the license terms and chose the install location.
  4. Install
  
### Custom Installation on Linux and Mac

If you are a Linux or Mac user and prefer to install the application to a different location, download a tar archive instead of the install script.

  * Linux users download `oneandone-linux-amd64-*.tar.gz`.
  * OS X users download `oneandone-darwin-amd64-*.tar.gz`.

Extract the files from the archive to a desired location in your $PATH and source auto-complete script `source bash_autocomplete`.

## Overview

Run `oneandone` or `oneandone --help` or `oneandone -h` to display available operations and global options.

```
$ oneandone
1&1 Cloud Server CLI

Version: 1.0.0

Usage: oneandone [OPTIONS] OPERATION COMMAND [arguments...]

Options:
   --about                      Show info about the application.
   --apikey                     The API token key. [$ONEANDONE_API_KEY]
   --baseurl                    The API base endpoint. Default: https://cloudpanel-api.1and1.com/v1 [$ONEANDONE_BASE_URL]
   --json                       Print output as JSON string. [$ONEANDONE_JSON_OUTPUT]
   --wrap                       Try to fit the screen display by wrapping long table cells' content. [$ONEANDONE_DISPLAY_WRAP]
   --help, -h                   Show help.
   --generate-bash-completion
   --version, -v                Print the version.

Operations:
   appliance            Server appliance operations.
   datacenter           Data center operations.
   dvdiso               DVD ISO operations.
   firewall             Firewall policy operations.
   image                Image operations.
   ip                   Public IP operations.
   loadbalancer         Load balancer operations.
   log                  Log operations.
   monitor              Monitoring center operations.
   monitorpolicy        Monitoring policy operations.
   ping                 Ping operations.
   pricing              Pricing operations.
   privatenet           Private network operations.
   role                 Role operations.
   server               Server operations.
   sharedstorage        Shared storage operations.
   usage                Usage operations.
   user                 User operations.
   vpn                  VPN operations.
   blockstorage         Block storage operations.
   help, h              Shows a list of commands or help for one command

Run 'oneandone OPERATION --help' for more information on an operation's commands.
```

## Configuration

Set `ONEANDONE_API_KEY` environment variable before using the 1&amp;1 Cloud Server CLI:

`export ONEANDONE_API_KEY=mytokenkey`

Alternatively, use `--apikey` global flag when performing any operation.

# How To's

## Firewall Policy Basics

Let's start from scratch and create a firewall policy first. For this operation, we need to provide a `--name` and at least one value for each of `--portfrom`, `--portto` and `--protocol` options.

```
oneandone firewall create --name "CLI Express Policy" --portfrom 8080 --portto 8080 --protocol TCP
OK, wait for the action to complete.
```

To see the actual response, not just the confirmation message, use `--json`  or `--json=true` option.

```
oneandone --json firewall create --name "CLI Express Policy" --portfrom 8080 --portto 8080 --protocol TCP
{
    "id": "78C1CCBAB64ECA846732AF37CA041C24",
    "name": "CLI Express Policy",
    "default": 0,
    "cloudpanel_id": "FWFF88C_29",
    "creation_date": "2016-03-23T12:43:59+00:00",
    "state": "CONFIGURING",
    "rules": [
        {
            "id": "6FD06679F5A406201F6DBCD90127B43D",
            "protocol": "TCP",
            "port_from": 8080,
            "port_to": 8080,
            "source": "0.0.0.0"
        }
    ]
}
OK, wait for the action to complete.
```

If you have not specified `--json` flag, you can alway list all available firewall policies to find the ID your policy.

```
oneandone firewall list
+----------------------------------+--------------------+--------+
|                ID                |        NAME        | STATE  |
+----------------------------------+--------------------+--------+
| C76E08E8E689330132A67E8DAF163072 | Windows            | ACTIVE |
| 34A7E423DA3253E6D38563ED06F1041F | Linux              | ACTIVE |
| 78C1CCBAB64ECA846732AF37CA041C24 | CLI Express Policy | ACTIVE |
+----------------------------------+--------------------+--------+
```

Since some actions, such as creating a server or assigning a firewall policy to a server, might take a couple of minutes you can check the status of an instance using `info` command and ID of the object. In case of the firewall policy, the full command statement is as follows.

```
oneandone firewall info --id 78C1CCBAB64ECA846732AF37CA041C24
{
    "id": "78C1CCBAB64ECA846732AF37CA041C24",
    "name": "CLI Express Policy",
    "default": 0,
    "cloudpanel_id": "FWFF88C_29",
    "creation_date": "2016-03-23T12:43:59+00:00",
    "state": "ACTIVE",
    "rules": [
        {
            "id": "6FD06679F5A406201F6DBCD90127B43D",
            "protocol": "TCP",
            "port_from": 8080,
            "port_to": 8080,
            "source": "0.0.0.0"
        }
    ]
}
```

## Create Server

Servers deployed in 1&amp;1 Cloud environment might have a fixed-size or a flex configuration. Creating a fixed-size server requires a flavor ID to be provided. The ID and the configuration details of the desired size can be found using the following command:

```
oneandone server fixedsizes
+----------------------------------+------+----------+---------------+---------------------+----------------+
|                ID                | NAME | RAM (GB) | PROCESSOR NO  | CORES PER PROCESSOR | DISK SIZE (GB) |
+----------------------------------+------+----------+---------------+---------------------+----------------+
| 65929629F35BBFBA63022008F773F3EB | M    | 1        | 1             | 1                   | 40             |
| 591A7FEF641A98B38D1C4F7C99910121 | L    | 2        | 2             | 1                   | 80             |
| E903FA4F907B5AAF17A7E987FFCDCC6B | XL   | 4        | 2             | 1                   | 120            |
| 57862AE452473D551B1673938DD3DFFE | XXL  | 8        | 4             | 1                   | 160            |
| 3D4C49EAEDD42FBC23DB58FE3DEF464F | S    | 0.5      | 1             | 1                   | 30             |
| 6A2383038420110058C77057D261A07C | 3XL  | 16       | 8             | 1                   | 240            |
| EED49B709368C3715382730A604E9F6A | 4XL  | 32       | 12            | 1                   | 360            |
| EE48ACD55FEFE57E2651862A348D1254 | 5XL  | 48       | 16            | 1                   | 500            |
+----------------------------------+------+----------+---------------+---------------------+----------------+
```

The next command illustrates how to create and power on a server of the "L" size that utilizes the firewall policy we created above.

```
oneandone server create --name "CLI Demo L Server" --fixsizeid 591A7FEF641A98B38D1C4F7C99910121 \
  --poweron=true --password MyStrongPass123 --firewallid 78C1CCBAB64ECA846732AF37CA041C24 \
  --osid 72A90ECC29F718404AC3093A3D78327C
```
The required options are `--name`, `--fixsizeid` and `--osid`. 

Creating a flex server configuration is fairly simple as well.
```
oneandone server create --name "CLI Flex Server" --cpu 2 --cores 1 --ram 4 --hdsize 80 \
  --firewallid 78C1CCBAB64ECA846732AF37CA041C24 --osid B77E19E062D5818532EFF11C747BD104
```

## Create Baremetal Server

Baremetal servers deployed in 1&amp;1 Cloud environment must have a defined baremetal model ID. The ID and the configuration details of the desired model can be found using the following command:

```
oneandone baremetal server models
+----------------------------------+-------------+----------+---------------+---------------------+----------------+
|                ID                |    NAME     | RAM (GB) | PROCESSOR NO  | CORES PER PROCESSOR | DISK SIZE (GB) |
+----------------------------------+-------------+----------+---------------+---------------------+----------------+
| B77E19E062D5818532EFF11C747BD104 | BMC_S       | 16       | 1             | 4                   | 480            |
| 7C5FA1D21B98DE39D7516333AAB7DA54 | BMC_S_HDD   | 16       | 1             | 4                   | 1000           |
| 81504C620D98BCEBAA5202D145203B4B | BMC_L       | 32       | 1             | 4                   | 800            |
| D2127B1C773877A693D718C78181D430 | BMC_L       | 32       | 1             | 4                   | 960            |
| EB231935B1CFAC3D98D6FF4FBE74F6F6 | BMC_L_HDD   | 32       | 1             | 4                   | 2000           |
| 6E1F2C70CCD3EE44ED194F4FFC47C4C9 | BMC_XL      | 64       | 1             | 4                   | 800            |
| 8CC97EC5F18722F3F0263E5FB955D9FC | BMC_XL      | 64       | 1             | 4                   | 960            |
| 758222E9C1806C144559AB3A14E58A83 | BMC_XL_HDD  | 64       | 1             | 4                   | 2000           |
+----------------------------------+-------------+----------+---------------+---------------------+----------------+
```

The next command illustrates how to create and power on a baremetal server of the "BMC_L_HDD" size .

```
oneandone server create --name "CLI Demo baremetal Server" --modelid EB231935B1CFAC3D98D6FF4FBE74F6F6 \
  --poweron=true --password MyStrongPass123 --osid 33352CCE1E710AF200CD1234BFD18862
```
The required options are `--name`, `--fixsizeid` and `--osid`. 

## Clone Server

To deploy exactly the same configuration you just need to supply the server ID and a name of the new server.

`oneandone server clone --id 39F707798F1B7FFC1F439352CF724441 --name "Flex Server Clone"`

## List Servers

List all servers using a simple command:

```
oneandone server list
+----------------------------------+-----------------------------------------+-------------+-------------+
|                ID                |                  NAME                   |    STATE    | DATA CENTER |
+----------------------------------+-----------------------------------------+-------------+-------------+
| 5ED3763CE328CB8DF1961A0550EE8CA8 | CLI Demo L Server                       | POWERED_ON  | US          |
| 27D08CBEE645A0633C959B3E034C8AD2 | Flex Server Clone                       | DEPLOYING   | DE          |
| 39F707798F1B7FFC1F439352CF724441 | CLI Flex Server                         | POWERED_OFF | GB          |
+----------------------------------+-----------------------------------------+-------------+-------------+
```

## Hardware Update

You have created a server but have not allocated enough resource to it. No problem, this command may help you to provision the server and avoid recreating it.

`oneandone server hwupdate --id 27D08CBEE645A0633C959B3E034C8AD2 --cpu 4 --cores 2 --ram 8`

## Restart Server

If you need to restart a server, provide the correct server ID and run the command:

`oneandone server reboot --id 27D08CBEE645A0633C959B3E034C8AD2`

You may use `--force` option to force hardware reboot.

## Create Snapshot

It might be a good idea to create a snapshot for some of your servers. With 1&amp;1 Cloud Server CLI that's an easy task.

```
oneandone --json server snapshotmake --id 27D08CBEE645A0633C959B3E034C8AD2
{
    "id": "27D08CBEE645A0633C959B3E034C8AD2",
    "name": "Flex Server Clone",
    "cloudpanel_id": "9689C9A",
    "creation_date": "2016-03-23T15:08:08+00:00",
    "status": {
        "state": "POWERED_ON",
        "percent": 0
    },
    "hardware": {
        "vcore": 4,
        "cores_per_processor": 2,
        "ram": 8,
        "hdds": [
            {
                "id": "94C2450C52608F8F0E6FF9BB12C7E55B",
                "size": 80,
                "is_main": true
            }
        ]
    },
    "image": {
        "id": "B77E19E062D5818532EFF11C747BD104",
        "name": "w2012r2datacenter64std"
    },
    "snapshot": {
        "id": "4C6650D152E4CC969E48BD2267CC06D4",
        "creation_date": "2016-03-23T18:49:26+00:00",
        "deletion_date": "2016-03-26T18:49:26+00:00"
    },
    "ips": [
        {
            "id": "9950CDD5A9FA30372C7EE9AA61A0CC82",
            "type": "IPV4",
            "ip": "70.35.200.149",
            "firewall_policy": {
                "id": "78C1CCBAB64ECA846732AF37CA041C24",
                "name": "CLI Express Policy"
            }
        }
    ],
    "alerts": []
}
OK, wait for the action to complete.
```

## Delete Snapshot

Once you don't need it, you might be willing to delete a server's snapshot.

```
oneandone server snapshotrm --id 27D08CBEE645A0633C959B3E034C8AD2 --snapshotid 4C6650D152E4CC969E48BD2267CC06D4
```

## Allocate Public IP

Anytime you need a new IP address for one of your server, just allocate a new one.

`oneandone server ipadd --id 27D08CBEE645A0633C959B3E034C8AD2`

## Create Load Balancer

To create a new, persistence enabled, load balancer you need to provide at least the following parameters:

```
oneandone loadbalancer create -n "Demo Load Balancer" --hctest TCP --hctime 120 --method RR \
  --persistence=true --persint 60 --portbalancer 80 --portserver 8080 --protocol TCP
```

It is always a good idea to check on a command arguments and usage by running `--help` option. For instance `oneandone loadbalancer create -h`.

## Assign Load Balancer

Assigning a server's IP address to a load balancer requires the IP and load balancer IDs to be provided. You can always obtain the ID of your balancer using the list command.

`oneandone loadbalancer list`

Also, a server's IP addresses can be listed.

```
oneandone server iplist --id 27D08CBEE645A0633C959B3E034C8AD2
+----------------------------------+---------------+-------------+
|                ID                |  IP ADDRESS   | REVERSE DNS |
+----------------------------------+---------------+-------------+
| 9950CDD5A9FA30372C7EE9AA61A0CC82 | 70.35.200.149 |             |
| 0D35557C5436CD24FE3D978CC0E7FF58 | 70.35.201.178 |             |
+----------------------------------+---------------+-------------+
```

Finally, linking the server's IP and the load balancer:

`oneandone loadbalancer assign --id 22CC4139BBF8CC72F5BE5E749BDCE72A --ipid 9950CDD5A9FA30372C7EE9AA61A0CC82`

The `--id` parameter always represents the operation instance ID. In the preceding example, the load balancer ID.

## Create Image

Here is an example how to create a server's image.

```
oneandone image create --serverid 27D08CBEE645A0633C959B3E034C8AD2 \
  --name "Demo CLI Image" --frequency ONCE --num 1
```

## Download VPN Configuration

When downloading a VPN's configuration file, only the VPN ID is a required parameter.

```
oneandone vpn configfile --id 4ADC7A1550FBF4F9A75E16D1BF483273
VPN configuration written to: "/home/nb/workspace/vpnDE88C_2.zip"
```
Run `oneandone vpn configfile --help` to learn how to change the name and location of the downloaded file.

## Summary

As we can see from the [How To's](#how-tos) examples, using 1&amp;1 Cloud Server CLI is quite simple. Help option provides more information on an operation, command or argument options, as well as the reference section below.

# References

## Server

**List all servers:**

`oneandone server list`

**Retrieve a single server:**

`oneandone server info --id [server ID]`

**List fixed-size server templates:**

`oneandone server flavors`

**Retrieve information about a fixed-size server template:**

`oneandone server flavor --id [flavor ID]`

**List Baremetal server models:**

`oneandone baremetalmodels`

**Retrieve information about a baremetal server model:**

`oneandone baremetalmodel --id [model ID]`

**Retrieve information about a server's hardware:**

`oneandone server hwinfo --id [server ID]`

**List a server's HDDs:**

`oneandone server hddlist --id [server ID]`

**Retrieve a single server HDD:**

`oneandone server hddinfo --id [server ID] --hddid [hard disk ID]`

**Retrieve information about a server's image:**

`oneandone server imginfo --id [server ID]`

**List a server's IPs:**

`oneandone server iplist --id [server ID]`

**Retrieve information about a single server IP:**

`oneandone server ipinfo --id [server ID] --ipid [IP ID]`

**Retrieve information about a server's firewall policy:**

`oneandone server fwinfo --id [server ID] --ipid [IP ID]`

**List all load balancers assigned to a server IP:**

`oneandone server lblist --id [server ID] --ipid [IP ID]`

**Retrieve information about a server's status:**

`oneandone server status --id [server ID]`

**Retrieve information about the DVD loaded into the virtual DVD unit of a server:**

`oneandone server dvdinfo --id [server ID]`

**List a server's private networks:**

`oneandone server pnlist --id [server ID]`

**Retrieve information about a server's private network:**

`oneandone server pninfo --id [server ID] --pnetid [network ID]`

**Retrieve information about a server's snapshot:**

`oneandone server snapshotinfo --id [server ID]`

**Create a server:**

```
oneandone server create \
   --cpu             [Number of processors] \
   --cores           [Number of cores per processor] \
   --datacenterid    [Data center ID] \
   --fixsizeid       [Fixed-instance size ID desired for the server] \
   --hdsize          [Size of the hard disk in GB] \
   --ram             [Size of RAM memory in GB] \
   --name            [Name of the server] \
   --desc            [Description of the server] \
   --password        [Password of the server] \
   --sshkeypath      [Path to SSH public key file] \
   --poweron         [Power on the server after creating] \
   --osid            [Server appliance ID] \
   --ipid            [ID of the IP] \
   --regionid        [Datacenter region ID] \
   --firewallid      [ID of the firewall policy] \
   --loadbalancerid  [ID of the load balancer] \
   --monitorpolicyid [Monitoring policy ID to use with the server]
```

**Create a baremetal server:**

```
oneandone createbaremetalServer create \
   --datacenterid    [Data center ID] \
   --modelid       [Fixed-instance size ID desired for the server] \
   --name            [Name of the server] \
   --desc            [Description of the server] \
   --password        [Password of the server] \
   --sshkeypath      [Path to SSH public key file] \
   --poweron         [Power on the server after creating] \
   --osid            [Server appliance ID] \
   --ipid            [ID of the IP] \
   --regionid        [Datacenter region ID] \
   --firewallid      [ID of the firewall policy] \
   --loadbalancerid  [ID of the load balancer] \
   --monitorpolicyid [Monitoring policy ID to use with the server]
```


**Update a server:**

`oneandone server update --id [server ID] --name [new name] --desc [new description]`

**Delete a server:**

`oneandone server rm --id [server ID] --keepips=[true|false]`

Set `--keepips` option to `true` for keeping server IPs after deleting a server.

**Update a server's hardware:**

```
oneandone server hwupdate \
   --cpu             [Number of processors] \
   --cores           [Number of cores per processor] \
   --fixsizeid        [Flavor size ID desired for the server] \
   --ram             [Size of RAM memory in GB] \
```

**Add new hard disk(s) to a server:**

`oneandone server hddadd --id [server ID] {--size [HDD size in GB] --size [HDD size in GB]}`

**Resize a server's hard disk:**

`oneandone server hddupdate --id [server ID] --hddid [hard disk ID] --newsize [new size in GB]`

**Remove a server's hard disk:**

`oneandone server hddrm --id [server ID] --hddid [hard disk ID]`

**Load a DVD into the virtual DVD unit of a server:**

`oneandone server dvdload --id [server ID] --dvdid [DVD ISO ID]`

**Unload a DVD from the virtual DVD unit of a server:**

`oneandone server dvdrm --id [server ID]`

**Reinstall a new image into a server:**

```
oneandone server imgupdate --id [server ID] --imgid [image ID] \
  --password [new server's password] --firewallid [firewall policy ID]
```

**Assign a new IP to a server:**

`oneandone server ipadd --id [server ID] --type [IPV4 or IPV6]`

**Release an IP from a server and optionally remove it:**

`oneandone server iprm --id [server ID] --ipid [IP ID] --keepip=[true|false]`

**Assign a new firewall policy to a server's IP:**

`oneandone server fwadd --id [server ID] --ipid [IP ID] --firewallid [firewall policy ID]`

**Assign a new load balancer to a server's IP:**

`oneandone server lbadd --id [server ID] --ipid [IP ID] --loadbalancerid [load balancer ID]`

**Remove a load balancer from a server's IP:**

`oneandone server lbrm --id [server ID] --ipid [IP ID] --loadbalancerid [load balancer ID]`

**Start a server:**

`oneandone server start --id [server ID]`

**Reboot a server:**

`oneandone server reboot --id [server ID]  --force=[true|false]`

Set `--force` to true to force HARDWARE method of rebooting.

**Shutdown a server:**

`oneandone server stop --id [server ID]  --force=[true|false]`

Set `--force` to true to force HARDWARE method of powering off.

**Assign a private network to a server:**

`oneandone server pnadd --id [server ID] --pnetid [private network ID]`

**Remove a server's private network:**

`oneandone server pnrm --id [server ID] --pnetid [private network ID]`

**Create a new server's snapshot:**

`oneandone server snapshotmake --id [server ID]`

**Restore a server's snapshot:**

`oneandone server snapshotrest --id [server ID] --snapshotid [snapshot ID]`

**Remove a server's snapshot:**

`oneandone server snapshotrm --id [server ID] --snapshotid [snapshot ID]`

**Clone a server:**

`server, err := api.CloneServer(server_id, new_name)`

## Image

**List all images:**

`oneandone image list`

**List available image OSes:**

`oneandone image os`

**Retrieve a single image:**

`oneandone image info --id [image ID]`

**Create an image:**

```
oneandone image create --serverid [server ID] --name [image name] --desc [image description] \
  --frequency  [ONCE|DAILY|WEEKLY] --num [number of images, 1 - 50] --datacenterid [data center ID]
```
`--datacenterid` and `--desc` are optional.

**Import a private image:**

```
oneandone image create --name [image name] -s [image|iso] -t [os|app] --osid [image OS ID] \
  --url [url to import the image from]
```
`--osid` is required if the image source is `image`, or the source is `iso` and the type is `os`.

**Update an image:**

```
oneandone image update --id [image ID] --name [new name] --desc [new description] \
  --nocp=[remove creation policy, true|false]
```

**Delete an image:**

`oneandone image rm --id [image ID]`

## Shared Storage

`oneandone sharedstorage list`

**Retrieve a shared storage:**

`oneandone sharedstorage info --id [shared storage ID]`

**Create a shared storage:**

```
oneandone sharedstorage create --datacenterid [data center ID] --name [shared storage name] \
  --desc [shared storage description] --size [shared storage size]
```

**Update a shared storage:**

```
oneandone sharedstorage update --id [shared storage ID] --name [new name] \
  --desc [new description] --size [new size]
```

**Remove a shared storage:**

`oneandone sharedstorage rm --id [shared storage ID]`

**List shared storage servers:**

`oneandone sharedstorage serverlist --id [shared storage ID]`

**Retrieve a shared storage server:**

`oneandone sharedstorage serverinfo --id [shared storage ID] --serverid [server ID]`

**Add servers to a shared storage:**

```
oneandone sharedstorage attach --id [shared storage ID] {--serverid [server ID] --serverid [server ID]} \
  {--perm [permission, R|RW] --perm [permission, R|RW]}
```

**Remove a server from a shared storage:**

`oneandone sharedstorage detach --id [shared storage ID] --serverid [server ID]`

**Retrieve the credentials for accessing the shared storages:**

`oneandone sharedstorage access`

**Change the password for accessing the shared storages:**

`oneandone sharedstorage access --newpass [new password]`

## Firewall Policy

**List firewall policies:**

`oneandone firewall list`

**Create a firewall policy:**

```
oneandone firewall create --name [firewall name] --desc [firewall description] \
  {--portfrom [first port] --portto [last port] --protocol [TCP|UDP|TCP/UDP|ICMP|IPSEC|GRE] --source [source IP]}
```

**Update a firewall policy:**

`oneandone firewall update --id [firewall policy ID] --name [new name] --desc [new description]`

**Delete a firewall policy:**

`oneandone firewall rm --id [firewall policy ID]`

**List servers/IPs attached to a firewall policy:**

`oneandone firewall servers --id [firewall policy ID]`

**Retrieve information about a server/IP assigned to a firewall policy:**

`oneandone firewall server --id [firewall policy ID] --ipid [IP ID]`

**Add servers/IPs to a firewall policy:**

`oneandone firewall assign --id [firewall policy ID] {--ipid [IP ID] --ipid [IP ID]}`

**List rules of a firewall policy:**

`oneandone firewall rules --id [firewall policy ID]`

**Retrieve information about a rule of a firewall policy:**

`oneandone firewall rule --id [firewall policy ID] --ruleid [rule ID]`

**Adds new rules to a firewall policy:**

```
oneandone firewall ruleadd --id [firewall policy ID] \
  {--portfrom [first port] --portto [last port] --protocol [TCP|UDP|TCP/UDP|ICMP|IPSEC|GRE] --source [source IP]}
```

**Remove a rule from a firewall policy:**

`oneandone firewall rulerm --id [firewall policy ID] --ruleid [rule ID]`

## Load Balancer

**List load balancers:**

`oneandone loadbalancer list`

**Create a load balancer:**

```
oneandone loadbalancer create --name [load balancer name] --desc [load balancer description] --method [ROUND_ROBIN|LEAST_CONNECTIONS (RR|LC)]\
  --hctest [health check test, NONE|TCP|ICMP|HTTP] --hctime [health check time (s), 5 - 300] --hcpath [health check URL] \
  --hcregex [health check regex] --persistence=[true|false] --persint [persistence time (s), 30 - 1200] \
  {--portbalancer [load balancer port] --portserver [server port] --protocol [TCP|UDP] --source [source IP]} \
  --datacenterid [data center ID]
```

**Update a load balancer:**

```
oneandone loadbalancer update --id [load balancer ID] --name [load balancer name] --desc [load balancer description] --method [ROUND_ROBIN|LEAST_CONNECTIONS (RR|LC)]\
  --hctest [health check test, NONE|TCP|ICMP|HTTP] --hctime [health check time (s), 5 - 300] --hcpath [health check URL] \
  --hcregex [health check regex] --persistence=[true|false] --persint [persistence time (s), 30 - 1200]
```

**Delete a load balancer:**

`oneandone loadbalancer rm --id [load balancer ID]`

**List servers/IPs attached to a load balancer:**

`oneandone loadbalancer servers --id [load balancer ID]`

**Retrieve information about a server/IP assigned to a load balancer:**

`oneandone loadbalancer server --id [load balancer ID] --ipid [IP ID]`

**Add servers/IPs to a load balancer:**

`oneandone loadbalancer assign --id [load balancer ID] {--ipid [IP ID] --ipid [IP ID]}`

**Remove a server/IP from a load balancer:**

`oneandone loadbalancer unassign --id [load balancer ID] --ipid [IP ID]`

**List rules of a load balancer:**

`oneandone loadbalancer rules --id [load balancer ID]`

**Retrieve information about a rule of a load balancer:**

`oneandone loadbalancer rule --id [load balancer ID] --ruleid [rule ID]`

**Adds new rules to a load balancer:**

```
oneandone loadbalancer ruleadd --id [load balancer ID] --ruleid [rule ID] \
  {--portbalancer [load balancer port] --portserver [server port] --protocol [TCP|UDP] --source [source IP]}
```

**Remove a rule from a load balancer:**

`oneandone loadbalancer rulerm --id [load balancer ID] --ruleid [rule ID]`

## Public IP

**Retrieve a list of your public IPs:**

`oneandone ip list`

**Retrieve a single public IP:**

`oneandone ip info --id [IP ID]`

**Create a public IP:**

`oneandone ip create --type [IPV4|IPV6] --dns [reverse DNS] --datacenterid [data center ID]`

**Update the reverse DNS of a public IP:**

`oneandone ip update --id [IP ID] --dns [new reverse DNS|""]`

**Remove a public IP:**

`oneandone ip rm --id [IP ID]`

## Private Network

**List all private networks:**

`oneandone privatenet list`

**Create a new private network:**

```
oneandone privatenet create --name [private net name] --desc [private net description] \
  --netip [network IP] --netmask [subnet mask] --datacenterid [data center ID]
```

**Modify a private network:**

```
oneandone privatenet update --id [private net ID] --name [private net name] 
  --desc [private net description] --netip [network IP] --netmask [subnet mask]
```

**Delete a private network:**

`oneandone privatenet rm --id [private net ID]`

**List all servers attached to a private network:**

`oneandone privatenet servers --id [private net ID]`

**Retrieve a server attached to a private network:**

`oneandone privatenet server --id [private net ID] --serverid [server ID]`

**Attach servers to a private network:**

`oneandone privatenet assign --id [private net ID] {--serverid [server ID] --serverid [server ID]}`

**Remove a server from a private network:**

`oneandone privatenet unassign --id [private net ID] --serverid [server ID]`

## VPN

**List all VPNs:**

`oneandone vpn list`

**Retrieve information about a VPN:**

`oneandone vpn info --id [VPN ID]`

**Create a VPN:**

`oneandone vpn create --name [VPN name] --desc [VPN description] --datacenterid [data center ID]`

**Modify a VPN:**

`oneandone vpn modify --id [VPN ID] --name [VPN name] --desc [VPN description]`

**Delete a VPN:**

`oneandone vpn rm --id [VPN ID]`

**Retrieve a VPN's configuration file:**

`oneandone vpn configfile --id [VPN ID] --dir [store location if not current dir] --name [file name if not default]`

## Monitoring Center

**List all usages and alerts of monitoring servers:**

```
oneandone monitor list --cpu=[true|false] --disk=[true|false] --ram=[true|false] \
  --ping=[true|false] --transfer=[true|false]
```

**Retrieve the usages and alerts for a monitoring server:**

```
oneandone monitor info --id [server ID] --period [LAST_HOUR|LAST_24H|LAST_7D|LAST_30D|LAST_365D|CUSTOM] \
  --startdate [custom start date] --enddate [custom end date]
```

## Monitoring Policy

**List all monitoring policies:**

`oneandone monitorpolicy list`

**Retrieve a single monitoring policy:**

`oneandone monitorpolicy info --id [monitor policy ID]`

**Create a monitoring policy:**

```
oneandone monitorpolicy create --agent=[true|false] --name [monitor policy name] \
  --desc [monitor policy description] --email [user's e-mail] \
  --cpuwa=[true|false] --cpuwv [1-95] --cpuca=[true|false] --cpucv [max. 100] \
  --diskwa=[true|false] -diskwv [1-95] --diskca=[true|false] --diskcv [max. 100] \
  --pingwa=[true|false] --pingwv [min. 1] --pingca=[true|false] --pingcv [max. 100] \
  --ramwa=[true|false] --ramwv [1-95] --ramca=[true|false] --ramcv [max. 100] \
  --transferwa=[true|false] --transferwv [min. 1] --transferca=[true|false] --transfercv [max. 2000] \
  {--port [1-65535] --port [1-65535]} {--protocol [TCP|UDP] --protocol [TCP|UDP]} \
  {--ptalert [RESPONDING|NOT_RESPONDING (R|NR)] --ptalert [RESPONDING|NOT_RESPONDING (R|NR)]} \
  {--ptnotify=[true|false] --ptnotify=[true|false]} \
  {--process [process name] --process [process name]} \
  {--pcalert [RUNNING|NOT_RUNNING (R|NR)] --pcalert [RUNNING|NOT_RUNNING (R|NR)]} \
  {--pcnotify=[true|false] --pcnotify=[true|false]}
```
Run `oneandone monitorpolicy create --help` for more details on available options.

**Update a monitoring policy:**

```
oneandone monitorpolicy update --id [monitor policy ID] --agent=[true|false] \
  --name [monitor policy name] --desc [monitor policy description] --email [user's e-mail] \
  --cpuwa=[true|false] --cpuwv [1-95] --cpuca=[true|false] --cpucv [max. 100] \
  --diskwa=[true|false] -diskwv [1-95] --diskca=[true|false] --diskcv [max. 100] \
  --pingwa=[true|false] --pingwv [min. 1] --pingca=[true|false] --pingcv [max. 100] \
  --ramwa=[true|false] --ramwv [1-95] --ramca=[true|false] --ramcv [max. 100] \
  --transferwa=[true|false] --transferwv [min. 1] --transferca=[true|false] --transfercv [max. 2000] \
```
Run `oneandone monitorpolicy update --help` for more details on available options.

**Delete a monitoring policy:**

`oneandone monitorpolicy rm --id [monitor policy ID]`

**List all ports of a monitoring policy:**

`oneandone monitorpolicy ports --id [monitor policy ID]`

**Retrieve information about a port of a monitoring policy:**

`oneandone monitorpolicy port --id [monitor policy ID] --portid [port ID]`

**Add new ports to a monitoring policy:**

```
oneandone monitorpolicy portadd --id [monitor policy ID] \
  {--port [1-65535] --port [1-65535]} {--protocol [TCP|UDP] --protocol [TCP|UDP]}
  {--ptalert [RESPONDING|NOT_RESPONDING (R|NR)] --ptalert [RESPONDING|NOT_RESPONDING (R|NR)]} \
  {--ptnotify=[true|false] --ptnotify=[true|false]}
```

**Modify a port of a monitoring policy:**

```
oneandone monitorpolicy portmod --id [monitor policy ID] --portid [port ID] \
  --alertif [RESPONDING|NOT_RESPONDING (R|NR)] --notify=[true|false]
```

**Remove a port from a monitoring policy:**

`oneandone monitorpolicy portrm --id [monitor policy ID] --portid [port ID]`

**List the processes of a monitoring policy:**

`oneandone monitorpolicy processes --id [monitor policy ID]`

**Retrieve information about a process of a monitoring policy:**

`oneandone monitorpolicy process --id [monitor policy ID] --processid [process ID]`

**Add new processes to a monitoring policy:**

```
oneandone monitorpolicy process --id [monitor policy ID] \
  {--process [process name] --process [process name]} \
  {--pcalert [RUNNING|NOT_RUNNING (R|NR)] --pcalert [RUNNING|NOT_RUNNING (R|NR)]} \
  {--pcnotify=[true|false] --pcnotify=[true|false]}
```

**Modify a process of a monitoring policy:**

```
oneandone monitorpolicy process --id [monitor policy ID] --processid [process ID] \
  --alertif [RUNNING|NOT_RUNNING (R|NR)] --notify=[true|false]
```

**Remove a process from a monitoring policy:**

```oneandone monitorpolicy processrm --id [monitor policy ID] --processid [process ID]```

**List all servers attached to a monitoring policy:**

`oneandone monitorpolicy servers --id [monitor policy ID]`

**Retrieve information about a server attached to a monitoring policy:**

`oneandone monitorpolicy server --id [monitor policy ID] --serverid [server ID]`

**Attach servers to a monitoring policy:**

```oneandone monitorpolicy assign --id [monitor policy ID] {--serverid [server ID] --serverid [server ID]}```

`server_ids` is a slice of server ID's.

**Remove a server from a monitoring policy:**

`oneandone monitorpolicy unassign --id [monitor policy ID] --serverid [server ID]`

## Log

**List all logs:**

```
oneandone log list --period [LAST_HOUR|LAST_24H|LAST_7D|LAST_30D|LAST_365D|CUSTOM] \
  --startdate [custom start date] --enddate [custom end date]
```

**Retrieve a single log:**

`oneandone log info --id [log ID]`

## User

**List all users:**

`oneandone user list`

**Retrieve information about a user:**

`oneandone user info --id [user ID]`

**Create a user:**

```
oneandone user create --name [username] --desc [user description] \
  --email [user's e-mail] --password [user's password.]
```

**Modify a user:**

```
oneandone user modify --id [user ID] --desc [new description] --email [new e-mail] \
  --password [new password.] --status [ACTIVE|DISABLED]
```

**Delete a user:**

`oneandone user rm --id [user ID]`

**Retrieve information about a user's API privileges:**

`oneandone user api --id [user ID]`

**Retrieve a user's API key:**

`oneandone user apitoken --id [user ID]`

**List IP's from which API access is allowed for a user:**

`oneandone user ips --id [user ID]`

**Add new IP's to a user:**

`oneandone user ipadd --id [user ID] {--ip [IP address] --ip [IP address]}`

**Remove an IP and forbid API access from it:**

`oneandone user iprm --id [user ID] --ip [IP address]`

**Enable or disable a user's API access:**

`oneandone user enableapi|disableapi --id [user ID]`

**Renew a user's API key:**

`oneandone user newtoken --id [user ID]`

**Retrieve current user permissions:**

`oneandone user permissions`

## Role

**List all roles:**

`oneandone role list`

**Retrieve information about a role:**

`oneandone role info --id [role ID]`

**Create a role:**

`oneandone role create --name [role name]`

**Clone a role:**

`oneandone role clone --id [role ID] --name [new role name]`

**Modify a role:**

`oneandone role modify --id [role ID] --name [new name] --desc [new description] --state [new state]`

`ACTIVE` and `DISABLE` are valid values for `--state` flag.

**Delete a role:**

`oneandone role rm --id [role ID]`

**Show all role's permissions:**

`oneandone role permissions info --id [role ID]`

**Show role permissions for backups:**

`oneandone role permissions backinfo --id [role ID]`

**Show role permissions for firewall policies:**

`oneandone role permissions fwinfo --id [role ID]`

**Show role permissions for images:**

`oneandone role permissions imginfo --id [role ID]` 

**Show role permissions for invoice:**

`oneandone role permissions invinfo --id [role ID]` 

**Show role permissions for IPs:**

`oneandone role permissions ipinfo --id [role ID]`  

**Show role permissions for load balancers:**

`oneandone role permissions lbinfo --id [role ID]`  

**Show role permissions for logs:**

`oneandone role permissions loginfo --id [role ID]` 

**Show role permissions for monitoring center:**

`oneandone role permissions mcinfo --id [role ID]`  

**Show role permissions for monitoring policies:**

`oneandone role permissions mpinfo --id [role ID]`  

**Show role permissions for private networks:**

`oneandone role permissions pninfo --id [role ID]`  

**Show role permissions for roles:**

`oneandone role permissions roleinfo --id [role ID]`

**Show role permissions for servers:**

`oneandone role permissions serinfo --id [role ID]` 

**Show role permissions for shared storages:**

`oneandone role permissions ssinfo --id [role ID]`  

**Show role permissions for usages:**

`oneandone role permissions usginfo --id [role ID]` 

**Show role permissions for users:**

`oneandone role permissions userinfo --id [role ID]`

**Show role permissions for VPNs:**

`oneandone role permissions vpninfo --id [role ID]` 

**Enable all role's permissions:**

`oneandone role permissions setall --id [role ID]`

**Disable all role's permissions:**

`oneandone role permissions unsetall --id [role ID]`

**Modify role permissions for backups:**

```
oneandone role permissions backmod --id [role ID] --all=[true|false] \
--create=[true|false] --delete=[true|false] --show=[true|false]  
```

**Modify role permissions for firewall policies:**

```
oneandone role permissions fwmod --id [role ID] --all=[true|false] \
--clone=[true|false] --create=[true|false] --delete=[true|false] \  
--manageip=[true|false] --managerule=[true|false] --setdesc=[true|false] \ 
--setname=[true|false] --show=[true|false]      
```

**Modify role permissions for images:**

```
oneandone role permissions imgmod --id [role ID] --all=[true|false] \
--create=[true|false] --delete=[true|false] --noautocreate=[true|false] \
--setdesc=[true|false] --setname=[true|false] --show=[true|false]  
```

**Modify role permissions for invoice:**

```
oneandone role permissions invmod --id [role ID] --all=[true|false] --show=[true|false]
```

**Modify role permissions for IPs:**

```
oneandone role permissions ipmod --id [role ID] --all=[true|false] \
--create=[true|false] --delete=[true|false] --release=[true|false] \
--setdns=[true|false] --show=[true|false]
```

**Modify role permissions for load balancers:**

```
oneandone role permissions lbmod --id [role ID] --all=[true|false] \
--create=[true|false] --delete=[true|false] --manageip=[true|false] \
--managerule=[true|false] --modify=[true|false] --setdesc=[true|false] \
--setname=[true|false] --show=[true|false]
```

**Modify role permissions for logs:**

```
oneandone role permissions logmod --id [role ID] --all=[true|false] --show=[true|false]
```

**Modify role permissions for monitoring center:**

```
oneandone role permissions mcmod --id [role ID] --all=[true|false] --show=[true|false]
```

**Modify role permissions for monitoring policies:**

```
oneandone role permissions mpmod --id [role ID] --all=[true|false] \
--clone=[true|false] --create=[true|false] --delete=[true|false] \
--manageserver=[true|false] --manageport=[true|false] --manageprocess=[true|false] \
--resources=[true|false] --setdesc=[true|false] --setemail=[true|false] \
--setname=[true|false] --show=[true|false]
```

**Modify role permissions for private networks:**

```
oneandone role permissions pnmod --id [role ID] --all=[true|false] \
--create=[true|false] --delete=[true|false] --manageserver=[true|false] \
--setdesc=[true|false] --setname=[true|false] --setnetinfo=[true|false] --show=[true|false]
```

**Modify role permissions for roles:**

```
oneandone role permissions rolemod --id [role ID] --all=[true|false] \
--clone=[true|false] --create=[true|false] --delete=[true|false] \
--manageuser=[true|false] --modify=[true|false] --setdesc=[true|false] \
--setname=[true|false] --show=[true|false]
```

**Modify role permissions for servers:**

```
oneandone role permissions sermod --id [role ID] --all=[true|false] \
--assignip=[true|false] --clone=[true|false] --create=[true|false] \
--delete=[true|false] --kvm=[true|false] --managedvd=[true|false] \
--managesnapshot=[true|false] --reinstall=[true|false] --resize=[true|false] \
--restart=[true|false] --setdesc=[true|false] --setname=[true|false] \
--show=[true|false] --shutdown=[true|false] --start=[true|false]
```

**Modify role permissions for shared storages:**

```
oneandone role permissions ssmod --id [role ID] --all=[true|false] \
--access=[true|false] --create=[true|false] --delete=[true|false] \
--manageserver=[true|false] --resize=[true|false] --setdesc=[true|false] \
--setname=[true|false] --show=[true|false]
```

**Modify role permissions for usages:**

```
oneandone role permissions usgmod --id [role ID] --all=[true|false] --show=[true|false]
```

**Modify role permissions for users:**

```
oneandone role permissions usermod --id [role ID] --all=[true|false] \
--changerole=[true|false] --create=[true|false] --delete=[true|false] \
--disable=[true|false] --enable=[true|false] --manageapi=[true|false] \
--setdesc=[true|false] --setemail=[true|false] --setpassword=[true|false] \
--show=[true|false]
```

**Modify role permissions for VPNs:**

```
oneandone role permissions vpnmod --id [role ID] --all=[true|false] \
--create=[true|false] --delete=[true|false] --downloadfile=[true|false] \
--setdesc=[true|false] --setname=[true|false] --show=[true|false]
```

**Assign users to a role:**

`oneandone role useradd --id [role ID] {--userid [user ID] --userid [user ID]}`

**List a role's users:**

`oneandone role userlist --id [role ID]`

**Retrieve information about a role's user:**

`oneandone role userinfo --id [role ID] --userid [user ID]`

**Remove a role's user:**

`oneandone role userrm --id [role ID] --userid [user ID]`

## Usage

**List your usages images, load balancers, IPs, servers or shared storages:**

```
oneandone usage images|loadbalancers|ips|servers|sharedstorages \
  --period [LAST_HOUR|LAST_24H|LAST_7D|LAST_30D|LAST_365D|CUSTOM] \
  --startdate [custom start date] --enddate [custom end date]
```
Only one command at the time is allowed, `images`, `loadbalancers`, `ips`, `servers` or `sharedstorages`.

## Server Appliance

**List all the appliances that you can use to create a server:**

`oneandone appliance list`

**Retrieve information about specific appliance:**

`oneandone appliance info --id [server appliance ID]`

## DVD ISO

**List all operative systems and tools that you can load into your virtual DVD unit:**

`oneandone dvdiso list`

**Retrieve a specific ISO image:**

`oneandone dvdiso info --id [DVD/ISO ID]`

## Ping

**Check if 1&amp;1 REST API is running:**

`oneandone ping api`

## Ping Authentication

**Validate if 1&amp;1 REST API is running and the authorization token is valid:**

`oneandone ping auth`

## Pricing

**Show all information about the pricing:**

`oneandone pricing info`

**Show information about image pricing:**

`oneandone pricing image`

**Show information about public IP pricing:**

`oneandone pricing ip`

**Show information about fixed server pricing:**

`oneandone pricing fixserver`

**Show information about flex server pricing:**

`oneandone pricing flexserver`

**Show information about shared storage pricing:**

`oneandone pricing sharedstorage`

**Show information about software license pricing:**

`oneandone pricing software`

## Data Center

**List all 1&amp;1 Cloud Server data centers:**

`oneandone datacenter list`

**Retrieve a specific data center:**

`oneandone datacenter info --id [data center ID]`

## Block Storage

**List block storages:**

`oneandone blockstorage list`

**Retrieve a block storage:**

`oneandone blockstorage info --id [block storage ID]`

**Create a block storage:**

```
oneandone blockstorage create --name [block storage name] --size [block storage size] \
  --desc [block storage description] --datacenterid [data center ID] \
  --serverid [server ID] 
```

**Update a block storage:**

`oneandone blockstorage update --id [block storage ID] --name [updated bs name] --desc [updated bs description]`

**Remove a block storage:**

`oneandone blockstorage rm --id [block storage ID]`

**Attach a block storage to a server:**

`oneandone blockstorage attach --id [block storage ID] --serverid [server ID]`

**Retrieve a block storage server:**

`oneandone blockstorage serverinfo --id [block storage ID]`

**Detach a block storage from a server:**

`oneandone blockstorage detach --id [block storage ID] --serverid [server ID]`

## SSH Key

**List all SSH Keys:**

`oneandone sshkey list`

**Retrieve information about an SSH Key:**

`oneandone sshkey info --id [SSH Key ID]`

**Create an SSH Key:**

`oneandone sshkey create --name [SSH Key name] --desc [SSH Key description] --publickey [SSH Key public key]`

**Modify an SSH Key:**

`oneandone sshkey modify --id [SSH Key ID] --name [SSH Key name] --desc [SSH Key description]`

**Delete an SSH Key:**

`oneandone sshkey rm --id [SSH Key ID]`
