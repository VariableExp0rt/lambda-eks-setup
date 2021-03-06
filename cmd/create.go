package cmd

import (
	"fmt"
	"os"

	"github.com/VariableExp0rt/lambda-and-fun/config/helper"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdCreate)
	cmdCreate.AddCommand(cmdCreateRole)
	cmdCreate.AddCommand(cmdAttachPolicy)
	cmdCreate.AddCommand(cmdCreateLambda)
	cmdCreate.AddCommand(cmdCreateGateway)
}

var cmdCreate = &cobra.Command{
	Use:   "create [resource to setup]",
	Short: "Create tells the program to create a resource",
	Long:  "Use this flag to signal the program to create the requested resource, which is specified in a subcommand",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) == 0 {
			fmt.Println("Supply at least one subcommand to create a resource")
		}
	},
}

var cmdCreateRole = &cobra.Command{
	Use:   "role [args]",
	Short: "Create an AWS IAM Role",
	Long:  "This command creates an AWS IAM Role, which can be used to attach policies and for deploying other services",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		role, err := helper.CreateRole(RoleArgs, sess)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println("Role created: ", role)
		}
	},
}

var cmdAttachPolicy = &cobra.Command{
	Use:   "attach-policy",
	Short: "Attach an IAM Policy to a Role",
	Long: `This command allows users to attach either managed or self created policies
				to an IAM role. This is useful for adding privileges to roles that need
				to perform certain functions.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := helper.AttachPolicy(AttachPolArgs, sess)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println("Attached policy to role: ", res)
		}
	},
}

var cmdCreateLambda = &cobra.Command{
	Use:   "lambda [args]",
	Short: "Create a Lambda function",
	Long: `Supplying the lambda argument to the create command creates a new lambda function
	with the given arguments which configures the function`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		lmb, err := helper.CreateLambda(LambdaArgs, sess)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println("Function created: ", lmb)
		}
	},
}

var cmdCreateGateway = &cobra.Command{
	Use:   "gateway [args]",
	Short: "Create an API Gateway resource",
	Long: `This subcommand creates a new API Gateway service that will be used to expose other services
	or to trigger our workloads through a Lambda, via HTTPS.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gwy, id, err := helper.CreateGateway(GatewayArgs, sess)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println("API Gateway created: ", gwy)
		}
		helper.ConfigureAPIEndpoint(id, gwy.Id, gwy.Name, LambdaArgs.FunctionName, sess)
	},
}
