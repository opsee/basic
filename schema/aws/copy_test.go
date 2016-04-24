package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	opsee_ec2 "github.com/opsee/basic/schema/aws/ec2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	assert := assert.New(t)

	awsStructSrc := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("hello"),
				Values: []*string{aws.String("hellovalue")},
			},
		},
		InstanceIds: []*string{aws.String("id1"), aws.String("id2")},
	}

	opseeStructDst := &opsee_ec2.DescribeInstancesInput{}
	CopyInto(opseeStructDst, awsStructSrc)

	assert.Equal(opseeStructDst.InstanceIds[0], "id1")
	assert.Equal(opseeStructDst.InstanceIds[1], "id2")
	assert.Equal(aws.StringValue(opseeStructDst.Filters[0].Name), "hello")
	assert.Equal(opseeStructDst.Filters[0].Values[0], "hellovalue")

	opseeStructSrc := &opsee_ec2.DescribeInstancesInput{
		Filters: []*opsee_ec2.Filter{
			{
				Name:   aws.String("hello"),
				Values: []string{"hellovalue"},
			},
		},
		InstanceIds: []string{"id1", "id2"},
	}

	awsStructDst := &ec2.DescribeInstancesInput{}
	CopyInto(awsStructDst, opseeStructSrc)

	assert.Equal(aws.StringValue(awsStructDst.InstanceIds[0]), "id1")
	assert.Equal(aws.StringValue(awsStructDst.InstanceIds[1]), "id2")
	assert.Equal(aws.StringValue(awsStructDst.Filters[0].Name), "hello")
	assert.Equal(aws.StringValue(awsStructDst.Filters[0].Values[0]), "hellovalue")
}
