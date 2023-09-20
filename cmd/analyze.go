package cmd

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
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
	getCredentials(&loginReader{})
	return nil
}

func getCredentials(determinator credentialReader) (*harborCredentials, error) {
	username, err := determinator.readUsername()
	if err != nil {
		return nil, err
	}

	password, err := determinator.readPassword()
	if err != nil {
		return nil, fmt.Errorf("failed to get password: %w", err)
	}

	return &harborCredentials{
		Username: username,
		Password: password,
	}, nil
}

type harborCredentials struct {
	Username string
	Password string
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
