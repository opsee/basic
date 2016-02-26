// Code generated by protoc-gen-gogo.
// source: vape.proto
// DO NOT EDIT!

package service

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/opsee/protobuf/opseeproto"
import _ "github.com/opsee/protobuf/opseeproto/types"
import opsee1 "github.com/opsee/basic/schema"

import github_com_graphql_go_graphql "github.com/graphql-go/graphql"
import github_com_opsee_protobuf_plugin_graphql_scalars "github.com/opsee/protobuf/plugin/graphql/scalars"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetUserRequest struct {
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,proto3" json:"customer_id,omitempty"`
	Id         int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}

type GetUserResponse struct {
	User       *opsee1.User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	BasicToken string       `protobuf:"bytes,2,opt,name=basic_token,proto3" json:"basic_token,omitempty"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}

func (m *GetUserResponse) GetUser() *opsee1.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListUsersRequest struct {
	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,proto3" json:"per_page,omitempty"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}

type ListUsersResponse struct {
	Users   []*opsee1.User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	Page    int32          `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32          `protobuf:"varint,3,opt,name=per_page,proto3" json:"per_page,omitempty"`
	Total   int32          `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}

func (m *ListUsersResponse) GetUsers() []*opsee1.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "opsee.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "opsee.GetUserResponse")
	proto.RegisterType((*ListUsersRequest)(nil), "opsee.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "opsee.ListUsersResponse")
}
func (this *GetUserRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetUserRequest)
	if !ok {
		that2, ok := that.(GetUserRequest)
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
	if this.CustomerId != that1.CustomerId {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	return true
}
func (this *GetUserResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetUserResponse)
	if !ok {
		that2, ok := that.(GetUserResponse)
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
	if this.BasicToken != that1.BasicToken {
		return false
	}
	return true
}
func (this *ListUsersRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ListUsersRequest)
	if !ok {
		that2, ok := that.(ListUsersRequest)
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
	if this.Page != that1.Page {
		return false
	}
	if this.PerPage != that1.PerPage {
		return false
	}
	return true
}
func (this *ListUsersResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ListUsersResponse)
	if !ok {
		that2, ok := that.(ListUsersResponse)
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
	if len(this.Users) != len(that1.Users) {
		return false
	}
	for i := range this.Users {
		if !this.Users[i].Equal(that1.Users[i]) {
			return false
		}
	}
	if this.Page != that1.Page {
		return false
	}
	if this.PerPage != that1.PerPage {
		return false
	}
	if this.Total != that1.Total {
		return false
	}
	return true
}

type GetUserRequestGetter interface {
	GetGetUserRequest() *GetUserRequest
}

var GraphQLGetUserRequestType *github_com_graphql_go_graphql.Object

type GetUserResponseGetter interface {
	GetGetUserResponse() *GetUserResponse
}

var GraphQLGetUserResponseType *github_com_graphql_go_graphql.Object

type ListUsersRequestGetter interface {
	GetListUsersRequest() *ListUsersRequest
}

var GraphQLListUsersRequestType *github_com_graphql_go_graphql.Object

type ListUsersResponseGetter interface {
	GetListUsersResponse() *ListUsersResponse
}

var GraphQLListUsersResponseType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLGetUserRequestType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "serviceGetUserRequest",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"customer_id": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetUserRequest)
						if ok {
							return obj.CustomerId, nil
						}
						inter, ok := p.Source.(GetUserRequestGetter)
						if ok {
							face := inter.GetGetUserRequest()
							if face == nil {
								return nil, nil
							}
							return face.CustomerId, nil
						}
						return nil, fmt.Errorf("field customer_id not resolved")
					},
				},
				"id": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetUserRequest)
						if ok {
							return obj.Id, nil
						}
						inter, ok := p.Source.(GetUserRequestGetter)
						if ok {
							face := inter.GetGetUserRequest()
							if face == nil {
								return nil, nil
							}
							return face.Id, nil
						}
						return nil, fmt.Errorf("field id not resolved")
					},
				},
				"email": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetUserRequest)
						if ok {
							return obj.Email, nil
						}
						inter, ok := p.Source.(GetUserRequestGetter)
						if ok {
							face := inter.GetGetUserRequest()
							if face == nil {
								return nil, nil
							}
							return face.Email, nil
						}
						return nil, fmt.Errorf("field email not resolved")
					},
				},
			}
		}),
	})
	GraphQLGetUserResponseType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "serviceGetUserResponse",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"user": &github_com_graphql_go_graphql.Field{
					Type:        github_com_opsee_protobuf_plugin_graphql_scalars.ByteString,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetUserResponse)
						if ok {
							if obj.User == nil {
								return nil, nil
							}
							return obj.GetUser(), nil
						}
						inter, ok := p.Source.(GetUserResponseGetter)
						if ok {
							face := inter.GetGetUserResponse()
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
				"basic_token": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*GetUserResponse)
						if ok {
							return obj.BasicToken, nil
						}
						inter, ok := p.Source.(GetUserResponseGetter)
						if ok {
							face := inter.GetGetUserResponse()
							if face == nil {
								return nil, nil
							}
							return face.BasicToken, nil
						}
						return nil, fmt.Errorf("field basic_token not resolved")
					},
				},
			}
		}),
	})
	GraphQLListUsersRequestType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "serviceListUsersRequest",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"page": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*ListUsersRequest)
						if ok {
							return obj.Page, nil
						}
						inter, ok := p.Source.(ListUsersRequestGetter)
						if ok {
							face := inter.GetListUsersRequest()
							if face == nil {
								return nil, nil
							}
							return face.Page, nil
						}
						return nil, fmt.Errorf("field page not resolved")
					},
				},
				"per_page": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*ListUsersRequest)
						if ok {
							return obj.PerPage, nil
						}
						inter, ok := p.Source.(ListUsersRequestGetter)
						if ok {
							face := inter.GetListUsersRequest()
							if face == nil {
								return nil, nil
							}
							return face.PerPage, nil
						}
						return nil, fmt.Errorf("field per_page not resolved")
					},
				},
			}
		}),
	})
	GraphQLListUsersResponseType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "serviceListUsersResponse",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"users": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.NewList(github_com_opsee_protobuf_plugin_graphql_scalars.ByteString),
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*ListUsersResponse)
						if ok {
							return obj.Users, nil
						}
						inter, ok := p.Source.(ListUsersResponseGetter)
						if ok {
							face := inter.GetListUsersResponse()
							if face == nil {
								return nil, nil
							}
							return face.Users, nil
						}
						return nil, fmt.Errorf("field users not resolved")
					},
				},
				"page": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*ListUsersResponse)
						if ok {
							return obj.Page, nil
						}
						inter, ok := p.Source.(ListUsersResponseGetter)
						if ok {
							face := inter.GetListUsersResponse()
							if face == nil {
								return nil, nil
							}
							return face.Page, nil
						}
						return nil, fmt.Errorf("field page not resolved")
					},
				},
				"per_page": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*ListUsersResponse)
						if ok {
							return obj.PerPage, nil
						}
						inter, ok := p.Source.(ListUsersResponseGetter)
						if ok {
							face := inter.GetListUsersResponse()
							if face == nil {
								return nil, nil
							}
							return face.PerPage, nil
						}
						return nil, fmt.Errorf("field per_page not resolved")
					},
				},
				"total": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*ListUsersResponse)
						if ok {
							return obj.Total, nil
						}
						inter, ok := p.Source.(ListUsersResponseGetter)
						if ok {
							face := inter.GetListUsersResponse()
							if face == nil {
								return nil, nil
							}
							return face.Total, nil
						}
						return nil, fmt.Errorf("field total not resolved")
					},
				},
			}
		}),
	})
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Vape service

type VapeClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type vapeClient struct {
	cc *grpc.ClientConn
}

func NewVapeClient(cc *grpc.ClientConn) VapeClient {
	return &vapeClient{cc}
}

func (c *vapeClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := grpc.Invoke(ctx, "/opsee.Vape/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vapeClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := grpc.Invoke(ctx, "/opsee.Vape/ListUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Vape service

type VapeServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
}

func RegisterVapeServer(s *grpc.Server, srv VapeServer) {
	s.RegisterService(&_Vape_serviceDesc, srv)
}

func _Vape_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(VapeServer).GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Vape_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(VapeServer).ListUsers(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Vape_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opsee.Vape",
	HandlerType: (*VapeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Vape_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Vape_ListUsers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func NewPopulatedGetUserRequest(r randyVape, easy bool) *GetUserRequest {
	this := &GetUserRequest{}
	this.CustomerId = randStringVape(r)
	this.Id = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Id *= -1
	}
	this.Email = randStringVape(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGetUserResponse(r randyVape, easy bool) *GetUserResponse {
	this := &GetUserResponse{}
	if r.Intn(10) != 0 {
		this.User = opsee1.NewPopulatedUser(r, easy)
	}
	this.BasicToken = randStringVape(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListUsersRequest(r randyVape, easy bool) *ListUsersRequest {
	this := &ListUsersRequest{}
	this.Page = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Page *= -1
	}
	this.PerPage = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.PerPage *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListUsersResponse(r randyVape, easy bool) *ListUsersResponse {
	this := &ListUsersResponse{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Users = make([]*opsee1.User, v1)
		for i := 0; i < v1; i++ {
			this.Users[i] = opsee1.NewPopulatedUser(r, easy)
		}
	}
	this.Page = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Page *= -1
	}
	this.PerPage = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.PerPage *= -1
	}
	this.Total = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Total *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyVape interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneVape(r randyVape) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringVape(r randyVape) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneVape(r)
	}
	return string(tmps)
}
func randUnrecognizedVape(r randyVape, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldVape(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldVape(data []byte, r randyVape, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateVape(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateVape(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateVape(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateVape(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateVape(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateVape(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateVape(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
