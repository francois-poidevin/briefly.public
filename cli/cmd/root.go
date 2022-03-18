package cmd

import (
	"fmt"
	"os"

	"github.com/francois-poidevin/briefly.public/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	log *logrus.Logger
	// cfgFile string
	conf = &config.Configuration{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Briefly.public",
	Short: "Briefly.public application allow to redirect a shortcode URL",
	Long:  `Briefly.public application allow to redirect a shortcode URL.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//log handling
	log = logrus.New()
	// log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                     //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	rootCmd.AddCommand(startHttpCmd)
	rootCmd.AddCommand(configCmd)
}
