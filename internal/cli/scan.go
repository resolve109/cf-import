package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan AWS account for existing resources",
	Long: `Scan your AWS account to discover existing resources that can be
imported into CloudFormation.

This initiates a resource scan using the CloudFormation IaC generator.
Scans are region-wide and expire after 30 days. You can run up to
3 scans per day in an account.

The scan discovers resources that:
  - Exist in the specified AWS region
  - Are not already managed by CloudFormation
  - Support CloudFormation import

Examples:
  cfimport scan --region us-east-1
  cfimport scan --region us-west-2 --output scan-results.json
  cfimport scan --types AWS::EC2::Instance,AWS::S3::Bucket`,
	RunE: runScan,
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringP("output", "o", "", "Output file for scan results (JSON)")
	scanCmd.Flags().StringSlice("types", []string{}, "Resource types to scan for (default: all supported)")
	scanCmd.Flags().Bool("wait", true, "Wait for scan to complete")
}

func runScan(cmd *cobra.Command, args []string) error {
	output, _ := cmd.Flags().GetString("output")
	types, _ := cmd.Flags().GetStringSlice("types")
	wait, _ := cmd.Flags().GetBool("wait")

	if verbose {
		fmt.Printf("Initiating resource scan in region: %s\n", region)
		if len(types) > 0 {
			fmt.Printf("Resource types: %v\n", types)
		}
	}

	// TODO: Implement using AWS SDK
	// aws cloudformation start-resource-scan
	fmt.Println("Resource scan initiated...")

	if wait {
		fmt.Println("Waiting for scan to complete...")
	}

	if output != "" {
		fmt.Printf("Results will be saved to: %s\n", output)
	}

	return nil
}
