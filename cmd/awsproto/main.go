package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/dedis/protobuf"
	"io/ioutil"
	"os"
	"path"
	"reflect"
)

type protonamer struct{}

func (p *protonamer) FieldName(f protobuf.ProtoField) string {
	if f.Name != "" {
		return f.Name
	}
	return f.Field.Name
}

func (p *protonamer) TypeName(name string) string {
	return name
}

func (p *protonamer) ConstName(name string) string {
	return name
}

var (
	topTypes = map[string][]interface{}{
		"autoscaling": []interface{}{
			autoscaling.DescribeAutoScalingGroupsInput{},
			autoscaling.DescribeAutoScalingGroupsOutput{},
		},
		"cloudwatch": []interface{}{
			cloudwatch.ListMetricsInput{},
			cloudwatch.ListMetricsOutput{},
			cloudwatch.GetMetricStatisticsInput{},
			cloudwatch.GetMetricStatisticsOutput{},
			cloudwatch.DescribeAlarmsInput{},
			cloudwatch.DescribeAlarmsOutput{},
			cloudwatch.DescribeAlarmsForMetricInput{},
			cloudwatch.DescribeAlarmsForMetricOutput{},
		},
		"credentials": []interface{}{
			credentials.Value{},
		},
		"ec2": []interface{}{
			ec2.DescribeInstancesInput{},
			ec2.DescribeInstancesOutput{},
			ec2.DescribeSecurityGroupsInput{},
			ec2.DescribeSecurityGroupsOutput{},
			ec2.DescribeVpcsInput{},
			ec2.DescribeVpcsOutput{},
			ec2.DescribeSubnetsInput{},
			ec2.DescribeSubnetsOutput{},
			ec2.DescribeRouteTablesInput{},
			ec2.DescribeRouteTablesOutput{},
		},
		"ecs": []interface{}{
			ecs.ListTasksInput{},
			ecs.ListTasksOutput{},
			ecs.ListServicesInput{},
			ecs.ListServicesOutput{},
			ecs.DescribeServicesInput{},
			ecs.DescribeServicesOutput{},
			ecs.ListClustersInput{},
			ecs.ListClustersOutput{},
			ecs.DescribeTasksInput{},
			ecs.DescribeTasksOutput{},
			ecs.DescribeContainerInstancesInput{},
			ecs.DescribeContainerInstancesOutput{},
			ecs.ListContainerInstancesInput{},
			ecs.ListContainerInstancesOutput{},
			ecs.DescribeTaskDefinitionInput{},
			ecs.DescribeTaskDefinitionOutput{},
		},
		"elb": []interface{}{
			elb.DescribeLoadBalancersInput{},
			elb.DescribeLoadBalancersOutput{},
		},
		"rds": []interface{}{
			rds.DescribeDBInstancesInput{},
			rds.DescribeDBInstancesOutput{},
		},
	}

	basePath = flag.String("basepath", "", "the base path for aws proto")
	header   = `syntax = "proto2";
import "github.com/gogo/protobuf/gogoproto/gogo.proto";
import "github.com/opsee/protobuf/opseeproto/opsee.proto";
import "github.com/opsee/protobuf/opseeproto/types/timestamp.proto";

package opsee.aws.%s;

option go_package = "%s";
option (opseeproto.graphql) = true;

`
)

func main() {
	flag.Parse()

	for pkg, tlist := range topTypes {
		var (
			buf     = bytes.NewBuffer([]byte{})
			typeMap = make(map[string]interface{})
		)

		for _, t := range tlist {
			typeMap[reflect.TypeOf(t).String()] = t
			findTypes(t, typeMap)
		}

		allTypes := make([]interface{}, 0, len(typeMap))
		for _, t := range typeMap {
			allTypes = append(allTypes, t)
		}

		buf.WriteString(fmt.Sprintf(header, pkg, pkg))

		protobuf.GenerateProtobufDefinition(buf, allTypes, nil, &protonamer{})

		sanitized := bytes.Replace(buf.Bytes(), []byte("  required  _ = 1;\n"), []byte{}, -1)
		sanitized = bytes.Replace(sanitized, []byte("  required"), []byte("  optional"), -1)
		sanitized = bytes.Replace(sanitized, []byte("sfixed64"), []byte("opsee.types.Timestamp"), -1)
		sanitized = bytes.Replace(sanitized, []byte("optional map"), []byte("map"), -1)

		p := path.Join(*basePath, pkg)
		if err := os.MkdirAll(p, 0777); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		if err := ioutil.WriteFile(path.Join(p, "types.proto"), sanitized, 0644); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
}

func findTypes(t interface{}, typeMap map[string]interface{}) {
	reflectValue := reflect.ValueOf(t)
	for i := 0; i < reflectValue.NumField(); i++ {
		fieldValue := reflectValue.FieldByIndex([]int{i})
		var fieldType reflect.Type

		switch fieldValue.Kind() {
		case reflect.Ptr, reflect.Slice:
			fieldType = fieldValue.Type().Elem()

			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}

			switch fieldType.Kind() {
			case reflect.Struct:
				if fieldType.PkgPath() != "time" {
					tname := fieldType.String()

					if _, ok := typeMap[tname]; !ok {
						typeMap[tname] = reflect.Zero(fieldType).Interface()
						findTypes(typeMap[tname], typeMap)
					}
				}
			}
		}
	}
}
