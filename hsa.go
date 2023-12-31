package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/ppxl/harbor-space-analyzer/cmd"
	"github.com/ppxl/harbor-space-analyzer/logging"
)

const flagLogLevel = "log-level"

var (
	version string
	log            = logging.GetInstance()
	osExit  anExit = &realOsExit{}
)

func configureLogging(cliCtx *cli.Context) error {
	logLevelStr := cliCtx.String(flagLogLevel)
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		return fmt.Errorf("failed to parse log level %s: %w", logLevelStr, err)
	}
	err = logging.Init(logLevel)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "hsa"
	app.Usage = "analyze a harbor instance in terms of space"
	app.Commands = []*cli.Command{cmd.Analyze}
	app.HideHelpCommand = true

	app.Flags = parseGlobalFlags()
	app.Before = func(cliCtx *cli.Context) error {
		err := configureLogging(cliCtx)
		exitOnError(err)
		return nil
	}

	err := app.Run(os.Args)
	exitOnError(err)
}

func parseGlobalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  flagLogLevel,
			Usage: "a log level of lower than error will produce all kind of silly output",
			Value: "error",
		},
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("%#v\n", err)
		osExit.Exit(1)
	}
}

type anExit interface {
	Exit(exitCode int)
}

type realOsExit struct{}

func (ex *realOsExit) Exit(exitCode int) {
	os.Exit(exitCode)
}
