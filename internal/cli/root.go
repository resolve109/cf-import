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
	region  string
	profile string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "cfimport",
	Short: "AWS CloudFormation Import Tool",
	Long: `cfimport automates the process of importing existing AWS resources
into CloudFormation stacks.

It simplifies the manual steps required for CloudFormation resource imports,
including template generation, resource scanning, and stack operations.

Examples:
  cfimport scan --region us-east-1
  cfimport generate --resources ec2,s3 --output template.yaml
  cfimport import --stack mystack --template template.yaml --resources resources.json
  cfimport iac-generator --scan --region us-east-1`,
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
	rootCmd.PersistentFlags().StringVar(&region, "region", "", "AWS region (overrides AWS_REGION)")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "", "AWS profile to use")

	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
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

	viper.SetEnvPrefix("AWS")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}
