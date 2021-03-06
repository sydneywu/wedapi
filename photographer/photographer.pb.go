// Code generated by protoc-gen-go. DO NOT EDIT.
// source: photographer/photographer.proto

package wedapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Photographer struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartDate            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Photographer) Reset()         { *m = Photographer{} }
func (m *Photographer) String() string { return proto.CompactTextString(m) }
func (*Photographer) ProtoMessage()    {}
func (*Photographer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e03d7f32df560c2b, []int{0}
}

func (m *Photographer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Photographer.Unmarshal(m, b)
}
func (m *Photographer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Photographer.Marshal(b, m, deterministic)
}
func (m *Photographer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Photographer.Merge(m, src)
}
func (m *Photographer) XXX_Size() int {
	return xxx_messageInfo_Photographer.Size(m)
}
func (m *Photographer) XXX_DiscardUnknown() {
	xxx_messageInfo_Photographer.DiscardUnknown(m)
}

var xxx_messageInfo_Photographer proto.InternalMessageInfo

func (m *Photographer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Photographer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Photographer) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Photographer) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

// The request message containing the user's name.
type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e03d7f32df560c2b, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Photographer)(nil), "photographer.Photographer")
	proto.RegisterType((*GetRequest)(nil), "photographer.GetRequest")
}

func init() {
	proto.RegisterFile("photographer/photographer.proto", fileDescriptor_e03d7f32df560c2b)
}

var fileDescriptor_e03d7f32df560c2b = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x6a, 0xf3, 0x30,
	0x14, 0x85, 0x7f, 0x25, 0x21, 0x90, 0x9b, 0xf0, 0x0f, 0x6a, 0x07, 0xe3, 0x96, 0xc6, 0x64, 0xca,
	0x24, 0x43, 0xba, 0x74, 0xe9, 0x92, 0x16, 0xb2, 0x06, 0xb7, 0xa5, 0xd0, 0x4d, 0xb6, 0x6e, 0x6d,
	0x41, 0x6c, 0xa9, 0xf2, 0x75, 0x4d, 0x5e, 0xa2, 0x73, 0x1f, 0xb7, 0xc4, 0x21, 0x44, 0x5e, 0xb2,
	0x49, 0xf7, 0x9c, 0x0f, 0x3e, 0x0e, 0xcc, 0x6d, 0x61, 0xc8, 0xe4, 0x4e, 0xda, 0x02, 0x5d, 0xec,
	0x7f, 0x84, 0x75, 0x86, 0x0c, 0x9f, 0xf9, 0xb7, 0x70, 0x9e, 0x1b, 0x93, 0xef, 0x30, 0xee, 0xb2,
	0xb4, 0xf9, 0x8c, 0x49, 0x97, 0x58, 0x93, 0x2c, 0xed, 0xb1, 0xbe, 0xf8, 0x61, 0x30, 0xdb, 0x7a,
	0x04, 0xff, 0x0f, 0x03, 0xad, 0x02, 0x16, 0xb1, 0xe5, 0x30, 0x19, 0x68, 0xc5, 0x39, 0x8c, 0x2a,
	0x59, 0x62, 0x30, 0x88, 0xd8, 0x72, 0x92, 0x74, 0x6f, 0x1e, 0xc1, 0x54, 0x61, 0x9d, 0x39, 0x6d,
	0x49, 0x9b, 0x2a, 0x18, 0x76, 0x91, 0x7f, 0xe2, 0x0f, 0x30, 0xa9, 0x49, 0x3a, 0x7a, 0x96, 0x84,
	0xc1, 0x28, 0x62, 0xcb, 0xe9, 0x2a, 0x14, 0x47, 0x17, 0x71, 0x72, 0x11, 0xaf, 0x27, 0x97, 0xe4,
	0x5c, 0x5e, 0xdc, 0x02, 0x6c, 0x90, 0x12, 0xfc, 0x6a, 0xb0, 0x26, 0xcf, 0x66, 0x72, 0xb0, 0x59,
	0xfd, 0x32, 0xb8, 0xf2, 0x75, 0x5f, 0xd0, 0x7d, 0xeb, 0x0c, 0xf9, 0x1a, 0xc6, 0x4f, 0x0e, 0x25,
	0x21, 0x0f, 0x45, 0x6f, 0x14, 0xbf, 0x1c, 0x5e, 0xc8, 0x16, 0xff, 0xf8, 0x23, 0x0c, 0x37, 0x48,
	0x3c, 0xe8, 0x97, 0xce, 0x32, 0x97, 0xf1, 0xf5, 0x1b, 0xdc, 0x65, 0xa6, 0x14, 0x2d, 0x2a, 0xa5,
	0xab, 0xdc, 0x62, 0x95, 0xe9, 0x5d, 0x0f, 0x58, 0x5f, 0xbf, 0xa3, 0xf2, 0xa1, 0xed, 0x61, 0x88,
	0x2d, 0xfb, 0xb8, 0xc9, 0x35, 0x15, 0x45, 0x93, 0x8a, 0xcc, 0x94, 0x71, 0xbd, 0x57, 0x15, 0xee,
	0xdb, 0x26, 0x6e, 0x51, 0x49, 0xab, 0xd3, 0x71, 0x37, 0xd7, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x13, 0xeb, 0x24, 0x47, 0xf9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PhotographerServiceClient is the client API for PhotographerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhotographerServiceClient interface {
	Create(ctx context.Context, in *Photographer, opts ...grpc.CallOption) (*Photographer, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Photographer, error)
}

type photographerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotographerServiceClient(cc grpc.ClientConnInterface) PhotographerServiceClient {
	return &photographerServiceClient{cc}
}

func (c *photographerServiceClient) Create(ctx context.Context, in *Photographer, opts ...grpc.CallOption) (*Photographer, error) {
	out := new(Photographer)
	err := c.cc.Invoke(ctx, "/photographer.PhotographerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photographerServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Photographer, error) {
	out := new(Photographer)
	err := c.cc.Invoke(ctx, "/photographer.PhotographerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotographerServiceServer is the server API for PhotographerService service.
type PhotographerServiceServer interface {
	Create(context.Context, *Photographer) (*Photographer, error)
	Get(context.Context, *GetRequest) (*Photographer, error)
}

// UnimplementedPhotographerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPhotographerServiceServer struct {
}

func (*UnimplementedPhotographerServiceServer) Create(ctx context.Context, req *Photographer) (*Photographer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPhotographerServiceServer) Get(ctx context.Context, req *GetRequest) (*Photographer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterPhotographerServiceServer(s *grpc.Server, srv PhotographerServiceServer) {
	s.RegisterService(&_PhotographerService_serviceDesc, srv)
}

func _PhotographerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Photographer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotographerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photographer.PhotographerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotographerServiceServer).Create(ctx, req.(*Photographer))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotographerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotographerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photographer.PhotographerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotographerServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhotographerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "photographer.PhotographerService",
	HandlerType: (*PhotographerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PhotographerService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PhotographerService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "photographer/photographer.proto",
}
