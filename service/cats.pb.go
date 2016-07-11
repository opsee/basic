// Code generated by protoc-gen-gogo.
// source: cats.proto
// DO NOT EDIT!

package service

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/opsee/protobuf/opseeproto"
import _ "github.com/opsee/protobuf/opseeproto/types"
import _ "github.com/opsee/basic/schema/aws/credentials"
import opsee1 "github.com/opsee/basic/schema"

import github_com_graphql_go_graphql "github.com/graphql-go/graphql"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetCheckCountRequest struct {
	User     *opsee1.User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Prorated bool         `protobuf:"varint,2,opt,name=prorated,proto3" json:"prorated,omitempty"`
}

func (m *GetCheckCountRequest) Reset()                    { *m = GetCheckCountRequest{} }
func (m *GetCheckCountRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCheckCountRequest) ProtoMessage()               {}
func (*GetCheckCountRequest) Descriptor() ([]byte, []int) { return fileDescriptorCats, []int{0} }

func (m *GetCheckCountRequest) GetUser() *opsee1.User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetCheckCountResponse struct {
	Count    float32 `protobuf:"fixed32,1,opt,name=count,proto3" json:"count,omitempty"`
	Prorated bool    `protobuf:"varint,2,opt,name=prorated,proto3" json:"prorated,omitempty"`
}

func (m *GetCheckCountResponse) Reset()                    { *m = GetCheckCountResponse{} }
func (m *GetCheckCountResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCheckCountResponse) ProtoMessage()               {}
func (*GetCheckCountResponse) Descriptor() ([]byte, []int) { return fileDescriptorCats, []int{1} }

func init() {
	proto.RegisterType((*GetCheckCountRequest)(nil), "opsee.GetCheckCountRequest")
	proto.RegisterType((*GetCheckCountResponse)(nil), "opsee.GetCheckCountResponse")
}
func (this *GetCheckCountRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetCheckCountRequest)
	if !ok {
		that2, ok := that.(GetCheckCountRequest)
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
	if !this.User.Equal(that1.User) {
		return false
	}
	if this.Prorated != that1.Prorated {
		return false
	}
	return true
}
func (this *GetCheckCountResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetCheckCountResponse)
	if !ok {
		that2, ok := that.(GetCheckCountResponse)
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
	if this.Count != that1.Count {
		return false
	}
	if this.Prorated != that1.Prorated {
		return false
	}
	return true
}

type GetCheckCountRequestGetter interface {
	GetGetCheckCountRequest() *GetCheckCountRequest
}

var GraphQLGetCheckCountRequestType *github_com_graphql_go_graphql.Object

type GetCheckCountResponseGetter interface {
	GetGetCheckCountResponse() *GetCheckCountResponse
}

var GraphQLGetCheckCountResponseType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLGetCheckCountRequestType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "serviceGetCheckCountRequest",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"user": &github_com_graphql_go_graphql.Field{
					Type:        opsee1.GraphQLUserType,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetCheckCountRequest)
						if ok {
							if obj.User == nil {
								return nil, nil
							}
							return obj.GetUser(), nil
						}
						inter, ok := p.Source.(GetCheckCountRequestGetter)
						if ok {
							face := inter.GetGetCheckCountRequest()
							if face == nil {
								return nil, nil
							}
							if face.User == nil {
								return nil, nil
							}
							return face.GetUser(), nil
						}
						return nil, fmt.Errorf("field user not resolved")
					},
				},
				"prorated": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetCheckCountRequest)
						if ok {
							return obj.Prorated, nil
						}
						inter, ok := p.Source.(GetCheckCountRequestGetter)
						if ok {
							face := inter.GetGetCheckCountRequest()
							if face == nil {
								return nil, nil
							}
							return face.Prorated, nil
						}
						return nil, fmt.Errorf("field prorated not resolved")
					},
				},
			}
		}),
	})
	GraphQLGetCheckCountResponseType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "serviceGetCheckCountResponse",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"count": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Float,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetCheckCountResponse)
						if ok {
							return obj.Count, nil
						}
						inter, ok := p.Source.(GetCheckCountResponseGetter)
						if ok {
							face := inter.GetGetCheckCountResponse()
							if face == nil {
								return nil, nil
							}
							return face.Count, nil
						}
						return nil, fmt.Errorf("field count not resolved")
					},
				},
				"prorated": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetCheckCountResponse)
						if ok {
							return obj.Prorated, nil
						}
						inter, ok := p.Source.(GetCheckCountResponseGetter)
						if ok {
							face := inter.GetGetCheckCountResponse()
							if face == nil {
								return nil, nil
							}
							return face.Prorated, nil
						}
						return nil, fmt.Errorf("field prorated not resolved")
					},
				},
			}
		}),
	})
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Cats service

type CatsClient interface {
	GetCheckCount(ctx context.Context, in *GetCheckCountRequest, opts ...grpc.CallOption) (*GetCheckCountResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserTokenResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error)
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error)
}

type catsClient struct {
	cc *grpc.ClientConn
}

func NewCatsClient(cc *grpc.ClientConn) CatsClient {
	return &catsClient{cc}
}

func (c *catsClient) GetCheckCount(ctx context.Context, in *GetCheckCountRequest, opts ...grpc.CallOption) (*GetCheckCountResponse, error) {
	out := new(GetCheckCountResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/GetCheckCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserTokenResponse, error) {
	out := new(UserTokenResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/ListUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error) {
	out := new(InviteUserResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/InviteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error) {
	out := new(GetTeamResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/GetTeam", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error) {
	out := new(UpdateTeamResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/UpdateTeam", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catsClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error) {
	out := new(DeleteTeamResponse)
	err := grpc.Invoke(ctx, "/opsee.Cats/DeleteTeam", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cats service

type CatsServer interface {
	GetCheckCount(context.Context, *GetCheckCountRequest) (*GetCheckCountResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserTokenResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamResponse, error)
	UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error)
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error)
}

func RegisterCatsServer(s *grpc.Server, srv CatsServer) {
	s.RegisterService(&_Cats_serviceDesc, srv)
}

func _Cats_GetCheckCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).GetCheckCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/GetCheckCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).GetCheckCount(ctx, req.(*GetCheckCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_InviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).InviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/InviteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).InviteUser(ctx, req.(*InviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).GetTeam(ctx, req.(*GetTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).UpdateTeam(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cats_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatsServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opsee.Cats/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatsServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opsee.Cats",
	HandlerType: (*CatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCheckCount",
			Handler:    _Cats_GetCheckCount_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Cats_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Cats_UpdateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Cats_ListUsers_Handler,
		},
		{
			MethodName: "InviteUser",
			Handler:    _Cats_InviteUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Cats_DeleteUser_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _Cats_GetTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Cats_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Cats_DeleteTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorCats,
}

func (m *GetCheckCountRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetCheckCountRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.User != nil {
		data[i] = 0xa
		i++
		i = encodeVarintCats(data, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Prorated {
		data[i] = 0x10
		i++
		if m.Prorated {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *GetCheckCountResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetCheckCountResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		data[i] = 0xd
		i++
		i = encodeFixed32Cats(data, i, uint32(math.Float32bits(float32(m.Count))))
	}
	if m.Prorated {
		data[i] = 0x10
		i++
		if m.Prorated {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Cats(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Cats(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCats(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGetCheckCountRequest(r randyCats, easy bool) *GetCheckCountRequest {
	this := &GetCheckCountRequest{}
	if r.Intn(10) != 0 {
		this.User = opsee1.NewPopulatedUser(r, easy)
	}
	this.Prorated = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGetCheckCountResponse(r randyCats, easy bool) *GetCheckCountResponse {
	this := &GetCheckCountResponse{}
	this.Count = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Count *= -1
	}
	this.Prorated = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyCats interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneCats(r randyCats) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringCats(r randyCats) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneCats(r)
	}
	return string(tmps)
}
func randUnrecognizedCats(r randyCats, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldCats(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldCats(data []byte, r randyCats, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateCats(data, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		data = encodeVarintPopulateCats(data, uint64(v2))
	case 1:
		data = encodeVarintPopulateCats(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateCats(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateCats(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateCats(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateCats(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *GetCheckCountRequest) Size() (n int) {
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovCats(uint64(l))
	}
	if m.Prorated {
		n += 2
	}
	return n
}

func (m *GetCheckCountResponse) Size() (n int) {
	var l int
	_ = l
	if m.Count != 0 {
		n += 5
	}
	if m.Prorated {
		n += 2
	}
	return n
}

func sovCats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCats(x uint64) (n int) {
	return sovCats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetCheckCountRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCheckCountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCheckCountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &opsee1.User{}
			}
			if err := m.User.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prorated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Prorated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCats(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetCheckCountResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCheckCountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCheckCountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Count = float32(math.Float32frombits(v))
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prorated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Prorated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCats(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCats(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCats
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCats
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCats(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCats   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorCats = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0xf1, 0x29, 0xc7, 0x1d, 0x73, 0xa2, 0x59, 0xdd, 0x81, 0xcf, 0xa0, 0x04, 0x45, 0x42,
	0x4a, 0x15, 0xa3, 0x50, 0x91, 0x0a, 0xc5, 0x48, 0x28, 0x52, 0x2a, 0x93, 0x34, 0x74, 0xeb, 0xcd,
	0x40, 0xac, 0xc4, 0x7f, 0xf0, 0xae, 0x83, 0x78, 0x1d, 0x2a, 0x1e, 0x81, 0x92, 0x92, 0x92, 0x47,
	0x00, 0x0a, 0x9e, 0x81, 0x92, 0xf5, 0xae, 0x63, 0xaf, 0x1d, 0x2b, 0xa2, 0x58, 0x29, 0xf3, 0xcd,
	0x7c, 0x3f, 0x7d, 0x33, 0x91, 0x01, 0x18, 0x15, 0x7c, 0x9c, 0x66, 0x89, 0x48, 0xc8, 0x79, 0x92,
	0x72, 0x44, 0xe7, 0xd9, 0xfb, 0x50, 0x6c, 0xf2, 0x60, 0xcc, 0x92, 0xc8, 0x55, 0x8a, 0xab, 0xda,
	0x41, 0xfe, 0x4e, 0x97, 0xaa, 0xd2, 0x3f, 0xb5, 0xd1, 0x99, 0xfe, 0x97, 0x43, 0x7c, 0x4a, 0x91,
	0xbb, 0x22, 0x8c, 0x90, 0x0b, 0x1a, 0xa5, 0xa5, 0xf7, 0xc5, 0x91, 0x37, 0xa0, 0x3c, 0x64, 0x2e,
	0x67, 0x1b, 0x8c, 0xa8, 0x4b, 0x3f, 0x72, 0x97, 0x65, 0xb8, 0xc6, 0x58, 0x84, 0x74, 0xc7, 0x35,
	0xa4, 0xb4, 0x8e, 0x4e, 0x5b, 0x73, 0x8e, 0x59, 0x39, 0x09, 0x7b, 0x9a, 0x96, 0x61, 0x87, 0x6f,
	0xe0, 0xfa, 0x35, 0x0a, 0x6f, 0x83, 0x6c, 0xeb, 0x25, 0x79, 0x2c, 0x7c, 0xfc, 0x90, 0xcb, 0x44,
	0x64, 0x00, 0xbd, 0xc2, 0x61, 0x5b, 0x4f, 0xac, 0xd1, 0xd5, 0xe4, 0x6a, 0xac, 0x17, 0x5c, 0x49,
	0xc9, 0x57, 0x0d, 0xe2, 0xc0, 0xa5, 0x24, 0x64, 0x54, 0xe0, 0xda, 0x3e, 0x93, 0x43, 0x97, 0x7e,
	0x55, 0x0f, 0xe7, 0x70, 0xd3, 0x82, 0xf2, 0x34, 0x89, 0x39, 0x92, 0x6b, 0x38, 0x67, 0x85, 0xa0,
	0xb0, 0x67, 0xbe, 0x2e, 0x4e, 0xa1, 0x26, 0x7f, 0x7a, 0xd0, 0xf3, 0xe4, 0x9f, 0x42, 0x16, 0x70,
	0xbf, 0xc1, 0x24, 0x8f, 0xca, 0x4c, 0x5d, 0xf1, 0x9d, 0xc7, 0xdd, 0x4d, 0x1d, 0x63, 0x78, 0x87,
	0x4c, 0xe1, 0x42, 0xb6, 0x8a, 0x75, 0xc8, 0x4d, 0x3d, 0xaa, 0xd6, 0x2b, 0x09, 0x0f, 0xda, 0x72,
	0xe5, 0x9d, 0x01, 0xac, 0xd2, 0xb5, 0x4c, 0xa7, 0xec, 0xf6, 0xe1, 0x34, 0x95, 0x74, 0x20, 0xd8,
	0xc6, 0xd1, 0x96, 0xc9, 0x16, 0x63, 0x83, 0xf1, 0x12, 0xee, 0x2d, 0x42, 0xae, 0xc8, 0x9c, 0x3c,
	0x2c, 0x07, 0x2b, 0xa5, 0x4d, 0x30, 0x1a, 0x15, 0xc1, 0x03, 0x98, 0xc7, 0xfb, 0xb0, 0x95, 0xa2,
	0x96, 0x0e, 0x8c, 0xdb, 0x8e, 0x8e, 0x09, 0x79, 0x85, 0x3b, 0x6c, 0x41, 0x6a, 0xa9, 0x0d, 0x31,
	0x3b, 0xad, 0x5b, 0x2e, 0x91, 0x46, 0xe6, 0x2d, 0x8b, 0xba, 0xe3, 0x96, 0x5a, 0x36, 0x03, 0xe8,
	0xc3, 0x29, 0x7b, 0xf3, 0x96, 0x26, 0xe1, 0xb6, 0xa3, 0x73, 0xbc, 0x45, 0x03, 0x52, 0x4b, 0xdd,
	0x5b, 0x34, 0x21, 0xb3, 0xa7, 0x7f, 0x7f, 0xf5, 0xad, 0x2f, 0xbf, 0xfb, 0xd6, 0x57, 0xf9, 0xbe,
	0xcb, 0xf7, 0x43, 0xbe, 0x9f, 0xf2, 0x7d, 0xfb, 0x3c, 0xb0, 0xde, 0x5e, 0xc8, 0x95, 0xf7, 0x21,
	0xc3, 0xe0, 0xae, 0xfa, 0x6c, 0x9e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x2a, 0xd9, 0xe3,
	0x2a, 0x04, 0x00, 0x00,
}
