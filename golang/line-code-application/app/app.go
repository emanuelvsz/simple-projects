package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Buscador de IP por linha de comando"
	app.Author = "@emanuelvsz"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ip's na web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: getIP,
		},
	}

	return app
}

func getIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips{
		fmt.Println(ip)
		
	}
}
