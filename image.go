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
							Usage: "ID of the server to create the image for",
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
							Name:  "frequency",
							Usage: "Creation policy frequency: ONCE, DAILY or WEEKLY",
						},
						cli.StringFlag{
							Name:  "num",
							Usage: "Maximum number of images, 1 - 50",
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
	serverId := getRequiredOption(ctx, "serverid")
	imgName := getRequiredOption(ctx, "name")
	imgFreq := getRequiredOption(ctx, "frequency")
	imgNo := getIntOptionInRange(ctx, "num", 1, 50)
	imgDesc := ctx.String("desc")
	req := oneandone.ImageConfig{
		ServerId:    serverId,
		Name:        imgName,
		NumImages:   imgNo,
		Frequency:   imgFreq,
		Description: imgDesc,
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
		data[i] = []string{
			image.Id,
			image.Name,
			image.OsVersion,
			strconv.Itoa(*image.Architecture),
			getDatacenter(image.Datacenter),
		}
	}
	header := []string{"ID", "Name", "OS", "Architecture", "Data Center"}
	output(ctx, images, "", false, &header, &data)
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
