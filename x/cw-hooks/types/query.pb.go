// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: juno/cwhooks/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParams is the request type to get all module params.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b0c5bc2d2dc51, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryClockContractsResponse is the response type for the Query/ClockContracts RPC method.
type QueryParamsResponse struct {
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b0c5bc2d2dc51, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

// QueryStakingContractsRequest
type QueryStakingContractsRequest struct {
}

func (m *QueryStakingContractsRequest) Reset()         { *m = QueryStakingContractsRequest{} }
func (m *QueryStakingContractsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStakingContractsRequest) ProtoMessage()    {}
func (*QueryStakingContractsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b0c5bc2d2dc51, []int{2}
}
func (m *QueryStakingContractsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStakingContractsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStakingContractsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStakingContractsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStakingContractsRequest.Merge(m, src)
}
func (m *QueryStakingContractsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStakingContractsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStakingContractsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStakingContractsRequest proto.InternalMessageInfo

// QueryStakingContractsResponse
type QueryStakingContractsResponse struct {
	Contracts []string `protobuf:"bytes,1,rep,name=contracts,proto3" json:"contracts" yaml:"contracts"`
}

func (m *QueryStakingContractsResponse) Reset()         { *m = QueryStakingContractsResponse{} }
func (m *QueryStakingContractsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStakingContractsResponse) ProtoMessage()    {}
func (*QueryStakingContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b0c5bc2d2dc51, []int{3}
}
func (m *QueryStakingContractsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStakingContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStakingContractsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStakingContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStakingContractsResponse.Merge(m, src)
}
func (m *QueryStakingContractsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStakingContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStakingContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStakingContractsResponse proto.InternalMessageInfo

func (m *QueryStakingContractsResponse) GetContracts() []string {
	if m != nil {
		return m.Contracts
	}
	return nil
}

// QueryStakingContractsRequest
type QueryGovernanceContractsRequest struct {
}

func (m *QueryGovernanceContractsRequest) Reset()         { *m = QueryGovernanceContractsRequest{} }
func (m *QueryGovernanceContractsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGovernanceContractsRequest) ProtoMessage()    {}
func (*QueryGovernanceContractsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b0c5bc2d2dc51, []int{4}
}
func (m *QueryGovernanceContractsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGovernanceContractsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGovernanceContractsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGovernanceContractsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGovernanceContractsRequest.Merge(m, src)
}
func (m *QueryGovernanceContractsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGovernanceContractsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGovernanceContractsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGovernanceContractsRequest proto.InternalMessageInfo

// QueryGovernanceContractsResponse
type QueryGovernanceContractsResponse struct {
	Contracts []string `protobuf:"bytes,1,rep,name=contracts,proto3" json:"contracts" yaml:"contracts"`
}

func (m *QueryGovernanceContractsResponse) Reset()         { *m = QueryGovernanceContractsResponse{} }
func (m *QueryGovernanceContractsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGovernanceContractsResponse) ProtoMessage()    {}
func (*QueryGovernanceContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b0c5bc2d2dc51, []int{5}
}
func (m *QueryGovernanceContractsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGovernanceContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGovernanceContractsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGovernanceContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGovernanceContractsResponse.Merge(m, src)
}
func (m *QueryGovernanceContractsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGovernanceContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGovernanceContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGovernanceContractsResponse proto.InternalMessageInfo

func (m *QueryGovernanceContractsResponse) GetContracts() []string {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "juno.cwhooks.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "juno.cwhooks.v1.QueryParamsResponse")
	proto.RegisterType((*QueryStakingContractsRequest)(nil), "juno.cwhooks.v1.QueryStakingContractsRequest")
	proto.RegisterType((*QueryStakingContractsResponse)(nil), "juno.cwhooks.v1.QueryStakingContractsResponse")
	proto.RegisterType((*QueryGovernanceContractsRequest)(nil), "juno.cwhooks.v1.QueryGovernanceContractsRequest")
	proto.RegisterType((*QueryGovernanceContractsResponse)(nil), "juno.cwhooks.v1.QueryGovernanceContractsResponse")
}

func init() { proto.RegisterFile("juno/cwhooks/v1/query.proto", fileDescriptor_c08b0c5bc2d2dc51) }

var fileDescriptor_c08b0c5bc2d2dc51 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0x26, 0x2a, 0xcd, 0x08, 0x31, 0x79, 0x93, 0x36, 0xb2, 0x2d, 0xe9, 0xcc, 0x10,
	0x13, 0x52, 0xe3, 0x65, 0xdc, 0xb8, 0x20, 0x75, 0x07, 0x24, 0xc4, 0x01, 0xca, 0x8d, 0x0b, 0xb8,
	0x96, 0xe5, 0x85, 0xad, 0x7e, 0x59, 0xec, 0x76, 0xf4, 0xca, 0x27, 0x40, 0xe2, 0xcc, 0x07, 0xe0,
	0x9b, 0xc0, 0x6d, 0x12, 0x17, 0x4e, 0x15, 0x6a, 0x39, 0xed, 0xd8, 0x4f, 0x80, 0x6a, 0x9b, 0x54,
	0x34, 0x2d, 0x02, 0x89, 0x5b, 0xf2, 0xfe, 0xef, 0xf9, 0xff, 0xcb, 0xfb, 0xc7, 0x68, 0xfb, 0x4d,
	0x4f, 0x01, 0xe5, 0x17, 0x27, 0x00, 0xa7, 0x9a, 0xf6, 0x53, 0x7a, 0xde, 0x13, 0xc5, 0x20, 0xc9,
	0x0b, 0x30, 0x80, 0x6f, 0x4d, 0xc5, 0xc4, 0x8b, 0x49, 0x3f, 0x0d, 0x37, 0x24, 0x48, 0xb0, 0x1a,
	0x9d, 0x3e, 0xb9, 0xb6, 0x70, 0x47, 0x02, 0xc8, 0x33, 0x41, 0x59, 0x9e, 0x51, 0xa6, 0x14, 0x18,
	0x66, 0x32, 0x50, 0xda, 0xab, 0x11, 0x07, 0xdd, 0x05, 0x4d, 0x3b, 0x4c, 0x0b, 0xda, 0x4f, 0x3b,
	0xc2, 0xb0, 0x94, 0x72, 0xc8, 0x94, 0xd7, 0x77, 0xe7, 0x09, 0xa4, 0x50, 0x42, 0x67, 0x7e, 0x9c,
	0x6c, 0x20, 0xfc, 0x7c, 0x8a, 0xf4, 0x8c, 0x15, 0xac, 0xab, 0xdb, 0xe2, 0xbc, 0x27, 0xb4, 0x21,
	0x1c, 0xad, 0xff, 0x56, 0xd5, 0x39, 0x28, 0x2d, 0xf0, 0x53, 0x54, 0xcf, 0x6d, 0x65, 0x2b, 0x68,
	0x04, 0x07, 0x37, 0x8e, 0x36, 0x93, 0xb9, 0x2f, 0x48, 0xdc, 0x40, 0x6b, 0xfb, 0x6a, 0x18, 0xfb,
	0xd6, 0xc9, 0x30, 0xbe, 0x39, 0x60, 0xdd, 0xb3, 0x87, 0xc4, 0xbd, 0x93, 0xb6, 0x17, 0x48, 0x84,
	0x76, 0xac, 0xc9, 0x0b, 0xc3, 0x4e, 0x33, 0x25, 0x8f, 0x41, 0x99, 0x82, 0x71, 0x53, 0x42, 0xbc,
	0x46, 0xbb, 0x4b, 0x74, 0x8f, 0xf3, 0x08, 0xad, 0xf2, 0x5f, 0xc5, 0xad, 0xa0, 0xb1, 0x72, 0xb0,
	0xda, 0xda, 0xbb, 0x1a, 0xc6, 0xb3, 0xe2, 0x64, 0x18, 0xaf, 0x39, 0xef, 0xb2, 0x44, 0xda, 0x33,
	0x99, 0xec, 0xa1, 0xd8, 0x3a, 0x3c, 0x86, 0xbe, 0x28, 0x14, 0x53, 0x5c, 0x54, 0x20, 0x38, 0x6a,
	0x2c, 0x6f, 0xf9, 0x4f, 0x1c, 0x47, 0x5f, 0x56, 0xd0, 0x75, 0xeb, 0x82, 0x0d, 0xaa, 0xbb, 0x15,
	0xe2, 0x3b, 0x95, 0xdd, 0x56, 0x73, 0x0a, 0xf7, 0xff, 0xdc, 0xe4, 0xf8, 0x48, 0xfc, 0xee, 0xeb,
	0x8f, 0x0f, 0xd7, 0x6e, 0xe3, 0x4d, 0x3a, 0xff, 0x2f, 0xb8, 0x24, 0xf0, 0xc7, 0x00, 0xad, 0xcd,
	0x6f, 0x19, 0x37, 0x17, 0x9f, 0xbd, 0x24, 0xad, 0x30, 0xf9, 0xdb, 0x76, 0x0f, 0x75, 0xdf, 0x42,
	0xed, 0x63, 0x52, 0x81, 0xd2, 0x6e, 0xe4, 0x55, 0xb9, 0x1f, 0xfc, 0x29, 0x40, 0xeb, 0x0b, 0x02,
	0xc0, 0x87, 0x8b, 0x3d, 0x97, 0xc7, 0x19, 0xa6, 0xff, 0x30, 0xe1, 0x41, 0x9b, 0x16, 0xf4, 0x1e,
	0xbe, 0x5b, 0x01, 0x95, 0xe5, 0xd4, 0x8c, 0xb5, 0xf5, 0xe4, 0xf3, 0x28, 0x0a, 0x2e, 0x47, 0x51,
	0xf0, 0x7d, 0x14, 0x05, 0xef, 0xc7, 0x51, 0xed, 0x72, 0x1c, 0xd5, 0xbe, 0x8d, 0xa3, 0xda, 0xcb,
	0x43, 0x99, 0x99, 0x93, 0x5e, 0x27, 0xe1, 0xd0, 0xa5, 0xc7, 0xf6, 0xd2, 0x96, 0x66, 0xee, 0xe8,
	0xb7, 0x94, 0x5f, 0x34, 0xdd, 0xe9, 0x66, 0x90, 0x0b, 0xdd, 0xa9, 0xdb, 0x3b, 0xfa, 0xe0, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x28, 0x8e, 0xa2, 0x46, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// StakingContracts
	StakingContracts(ctx context.Context, in *QueryStakingContractsRequest, opts ...grpc.CallOption) (*QueryStakingContractsResponse, error)
	// GovernanceContracts
	GovernanceContracts(ctx context.Context, in *QueryGovernanceContractsRequest, opts ...grpc.CallOption) (*QueryGovernanceContractsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/juno.cwhooks.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StakingContracts(ctx context.Context, in *QueryStakingContractsRequest, opts ...grpc.CallOption) (*QueryStakingContractsResponse, error) {
	out := new(QueryStakingContractsResponse)
	err := c.cc.Invoke(ctx, "/juno.cwhooks.v1.Query/StakingContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GovernanceContracts(ctx context.Context, in *QueryGovernanceContractsRequest, opts ...grpc.CallOption) (*QueryGovernanceContractsResponse, error) {
	out := new(QueryGovernanceContractsResponse)
	err := c.cc.Invoke(ctx, "/juno.cwhooks.v1.Query/GovernanceContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// StakingContracts
	StakingContracts(context.Context, *QueryStakingContractsRequest) (*QueryStakingContractsResponse, error)
	// GovernanceContracts
	GovernanceContracts(context.Context, *QueryGovernanceContractsRequest) (*QueryGovernanceContractsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) StakingContracts(ctx context.Context, req *QueryStakingContractsRequest) (*QueryStakingContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakingContracts not implemented")
}
func (*UnimplementedQueryServer) GovernanceContracts(ctx context.Context, req *QueryGovernanceContractsRequest) (*QueryGovernanceContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovernanceContracts not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juno.cwhooks.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StakingContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStakingContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StakingContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juno.cwhooks.v1.Query/StakingContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StakingContracts(ctx, req.(*QueryStakingContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GovernanceContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGovernanceContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GovernanceContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juno.cwhooks.v1.Query/GovernanceContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GovernanceContracts(ctx, req.(*QueryGovernanceContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "juno.cwhooks.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "StakingContracts",
			Handler:    _Query_StakingContracts_Handler,
		},
		{
			MethodName: "GovernanceContracts",
			Handler:    _Query_GovernanceContracts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "juno/cwhooks/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStakingContractsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStakingContractsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStakingContractsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryStakingContractsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStakingContractsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStakingContractsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contracts[iNdEx])
			copy(dAtA[i:], m.Contracts[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Contracts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGovernanceContractsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGovernanceContractsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGovernanceContractsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGovernanceContractsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGovernanceContractsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGovernanceContractsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contracts[iNdEx])
			copy(dAtA[i:], m.Contracts[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Contracts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStakingContractsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryStakingContractsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Contracts) > 0 {
		for _, s := range m.Contracts {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryGovernanceContractsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGovernanceContractsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Contracts) > 0 {
		for _, s := range m.Contracts {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryStakingContractsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStakingContractsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStakingContractsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryStakingContractsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStakingContractsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStakingContractsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGovernanceContractsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGovernanceContractsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGovernanceContractsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGovernanceContractsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGovernanceContractsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGovernanceContractsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
