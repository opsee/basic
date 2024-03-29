package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	opsee_cloudwatch "github.com/opsee/basic/schema/aws/cloudwatch"
	opsee_ec2 "github.com/opsee/basic/schema/aws/ec2"
	opsee_ecs "github.com/opsee/basic/schema/aws/ecs"
	opsee_types "github.com/opsee/protobuf/opseeproto/types"
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

	opseeStructSrc2 := &opsee_cloudwatch.GetMetricStatisticsInput{
		StartTime: &opsee_types.Timestamp{
			Seconds: int64(666),
			Nanos:   int32(666),
		},
	}

	awsStructDst2 := &cloudwatch.GetMetricStatisticsInput{}
	CopyInto(awsStructDst2, opseeStructSrc2)
	assert.NotZero(aws.TimeValue(awsStructDst2.StartTime))

	awsStructSrcEcs := &ecs.ContainerDefinition{
		DockerLabels: map[string]*string{
			"hello": aws.String("there"),
		},
		DockerSecurityOptions: []*string{
			aws.String("v secure"),
		},
	}
	opseeStructDstEcs := &opsee_ecs.ContainerDefinition{}
	CopyInto(opseeStructDstEcs, awsStructSrcEcs)
	assert.Equal("v secure", opseeStructDstEcs.DockerSecurityOptions[0])
	assert.Equal("there", opseeStructDstEcs.DockerLabels["hello"])

	// TODO: implement me
	// opseeStructSrcEcs := &opsee_ecs.ContainerDefinition{
	// 	DockerLabels: map[string]string{
	// 		"hello": "there",
	// 	},
	// 	DockerSecurityOptions: []string{
	// 		"v secure",
	// 	},
	// }
	// awsStructDstEcs := &ecs.ContainerDefinition{}
	// CopyInto(awsStructDstEcs, opseeStructSrcEcs)
	// assert.Equal(aws.String("v secure"), awsStructDstEcs.DockerSecurityOptions[0])
	// assert.Equal(aws.String("there"), awsStructDstEcs.DockerLabels["hello"])
}
