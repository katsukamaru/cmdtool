package cmd

import (
	"github.com/spf13/cobra"

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func init() {
	RootCmd.AddCommand(ec2Cmd)
}

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "List of instances.",
	Long:  "List of instances.",
	Run: func(cmd *cobra.Command, args []string) {
		doSomething()
	},
}

func doSomething() {
	svc := ec2.New(
		session.New(),
		&aws.Config{
			Region: aws.String("ap-northeast-1"),
		},
	)

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
				},
			},
		},
	}

	resp, err := svc.DescribeInstances(params)
	if err != nil {
		panic(err)
	}

	for _, res := range resp.Reservations {
		for _, inst := range res.Instances {
			fmt.Println(inst)
		}
	}
}
