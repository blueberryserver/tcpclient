// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package MSG

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// keep alive
type PingReq struct {
	SessionKey       *string `protobuf:"bytes,1,req,name=sessionKey" json:"sessionKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PingReq) Reset()                    { *m = PingReq{} }
func (m *PingReq) String() string            { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()               {}
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PingReq) GetSessionKey() string {
	if m != nil && m.SessionKey != nil {
		return *m.SessionKey
	}
	return ""
}

type PongAns struct {
	Err              *ErrorCode `protobuf:"varint,1,req,name=err,enum=MSG.ErrorCode" json:"err,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *PongAns) Reset()                    { *m = PongAns{} }
func (m *PongAns) String() string            { return proto.CompactTextString(m) }
func (*PongAns) ProtoMessage()               {}
func (*PongAns) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PongAns) GetErr() ErrorCode {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ErrorCode_ERR_SUCCESS
}

// login request
type LoginReq struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *LoginReq) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// login answer
type LoginAns struct {
	Err              *ErrorCode `protobuf:"varint,1,req,name=err,enum=MSG.ErrorCode" json:"err,omitempty"`
	SessionKey       *string    `protobuf:"bytes,2,opt,name=sessionKey" json:"sessionKey,omitempty"`
	Data             *UserData_ `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *LoginAns) Reset()                    { *m = LoginAns{} }
func (m *LoginAns) String() string            { return proto.CompactTextString(m) }
func (*LoginAns) ProtoMessage()               {}
func (*LoginAns) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *LoginAns) GetErr() ErrorCode {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ErrorCode_ERR_SUCCESS
}

func (m *LoginAns) GetSessionKey() string {
	if m != nil && m.SessionKey != nil {
		return *m.SessionKey
	}
	return ""
}

func (m *LoginAns) GetData() *UserData_ {
	if m != nil {
		return m.Data
	}
	return nil
}

// regist request
type RegistReq struct {
	Data             *UserData_ `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RegistReq) Reset()                    { *m = RegistReq{} }
func (m *RegistReq) String() string            { return proto.CompactTextString(m) }
func (*RegistReq) ProtoMessage()               {}
func (*RegistReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *RegistReq) GetData() *UserData_ {
	if m != nil {
		return m.Data
	}
	return nil
}

// regist answer
type RegistAns struct {
	Err              *ErrorCode `protobuf:"varint,1,req,name=err,enum=MSG.ErrorCode" json:"err,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RegistAns) Reset()                    { *m = RegistAns{} }
func (m *RegistAns) String() string            { return proto.CompactTextString(m) }
func (*RegistAns) ProtoMessage()               {}
func (*RegistAns) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *RegistAns) GetErr() ErrorCode {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ErrorCode_ERR_SUCCESS
}

// version
type VersionReq struct {
	Version          *string `protobuf:"bytes,1,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VersionReq) Reset()                    { *m = VersionReq{} }
func (m *VersionReq) String() string            { return proto.CompactTextString(m) }
func (*VersionReq) ProtoMessage()               {}
func (*VersionReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *VersionReq) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type VersionAns struct {
	Err *ErrorCode `protobuf:"varint,1,req,name=err,enum=MSG.ErrorCode" json:"err,omitempty"`
	// available contents list
	Contents         []*Contents_ `protobuf:"bytes,2,rep,name=contents" json:"contents,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *VersionAns) Reset()                    { *m = VersionAns{} }
func (m *VersionAns) String() string            { return proto.CompactTextString(m) }
func (*VersionAns) ProtoMessage()               {}
func (*VersionAns) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *VersionAns) GetErr() ErrorCode {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ErrorCode_ERR_SUCCESS
}

func (m *VersionAns) GetContents() []*Contents_ {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*PingReq)(nil), "MSG.PingReq")
	proto.RegisterType((*PongAns)(nil), "MSG.PongAns")
	proto.RegisterType((*LoginReq)(nil), "MSG.LoginReq")
	proto.RegisterType((*LoginAns)(nil), "MSG.LoginAns")
	proto.RegisterType((*RegistReq)(nil), "MSG.RegistReq")
	proto.RegisterType((*RegistAns)(nil), "MSG.RegistAns")
	proto.RegisterType((*VersionReq)(nil), "MSG.VersionReq")
	proto.RegisterType((*VersionAns)(nil), "MSG.VersionAns")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x4f, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xa6, 0xed, 0xa0, 0xdb, 0x9b, 0xec, 0x90, 0x53, 0xf1, 0x30, 0x42, 0x0e, 0x32, 0x15, 0x2b,
	0xf8, 0x1f, 0xc8, 0x94, 0x1d, 0x74, 0x30, 0x32, 0xf4, 0xe0, 0x65, 0x84, 0xed, 0x51, 0x0a, 0x36,
	0x6f, 0xbe, 0x04, 0xc1, 0xff, 0x5e, 0x52, 0xd3, 0xa2, 0xc2, 0xa0, 0xb7, 0xe4, 0xfb, 0xde, 0xf7,
	0x0b, 0xa6, 0xef, 0x54, 0xd5, 0xb6, 0x3c, 0x32, 0x79, 0x12, 0xd9, 0x7a, 0xbb, 0x3a, 0x3f, 0xdb,
	0x53, 0xd3, 0x50, 0x84, 0xd4, 0x25, 0xe4, 0x9b, 0xda, 0x56, 0x1a, 0x3f, 0xc4, 0x1c, 0xc0, 0xa1,
	0x73, 0x35, 0xd9, 0x27, 0xfc, 0x2a, 0x12, 0x99, 0x2e, 0x26, 0xfa, 0x17, 0xa2, 0xae, 0x21, 0xdf,
	0x90, 0xad, 0xee, 0xad, 0x13, 0x12, 0x32, 0x64, 0x6e, 0x6f, 0x66, 0x77, 0xb3, 0x72, 0xbd, 0x5d,
	0x95, 0x8f, 0xcc, 0xc4, 0x4b, 0x3a, 0xa0, 0x0e, 0x94, 0x9a, 0xc3, 0xf8, 0x39, 0x24, 0x07, 0x63,
	0x01, 0x23, 0x6b, 0x1a, 0x8c, 0x96, 0xed, 0x5b, 0x1d, 0x23, 0x3f, 0xc8, 0xed, 0x5f, 0xb5, 0x54,
	0x26, 0x7f, 0xab, 0x09, 0x05, 0xa3, 0x83, 0xf1, 0xa6, 0xc8, 0x64, 0xb2, 0x98, 0x46, 0x8b, 0x17,
	0x87, 0xfc, 0x60, 0xbc, 0xd9, 0xe9, 0x96, 0x53, 0xb7, 0x30, 0xd1, 0x58, 0xd5, 0xce, 0x87, 0x4a,
	0x9d, 0x20, 0x64, 0x9e, 0x12, 0xdc, 0x74, 0x82, 0x61, 0x8b, 0x2f, 0x00, 0x5e, 0x91, 0x43, 0xa3,
	0x10, 0x50, 0x40, 0xfe, 0xf9, 0xf3, 0x8b, 0xb3, 0xbb, 0xaf, 0x7a, 0xeb, 0xef, 0x86, 0x6d, 0xbf,
	0x82, 0xf1, 0x9e, 0xac, 0x47, 0xeb, 0x5d, 0x91, 0xca, 0xac, 0xaf, 0xbb, 0x8c, 0xe0, 0x4e, 0xf7,
	0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xa5, 0x0f, 0xdc, 0xee, 0x01, 0x00, 0x00,
}
