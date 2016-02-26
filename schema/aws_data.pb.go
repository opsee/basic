// Code generated by protoc-gen-gogo.
// source: aws_data.proto
// DO NOT EDIT!

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	aws_data.proto
	checks.proto
	user.proto

It has these top-level messages:
	Group
	Instance
	Target
	Check
	Assertion
	Header
	HttpCheck
	CloudWatchCheck
	Metric
	HttpResponse
	CheckResponse
	CheckResult
	User
*/
package schema

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/opsee/protobuf/opseeproto"
import opsee_types "github.com/opsee/protobuf/opseeproto/types"
import opsee_aws_autoscaling "github.com/opsee/basic/schema/aws/autoscaling"
import opsee_aws_ec2 "github.com/opsee/basic/schema/aws/ec2"
import opsee_aws_elb "github.com/opsee/basic/schema/aws/elb"
import opsee_aws_rds "github.com/opsee/basic/schema/aws/rds"

import github_com_graphql_go_graphql "github.com/graphql-go/graphql"
import github_com_opsee_protobuf_plugin_graphql_scalars "github.com/opsee/protobuf/plugin/graphql/scalars"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An Opsee group target representing an AWS group resource.
type Group struct {
	// The AWS group identifier.
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,proto3" json:"customer_id,omitempty"`
	// The type of AWS resource.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// The raw AWS resource data.
	//
	// Types that are valid to be assigned to Resource:
	//	*Group_SecurityGroup
	//	*Group_LoadBalancer
	//	*Group_AutoscalingGroup
	Resource isGroup_Resource `protobuf_oneof:"resource"`
	// The last seen number of instances in the group target. This value is cached, so it may not be consistent.
	InstanceCount int32 `protobuf:"varint,4,opt,name=instance_count,proto3" json:"instance_count,omitempty"`
	// The last seen instances in the group target. This value is cached, so it may not be consistent.
	Instances []*Instance            `protobuf:"bytes,5,rep,name=instances" json:"instances,omitempty"`
	CreatedAt *opsee_types.Timestamp `protobuf:"bytes,6,opt,name=created_at" json:"created_at,omitempty"`
	UpdatedAt *opsee_types.Timestamp `protobuf:"bytes,7,opt,name=updated_at" json:"updated_at,omitempty"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}

type isGroup_Resource interface {
	isGroup_Resource()
	Equal(interface{}) bool
}

type Group_SecurityGroup struct {
	SecurityGroup *opsee_aws_ec2.SecurityGroup `protobuf:"bytes,101,opt,name=security_group,oneof"`
}
type Group_LoadBalancer struct {
	LoadBalancer *opsee_aws_elb.LoadBalancerDescription `protobuf:"bytes,102,opt,name=load_balancer,oneof"`
}
type Group_AutoscalingGroup struct {
	AutoscalingGroup *opsee_aws_autoscaling.Group `protobuf:"bytes,103,opt,name=autoscaling_group,oneof"`
}

func (*Group_SecurityGroup) isGroup_Resource()    {}
func (*Group_LoadBalancer) isGroup_Resource()     {}
func (*Group_AutoscalingGroup) isGroup_Resource() {}

func (m *Group) GetResource() isGroup_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Group) GetSecurityGroup() *opsee_aws_ec2.SecurityGroup {
	if x, ok := m.GetResource().(*Group_SecurityGroup); ok {
		return x.SecurityGroup
	}
	return nil
}

func (m *Group) GetLoadBalancer() *opsee_aws_elb.LoadBalancerDescription {
	if x, ok := m.GetResource().(*Group_LoadBalancer); ok {
		return x.LoadBalancer
	}
	return nil
}

func (m *Group) GetAutoscalingGroup() *opsee_aws_autoscaling.Group {
	if x, ok := m.GetResource().(*Group_AutoscalingGroup); ok {
		return x.AutoscalingGroup
	}
	return nil
}

func (m *Group) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *Group) GetCreatedAt() *opsee_types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Group) GetUpdatedAt() *opsee_types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Group) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Group_OneofMarshaler, _Group_OneofUnmarshaler, []interface{}{
		(*Group_SecurityGroup)(nil),
		(*Group_LoadBalancer)(nil),
		(*Group_AutoscalingGroup)(nil),
	}
}

func _Group_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Group)
	// resource
	switch x := m.Resource.(type) {
	case *Group_SecurityGroup:
		_ = b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SecurityGroup); err != nil {
			return err
		}
	case *Group_LoadBalancer:
		_ = b.EncodeVarint(102<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoadBalancer); err != nil {
			return err
		}
	case *Group_AutoscalingGroup:
		_ = b.EncodeVarint(103<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AutoscalingGroup); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Group.Resource has unexpected type %T", x)
	}
	return nil
}

func _Group_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Group)
	switch tag {
	case 101: // resource.security_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(opsee_aws_ec2.SecurityGroup)
		err := b.DecodeMessage(msg)
		m.Resource = &Group_SecurityGroup{msg}
		return true, err
	case 102: // resource.load_balancer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(opsee_aws_elb.LoadBalancerDescription)
		err := b.DecodeMessage(msg)
		m.Resource = &Group_LoadBalancer{msg}
		return true, err
	case 103: // resource.autoscaling_group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(opsee_aws_autoscaling.Group)
		err := b.DecodeMessage(msg)
		m.Resource = &Group_AutoscalingGroup{msg}
		return true, err
	default:
		return false, nil
	}
}

// An Opsee instance target representing an AWS instance resource.
type Instance struct {
	// The AWS instance identifier.
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,proto3" json:"customer_id,omitempty"`
	// The type of AWS resource.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// The raw AWS resource data.
	//
	// Types that are valid to be assigned to Resource:
	//	*Instance_Instance
	//	*Instance_DbInstance
	Resource isInstance_Resource `protobuf_oneof:"resource"`
	// The last seen group targets that the instance belongs to. This value is cached, so it may not be consistent.
	Groups    []*Group               `protobuf:"bytes,4,rep,name=groups" json:"groups,omitempty"`
	CreatedAt *opsee_types.Timestamp `protobuf:"bytes,5,opt,name=created_at" json:"created_at,omitempty"`
	UpdatedAt *opsee_types.Timestamp `protobuf:"bytes,6,opt,name=updated_at" json:"updated_at,omitempty"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}

type isInstance_Resource interface {
	isInstance_Resource()
	Equal(interface{}) bool
}

type Instance_Instance struct {
	Instance *opsee_aws_ec2.Instance `protobuf:"bytes,101,opt,name=instance,oneof"`
}
type Instance_DbInstance struct {
	DbInstance *opsee_aws_rds.DBInstance `protobuf:"bytes,102,opt,name=db_instance,oneof"`
}

func (*Instance_Instance) isInstance_Resource()   {}
func (*Instance_DbInstance) isInstance_Resource() {}

func (m *Instance) GetResource() isInstance_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Instance) GetInstance() *opsee_aws_ec2.Instance {
	if x, ok := m.GetResource().(*Instance_Instance); ok {
		return x.Instance
	}
	return nil
}

func (m *Instance) GetDbInstance() *opsee_aws_rds.DBInstance {
	if x, ok := m.GetResource().(*Instance_DbInstance); ok {
		return x.DbInstance
	}
	return nil
}

func (m *Instance) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *Instance) GetCreatedAt() *opsee_types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Instance) GetUpdatedAt() *opsee_types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Instance) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Instance_OneofMarshaler, _Instance_OneofUnmarshaler, []interface{}{
		(*Instance_Instance)(nil),
		(*Instance_DbInstance)(nil),
	}
}

func _Instance_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Instance)
	// resource
	switch x := m.Resource.(type) {
	case *Instance_Instance:
		_ = b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Instance); err != nil {
			return err
		}
	case *Instance_DbInstance:
		_ = b.EncodeVarint(102<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DbInstance); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Instance.Resource has unexpected type %T", x)
	}
	return nil
}

func _Instance_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Instance)
	switch tag {
	case 101: // resource.instance
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(opsee_aws_ec2.Instance)
		err := b.DecodeMessage(msg)
		m.Resource = &Instance_Instance{msg}
		return true, err
	case 102: // resource.db_instance
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(opsee_aws_rds.DBInstance)
		err := b.DecodeMessage(msg)
		m.Resource = &Instance_DbInstance{msg}
		return true, err
	default:
		return false, nil
	}
}

func init() {
	proto.RegisterType((*Group)(nil), "opsee.Group")
	proto.RegisterType((*Instance)(nil), "opsee.Instance")
}
func (this *Group) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Group)
	if !ok {
		that2, ok := that.(Group)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.CustomerId != that1.CustomerId {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if that1.Resource == nil {
		if this.Resource != nil {
			return false
		}
	} else if this.Resource == nil {
		return false
	} else if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if this.InstanceCount != that1.InstanceCount {
		return false
	}
	if len(this.Instances) != len(that1.Instances) {
		return false
	}
	for i := range this.Instances {
		if !this.Instances[i].Equal(that1.Instances[i]) {
			return false
		}
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return false
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return false
	}
	return true
}
func (this *Group_SecurityGroup) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Group_SecurityGroup)
	if !ok {
		that2, ok := that.(Group_SecurityGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.SecurityGroup.Equal(that1.SecurityGroup) {
		return false
	}
	return true
}
func (this *Group_LoadBalancer) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Group_LoadBalancer)
	if !ok {
		that2, ok := that.(Group_LoadBalancer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.LoadBalancer.Equal(that1.LoadBalancer) {
		return false
	}
	return true
}
func (this *Group_AutoscalingGroup) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Group_AutoscalingGroup)
	if !ok {
		that2, ok := that.(Group_AutoscalingGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.AutoscalingGroup.Equal(that1.AutoscalingGroup) {
		return false
	}
	return true
}
func (this *Instance) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Instance)
	if !ok {
		that2, ok := that.(Instance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.CustomerId != that1.CustomerId {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if that1.Resource == nil {
		if this.Resource != nil {
			return false
		}
	} else if this.Resource == nil {
		return false
	} else if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if len(this.Groups) != len(that1.Groups) {
		return false
	}
	for i := range this.Groups {
		if !this.Groups[i].Equal(that1.Groups[i]) {
			return false
		}
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return false
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return false
	}
	return true
}
func (this *Instance_Instance) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Instance_Instance)
	if !ok {
		that2, ok := that.(Instance_Instance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Instance.Equal(that1.Instance) {
		return false
	}
	return true
}
func (this *Instance_DbInstance) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Instance_DbInstance)
	if !ok {
		that2, ok := that.(Instance_DbInstance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.DbInstance.Equal(that1.DbInstance) {
		return false
	}
	return true
}

type GroupGetter interface {
	GetGroup() *Group
}

var GraphQLGroupType *github_com_graphql_go_graphql.Object
var GraphQLGroupResourceUnion *github_com_graphql_go_graphql.Union

type InstanceGetter interface {
	GetInstance() *Instance
}

var GraphQLInstanceType *github_com_graphql_go_graphql.Object
var GraphQLInstanceResourceUnion *github_com_graphql_go_graphql.Union

func (g *Group_SecurityGroup) GetSecurityGroup() *opsee_aws_ec2.SecurityGroup {
	return g.SecurityGroup
}
func (g *Group_LoadBalancer) GetLoadBalancerDescription() *opsee_aws_elb.LoadBalancerDescription {
	return g.LoadBalancer
}
func (g *Group_AutoscalingGroup) GetGroup() *opsee_aws_autoscaling.Group {
	return g.AutoscalingGroup
}
func (g *Instance_Instance) GetInstance() *opsee_aws_ec2.Instance {
	return g.Instance
}
func (g *Instance_DbInstance) GetDBInstance() *opsee_aws_rds.DBInstance {
	return g.DbInstance
}

func init() {
	GraphQLGroupType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "schemaGroup",
		Description: "An Opsee group target representing an AWS group resource.",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"id": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The AWS group identifier.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if ok {
							return obj.Id, nil
						}
						inter, ok := p.Source.(GroupGetter)
						if ok {
							face := inter.GetGroup()
							if face == nil {
								return nil, nil
							}
							return face.Id, nil
						}
						return nil, fmt.Errorf("field id not resolved")
					},
				},
				"customer_id": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if ok {
							return obj.CustomerId, nil
						}
						inter, ok := p.Source.(GroupGetter)
						if ok {
							face := inter.GetGroup()
							if face == nil {
								return nil, nil
							}
							return face.CustomerId, nil
						}
						return nil, fmt.Errorf("field customer_id not resolved")
					},
				},
				"type": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The type of AWS resource.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if ok {
							return obj.Type, nil
						}
						inter, ok := p.Source.(GroupGetter)
						if ok {
							face := inter.GetGroup()
							if face == nil {
								return nil, nil
							}
							return face.Type, nil
						}
						return nil, fmt.Errorf("field type not resolved")
					},
				},
				"instance_count": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "The last seen number of instances in the group target. This value is cached, so it may not be consistent.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if ok {
							return obj.InstanceCount, nil
						}
						inter, ok := p.Source.(GroupGetter)
						if ok {
							face := inter.GetGroup()
							if face == nil {
								return nil, nil
							}
							return face.InstanceCount, nil
						}
						return nil, fmt.Errorf("field instance_count not resolved")
					},
				},
				"instances": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.NewList(GraphQLInstanceType),
					Description: "The last seen instances in the group target. This value is cached, so it may not be consistent.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if ok {
							return obj.Instances, nil
						}
						inter, ok := p.Source.(GroupGetter)
						if ok {
							face := inter.GetGroup()
							if face == nil {
								return nil, nil
							}
							return face.Instances, nil
						}
						return nil, fmt.Errorf("field instances not resolved")
					},
				},
				"created_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.Timestamp,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if ok {
							if obj.CreatedAt == nil {
								return nil, nil
							}
							return obj.GetCreatedAt(), nil
						}
						inter, ok := p.Source.(GroupGetter)
						if ok {
							face := inter.GetGroup()
							if face == nil {
								return nil, nil
							}
							if face.CreatedAt == nil {
								return nil, nil
							}
							return face.GetCreatedAt(), nil
						}
						return nil, fmt.Errorf("field created_at not resolved")
					},
				},
				"updated_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.Timestamp,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if ok {
							if obj.UpdatedAt == nil {
								return nil, nil
							}
							return obj.GetUpdatedAt(), nil
						}
						inter, ok := p.Source.(GroupGetter)
						if ok {
							face := inter.GetGroup()
							if face == nil {
								return nil, nil
							}
							if face.UpdatedAt == nil {
								return nil, nil
							}
							return face.GetUpdatedAt(), nil
						}
						return nil, fmt.Errorf("field updated_at not resolved")
					},
				},
				"resource": &github_com_graphql_go_graphql.Field{
					Type:        GraphQLGroupResourceUnion,
					Description: "The raw AWS resource data.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Group)
						if !ok {
							return nil, fmt.Errorf("field resource not resolved")
						}
						return obj.GetResource(), nil
					},
				},
			}
		}),
	})
	GraphQLInstanceType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "schemaInstance",
		Description: "An Opsee instance target representing an AWS instance resource.",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"id": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The AWS instance identifier.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Instance)
						if ok {
							return obj.Id, nil
						}
						inter, ok := p.Source.(InstanceGetter)
						if ok {
							face := inter.GetInstance()
							if face == nil {
								return nil, nil
							}
							return face.Id, nil
						}
						return nil, fmt.Errorf("field id not resolved")
					},
				},
				"customer_id": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Instance)
						if ok {
							return obj.CustomerId, nil
						}
						inter, ok := p.Source.(InstanceGetter)
						if ok {
							face := inter.GetInstance()
							if face == nil {
								return nil, nil
							}
							return face.CustomerId, nil
						}
						return nil, fmt.Errorf("field customer_id not resolved")
					},
				},
				"type": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "The type of AWS resource.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Instance)
						if ok {
							return obj.Type, nil
						}
						inter, ok := p.Source.(InstanceGetter)
						if ok {
							face := inter.GetInstance()
							if face == nil {
								return nil, nil
							}
							return face.Type, nil
						}
						return nil, fmt.Errorf("field type not resolved")
					},
				},
				"groups": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.NewList(GraphQLGroupType),
					Description: "The last seen group targets that the instance belongs to. This value is cached, so it may not be consistent.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Instance)
						if ok {
							return obj.Groups, nil
						}
						inter, ok := p.Source.(InstanceGetter)
						if ok {
							face := inter.GetInstance()
							if face == nil {
								return nil, nil
							}
							return face.Groups, nil
						}
						return nil, fmt.Errorf("field groups not resolved")
					},
				},
				"created_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.Timestamp,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Instance)
						if ok {
							if obj.CreatedAt == nil {
								return nil, nil
							}
							return obj.GetCreatedAt(), nil
						}
						inter, ok := p.Source.(InstanceGetter)
						if ok {
							face := inter.GetInstance()
							if face == nil {
								return nil, nil
							}
							if face.CreatedAt == nil {
								return nil, nil
							}
							return face.GetCreatedAt(), nil
						}
						return nil, fmt.Errorf("field created_at not resolved")
					},
				},
				"updated_at": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.Timestamp,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Instance)
						if ok {
							if obj.UpdatedAt == nil {
								return nil, nil
							}
							return obj.GetUpdatedAt(), nil
						}
						inter, ok := p.Source.(InstanceGetter)
						if ok {
							face := inter.GetInstance()
							if face == nil {
								return nil, nil
							}
							if face.UpdatedAt == nil {
								return nil, nil
							}
							return face.GetUpdatedAt(), nil
						}
						return nil, fmt.Errorf("field updated_at not resolved")
					},
				},
				"resource": &github_com_graphql_go_graphql.Field{
					Type:        GraphQLInstanceResourceUnion,
					Description: "The raw AWS resource data.",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Instance)
						if !ok {
							return nil, fmt.Errorf("field resource not resolved")
						}
						return obj.GetResource(), nil
					},
				},
			}
		}),
	})
	GraphQLGroupResourceUnion = github_com_graphql_go_graphql.NewUnion(github_com_graphql_go_graphql.UnionConfig{
		Name:        "GroupResource",
		Description: "The raw AWS resource data.",
		Types: []*github_com_graphql_go_graphql.Object{
			opsee_aws_ec2.GraphQLSecurityGroupType,
			opsee_aws_elb.GraphQLLoadBalancerDescriptionType,
			opsee_aws_autoscaling.GraphQLGroupType,
		},
		ResolveType: func(value interface{}, info github_com_graphql_go_graphql.ResolveInfo) *github_com_graphql_go_graphql.Object {
			switch value.(type) {
			case *Group_SecurityGroup:
				return opsee_aws_ec2.GraphQLSecurityGroupType
			case *Group_LoadBalancer:
				return opsee_aws_elb.GraphQLLoadBalancerDescriptionType
			case *Group_AutoscalingGroup:
				return opsee_aws_autoscaling.GraphQLGroupType
			}
			return nil
		},
	})
	GraphQLInstanceResourceUnion = github_com_graphql_go_graphql.NewUnion(github_com_graphql_go_graphql.UnionConfig{
		Name:        "InstanceResource",
		Description: "The raw AWS resource data.",
		Types: []*github_com_graphql_go_graphql.Object{
			opsee_aws_ec2.GraphQLInstanceType,
			opsee_aws_rds.GraphQLDBInstanceType,
		},
		ResolveType: func(value interface{}, info github_com_graphql_go_graphql.ResolveInfo) *github_com_graphql_go_graphql.Object {
			switch value.(type) {
			case *Instance_Instance:
				return opsee_aws_ec2.GraphQLInstanceType
			case *Instance_DbInstance:
				return opsee_aws_rds.GraphQLDBInstanceType
			}
			return nil
		},
	})
}
func NewPopulatedGroup(r randyAwsData, easy bool) *Group {
	this := &Group{}
	this.Id = randStringAwsData(r)
	this.CustomerId = randStringAwsData(r)
	this.Type = randStringAwsData(r)
	oneofNumber_Resource := []int32{101, 102, 103}[r.Intn(3)]
	switch oneofNumber_Resource {
	case 101:
		this.Resource = NewPopulatedGroup_SecurityGroup(r, easy)
	case 102:
		this.Resource = NewPopulatedGroup_LoadBalancer(r, easy)
	case 103:
		this.Resource = NewPopulatedGroup_AutoscalingGroup(r, easy)
	}
	this.InstanceCount = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.InstanceCount *= -1
	}
	if r.Intn(10) == 0 {
		v1 := r.Intn(5)
		this.Instances = make([]*Instance, v1)
		for i := 0; i < v1; i++ {
			this.Instances[i] = NewPopulatedInstance(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		this.CreatedAt = opsee_types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.UpdatedAt = opsee_types.NewPopulatedTimestamp(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGroup_SecurityGroup(r randyAwsData, easy bool) *Group_SecurityGroup {
	this := &Group_SecurityGroup{}
	this.SecurityGroup = opsee_aws_ec2.NewPopulatedSecurityGroup(r, easy)
	return this
}
func NewPopulatedGroup_LoadBalancer(r randyAwsData, easy bool) *Group_LoadBalancer {
	this := &Group_LoadBalancer{}
	this.LoadBalancer = opsee_aws_elb.NewPopulatedLoadBalancerDescription(r, easy)
	return this
}
func NewPopulatedGroup_AutoscalingGroup(r randyAwsData, easy bool) *Group_AutoscalingGroup {
	this := &Group_AutoscalingGroup{}
	this.AutoscalingGroup = opsee_aws_autoscaling.NewPopulatedGroup(r, easy)
	return this
}
func NewPopulatedInstance(r randyAwsData, easy bool) *Instance {
	this := &Instance{}
	this.Id = randStringAwsData(r)
	this.CustomerId = randStringAwsData(r)
	this.Type = randStringAwsData(r)
	oneofNumber_Resource := []int32{101, 102}[r.Intn(2)]
	switch oneofNumber_Resource {
	case 101:
		this.Resource = NewPopulatedInstance_Instance(r, easy)
	case 102:
		this.Resource = NewPopulatedInstance_DbInstance(r, easy)
	}
	if r.Intn(10) == 0 {
		v2 := r.Intn(5)
		this.Groups = make([]*Group, v2)
		for i := 0; i < v2; i++ {
			this.Groups[i] = NewPopulatedGroup(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		this.CreatedAt = opsee_types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.UpdatedAt = opsee_types.NewPopulatedTimestamp(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedInstance_Instance(r randyAwsData, easy bool) *Instance_Instance {
	this := &Instance_Instance{}
	this.Instance = opsee_aws_ec2.NewPopulatedInstance(r, easy)
	return this
}
func NewPopulatedInstance_DbInstance(r randyAwsData, easy bool) *Instance_DbInstance {
	this := &Instance_DbInstance{}
	this.DbInstance = opsee_aws_rds.NewPopulatedDBInstance(r, easy)
	return this
}

type randyAwsData interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAwsData(r randyAwsData) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAwsData(r randyAwsData) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneAwsData(r)
	}
	return string(tmps)
}
func randUnrecognizedAwsData(r randyAwsData, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldAwsData(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldAwsData(data []byte, r randyAwsData, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateAwsData(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateAwsData(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateAwsData(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateAwsData(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateAwsData(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateAwsData(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateAwsData(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
