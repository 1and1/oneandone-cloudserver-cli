package main

import (
	"fmt"
	"strconv"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var pricingOps []cli.Command

func init() {
	pricingOps = []cli.Command{
		{
			Name:        "pricing",
			Description: "1&1 pricing operations",
			Usage:       "Pricing operations.",
			Subcommands: []cli.Command{
				{
					Name:   "info",
					Usage:  "Shows all information about the pricing.",
					Action: showPricing,
				},
				{
					Name:   "image",
					Usage:  "Shows information about image pricing.",
					Action: imagePlan,
				},
				{
					Name:   "ip",
					Usage:  "Shows information about public IP pricing.",
					Action: publicIPPlan,
				},
				{
					Name:   "fixserver",
					Usage:  "Shows information about fixed server pricing.",
					Action: fixedServerPlan,
				},
				{
					Name:   "flexserver",
					Usage:  "Shows information about flex server pricing.",
					Action: flexServerPlan,
				},
				{
					Name:   "sharedstorage",
					Usage:  "Shows information about shared storage pricing.",
					Action: sharedStoragePlan,
				},
				{
					Name:   "software",
					Usage:  "Shows information about software license pricing.",
					Action: softwareLicensePlan,
				},
			},
		},
	}
}

func getPricing() *oneandone.Pricing {
	pricing, err := api.GetPricing()
	exitOnError(err)
	return pricing
}

func displayPriceTable(ctx *cli.Context, pricing *oneandone.Pricing, data [][]string) {
	header := []string{
		"Name",
		fmt.Sprintf("Gross Price (%s)", pricing.Currency),
		fmt.Sprintf("Net Price (%s)", pricing.Currency),
		"Unit",
	}
	output(ctx, pricing, "", false, &header, &data)
}

func showPricing(ctx *cli.Context) {
	output(ctx, getPricing(), "", true, nil, nil)
}

func imagePlan(ctx *cli.Context) {
	pricing := getPricing()
	data := make([][]string, 1)
	data[0] = []string{
		pricing.Plan.Image.Name,
		strconv.FormatFloat(pricing.Plan.Image.GrossPrice, 'f', -1, 64),
		strconv.FormatFloat(pricing.Plan.Image.NetPrice, 'f', -1, 64),
		pricing.Plan.Image.Unit,
	}
	displayPriceTable(ctx, pricing, data)
}

func publicIPPlan(ctx *cli.Context) {
	pricing := getPricing()
	data := make([][]string, len(pricing.Plan.PublicIPs))
	for i, ip := range pricing.Plan.PublicIPs {
		data[i] = []string{
			ip.Name,
			strconv.FormatFloat(ip.GrossPrice, 'f', -1, 64),
			strconv.FormatFloat(ip.NetPrice, 'f', -1, 64),
			ip.Unit,
		}
	}
	displayPriceTable(ctx, pricing, data)
}

func fixedServerPlan(ctx *cli.Context) {
	pricing := getPricing()
	data := make([][]string, len(pricing.Plan.Servers.FixedServers))
	for i, s := range pricing.Plan.Servers.FixedServers {
		data[i] = []string{
			s.Name,
			strconv.FormatFloat(s.GrossPrice, 'f', -1, 64),
			strconv.FormatFloat(s.NetPrice, 'f', -1, 64),
			s.Unit,
		}
	}
	displayPriceTable(ctx, pricing, data)
}

func flexServerPlan(ctx *cli.Context) {
	pricing := getPricing()
	data := make([][]string, len(pricing.Plan.Servers.FlexServers))
	for i, s := range pricing.Plan.Servers.FlexServers {
		data[i] = []string{
			s.Name,
			strconv.FormatFloat(s.GrossPrice, 'f', -1, 64),
			strconv.FormatFloat(s.NetPrice, 'f', -1, 64),
			s.Unit,
		}
	}
	displayPriceTable(ctx, pricing, data)
}

func sharedStoragePlan(ctx *cli.Context) {
	pricing := getPricing()
	data := make([][]string, 1)
	data[0] = []string{
		pricing.Plan.SharedStorage.Name,
		strconv.FormatFloat(pricing.Plan.SharedStorage.GrossPrice, 'f', -1, 64),
		strconv.FormatFloat(pricing.Plan.SharedStorage.NetPrice, 'f', -1, 64),
		pricing.Plan.SharedStorage.Unit,
	}
	displayPriceTable(ctx, pricing, data)
}

func softwareLicensePlan(ctx *cli.Context) {
	pricing := getPricing()
	data := make([][]string, len(pricing.Plan.SoftwareLicenses))
	for i, lic := range pricing.Plan.SoftwareLicenses {
		data[i] = []string{
			lic.Name,
			strconv.FormatFloat(lic.GrossPrice, 'f', -1, 64),
			strconv.FormatFloat(lic.NetPrice, 'f', -1, 64),
			lic.Unit,
		}
	}
	displayPriceTable(ctx, pricing, data)
}
