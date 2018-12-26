// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gpio.proto

package rpi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type InitReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitReq) Reset()         { *m = InitReq{} }
func (m *InitReq) String() string { return proto.CompactTextString(m) }
func (*InitReq) ProtoMessage()    {}
func (*InitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{0}
}

func (m *InitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitReq.Unmarshal(m, b)
}
func (m *InitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitReq.Marshal(b, m, deterministic)
}
func (m *InitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitReq.Merge(m, src)
}
func (m *InitReq) XXX_Size() int {
	return xxx_messageInfo_InitReq.Size(m)
}
func (m *InitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InitReq.DiscardUnknown(m)
}

var xxx_messageInfo_InitReq proto.InternalMessageInfo

type CloseReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseReq) Reset()         { *m = CloseReq{} }
func (m *CloseReq) String() string { return proto.CompactTextString(m) }
func (*CloseReq) ProtoMessage()    {}
func (*CloseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{1}
}

func (m *CloseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseReq.Unmarshal(m, b)
}
func (m *CloseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseReq.Marshal(b, m, deterministic)
}
func (m *CloseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseReq.Merge(m, src)
}
func (m *CloseReq) XXX_Size() int {
	return xxx_messageInfo_CloseReq.Size(m)
}
func (m *CloseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseReq proto.InternalMessageInfo

type SetDirectionReq struct {
	Pin                  string   `protobuf:"bytes,1,opt,name=pin,proto3" json:"pin,omitempty"`
	Direction            int32    `protobuf:"varint,2,opt,name=direction,proto3" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDirectionReq) Reset()         { *m = SetDirectionReq{} }
func (m *SetDirectionReq) String() string { return proto.CompactTextString(m) }
func (*SetDirectionReq) ProtoMessage()    {}
func (*SetDirectionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{2}
}

func (m *SetDirectionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDirectionReq.Unmarshal(m, b)
}
func (m *SetDirectionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDirectionReq.Marshal(b, m, deterministic)
}
func (m *SetDirectionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDirectionReq.Merge(m, src)
}
func (m *SetDirectionReq) XXX_Size() int {
	return xxx_messageInfo_SetDirectionReq.Size(m)
}
func (m *SetDirectionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDirectionReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetDirectionReq proto.InternalMessageInfo

func (m *SetDirectionReq) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

func (m *SetDirectionReq) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type PinInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias                []string `protobuf:"bytes,2,rep,name=alias,proto3" json:"alias,omitempty"`
	Caps                 int32    `protobuf:"varint,3,opt,name=caps,proto3" json:"caps,omitempty"`
	DigitalLogical       int32    `protobuf:"varint,4,opt,name=digitalLogical,proto3" json:"digitalLogical,omitempty"`
	AnalogLogical        int32    `protobuf:"varint,5,opt,name=analogLogical,proto3" json:"analogLogical,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PinInfo) Reset()         { *m = PinInfo{} }
func (m *PinInfo) String() string { return proto.CompactTextString(m) }
func (*PinInfo) ProtoMessage()    {}
func (*PinInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{3}
}

func (m *PinInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PinInfo.Unmarshal(m, b)
}
func (m *PinInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PinInfo.Marshal(b, m, deterministic)
}
func (m *PinInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PinInfo.Merge(m, src)
}
func (m *PinInfo) XXX_Size() int {
	return xxx_messageInfo_PinInfo.Size(m)
}
func (m *PinInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PinInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PinInfo proto.InternalMessageInfo

func (m *PinInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PinInfo) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *PinInfo) GetCaps() int32 {
	if m != nil {
		return m.Caps
	}
	return 0
}

func (m *PinInfo) GetDigitalLogical() int32 {
	if m != nil {
		return m.DigitalLogical
	}
	return 0
}

func (m *PinInfo) GetAnalogLogical() int32 {
	if m != nil {
		return m.AnalogLogical
	}
	return 0
}

type InfoRes struct {
	Info                 []*PinInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	In                   int32      `protobuf:"varint,2,opt,name=in,proto3" json:"in,omitempty"`
	Out                  int32      `protobuf:"varint,3,opt,name=out,proto3" json:"out,omitempty"`
	Low                  int32      `protobuf:"varint,4,opt,name=low,proto3" json:"low,omitempty"`
	High                 int32      `protobuf:"varint,5,opt,name=high,proto3" json:"high,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InfoRes) Reset()         { *m = InfoRes{} }
func (m *InfoRes) String() string { return proto.CompactTextString(m) }
func (*InfoRes) ProtoMessage()    {}
func (*InfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{4}
}

func (m *InfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoRes.Unmarshal(m, b)
}
func (m *InfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoRes.Marshal(b, m, deterministic)
}
func (m *InfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRes.Merge(m, src)
}
func (m *InfoRes) XXX_Size() int {
	return xxx_messageInfo_InfoRes.Size(m)
}
func (m *InfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRes proto.InternalMessageInfo

func (m *InfoRes) GetInfo() []*PinInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *InfoRes) GetIn() int32 {
	if m != nil {
		return m.In
	}
	return 0
}

func (m *InfoRes) GetOut() int32 {
	if m != nil {
		return m.Out
	}
	return 0
}

func (m *InfoRes) GetLow() int32 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *InfoRes) GetHigh() int32 {
	if m != nil {
		return m.High
	}
	return 0
}

type DigitalWriteReq struct {
	Pin                  string   `protobuf:"bytes,1,opt,name=pin,proto3" json:"pin,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DigitalWriteReq) Reset()         { *m = DigitalWriteReq{} }
func (m *DigitalWriteReq) String() string { return proto.CompactTextString(m) }
func (*DigitalWriteReq) ProtoMessage()    {}
func (*DigitalWriteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{5}
}

func (m *DigitalWriteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DigitalWriteReq.Unmarshal(m, b)
}
func (m *DigitalWriteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DigitalWriteReq.Marshal(b, m, deterministic)
}
func (m *DigitalWriteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DigitalWriteReq.Merge(m, src)
}
func (m *DigitalWriteReq) XXX_Size() int {
	return xxx_messageInfo_DigitalWriteReq.Size(m)
}
func (m *DigitalWriteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DigitalWriteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DigitalWriteReq proto.InternalMessageInfo

func (m *DigitalWriteReq) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

func (m *DigitalWriteReq) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type PinReadReq struct {
	Pin                  string   `protobuf:"bytes,1,opt,name=pin,proto3" json:"pin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PinReadReq) Reset()         { *m = PinReadReq{} }
func (m *PinReadReq) String() string { return proto.CompactTextString(m) }
func (*PinReadReq) ProtoMessage()    {}
func (*PinReadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{6}
}

func (m *PinReadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PinReadReq.Unmarshal(m, b)
}
func (m *PinReadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PinReadReq.Marshal(b, m, deterministic)
}
func (m *PinReadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PinReadReq.Merge(m, src)
}
func (m *PinReadReq) XXX_Size() int {
	return xxx_messageInfo_PinReadReq.Size(m)
}
func (m *PinReadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PinReadReq.DiscardUnknown(m)
}

var xxx_messageInfo_PinReadReq proto.InternalMessageInfo

func (m *PinReadReq) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

type PinReadRes struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PinReadRes) Reset()         { *m = PinReadRes{} }
func (m *PinReadRes) String() string { return proto.CompactTextString(m) }
func (*PinReadRes) ProtoMessage()    {}
func (*PinReadRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_59fedb88b556689a, []int{7}
}

func (m *PinReadRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PinReadRes.Unmarshal(m, b)
}
func (m *PinReadRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PinReadRes.Marshal(b, m, deterministic)
}
func (m *PinReadRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PinReadRes.Merge(m, src)
}
func (m *PinReadRes) XXX_Size() int {
	return xxx_messageInfo_PinReadRes.Size(m)
}
func (m *PinReadRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PinReadRes.DiscardUnknown(m)
}

var xxx_messageInfo_PinReadRes proto.InternalMessageInfo

func (m *PinReadRes) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*InitReq)(nil), "rpi.InitReq")
	proto.RegisterType((*CloseReq)(nil), "rpi.CloseReq")
	proto.RegisterType((*SetDirectionReq)(nil), "rpi.SetDirectionReq")
	proto.RegisterType((*PinInfo)(nil), "rpi.PinInfo")
	proto.RegisterType((*InfoRes)(nil), "rpi.InfoRes")
	proto.RegisterType((*DigitalWriteReq)(nil), "rpi.DigitalWriteReq")
	proto.RegisterType((*PinReadReq)(nil), "rpi.PinReadReq")
	proto.RegisterType((*PinReadRes)(nil), "rpi.PinReadRes")
}

func init() { proto.RegisterFile("gpio.proto", fileDescriptor_59fedb88b556689a) }

var fileDescriptor_59fedb88b556689a = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x9d, 0x7c, 0xb9, 0xa6, 0x76, 0xdc, 0x91, 0x62, 0x0e, 0x61, 0x90, 0x35, 0x36, 0x22, 0x39,
	0x0d, 0xec, 0x7a, 0xf2, 0x28, 0x2e, 0xc8, 0x80, 0xe0, 0xd0, 0x82, 0x9e, 0x7b, 0x27, 0x9d, 0x6c,
	0x41, 0xe8, 0xce, 0xa6, 0xb3, 0xfa, 0x03, 0xfc, 0x05, 0xfe, 0x63, 0xe9, 0x4e, 0x87, 0xc9, 0x84,
	0xd9, 0xdb, 0xeb, 0xf7, 0xaa, 0xaa, 0xab, 0x5e, 0x15, 0x40, 0xdd, 0x92, 0xde, 0xb6, 0x9d, 0xee,
	0x35, 0x46, 0x5d, 0x4b, 0x1b, 0xb8, 0x17, 0x46, 0x0e, 0x04, 0x4b, 0xe1, 0x62, 0xa7, 0xa8, 0xe7,
	0xf2, 0x91, 0x01, 0xbc, 0xfc, 0xd2, 0x68, 0x23, 0x2d, 0xfe, 0x0c, 0xab, 0x1f, 0xb2, 0xbf, 0xa3,
	0x4e, 0x1e, 0x7a, 0xd2, 0x8a, 0xcb, 0x47, 0x7c, 0x0d, 0x51, 0x4b, 0x2a, 0x0b, 0xf2, 0xa0, 0x48,
	0xb9, 0x85, 0xf8, 0x06, 0xd2, 0x72, 0x8c, 0xc8, 0xc2, 0x3c, 0x28, 0x12, 0x7e, 0x24, 0xd8, 0xbf,
	0x00, 0x2e, 0xf6, 0xa4, 0x76, 0xaa, 0xd2, 0x78, 0x05, 0x21, 0x95, 0x3e, 0x35, 0xa4, 0x12, 0xd7,
	0x90, 0x88, 0x86, 0x84, 0xc9, 0xc2, 0x3c, 0x2a, 0x52, 0x3e, 0x3c, 0x10, 0x21, 0x3e, 0x88, 0xd6,
	0x64, 0x91, 0x2b, 0xe5, 0x30, 0x7e, 0x80, 0xab, 0x92, 0x6a, 0xea, 0x45, 0xf3, 0x4d, 0xd7, 0x74,
	0x10, 0x4d, 0x16, 0x3b, 0x75, 0xc6, 0xe2, 0x7b, 0x78, 0x25, 0x94, 0x68, 0x74, 0x3d, 0x86, 0x25,
	0x2e, 0xec, 0x94, 0x64, 0xc6, 0x4e, 0x5b, 0x69, 0x2e, 0x0d, 0xe6, 0x10, 0x93, 0xaa, 0x74, 0x16,
	0xe4, 0x51, 0x71, 0x79, 0xbb, 0xdc, 0x76, 0x2d, 0x6d, 0x7d, 0xbb, 0xdc, 0x29, 0xae, 0xe9, 0x71,
	0xae, 0x90, 0x94, 0x35, 0x40, 0x3f, 0xf5, 0xbe, 0x3b, 0x0b, 0x2d, 0xd3, 0xe8, 0x3f, 0xbe, 0x23,
	0x0b, 0xed, 0x08, 0x0f, 0x54, 0x3f, 0xf8, 0xdf, 0x1d, 0x66, 0x9f, 0x60, 0x75, 0x37, 0x34, 0xfb,
	0xab, 0xa3, 0x5e, 0x9e, 0xf7, 0x72, 0x0d, 0xc9, 0x6f, 0xd1, 0x3c, 0x49, 0xff, 0xdf, 0xf0, 0x60,
	0xd7, 0x00, 0x7b, 0x52, 0x5c, 0x8a, 0xf2, 0x6c, 0x16, 0x63, 0x13, 0xdd, 0x1c, 0x6b, 0x04, 0x93,
	0x1a, 0xb7, 0x7f, 0x43, 0x88, 0xbf, 0xee, 0x77, 0xdf, 0xf1, 0x1a, 0x62, 0xbb, 0x6a, 0x4c, 0xdd,
	0xac, 0x3f, 0x35, 0x95, 0x9b, 0x23, 0x64, 0x0b, 0x7c, 0x0b, 0x89, 0xdb, 0xff, 0xb3, 0x01, 0xef,
	0x6c, 0x81, 0x4a, 0x4f, 0xf5, 0xc1, 0x37, 0xef, 0x29, 0x5b, 0xe0, 0x0d, 0x2c, 0xa7, 0x77, 0x83,
	0x6b, 0xa7, 0xcf, 0x4e, 0xe9, 0xb4, 0xea, 0x0d, 0x2c, 0xa7, 0xf6, 0xf8, 0x94, 0x99, 0x63, 0xf3,
	0x94, 0x4b, 0xaf, 0xdb, 0xd1, 0x71, 0x35, 0x2e, 0xcf, 0x1b, 0xb5, 0x99, 0x11, 0x86, 0x2d, 0xee,
	0x5f, 0xb8, 0x73, 0xff, 0xf8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x38, 0x82, 0xb0, 0x0d, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GPIOClient is the client API for GPIO service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GPIOClient interface {
	Init(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	Close(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	Info(ctx context.Context, in *Void, opts ...grpc.CallOption) (*InfoRes, error)
	SetDirection(ctx context.Context, in *SetDirectionReq, opts ...grpc.CallOption) (*Void, error)
	DigitalWrite(ctx context.Context, in *DigitalWriteReq, opts ...grpc.CallOption) (*Void, error)
	DigitalRead(ctx context.Context, in *PinReadReq, opts ...grpc.CallOption) (*PinReadRes, error)
}

type gPIOClient struct {
	cc *grpc.ClientConn
}

func NewGPIOClient(cc *grpc.ClientConn) GPIOClient {
	return &gPIOClient{cc}
}

func (c *gPIOClient) Init(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/rpi.GPIO/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPIOClient) Close(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/rpi.GPIO/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPIOClient) Info(ctx context.Context, in *Void, opts ...grpc.CallOption) (*InfoRes, error) {
	out := new(InfoRes)
	err := c.cc.Invoke(ctx, "/rpi.GPIO/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPIOClient) SetDirection(ctx context.Context, in *SetDirectionReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/rpi.GPIO/SetDirection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPIOClient) DigitalWrite(ctx context.Context, in *DigitalWriteReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/rpi.GPIO/DigitalWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPIOClient) DigitalRead(ctx context.Context, in *PinReadReq, opts ...grpc.CallOption) (*PinReadRes, error) {
	out := new(PinReadRes)
	err := c.cc.Invoke(ctx, "/rpi.GPIO/DigitalRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GPIOServer is the server API for GPIO service.
type GPIOServer interface {
	Init(context.Context, *Void) (*Void, error)
	Close(context.Context, *Void) (*Void, error)
	Info(context.Context, *Void) (*InfoRes, error)
	SetDirection(context.Context, *SetDirectionReq) (*Void, error)
	DigitalWrite(context.Context, *DigitalWriteReq) (*Void, error)
	DigitalRead(context.Context, *PinReadReq) (*PinReadRes, error)
}

func RegisterGPIOServer(s *grpc.Server, srv GPIOServer) {
	s.RegisterService(&_GPIO_serviceDesc, srv)
}

func _GPIO_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPIOServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.GPIO/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPIOServer).Init(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPIO_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPIOServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.GPIO/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPIOServer).Close(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPIO_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPIOServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.GPIO/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPIOServer).Info(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPIO_SetDirection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDirectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPIOServer).SetDirection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.GPIO/SetDirection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPIOServer).SetDirection(ctx, req.(*SetDirectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPIO_DigitalWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DigitalWriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPIOServer).DigitalWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.GPIO/DigitalWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPIOServer).DigitalWrite(ctx, req.(*DigitalWriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPIO_DigitalRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PinReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPIOServer).DigitalRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpi.GPIO/DigitalRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPIOServer).DigitalRead(ctx, req.(*PinReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GPIO_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpi.GPIO",
	HandlerType: (*GPIOServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _GPIO_Init_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _GPIO_Close_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _GPIO_Info_Handler,
		},
		{
			MethodName: "SetDirection",
			Handler:    _GPIO_SetDirection_Handler,
		},
		{
			MethodName: "DigitalWrite",
			Handler:    _GPIO_DigitalWrite_Handler,
		},
		{
			MethodName: "DigitalRead",
			Handler:    _GPIO_DigitalRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gpio.proto",
}
