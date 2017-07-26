// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oidmc_moclient.proto

/*
Package oidmc_moclient is a generated protocol buffer package.

It is generated from these files:
	oidmc_moclient.proto

It has these top-level messages:
	MoClientNotifyOrderReq
	MoClientNotifyOrderRes
*/
package oidmc_moclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MoClientNotifyOrderReq struct {
	OrderId  uint32 `protobuf:"varint,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	UserId   uint32 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Sign     string `protobuf:"bytes,4,opt,name=sign" json:"sign,omitempty"`
}

func (m *MoClientNotifyOrderReq) Reset()                    { *m = MoClientNotifyOrderReq{} }
func (m *MoClientNotifyOrderReq) String() string            { return proto.CompactTextString(m) }
func (*MoClientNotifyOrderReq) ProtoMessage()               {}
func (*MoClientNotifyOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MoClientNotifyOrderReq) GetOrderId() uint32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *MoClientNotifyOrderReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MoClientNotifyOrderReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *MoClientNotifyOrderReq) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

type MoClientNotifyOrderRes struct {
	Err    uint32 `protobuf:"varint,1,opt,name=err" json:"err,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
}

func (m *MoClientNotifyOrderRes) Reset()                    { *m = MoClientNotifyOrderRes{} }
func (m *MoClientNotifyOrderRes) String() string            { return proto.CompactTextString(m) }
func (*MoClientNotifyOrderRes) ProtoMessage()               {}
func (*MoClientNotifyOrderRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MoClientNotifyOrderRes) GetErr() uint32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *MoClientNotifyOrderRes) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*MoClientNotifyOrderReq)(nil), "oidmc.moclient.MoClientNotifyOrderReq")
	proto.RegisterType((*MoClientNotifyOrderRes)(nil), "oidmc.moclient.MoClientNotifyOrderRes")
}

func init() { proto.RegisterFile("oidmc_moclient.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xcf, 0x4c, 0xc9,
	0x4d, 0x8e, 0xcf, 0xcd, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x03, 0x8b, 0xea, 0xc1, 0x44, 0x95, 0x6a, 0xb9, 0xc4, 0x7c, 0xf3, 0x9d, 0xc1, 0x6c,
	0xbf, 0xfc, 0x92, 0xcc, 0xb4, 0x4a, 0xff, 0xa2, 0x94, 0xd4, 0xa2, 0xa0, 0xd4, 0x42, 0x21, 0x49,
	0x2e, 0x8e, 0x7c, 0x10, 0x3b, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x37, 0x88, 0x1d,
	0xcc, 0xf7, 0x4c, 0x11, 0x12, 0xe7, 0x62, 0x2f, 0x2d, 0x86, 0xc8, 0x30, 0x81, 0x65, 0xd8, 0x40,
	0x5c, 0xcf, 0x14, 0x21, 0x69, 0x2e, 0x4e, 0xb0, 0x44, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xb3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x07, 0x48, 0xc0, 0x2f, 0x31, 0x37, 0x55, 0x48, 0x88, 0x8b, 0xa5, 0x38,
	0x33, 0x3d, 0x4f, 0x82, 0x05, 0x2c, 0x0e, 0x66, 0x2b, 0x39, 0xe3, 0xb0, 0xbe, 0x58, 0x48, 0x80,
	0x8b, 0x39, 0xb5, 0xa8, 0x08, 0x6a, 0x33, 0x88, 0x09, 0xb2, 0x35, 0xb5, 0xa8, 0x28, 0x3e, 0xb7,
	0x38, 0x1d, 0x6c, 0x2b, 0x67, 0x10, 0x5b, 0x6a, 0x51, 0x91, 0x6f, 0x71, 0xba, 0x13, 0x93, 0x07,
	0x73, 0x12, 0x1b, 0xd8, 0x7b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x54, 0xd6, 0x65,
	0xf6, 0x00, 0x00, 0x00,
}
