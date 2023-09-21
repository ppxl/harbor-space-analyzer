package cmd

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
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
func Analyze() *cli.Command {
	return &cli.Command{
		Name:        "analyze",
		Usage:       "analyzes the configured harbor",
		Description: "analyzes the configured harbor",
		Action:      AnalyzeSpace,
		ArgsUsage:   "-e <https://registryUrl>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  flagAnalyzeEndpoint + ", e",
				Usage: "the Harbor endpoint URL, mandatory",
			},
			&cli.IntFlag{
				Name:  flagAnalyzePieChart + ", p",
				Usage: "print out a nice pie chart, optional (but you'll miss out)",
			},
		},
	}
}

// AnalyzeSpace does things
func AnalyzeSpace(cliCtx *cli.Context) error {
	creds, err := getCredentials(&loginReader{})
	if err != nil {
		return err
	}

	ctx := cliCtx.Context

	// testdata
	//karacters := []string{"A", "B", "x", "y", "z", "q", "p"}
	//values := []float64{0.5, 0.25, 0.05, 0.05, 0.09, 0.01, 0.06}

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
		return err
	}
	//fmt.Printf("%#v\n", projSum)

	karacters, values := service.CalculateValues(projSum)

	fmt.Printf("%#v\n", values)

	if radius == 0 {
		return nil
	}

	pie := gfx.PrintChart(karacters, values, radius)
	fmt.Println(pie)

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
	fmt.Print("Username: ")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read username: %w", err)
	}

	return strings.TrimSpace(username), nil
}

// readPassword reads the password, but no characters are displayed for added security.
func (*loginReader) readPassword() (string, error) {
	log.Print("Password: ")

	stdInFileHandle := 0
	bytePassword, err := terminal.ReadPassword(stdInFileHandle)
	if err != nil {
		return "", fmt.Errorf("failed to read password: %w", err)
	}

	fmt.Println()

	return string(bytePassword), nil
}
