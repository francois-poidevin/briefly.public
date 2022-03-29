package cmd

import (
	"fmt"
	"os"

	defaults "github.com/mcuadros/go-defaults"

	"github.com/francois-poidevin/briefly.public/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	log     *logrus.Logger
	cfgFile string
	conf    = &config.Configuration{}
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
	// // Initialize config
	// initConfig()

	// //log handling
	// log = logrus.New()

	// fmt.Println(fmt.Sprintf("log format json: %t ", conf.Log.JSONFormatter))
	// if conf.Log.JSONFormatter {
	// 	fmt.Println("log format: Json")
	// 	log.Formatter = new(logrus.JSONFormatter)
	// } else {
	// 	fmt.Println("log format: Text")
	// 	log.Formatter = new(logrus.TextFormatter) //default
	// }
	// log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	// log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	// log.Level = logrus.TraceLevel

	// lvl, err := logrus.ParseLevel(conf.Log.Level)
	// if err != nil {
	// 	log.WithFields(logrus.Fields{
	// 		"Error": err,
	// 	}).Fatal("Not success to parse logrus log level")
	// }
	// log.Level = lvl
	// log.Out = os.Stdout

	rootCmd.AddCommand(startHttpCmd)
	rootCmd.AddCommand(configCmd)
}

func initConfig() {
	// Apply defaults first
	defaults.SetDefaults(conf)

	if cfgFile != "" {
		// If the config file doesn't exists, let's exit
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			// log.WithFields(logrus.Fields{
			// 	"Error": err,
			// }).Fatal("File doesn't exists")
			fmt.Println(fmt.Sprintf("File doesn't exists : %s", err))
		}
		// log.WithFields(logrus.Fields{
		// 	"File": cfgFile,
		// }).Info("Reading configuration file")
		fmt.Println(fmt.Sprintf("Reading configuration file : %s", cfgFile))

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			// log.WithFields(logrus.Fields{
			// 	"Error": err,
			// }).Fatal("Unable to read config")
			fmt.Println(fmt.Sprintf("Unable to read config : %s", err))
		}
	}

	if err := viper.Unmarshal(conf); err != nil {
		// log.WithFields(logrus.Fields{
		// 	"Error": err,
		// }).Fatal("Unable to parse config")
		fmt.Println(fmt.Sprintf("Unable to parse config : %s", err))
	}
}
