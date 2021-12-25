// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: domain_service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Port struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	IsActive             bool     `protobuf:"varint,3,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	Company              string   `protobuf:"bytes,4,opt,name=Company,proto3" json:"Company,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Address              string   `protobuf:"bytes,7,opt,name=Address,proto3" json:"Address,omitempty"`
	About                string   `protobuf:"bytes,8,opt,name=About,proto3" json:"About,omitempty"`
	Registered           string   `protobuf:"bytes,9,opt,name=Registered,proto3" json:"Registered,omitempty"`
	Latitude             float32  `protobuf:"fixed32,10,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,11,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f7c7947346ea07, []int{0}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Port) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Port) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Port) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Port) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Port) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *Port) GetRegistered() string {
	if m != nil {
		return m.Registered
	}
	return ""
}

func (m *Port) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Port) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Ports struct {
	Ports                []*Port  `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ports) Reset()         { *m = Ports{} }
func (m *Ports) String() string { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()    {}
func (*Ports) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f7c7947346ea07, []int{1}
}
func (m *Ports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ports.Unmarshal(m, b)
}
func (m *Ports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ports.Marshal(b, m, deterministic)
}
func (m *Ports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ports.Merge(m, src)
}
func (m *Ports) XXX_Size() int {
	return xxx_messageInfo_Ports.Size(m)
}
func (m *Ports) XXX_DiscardUnknown() {
	xxx_messageInfo_Ports.DiscardUnknown(m)
}

var xxx_messageInfo_Ports proto.InternalMessageInfo

func (m *Ports) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type RequestToListPorts struct {
	Filters              *FilterVars `protobuf:"bytes,1,opt,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RequestToListPorts) Reset()         { *m = RequestToListPorts{} }
func (m *RequestToListPorts) String() string { return proto.CompactTextString(m) }
func (*RequestToListPorts) ProtoMessage()    {}
func (*RequestToListPorts) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f7c7947346ea07, []int{2}
}
func (m *RequestToListPorts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestToListPorts.Unmarshal(m, b)
}
func (m *RequestToListPorts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestToListPorts.Marshal(b, m, deterministic)
}
func (m *RequestToListPorts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestToListPorts.Merge(m, src)
}
func (m *RequestToListPorts) XXX_Size() int {
	return xxx_messageInfo_RequestToListPorts.Size(m)
}
func (m *RequestToListPorts) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestToListPorts.DiscardUnknown(m)
}

var xxx_messageInfo_RequestToListPorts proto.InternalMessageInfo

func (m *RequestToListPorts) GetFilters() *FilterVars {
	if m != nil {
		return m.Filters
	}
	return nil
}

type FilterVars struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdLower              int64    `protobuf:"varint,2,opt,name=idLower,proto3" json:"idLower,omitempty"`
	IdUpper              int64    `protobuf:"varint,3,opt,name=idUpper,proto3" json:"idUpper,omitempty"`
	FilterString         string   `protobuf:"bytes,4,opt,name=filterString,proto3" json:"filterString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterVars) Reset()         { *m = FilterVars{} }
func (m *FilterVars) String() string { return proto.CompactTextString(m) }
func (*FilterVars) ProtoMessage()    {}
func (*FilterVars) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f7c7947346ea07, []int{3}
}
func (m *FilterVars) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterVars.Unmarshal(m, b)
}
func (m *FilterVars) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterVars.Marshal(b, m, deterministic)
}
func (m *FilterVars) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterVars.Merge(m, src)
}
func (m *FilterVars) XXX_Size() int {
	return xxx_messageInfo_FilterVars.Size(m)
}
func (m *FilterVars) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterVars.DiscardUnknown(m)
}

var xxx_messageInfo_FilterVars proto.InternalMessageInfo

func (m *FilterVars) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FilterVars) GetIdLower() int64 {
	if m != nil {
		return m.IdLower
	}
	return 0
}

func (m *FilterVars) GetIdUpper() int64 {
	if m != nil {
		return m.IdUpper
	}
	return 0
}

func (m *FilterVars) GetFilterString() string {
	if m != nil {
		return m.FilterString
	}
	return ""
}

type DefaultResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	SavedIdRecord        []int64  `protobuf:"varint,2,rep,packed,name=savedIdRecord,proto3" json:"savedIdRecord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultResponse) Reset()         { *m = DefaultResponse{} }
func (m *DefaultResponse) String() string { return proto.CompactTextString(m) }
func (*DefaultResponse) ProtoMessage()    {}
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f7c7947346ea07, []int{4}
}
func (m *DefaultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultResponse.Unmarshal(m, b)
}
func (m *DefaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultResponse.Marshal(b, m, deterministic)
}
func (m *DefaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultResponse.Merge(m, src)
}
func (m *DefaultResponse) XXX_Size() int {
	return xxx_messageInfo_DefaultResponse.Size(m)
}
func (m *DefaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultResponse proto.InternalMessageInfo

func (m *DefaultResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DefaultResponse) GetSavedIdRecord() []int64 {
	if m != nil {
		return m.SavedIdRecord
	}
	return nil
}

func init() {
	proto.RegisterType((*Port)(nil), "proto.Port")
	proto.RegisterType((*Ports)(nil), "proto.Ports")
	proto.RegisterType((*RequestToListPorts)(nil), "proto.RequestToListPorts")
	proto.RegisterType((*FilterVars)(nil), "proto.FilterVars")
	proto.RegisterType((*DefaultResponse)(nil), "proto.DefaultResponse")
}

func init() { proto.RegisterFile("domain_service.proto", fileDescriptor_74f7c7947346ea07) }

var fileDescriptor_74f7c7947346ea07 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x14, 0x85, 0x95, 0xa4, 0x59, 0xdb, 0x9b, 0x01, 0xc2, 0x9a, 0x90, 0xa9, 0x10, 0x0a, 0x11, 0x0f,
	0x11, 0x48, 0x7b, 0x68, 0x7f, 0x41, 0x44, 0x41, 0xaa, 0x54, 0xa1, 0xe1, 0x31, 0x5e, 0x51, 0x16,
	0xdf, 0x15, 0x4b, 0x6d, 0x1c, 0x6c, 0xa7, 0x88, 0x3f, 0xc2, 0xef, 0x45, 0xbe, 0x4e, 0xb6, 0x15,
	0x9e, 0xea, 0xef, 0x1c, 0xdf, 0x2b, 0xf7, 0x9c, 0xc0, 0x85, 0xd4, 0x87, 0x5a, 0xb5, 0xdf, 0x2d,
	0x9a, 0xa3, 0x6a, 0xf0, 0xb2, 0x33, 0xda, 0x69, 0x96, 0xd2, 0x4f, 0xf1, 0x27, 0x86, 0xc9, 0x95,
	0x36, 0x8e, 0x3d, 0x85, 0x78, 0xb3, 0xe6, 0x51, 0x1e, 0x95, 0x89, 0x88, 0x37, 0x6b, 0xc6, 0x60,
	0xf2, 0xb9, 0x3e, 0x20, 0x8f, 0xf3, 0xa8, 0x9c, 0x0b, 0x3a, 0xb3, 0x05, 0xcc, 0x36, 0xb6, 0x6a,
	0x9c, 0x3a, 0x22, 0x4f, 0xf2, 0xa8, 0x9c, 0x89, 0x7b, 0x66, 0x1c, 0xa6, 0x1f, 0xf4, 0xa1, 0xab,
	0xdb, 0xdf, 0x7c, 0x42, 0x23, 0x23, 0xb2, 0x0b, 0x48, 0x3f, 0x1e, 0x6a, 0xb5, 0xe7, 0x29, 0xe9,
	0x01, 0xbc, 0x7a, 0xf5, 0x43, 0xb7, 0xc8, 0xcf, 0x82, 0x4a, 0xe0, 0xb7, 0x54, 0x52, 0x1a, 0xb4,
	0x96, 0x4f, 0xc3, 0x96, 0x01, 0xfd, 0xfd, 0xea, 0x56, 0xf7, 0x8e, 0xcf, 0xc2, 0x7d, 0x02, 0xf6,
	0x1a, 0x40, 0xe0, 0x4e, 0x59, 0x87, 0x06, 0x25, 0x9f, 0x93, 0xf5, 0x48, 0xf1, 0x2f, 0xde, 0xd6,
	0x4e, 0xb9, 0x5e, 0x22, 0x87, 0x3c, 0x2a, 0x63, 0x71, 0xcf, 0xec, 0x15, 0xcc, 0xb7, 0xba, 0xdd,
	0x05, 0x33, 0x23, 0xf3, 0x41, 0x28, 0xde, 0x41, 0xea, 0x73, 0xb1, 0xec, 0x0d, 0xa4, 0x9d, 0x3f,
	0xf0, 0x28, 0x4f, 0xca, 0x6c, 0x99, 0x85, 0xfc, 0x2e, 0xbd, 0x29, 0x82, 0x53, 0x54, 0xc0, 0x04,
	0xfe, 0xec, 0xd1, 0xba, 0xaf, 0x7a, 0xab, 0xac, 0x0b, 0x83, 0xef, 0x61, 0x7a, 0xa7, 0xf6, 0x0e,
	0x8d, 0xa5, 0x58, 0xb3, 0xe5, 0xf3, 0x61, 0xf4, 0x13, 0xa9, 0xdf, 0x6a, 0x63, 0xc5, 0x78, 0xa3,
	0x70, 0x00, 0x0f, 0xb2, 0x2f, 0x43, 0xc9, 0xb1, 0x0c, 0x25, 0x7d, 0x2c, 0x4a, 0x6e, 0xf5, 0x2f,
	0x34, 0xd4, 0x47, 0x22, 0x46, 0x0c, 0xce, 0x4d, 0xd7, 0xa1, 0xa1, 0x46, 0xc8, 0x21, 0x64, 0x05,
	0x9c, 0x87, 0xe5, 0xd7, 0xce, 0xa8, 0x76, 0x37, 0xb4, 0x72, 0xa2, 0x15, 0x5f, 0xe0, 0xd9, 0x1a,
	0xef, 0xea, 0x7e, 0xef, 0x04, 0xda, 0x4e, 0xb7, 0x96, 0x1a, 0xb0, 0x7d, 0xd3, 0xf8, 0x06, 0x22,
	0xaa, 0x78, 0x44, 0xf6, 0x16, 0x9e, 0xd8, 0xfa, 0x88, 0x72, 0x23, 0x05, 0x36, 0xda, 0x48, 0x1e,
	0xe7, 0x49, 0x99, 0x88, 0x53, 0x71, 0x79, 0x04, 0xa8, 0xa4, 0xbc, 0x0e, 0xdf, 0x1a, 0x5b, 0xc1,
	0x6c, 0x0c, 0x84, 0xbd, 0x1c, 0xfe, 0xfe, 0xff, 0x51, 0x2d, 0xce, 0x1f, 0x85, 0x6a, 0xd9, 0x0a,
	0xb2, 0x9b, 0xce, 0xa2, 0x19, 0x72, 0x3c, 0x31, 0x17, 0x2f, 0x06, 0xfa, 0xe7, 0xdd, 0xb7, 0x67,
	0x24, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xfe, 0x27, 0xa8, 0xee, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	ListPort(ctx context.Context, in *RequestToListPorts, opts ...grpc.CallOption) (*Ports, error)
	UpsertPorts(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) ListPort(ctx context.Context, in *RequestToListPorts, opts ...grpc.CallOption) (*Ports, error) {
	out := new(Ports)
	err := c.cc.Invoke(ctx, "/proto.AddService/ListPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) UpsertPorts(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/proto.AddService/UpsertPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	ListPort(context.Context, *RequestToListPorts) (*Ports, error)
	UpsertPorts(context.Context, *Ports) (*DefaultResponse, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) ListPort(ctx context.Context, req *RequestToListPorts) (*Ports, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPort not implemented")
}
func (*UnimplementedAddServiceServer) UpsertPorts(ctx context.Context, req *Ports) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertPorts not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_ListPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToListPorts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).ListPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/ListPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).ListPort(ctx, req.(*RequestToListPorts))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_UpsertPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ports)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).UpsertPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/UpsertPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).UpsertPorts(ctx, req.(*Ports))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPort",
			Handler:    _AddService_ListPort_Handler,
		},
		{
			MethodName: "UpsertPorts",
			Handler:    _AddService_UpsertPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain_service.proto",
}
