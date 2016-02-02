package main

import (
	"os"

	etcdapp "github.com/eris-ltd/etcd-tmsp/app"

	"github.com/codegangsta/cli"
)

const Version = "1.3.3-tmsp"

func main() {
	app := utils.NewApp(Version, "the etcd-tmsp command line interface")
	app.Action = run
	app.Flags = []cli.Flag{
		etcdapp.TendermintCoreHostFlag,
	}
	app.Run(os.Args)
}

func run(ctx *cli.Context) {
	// Start the tmsp listener
	_, err := server.StartListener("tcp://0.0.0.0:46658", etcdapp.NewETCDApplication(ctx))
	if err != nil {
		Exit(err.Error())
	}

	// Wait forever
	TrapSignal(func() {
		// TODO: Cleanup
	})
}
