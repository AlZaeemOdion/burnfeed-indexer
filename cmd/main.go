package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/burnfeed/indexer/cmd/flags"
	"github.com/burnfeed/indexer/cmd/logger"
	"github.com/burnfeed/indexer/indexer"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Name = "BurnFeed Actions indexer"
	app.Usage = "Actions indexer for BurnFeed protocol"
	app.EnableBashCompletion = true
	app.Flags = flags.CommonFlags
	app.Action = func(c *cli.Context) error {
		logger.InitLogger(c)

		ctx, ctxClose := context.WithCancel(context.Background())
		defer func() { ctxClose() }()

		cfg, err := indexer.NewConfigFromCliContext(c)
		if err != nil {
			return err
		}

		indexer, err := indexer.New(ctx, cfg)
		if err != nil {
			return err
		}

		go indexer.Start()

		log.Info("Action indexer started")

		defer func() {
			ctxClose()
			log.Info("Actions indexer stopped")
		}()

		quitCh := make(chan os.Signal, 1)
		signal.Notify(quitCh, []os.Signal{
			os.Interrupt,
			os.Kill,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		}...)
		<-quitCh

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("Failed to start BurnFeed Actions indexer", "error", err)
	}
}
