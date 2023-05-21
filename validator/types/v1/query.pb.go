// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chora/validator/v1/query.proto

package v1

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryValidatorRequest is the Query/Validator request type.
type QueryValidatorRequest struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryValidatorRequest) Reset()         { *m = QueryValidatorRequest{} }
func (m *QueryValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorRequest) ProtoMessage()    {}
func (*QueryValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7398e53758cca59c, []int{0}
}
func (m *QueryValidatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorRequest.Merge(m, src)
}
func (m *QueryValidatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorRequest proto.InternalMessageInfo

func (m *QueryValidatorRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// QueryValidatorResponse is the Query/Validator response type.
type QueryValidatorResponse struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// metadata is the metadata of the validator.
	Metadata string `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *QueryValidatorResponse) Reset()         { *m = QueryValidatorResponse{} }
func (m *QueryValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorResponse) ProtoMessage()    {}
func (*QueryValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7398e53758cca59c, []int{1}
}
func (m *QueryValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorResponse.Merge(m, src)
}
func (m *QueryValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorResponse proto.InternalMessageInfo

func (m *QueryValidatorResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryValidatorResponse) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// QueryValidatorsRequest is the Query/Validators request type.
type QueryValidatorsRequest struct {
	// pagination is the optional pagination of the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryValidatorsRequest) Reset()         { *m = QueryValidatorsRequest{} }
func (m *QueryValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsRequest) ProtoMessage()    {}
func (*QueryValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7398e53758cca59c, []int{2}
}
func (m *QueryValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsRequest.Merge(m, src)
}
func (m *QueryValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsRequest proto.InternalMessageInfo

func (m *QueryValidatorsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryValidatorsResponse is the Query/Validators response type.
type QueryValidatorsResponse struct {
	// validators is the list of validators.
	Validators []*QueryValidatorsResponse_Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	// pagination is the pagination of the response.
	Pagination *query.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryValidatorsResponse) Reset()         { *m = QueryValidatorsResponse{} }
func (m *QueryValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsResponse) ProtoMessage()    {}
func (*QueryValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7398e53758cca59c, []int{3}
}
func (m *QueryValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse.Merge(m, src)
}
func (m *QueryValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse proto.InternalMessageInfo

func (m *QueryValidatorsResponse) GetValidators() []*QueryValidatorsResponse_Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *QueryValidatorsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// Validator is the validator properties.
type QueryValidatorsResponse_Validator struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// metadata is the metadata of the validator.
	Metadata string `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *QueryValidatorsResponse_Validator) Reset()         { *m = QueryValidatorsResponse_Validator{} }
func (m *QueryValidatorsResponse_Validator) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorsResponse_Validator) ProtoMessage()    {}
func (*QueryValidatorsResponse_Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7398e53758cca59c, []int{3, 0}
}
func (m *QueryValidatorsResponse_Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorsResponse_Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorsResponse_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorsResponse_Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorsResponse_Validator.Merge(m, src)
}
func (m *QueryValidatorsResponse_Validator) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorsResponse_Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorsResponse_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorsResponse_Validator proto.InternalMessageInfo

func (m *QueryValidatorsResponse_Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryValidatorsResponse_Validator) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryValidatorRequest)(nil), "chora.validator.v1.QueryValidatorRequest")
	proto.RegisterType((*QueryValidatorResponse)(nil), "chora.validator.v1.QueryValidatorResponse")
	proto.RegisterType((*QueryValidatorsRequest)(nil), "chora.validator.v1.QueryValidatorsRequest")
	proto.RegisterType((*QueryValidatorsResponse)(nil), "chora.validator.v1.QueryValidatorsResponse")
	proto.RegisterType((*QueryValidatorsResponse_Validator)(nil), "chora.validator.v1.QueryValidatorsResponse.Validator")
}

func init() { proto.RegisterFile("chora/validator/v1/query.proto", fileDescriptor_7398e53758cca59c) }

var fileDescriptor_7398e53758cca59c = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0xaa, 0xd3, 0x40,
	0x1c, 0xc5, 0x3b, 0x2d, 0x7e, 0x74, 0xba, 0x1b, 0x50, 0x4b, 0x90, 0xa1, 0x64, 0xd1, 0x8f, 0x08,
	0x33, 0xa4, 0xe2, 0x03, 0x28, 0xa2, 0x3b, 0xd1, 0x82, 0x2e, 0x5c, 0x39, 0x6d, 0x86, 0x34, 0xd0,
	0x64, 0xd2, 0xcc, 0x34, 0x50, 0xc4, 0x8d, 0x7b, 0x41, 0xe8, 0x73, 0xf8, 0x1e, 0x2e, 0x0b, 0x6e,
	0x5c, 0x4a, 0xeb, 0x63, 0xdc, 0xc5, 0x25, 0x33, 0x69, 0x92, 0xdb, 0xde, 0xde, 0xf6, 0xde, 0xe5,
	0x9f, 0xff, 0x9c, 0x33, 0xbf, 0x73, 0x92, 0x81, 0x78, 0x32, 0x15, 0x09, 0xa3, 0x29, 0x9b, 0x05,
	0x1e, 0x53, 0x22, 0xa1, 0xa9, 0x4b, 0xe7, 0x0b, 0x9e, 0x2c, 0x49, 0x9c, 0x08, 0x25, 0x10, 0xd2,
	0x7b, 0x52, 0xec, 0x49, 0xea, 0x5a, 0xce, 0x44, 0xc8, 0x50, 0x48, 0x3a, 0x66, 0x92, 0x9b, 0xc3,
	0x34, 0x75, 0xc7, 0x5c, 0x31, 0x97, 0xc6, 0xcc, 0x0f, 0x22, 0xa6, 0x02, 0x11, 0x19, 0xbd, 0xf5,
	0xd4, 0x17, 0xc2, 0x9f, 0x71, 0xca, 0xe2, 0x80, 0xb2, 0x28, 0x12, 0x4a, 0x2f, 0xa5, 0xd9, 0xda,
	0x2e, 0x7c, 0xf4, 0x21, 0xd3, 0x7f, 0xda, 0xd9, 0x8f, 0xf8, 0x7c, 0xc1, 0xa5, 0x42, 0x6d, 0xf8,
	0x80, 0x79, 0x5e, 0xc2, 0xa5, 0x6c, 0x83, 0x0e, 0xe8, 0x37, 0x47, 0xbb, 0xd1, 0x7e, 0x07, 0x1f,
	0xef, 0x4b, 0x64, 0x2c, 0x22, 0xc9, 0x8f, 0x6b, 0x90, 0x05, 0x1f, 0x86, 0x5c, 0x31, 0x8f, 0x29,
	0xd6, 0xae, 0xeb, 0x55, 0x31, 0xdb, 0x5f, 0xf6, 0xfd, 0xe4, 0x8e, 0xe1, 0x0d, 0x84, 0x65, 0x1c,
	0x6d, 0xd9, 0x1a, 0x76, 0x89, 0xc9, 0x4e, 0xb2, 0xec, 0xc4, 0x14, 0x95, 0x67, 0x27, 0xef, 0x99,
	0xcf, 0x73, 0xed, 0xa8, 0xa2, 0xb4, 0x2f, 0x00, 0x7c, 0x72, 0x70, 0x45, 0xce, 0xfc, 0x11, 0xc2,
	0xa2, 0xda, 0x0c, 0xbb, 0xd1, 0x6f, 0x0d, 0x5f, 0x90, 0xc3, 0xce, 0xc9, 0x11, 0x03, 0x52, 0xd6,
	0x50, 0x31, 0x42, 0x6f, 0xaf, 0xa0, 0x37, 0x34, 0x7a, 0xef, 0x24, 0xba, 0xb1, 0xac, 0xb2, 0x5b,
	0x2f, 0x61, 0xb3, 0xb8, 0xe1, 0x6e, 0x05, 0x0f, 0x7f, 0xd5, 0xe1, 0x3d, 0x4d, 0x8f, 0x56, 0xa0,
	0xea, 0x36, 0x38, 0x1d, 0x33, 0x6f, 0xd3, 0x72, 0xce, 0x39, 0x6a, 0xe8, 0x6d, 0xfa, 0xfd, 0xcf,
	0xff, 0x55, 0x7d, 0x80, 0x7a, 0xf4, 0x9a, 0x3f, 0xbb, 0x1c, 0xbe, 0xe6, 0xe8, 0xdf, 0xd0, 0x0f,
	0x00, 0x61, 0x59, 0x2c, 0x72, 0xce, 0x6a, 0xdf, 0x70, 0x3d, 0xbb, 0xc5, 0x97, 0xb2, 0xbb, 0x1a,
	0xac, 0x83, 0xf0, 0x8d, 0x60, 0xf2, 0xd5, 0xeb, 0xdf, 0x1b, 0x0c, 0xd6, 0x1b, 0x0c, 0xfe, 0x6d,
	0x30, 0xf8, 0xb9, 0xc5, 0xb5, 0xf5, 0x16, 0xd7, 0xfe, 0x6e, 0x71, 0xed, 0xb3, 0xe3, 0x07, 0x6a,
	0xba, 0x18, 0x93, 0x89, 0x08, 0x8d, 0x47, 0x20, 0x68, 0x28, 0x3c, 0x59, 0xb1, 0x52, 0xcb, 0x98,
	0xcb, 0xec, 0x35, 0xde, 0xd7, 0x0f, 0xec, 0xf9, 0x65, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x46,
	0x99, 0x17, 0xe0, 0x03, 0x00, 0x00,
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
	// Validator queries a validator by address.
	Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error)
	// Validators queries all validators.
	Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error) {
	out := new(QueryValidatorResponse)
	err := c.cc.Invoke(ctx, "/chora.validator.v1.Query/Validator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error) {
	out := new(QueryValidatorsResponse)
	err := c.cc.Invoke(ctx, "/chora.validator.v1.Query/Validators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Validator queries a validator by address.
	Validator(context.Context, *QueryValidatorRequest) (*QueryValidatorResponse, error)
	// Validators queries all validators.
	Validators(context.Context, *QueryValidatorsRequest) (*QueryValidatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Validator(ctx context.Context, req *QueryValidatorRequest) (*QueryValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validator not implemented")
}
func (*UnimplementedQueryServer) Validators(ctx context.Context, req *QueryValidatorsRequest) (*QueryValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validators not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Validator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.validator.v1.Query/Validator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validator(ctx, req.(*QueryValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Validators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.validator.v1.Query/Validators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validators(ctx, req.(*QueryValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chora.validator.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validator",
			Handler:    _Query_Validator_Handler,
		},
		{
			MethodName: "Validators",
			Handler:    _Query_Validators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/validator/v1/query.proto",
}

func (m *QueryValidatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorsResponse_Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorsResponse_Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorsResponse_Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *QueryValidatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorsResponse_Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryValidatorRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryValidatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryValidatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			m.Metadata = string(dAtA[iNdEx:postIndex])
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
func (m *QueryValidatorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryValidatorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
			m.Validators = append(m.Validators, &QueryValidatorsResponse_Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryValidatorsResponse_Validator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			m.Metadata = string(dAtA[iNdEx:postIndex])
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