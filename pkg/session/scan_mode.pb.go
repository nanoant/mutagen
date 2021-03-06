// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/scan_mode.proto

package session

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// ScanMode specifies the mode for synchronization root scanning.
type ScanMode int32

const (
	// ScanMode_ScanModeDefault represents an unspecified scan mode. It should
	// be converted to one of the following values based on the desired default
	// behavior.
	ScanMode_ScanModeDefault ScanMode = 0
	// ScanMode_ScanModeFull specifies that full scans should be performed on
	// each synchronization cycle.
	ScanMode_ScanModeFull ScanMode = 1
	// ScanMode_ScanModeAccelerated specifies that scans should attempt to use
	// watch-based acceleration.
	ScanMode_ScanModeAccelerated ScanMode = 2
)

var ScanMode_name = map[int32]string{
	0: "ScanModeDefault",
	1: "ScanModeFull",
	2: "ScanModeAccelerated",
}

var ScanMode_value = map[string]int32{
	"ScanModeDefault":     0,
	"ScanModeFull":        1,
	"ScanModeAccelerated": 2,
}

func (x ScanMode) String() string {
	return proto.EnumName(ScanMode_name, int32(x))
}

func (ScanMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0fe26fc659b8fc39, []int{0}
}

func init() {
	proto.RegisterEnum("session.ScanMode", ScanMode_name, ScanMode_value)
}

func init() { proto.RegisterFile("session/scan_mode.proto", fileDescriptor_0fe26fc659b8fc39) }

var fileDescriptor_0fe26fc659b8fc39 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x4e, 0x4e, 0xcc, 0x8b, 0xcf, 0xcd, 0x4f, 0x49, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x68, 0x79, 0x71, 0x71, 0x04, 0x27, 0x27, 0xe6, 0xf9,
	0xe6, 0xa7, 0xa4, 0x0a, 0x09, 0x73, 0xf1, 0xc3, 0xd8, 0x2e, 0xa9, 0x69, 0x89, 0xa5, 0x39, 0x25,
	0x02, 0x0c, 0x42, 0x02, 0x5c, 0x3c, 0x30, 0x41, 0xb7, 0xd2, 0x9c, 0x1c, 0x01, 0x46, 0x21, 0x71,
	0x2e, 0x61, 0x98, 0x88, 0x63, 0x72, 0x72, 0x6a, 0x4e, 0x6a, 0x51, 0x62, 0x49, 0x6a, 0x8a, 0x00,
	0x93, 0x93, 0x66, 0x94, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x46, 0x62, 0x59, 0x7e, 0xb2, 0x6e, 0x66, 0xbe, 0x7e, 0x6e, 0x69, 0x49, 0x62, 0x7a, 0x6a, 0x9e,
	0x7e, 0x41, 0x76, 0xba, 0x3e, 0xd4, 0xda, 0x24, 0x36, 0xb0, 0x33, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0x25, 0x51, 0xdf, 0xa1, 0x00, 0x00, 0x00,
}
