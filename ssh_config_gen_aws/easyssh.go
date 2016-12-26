package main

import (
	"html/template"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type HostDefinition struct {
	Host, IpAddr string
}

const ssh_config_sample = `
{{range .}}
Host {{.Host}}
	HostName {{.IpAddr}}
	User arch
	Port 3417
{{end}}
`

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sess, err := session.NewSession()
	check_err(err)
	ec2obj := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{ // Required
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"), // Required
					// More values...
				},
			},
		},
	}
	resp, err := ec2obj.DescribeInstances(params)
	check_err(err)
	// fmt.Println(resp)
	all_hosts := []HostDefinition{}
	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			single_host := HostDefinition{}
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					single_host.Host = *tag.Value
				}
			}
			// if vpc use private ip
			if inst.VpcId != nil {
				single_host.IpAddr = *inst.PrivateIpAddress
			} else {
				single_host.IpAddr = *inst.PublicIpAddress
			}

			// fmt.Println(single_host)
			// append to array
			all_hosts = append(all_hosts, single_host)
		}
	}
	template_sample := template.Must(template.New("sample").Parse(ssh_config_sample))
	check_err(template_sample.Execute(os.Stdout, all_hosts))
}
