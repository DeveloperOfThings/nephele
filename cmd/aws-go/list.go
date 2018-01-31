package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/bharath-srinivas/aws-go/function"
	"github.com/bharath-srinivas/aws-go/spinner"
)

// list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available EC2 instances",
	Args:  cobra.NoArgs,
	Run:   listInstances,
}

func init() {
	Command.AddCommand(listCmd)
}

// run command.
func listInstances(cmd *cobra.Command, args []string) {
	sp := spinner.Default(spinnerPrefix[1])
	sp.Start()
	sess := ec2.New(Session)

	ec2Service := &function.EC2Service{
		Service: sess,
	}

	resp, err := ec2Service.GetInstances()

	if err != nil {
		sp.Stop()
		fmt.Println(err.Error())
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(true)
		table.SetHeader([]string{
			"Instance Name",
			"Instance ID",
			"Instance State",
			"Private IPv4 Address",
			"Public IPv4 Address",
			"Instance Type",
		})
		table.AppendBulk(resp)
		sp.Stop()
		table.Render()
	}
}
