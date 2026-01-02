package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "cfimport",
	Short: "Cloudflare Import Tool",
	Long: `cfimport is a CLI tool for importing Cloudflare configurations and resources.

It allows you to import DNS records, firewall rules, page rules,
and other Cloudflare configurations from various sources.

Examples:
  cfimport dns --zone example.com --file dns-records.json
  cfimport firewall --zone example.com --file rules.yaml
  cfimport export --zone example.com --output backup/`,
}

// Execute runs the root command
func Execute(version, commit, date string) error {
	rootCmd.Version = fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date)
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cfimport.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().String("api-token", "", "Cloudflare API token")
	rootCmd.PersistentFlags().String("api-key", "", "Cloudflare API key (legacy)")
	rootCmd.PersistentFlags().String("api-email", "", "Cloudflare account email")

	viper.BindPFlag("api_token", rootCmd.PersistentFlags().Lookup("api-token"))
	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("api_email", rootCmd.PersistentFlags().Lookup("api-email"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err == nil {
			viper.AddConfigPath(home)
			viper.SetConfigName(".cfimport")
		}
		viper.AddConfigPath("/etc/cfimport")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("CF")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}
