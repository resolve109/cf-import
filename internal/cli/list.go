package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources that support CloudFormation import",
	Long: `List AWS resources that support CloudFormation import operations.

This helps you understand which resource types can be imported
and shows the required identifiers for each type.

Examples:
  cfimport list --types           # List all supported resource types
  cfimport list --scan-id abc123  # List resources from a scan
  cfimport list --stack mystack   # List resources in a stack`,
	RunE: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().Bool("types", false, "List supported resource types")
	listCmd.Flags().String("scan-id", "", "List resources from a scan")
	listCmd.Flags().StringP("stack", "s", "", "List resources in a stack")
	listCmd.Flags().String("format", "table", "Output format (table, json, yaml)")
}

func runList(cmd *cobra.Command, args []string) error {
	showTypes, _ := cmd.Flags().GetBool("types")
	scanID, _ := cmd.Flags().GetString("scan-id")
	stack, _ := cmd.Flags().GetString("stack")
	format, _ := cmd.Flags().GetString("format")

	if verbose {
		fmt.Printf("Output format: %s\n", format)
	}

	if showTypes {
		fmt.Println("Supported resource types for import:")
		fmt.Println("  AWS::EC2::Instance")
		fmt.Println("  AWS::EC2::VPC")
		fmt.Println("  AWS::EC2::Subnet")
		fmt.Println("  AWS::EC2::SecurityGroup")
		fmt.Println("  AWS::S3::Bucket")
		fmt.Println("  AWS::RDS::DBInstance")
		fmt.Println("  AWS::Lambda::Function")
		fmt.Println("  AWS::DynamoDB::Table")
		fmt.Println("  ... (use --format json for full list)")
		return nil
	}

	if scanID != "" {
		// TODO: Implement using AWS SDK
		// aws cloudformation list-resource-scan-resources
		fmt.Printf("Listing resources from scan: %s\n", scanID)
		return nil
	}

	if stack != "" {
		// TODO: Implement using AWS SDK
		// aws cloudformation list-stack-resources
		fmt.Printf("Listing resources in stack: %s\n", stack)
		return nil
	}

	return fmt.Errorf("please specify --types, --scan-id, or --stack")
}
