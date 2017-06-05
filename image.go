package main

import (
	"strconv"

	"github.com/1and1/oneandone-cloudserver-sdk-go"
	"github.com/codegangsta/cli"
)

var imageOps []cli.Command

func init() {
	imageIdFlag := cli.StringFlag{
		Name:  "id, i",
		Usage: "ID of the image.",
	}
	imageOps = []cli.Command{
		{
			Name:        "image",
			Description: "1&1 image operations",
			Usage:       "Image operations.",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "Creates new image.",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "serverid",
							Usage: "ID of the server to create the image for.",
						},
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Name of the image.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "Description of the image.",
						},
						cli.StringFlag{
							Name:  "frequency, f",
							Usage: "Creation policy frequency: ONCE, DAILY or WEEKLY.",
						},
						cli.StringFlag{
							Name:  "num",
							Usage: "Maximum number of images, 1 - 50.",
						},
						cli.StringFlag{
							Name:  "datacenterid",
							Usage: "ID of the data center where the image will be created.",
						},
						cli.StringFlag{
							Name:  "source, s",
							Usage: "Source of the new image: 'server' (default), 'image' or 'iso'.",
						},
						cli.StringFlag{
							Name:  "url",
							Usage: "URL where the image can be downloaded from. Required if the source is 'image' or 'iso'.",
						},
						cli.StringFlag{
							Name:  "osid",
							Usage: "ID of the Operative System you want to import.",
						},
						cli.StringFlag{
							Name:  "type, t",
							Usage: "Type of the ISO you want to import: 'os' or 'app'. Required if the source is 'iso'.",
						},
					},
					Action: createImage,
				},
				{
					Name:   "info",
					Usage:  "Shows information about image.",
					Flags:  []cli.Flag{imageIdFlag},
					Action: showImage,
				},
				{
					Name:   "list",
					Usage:  "Lists all available images.",
					Flags:  queryFlags,
					Action: listImages,
				},
				{
					Name:   "os",
					Usage:  "Lists all available image OSes.",
					Flags:  queryFlags,
					Action: listImageOs,
				},
				{
					Name:   "rm",
					Usage:  "Deletes image.",
					Flags:  []cli.Flag{imageIdFlag},
					Action: deleteImage,
				},
				{
					Name:  "update",
					Usage: "Updates image.",
					Flags: []cli.Flag{
						imageIdFlag,
						cli.StringFlag{
							Name:  "name, n",
							Usage: "New name of the image.",
						},
						cli.StringFlag{
							Name:  "desc, d",
							Usage: "New description of the image.",
						},
						cli.BoolFlag{
							Name:  "nocp",
							Usage: "Remove creation policy, set frequency to ONCE",
						},
					},
					Action: updateImage,
				},
			},
		},
	}
}

func createImage(ctx *cli.Context) {
	imgName := getRequiredOption(ctx, "name")
	serverId := ctx.String("serverid")
	imgFreq := ctx.String("frequency")
	imgNo := getIntOrNil(ctx, "num", false)
	imgDesc := ctx.String("desc")
	dcId := ctx.String("datacenterid")
	source := ctx.String("source")
	url := ctx.String("url")
	osId := ctx.String("osid")
	isoType := ctx.String("type")
	req := oneandone.ImageRequest{
		ServerId:     serverId,
		DatacenterId: dcId,
		Source:       source,
		Url:          url,
		OsId:         osId,
		Type:         isoType,
		Name:         imgName,
		NumImages:    imgNo,
		Frequency:    imgFreq,
		Description:  imgDesc,
	}
	_, image, err := api.CreateImage(&req)
	exitOnError(err)
	output(ctx, image, okWaitMessage, false, nil, nil)
}

func listImages(ctx *cli.Context) {
	images, err := api.ListImages(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(images))
	for i, image := range images {
		osVersion := image.OsVersion
		if osVersion == "" {
			osVersion = image.Os
		}
		arch := ""
		if image.Architecture != nil {
			arch = strconv.Itoa(*image.Architecture)
		}
		data[i] = []string{
			image.Id,
			image.Name,
			osVersion,
			arch,
			getDatacenter(image.Datacenter),
		}
	}
	header := []string{"ID", "Name", "OS", "Architecture", "Data Center"}
	output(ctx, images, "", false, &header, &data)
}

func listImageOs(ctx *cli.Context) {
	imageOs, err := api.ListImageOs(getQueryParams(ctx))
	exitOnError(err)
	data := make([][]string, len(imageOs))
	for i, os := range imageOs {
		arch := ""
		if os.Architecture != nil {
			arch = strconv.Itoa(*os.Architecture)
		}
		data[i] = []string{
			os.Id,
			os.Os,
			os.OsFamily,
			os.OsVersion,
			arch,
		}
	}
	header := []string{"ID", "OS", "OS Family", "OS Version", "Architecture"}
	output(ctx, imageOs, "", false, &header, &data)
}

func showImage(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	image, err := api.GetImage(id)
	exitOnError(err)
	output(ctx, image, "", true, nil, nil)
}

func deleteImage(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	image, err := api.DeleteImage(id)
	exitOnError(err)
	output(ctx, image, okWaitMessage, false, nil, nil)
}

func updateImage(ctx *cli.Context) {
	id := getRequiredOption(ctx, "id")
	var freq string
	if ctx.Bool("nocp") {
		freq = "ONCE"
	}
	image, err := api.UpdateImage(id, ctx.String("name"), ctx.String("desc"), freq)
	exitOnError(err)
	output(ctx, image, okWaitMessage, false, nil, nil)
}
