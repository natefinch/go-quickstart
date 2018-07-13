package run

import (
	"fmt"
	"strings"

	"github.com/natefinch/go-quickstart/server/cmd/my-server/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Update these, they're used for the cobra command and description.
const (
	appname     = "My-Server"
	description = "A go-quickstart default server"
)

var (
	timestamp  = "(unknown)"
	gitTag     = "(unknown)"
	commitHash = "(unknown)"
)

// Run starts the server. It returns the code the application should exit with.
func Run() int {
	vals := config.Init(viper.New())
	logger := logrus.New()
	logger.Level = logrus.DebugLevel

	// rootCmd is the root cobra command for the daemon.
	rootCmd := &cobra.Command{
		Use:   strings.ToLower(appname),
		Short: description,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use the help command to list all commands.")
		},
	}
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "version",
			Short: "Print the version of the application",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("%s v%s\n", appname, gitTag)
				fmt.Println("Build Date:", timestamp)
				fmt.Println("Commit:", commitHash)
			},
		},
		&cobra.Command{
			Use:   "server",
			Short: "Start the server",
			RunE: func(cmd *cobra.Command, args []string) error {
				return runServer(logger, vals)
			},
		},
		// TODO: add more commands here as needed.
	)
	if err := rootCmd.Execute(); err != nil {
		logger.WithError(err).Error()
		return 1
	}
	return 0
}

func runServer(logger *logrus.Logger, v *config.Values) error {
	// TODO: Call your server implementation here.
	// Don't *implement* the server in this package.  This file should just
	// translate the config values into a format the server function requires,
	// and pass it those values and the logger.  Do *not* pass the config.Values
	// struct directly.  That's the definition of the environment config.  It is
	// constrained by the limited types available to configuration languages.
	// Your server package should define its own configuration struct that takes
	// data in the format that best represents the requirements (which may well
	// be impossible to directly express in a configuration language).
	return nil
}
