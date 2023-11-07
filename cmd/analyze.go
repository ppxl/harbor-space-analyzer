package cmd

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/ppxl/harbor-space-analyzer/pkg/core"
	"github.com/ppxl/harbor-space-analyzer/pkg/gfx"
	"github.com/ppxl/harbor-space-analyzer/pkg/service"
)

const (
	flagAnalyzeEndpoint = "endpoint"
	flagAnalyzePieChart = "pie-chart"
)

type credentialReader interface {
	readUsername() (string, error)
	readPassword() (string, error)
}

// Analyze is a cmd entrypoint.
var Analyze = &cli.Command{
	Name:        "analyze",
	Usage:       "analyzes the configured harbor",
	Description: "analyzes the configured harbor",
	Action:      analyzeSpace,
	ArgsUsage:   "-e <https://registryUrl> [-p [radius]]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    flagAnalyzeEndpoint,
			Aliases: []string{"e"},
			Usage:   "the Harbor endpoint URL, mandatory",
		},
		&cli.IntFlag{
			Name:    flagAnalyzePieChart,
			Aliases: []string{"p"},
			Usage:   "print out a nice pie chart, optional (but you'll miss out)",
		},
	},
}

// analyzeSpace does things
func analyzeSpace(cliCtx *cli.Context) error {
	creds, err := getCredentials(&loginReader{})
	if err != nil {
		return err
	}

	ctx := cliCtx.Context

	radius := cliCtx.Int(flagAnalyzePieChart)
	harborURL := cliCtx.String(flagAnalyzeEndpoint)

	//if term.IsTerminal(0) {
	//	println("in a term")
	//} else {
	//	println("not in a term")
	//}
	//width, height, err := term.GetSize(0)
	//if err != nil {
	//	return
	//}
	//println("width:", width, "height:", height)

	args := core.AnalyzerArgs{Credentials: creds, HarborURL: harborURL}

	projSum, err := service.New(args).GetProjectInfo(ctx)
	if err != nil {
		return fmt.Errorf("failed to collect project data: %w", err)
	}
	//fmt.Printf("%#v\n", projSum)

	karacters, usagePercents := service.CalculateValues(projSum)
	// fix hole in diagram consisting of dashes
	usagePercents[len(usagePercents)-1] += 0.001

	gfx.PrintLegend(projSum, karacters, usagePercents)

	if radius == 0 {
		return nil
	}

	pie := gfx.CreateChart(karacters, usagePercents, radius)
	gfx.PrintColorChart(karacters, pie)

	return nil
}

func getCredentials(determinator credentialReader) (*core.HarborCredentials, error) {
	username, err := determinator.readUsername()
	if err != nil {
		return nil, err
	}

	password, err := determinator.readPassword()
	if err != nil {
		return nil, fmt.Errorf("failed to get password: %w", err)
	}

	return &core.HarborCredentials{
		Username: username,
		Password: password,
	}, nil
}

type loginReader struct{}

func (*loginReader) readUsername() (string, error) {
	fmt.Print("\nUsername: ")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read username: %w", err)
	}

	return strings.TrimSpace(username), nil
}

// readPassword reads the password, but no characters are displayed for added security.
func (*loginReader) readPassword() (string, error) {
	fmt.Print("Password: ")

	stdInFileHandle := 0
	bytePassword, err := terminal.ReadPassword(stdInFileHandle)
	if err != nil {
		return "", fmt.Errorf("failed to read password: %w", err)
	}

	fmt.Println()

	return string(bytePassword), nil
}
