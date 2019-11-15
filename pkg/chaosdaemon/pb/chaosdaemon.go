// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chaosdaemon.proto

package chaosdaemon

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NetemRequest struct {
	Netem                *Netem   `protobuf:"bytes,1,opt,name=netem,proto3" json:"netem,omitempty"`
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetemRequest) Reset()         { *m = NetemRequest{} }
func (m *NetemRequest) String() string { return proto.CompactTextString(m) }
func (*NetemRequest) ProtoMessage()    {}
func (*NetemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4925ade7a16fc3f, []int{0}
}

func (m *NetemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetemRequest.Unmarshal(m, b)
}
func (m *NetemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetemRequest.Marshal(b, m, deterministic)
}
func (m *NetemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetemRequest.Merge(m, src)
}
func (m *NetemRequest) XXX_Size() int {
	return xxx_messageInfo_NetemRequest.Size(m)
}
func (m *NetemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetemRequest proto.InternalMessageInfo

func (m *NetemRequest) GetNetem() *Netem {
	if m != nil {
		return m.Netem
	}
	return nil
}

func (m *NetemRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

type Netem struct {
	Time                 uint32   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Jitter               uint32   `protobuf:"varint,2,opt,name=jitter,proto3" json:"jitter,omitempty"`
	DelayCorr            float32  `protobuf:"fixed32,3,opt,name=delay_corr,json=delayCorr,proto3" json:"delay_corr,omitempty"`
	Limit                uint32   `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Loss                 float32  `protobuf:"fixed32,5,opt,name=loss,proto3" json:"loss,omitempty"`
	LossCorr             float32  `protobuf:"fixed32,6,opt,name=loss_corr,json=lossCorr,proto3" json:"loss_corr,omitempty"`
	Gap                  uint32   `protobuf:"varint,7,opt,name=gap,proto3" json:"gap,omitempty"`
	Duplicate            float32  `protobuf:"fixed32,8,opt,name=duplicate,proto3" json:"duplicate,omitempty"`
	DuplicateCorr        float32  `protobuf:"fixed32,9,opt,name=duplicate_corr,json=duplicateCorr,proto3" json:"duplicate_corr,omitempty"`
	Reorder              float32  `protobuf:"fixed32,10,opt,name=reorder,proto3" json:"reorder,omitempty"`
	ReorderCorr          float32  `protobuf:"fixed32,11,opt,name=reorder_corr,json=reorderCorr,proto3" json:"reorder_corr,omitempty"`
	Corrupt              float32  `protobuf:"fixed32,12,opt,name=corrupt,proto3" json:"corrupt,omitempty"`
	CorruptCorr          float32  `protobuf:"fixed32,13,opt,name=corrupt_corr,json=corruptCorr,proto3" json:"corrupt_corr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Netem) Reset()         { *m = Netem{} }
func (m *Netem) String() string { return proto.CompactTextString(m) }
func (*Netem) ProtoMessage()    {}
func (*Netem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4925ade7a16fc3f, []int{1}
}

func (m *Netem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Netem.Unmarshal(m, b)
}
func (m *Netem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Netem.Marshal(b, m, deterministic)
}
func (m *Netem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Netem.Merge(m, src)
}
func (m *Netem) XXX_Size() int {
	return xxx_messageInfo_Netem.Size(m)
}
func (m *Netem) XXX_DiscardUnknown() {
	xxx_messageInfo_Netem.DiscardUnknown(m)
}

var xxx_messageInfo_Netem proto.InternalMessageInfo

func (m *Netem) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Netem) GetJitter() uint32 {
	if m != nil {
		return m.Jitter
	}
	return 0
}

func (m *Netem) GetDelayCorr() float32 {
	if m != nil {
		return m.DelayCorr
	}
	return 0
}

func (m *Netem) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Netem) GetLoss() float32 {
	if m != nil {
		return m.Loss
	}
	return 0
}

func (m *Netem) GetLossCorr() float32 {
	if m != nil {
		return m.LossCorr
	}
	return 0
}

func (m *Netem) GetGap() uint32 {
	if m != nil {
		return m.Gap
	}
	return 0
}

func (m *Netem) GetDuplicate() float32 {
	if m != nil {
		return m.Duplicate
	}
	return 0
}

func (m *Netem) GetDuplicateCorr() float32 {
	if m != nil {
		return m.DuplicateCorr
	}
	return 0
}

func (m *Netem) GetReorder() float32 {
	if m != nil {
		return m.Reorder
	}
	return 0
}

func (m *Netem) GetReorderCorr() float32 {
	if m != nil {
		return m.ReorderCorr
	}
	return 0
}

func (m *Netem) GetCorrupt() float32 {
	if m != nil {
		return m.Corrupt
	}
	return 0
}

func (m *Netem) GetCorruptCorr() float32 {
	if m != nil {
		return m.CorruptCorr
	}
	return 0
}

func init() {
	proto.RegisterType((*NetemRequest)(nil), "chaosdaemon.NetemRequest")
	proto.RegisterType((*Netem)(nil), "chaosdaemon.Netem")
}

func init() { proto.RegisterFile("chaosdaemon.proto", fileDescriptor_d4925ade7a16fc3f) }

var fileDescriptor_d4925ade7a16fc3f = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0x0a, 0x14, 0xda, 0x29, 0xe5, 0x33, 0x1b, 0x43, 0x36, 0xa0, 0x09, 0x92, 0x90, 0x70,
	0x2a, 0x09, 0x5e, 0x4d, 0x3c, 0x88, 0x07, 0x2f, 0x1e, 0xaa, 0x07, 0x6f, 0xa4, 0xb4, 0x23, 0x59,
	0xd3, 0x76, 0xeb, 0xb2, 0x3d, 0xf0, 0x13, 0xfc, 0xab, 0xfe, 0x0a, 0xb3, 0xb3, 0x6d, 0x0f, 0x1e,
	0x3d, 0xed, 0xbc, 0xf7, 0xe6, 0xbd, 0x4d, 0x66, 0x06, 0x26, 0x3a, 0xcd, 0x12, 0x2c, 0x64, 0x19,
	0x55, 0x4a, 0x6a, 0xc9, 0xbc, 0x16, 0xcf, 0xe6, 0x47, 0x29, 0x8f, 0x39, 0x6e, 0x88, 0x3f, 0xd4,
	0xef, 0x1b, 0x2c, 0x2a, 0x7d, 0xb6, 0x6d, 0xcb, 0x37, 0x18, 0x3f, 0xa3, 0xc6, 0x22, 0xc6, 0xcf,
	0x1a, 0x4f, 0x9a, 0xad, 0xc0, 0x2d, 0x0d, 0xe6, 0xce, 0xc2, 0x59, 0x07, 0xdb, 0xff, 0x51, 0x17,
	0x6b, 0xdb, 0xac, 0xca, 0x6e, 0x60, 0x9c, 0xca, 0x52, 0x27, 0xa2, 0x44, 0xb5, 0x17, 0x19, 0xef,
	0x2d, 0x9c, 0xb5, 0x1f, 0x07, 0x1d, 0xf7, 0x94, 0x2d, 0xbf, 0x7b, 0xe0, 0x92, 0x87, 0x31, 0x18,
	0x68, 0x51, 0x20, 0x45, 0x86, 0x31, 0xd5, 0x6c, 0x0a, 0xc3, 0x0f, 0xa1, 0x35, 0x2a, 0xb2, 0x86,
	0x71, 0x83, 0xd8, 0x35, 0x40, 0x86, 0x79, 0x72, 0xde, 0xa7, 0x52, 0x29, 0xde, 0x5f, 0x38, 0xeb,
	0x5e, 0xec, 0x13, 0xf3, 0x20, 0x95, 0x62, 0x97, 0xe0, 0xe6, 0xa2, 0x10, 0x9a, 0x0f, 0xc8, 0x65,
	0x81, 0xf9, 0x20, 0x97, 0xa7, 0x13, 0x77, 0xa9, 0x9d, 0x6a, 0x36, 0x07, 0xdf, 0xbc, 0x36, 0x67,
	0x48, 0x82, 0x67, 0x08, 0x8a, 0xb9, 0x80, 0xfe, 0x31, 0xa9, 0xf8, 0x88, 0x42, 0x4c, 0xc9, 0xae,
	0xc0, 0xcf, 0xea, 0x2a, 0x17, 0x69, 0xa2, 0x91, 0x7b, 0xcd, 0xb7, 0x2d, 0xc1, 0x56, 0x30, 0xe9,
	0x80, 0x4d, 0xf4, 0xa9, 0x25, 0xec, 0x58, 0x8a, 0xe5, 0x30, 0x52, 0x28, 0x55, 0x86, 0x8a, 0x03,
	0xe9, 0x2d, 0x34, 0xf3, 0x6a, 0x4a, 0x6b, 0x0f, 0x48, 0x0e, 0x1a, 0xae, 0x35, 0x1b, 0xa9, 0xae,
	0x34, 0x1f, 0x5b, 0x73, 0x03, 0xed, 0xb0, 0xa9, 0xb4, 0xe6, 0xd0, 0x9a, 0x1b, 0xce, 0x98, 0xb7,
	0x5f, 0x0e, 0x78, 0xaf, 0xe9, 0x8e, 0x36, 0xc5, 0xee, 0xc0, 0x7b, 0x41, 0x6d, 0x67, 0x3f, 0xfd,
	0xbd, 0x40, 0xbb, 0xe7, 0xd9, 0x34, 0xb2, 0x57, 0x11, 0xb5, 0x57, 0x11, 0x3d, 0x9a, 0xab, 0x58,
	0xfe, 0x63, 0xf7, 0x10, 0xec, 0x30, 0x47, 0x8d, 0x7f, 0x0c, 0x38, 0x0c, 0x89, 0xb9, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0x65, 0xc0, 0x25, 0x92, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChaosDaemonClient is the client API for ChaosDaemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChaosDaemonClient interface {
	SetNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type chaosDaemonClient struct {
	cc *grpc.ClientConn
}

func NewChaosDaemonClient(cc *grpc.ClientConn) ChaosDaemonClient {
	return &chaosDaemonClient{cc}
}

func (c *chaosDaemonClient) SetNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/SetNetem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosDaemonClient) DeleteNetem(ctx context.Context, in *NetemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chaosdaemon.ChaosDaemon/DeleteNetem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaosDaemonServer is the server API for ChaosDaemon service.
type ChaosDaemonServer interface {
	SetNetem(context.Context, *NetemRequest) (*empty.Empty, error)
	DeleteNetem(context.Context, *NetemRequest) (*empty.Empty, error)
}

// UnimplementedChaosDaemonServer can be embedded to have forward compatible implementations.
type UnimplementedChaosDaemonServer struct {
}

func (*UnimplementedChaosDaemonServer) SetNetem(ctx context.Context, req *NetemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNetem not implemented")
}
func (*UnimplementedChaosDaemonServer) DeleteNetem(ctx context.Context, req *NetemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetem not implemented")
}

func RegisterChaosDaemonServer(s *grpc.Server, srv ChaosDaemonServer) {
	s.RegisterService(&_ChaosDaemon_serviceDesc, srv)
}

func _ChaosDaemon_SetNetem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).SetNetem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/SetNetem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).SetNetem(ctx, req.(*NetemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosDaemon_DeleteNetem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosDaemonServer).DeleteNetem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chaosdaemon.ChaosDaemon/DeleteNetem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosDaemonServer).DeleteNetem(ctx, req.(*NetemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChaosDaemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chaosdaemon.ChaosDaemon",
	HandlerType: (*ChaosDaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetNetem",
			Handler:    _ChaosDaemon_SetNetem_Handler,
		},
		{
			MethodName: "DeleteNetem",
			Handler:    _ChaosDaemon_DeleteNetem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chaosdaemon.proto",
}