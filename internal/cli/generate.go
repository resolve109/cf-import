package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate CloudFormation template from existing resources",
	Long: `Generate a CloudFormation template from scanned resources using
the IaC generator.

This creates a template that describes your existing AWS resources
in CloudFormation format. The template includes:
  - Resource definitions with current configurations
  - DeletionPolicy set to Retain (required for import)
  - Logical IDs based on resource identifiers

Prerequisites:
  - A completed resource scan (use 'cfimport scan' first)
  - Resources must exist in the same region as the target stack

Examples:
  cfimport generate --scan-id abc123 --output template.yaml
  cfimport generate --resources i-1234567890abcdef0,bucket-name
  cfimport generate --types AWS::EC2::Instance --output ec2-template.yaml`,
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().String("scan-id", "", "Scan ID from a previous scan")
	generateCmd.Flags().StringSliceP("resources", "r", []string{}, "Specific resource IDs to include")
	generateCmd.Flags().StringSlice("types", []string{}, "Resource types to include")
	generateCmd.Flags().StringP("output", "o", "template.yaml", "Output template file")
	generateCmd.Flags().String("format", "yaml", "Output format (yaml, json)")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	scanID, _ := cmd.Flags().GetString("scan-id")
	resources, _ := cmd.Flags().GetStringSlice("resources")
	types, _ := cmd.Flags().GetStringSlice("types")
	output, _ := cmd.Flags().GetString("output")
	format, _ := cmd.Flags().GetString("format")

	if verbose {
		fmt.Printf("Generating CloudFormation template\n")
		if scanID != "" {
			fmt.Printf("Using scan ID: %s\n", scanID)
		}
		if len(resources) > 0 {
			fmt.Printf("Resources: %v\n", resources)
		}
		if len(types) > 0 {
			fmt.Printf("Types: %v\n", types)
		}
	}

	// TODO: Implement using AWS SDK
	// aws cloudformation create-generated-template
	fmt.Printf("Generating template in %s format...\n", format)
	fmt.Printf("Template will be saved to: %s\n", output)

	return nil
}
