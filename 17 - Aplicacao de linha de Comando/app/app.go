package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de linha de comando"
	app.Usage = "Busca IPs e nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de enderecos na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "cname",
			Usage:  "Busca CNAMES de enderecos na internet",
			Flags:  flags,
			Action: buscarCname,
		},
		{
			Name:   "servers",
			Usage:  "Busca o nome do servidores na internet",
			Flags:  flags,
			Action: buscarServers,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarCname(c *cli.Context) {
	host := c.String("host")
	cname, erro := net.LookupCNAME(host)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cname)
}

func buscarServers(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
