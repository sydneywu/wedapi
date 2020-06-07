// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/photographer/photographer.proto

package photographer

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
	DateFounded          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=dateFounded,proto3" json:"dateFounded,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Photographer) Reset()         { *m = Photographer{} }
func (m *Photographer) String() string { return proto.CompactTextString(m) }
func (*Photographer) ProtoMessage()    {}
func (*Photographer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa78fcc0892a986, []int{0}
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

func (m *Photographer) GetDateFounded() *timestamp.Timestamp {
	if m != nil {
		return m.DateFounded
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
	return fileDescriptor_1fa78fcc0892a986, []int{1}
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
	proto.RegisterFile("api/photographer/photographer.proto", fileDescriptor_1fa78fcc0892a986)
}

var fileDescriptor_1fa78fcc0892a986 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x6b, 0xf3, 0x30,
	0x10, 0x86, 0x3f, 0xc5, 0x21, 0x90, 0x73, 0xf8, 0x06, 0x75, 0x31, 0xa6, 0x50, 0xe3, 0x2e, 0x1e,
	0x8a, 0x0c, 0xe9, 0xda, 0x2e, 0x29, 0x34, 0x6b, 0x71, 0x3b, 0x75, 0x93, 0xad, 0xab, 0x2d, 0xa8,
	0x2d, 0x55, 0x3e, 0x37, 0xe4, 0x6f, 0x74, 0xea, 0xcf, 0x2d, 0x71, 0x08, 0x95, 0x3b, 0x64, 0x93,
	0xee, 0x7d, 0x0e, 0x1e, 0xee, 0x85, 0x6b, 0x69, 0x75, 0x6e, 0x1b, 0x43, 0xa6, 0x76, 0xd2, 0x36,
	0xe8, 0x26, 0x1f, 0x61, 0x9d, 0x21, 0xc3, 0x57, 0xfe, 0x2c, 0xbe, 0xaa, 0x8d, 0xa9, 0xdf, 0x31,
	0x1f, 0xb3, 0x72, 0x78, 0xcb, 0x49, 0xb7, 0xd8, 0x93, 0x6c, 0xed, 0x11, 0x4f, 0xbf, 0x18, 0xac,
	0x9e, 0xbc, 0x0d, 0xfe, 0x1f, 0x66, 0x5a, 0x45, 0x2c, 0x61, 0x59, 0x50, 0xcc, 0xb4, 0xe2, 0x1c,
	0xe6, 0x9d, 0x6c, 0x31, 0x9a, 0x25, 0x2c, 0x5b, 0x16, 0xe3, 0x9b, 0x27, 0x10, 0x2a, 0xec, 0x2b,
	0xa7, 0x2d, 0x69, 0xd3, 0x45, 0xc1, 0x18, 0xf9, 0x23, 0x7e, 0x07, 0xa1, 0x92, 0x84, 0x8f, 0x66,
	0xe8, 0x14, 0xaa, 0x68, 0x9e, 0xb0, 0x2c, 0x5c, 0xc7, 0xe2, 0x68, 0x23, 0x4e, 0x36, 0xe2, 0xe5,
	0x64, 0x53, 0xf8, 0x78, 0x7a, 0x09, 0xb0, 0x45, 0x2a, 0xf0, 0x63, 0xc0, 0x9e, 0x3c, 0xa3, 0xe5,
	0xc1, 0x68, 0xfd, 0xcd, 0xe0, 0xc2, 0x57, 0x7e, 0x46, 0xf7, 0xa9, 0x2b, 0xe4, 0x1b, 0x58, 0x3c,
	0x38, 0x94, 0x84, 0x3c, 0x16, 0x93, 0xc3, 0xf8, 0x70, 0x7c, 0x26, 0x4b, 0xff, 0xf1, 0x7b, 0x08,
	0xb6, 0x48, 0x3c, 0x9a, 0x42, 0xbf, 0x32, 0xe7, 0xd7, 0x37, 0xe2, 0xf5, 0xa6, 0xd6, 0xd4, 0x34,
	0x43, 0x29, 0x2a, 0xd3, 0xe6, 0xfd, 0x5e, 0x75, 0xb8, 0xdf, 0x0d, 0xf9, 0x0e, 0xd5, 0xa1, 0xbb,
	0xbf, 0xfd, 0x95, 0x8b, 0xf1, 0x12, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x31, 0x19,
	0xcd, 0xda, 0x01, 0x00, 0x00,
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
	Metadata: "api/photographer/photographer.proto",
}