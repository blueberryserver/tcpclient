// Code generated by protoc-gen-go.
// source: common.proto
// DO NOT EDIT!

package MSG

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// message type
type MsgId int32

const (
	MsgId_CLOSED             MsgId = 10000
	MsgId_LOGIN_REQ          MsgId = 10101
	MsgId_LOGIN_ANS          MsgId = 10102
	MsgId_PING_REQ           MsgId = 10103
	MsgId_PONG_ANS           MsgId = 10104
	MsgId_REGIST_REQ         MsgId = 10105
	MsgId_REGIST_ANS         MsgId = 10106
	MsgId_VERSION_REQ        MsgId = 10107
	MsgId_VERSION_ANS        MsgId = 10108
	MsgId_CHAT_REQ           MsgId = 20101
	MsgId_CHAT_ANS           MsgId = 20102
	MsgId_CHAT_NOT           MsgId = 20103
	MsgId_CREATECHATROOM_REQ MsgId = 20111
	MsgId_CREATECHATROOM_ANS MsgId = 20112
	MsgId_CREATECHATROOM_NOT MsgId = 20113
	MsgId_INVITECHATROOM_REQ MsgId = 20121
	MsgId_INVITECHATROOM_ANS MsgId = 20122
	MsgId_INVITECHATROOM_NOT MsgId = 20123
	MsgId_ENTERCHATROOM_REQ  MsgId = 20131
	MsgId_ENTERCHATROOM_ANS  MsgId = 20132
	MsgId_ENTERCHATROOM_NOT  MsgId = 20133
	MsgId_LEAVECHATROOM_REQ  MsgId = 20141
	MsgId_LEAVECHATROOM_ANS  MsgId = 20142
	MsgId_LEAVECHATROOM_NOT  MsgId = 20143
	MsgId_CREATECHAR_REQ     MsgId = 20151
	MsgId_CREATECHAR_ANS     MsgId = 20152
	MsgId_CONTENTS_NOT       MsgId = 20161
	MsgId_OWNEDCHARACTER_NOT MsgId = 20162
	MsgId_CURRENCY_NOT       MsgId = 20163
	MsgId_PLAYDUNGEON_REQ    MsgId = 20171
	MsgId_PLAYDUNGEON_ANS    MsgId = 20172
	MsgId_PLAYDUNGEON_NOT    MsgId = 20173
	MsgId_LEVELUPCHAR_REQ    MsgId = 20181
	MsgId_LEVELUPCHAR_ANS    MsgId = 20182
	MsgId_TIERUPCHAR_REQ     MsgId = 20191
	MsgId_TIERUPCHAR_ANS     MsgId = 20192
)

var MsgId_name = map[int32]string{
	10000: "CLOSED",
	10101: "LOGIN_REQ",
	10102: "LOGIN_ANS",
	10103: "PING_REQ",
	10104: "PONG_ANS",
	10105: "REGIST_REQ",
	10106: "REGIST_ANS",
	10107: "VERSION_REQ",
	10108: "VERSION_ANS",
	20101: "CHAT_REQ",
	20102: "CHAT_ANS",
	20103: "CHAT_NOT",
	20111: "CREATECHATROOM_REQ",
	20112: "CREATECHATROOM_ANS",
	20113: "CREATECHATROOM_NOT",
	20121: "INVITECHATROOM_REQ",
	20122: "INVITECHATROOM_ANS",
	20123: "INVITECHATROOM_NOT",
	20131: "ENTERCHATROOM_REQ",
	20132: "ENTERCHATROOM_ANS",
	20133: "ENTERCHATROOM_NOT",
	20141: "LEAVECHATROOM_REQ",
	20142: "LEAVECHATROOM_ANS",
	20143: "LEAVECHATROOM_NOT",
	20151: "CREATECHAR_REQ",
	20152: "CREATECHAR_ANS",
	20161: "CONTENTS_NOT",
	20162: "OWNEDCHARACTER_NOT",
	20163: "CURRENCY_NOT",
	20171: "PLAYDUNGEON_REQ",
	20172: "PLAYDUNGEON_ANS",
	20173: "PLAYDUNGEON_NOT",
	20181: "LEVELUPCHAR_REQ",
	20182: "LEVELUPCHAR_ANS",
	20191: "TIERUPCHAR_REQ",
	20192: "TIERUPCHAR_ANS",
}
var MsgId_value = map[string]int32{
	"CLOSED":             10000,
	"LOGIN_REQ":          10101,
	"LOGIN_ANS":          10102,
	"PING_REQ":           10103,
	"PONG_ANS":           10104,
	"REGIST_REQ":         10105,
	"REGIST_ANS":         10106,
	"VERSION_REQ":        10107,
	"VERSION_ANS":        10108,
	"CHAT_REQ":           20101,
	"CHAT_ANS":           20102,
	"CHAT_NOT":           20103,
	"CREATECHATROOM_REQ": 20111,
	"CREATECHATROOM_ANS": 20112,
	"CREATECHATROOM_NOT": 20113,
	"INVITECHATROOM_REQ": 20121,
	"INVITECHATROOM_ANS": 20122,
	"INVITECHATROOM_NOT": 20123,
	"ENTERCHATROOM_REQ":  20131,
	"ENTERCHATROOM_ANS":  20132,
	"ENTERCHATROOM_NOT":  20133,
	"LEAVECHATROOM_REQ":  20141,
	"LEAVECHATROOM_ANS":  20142,
	"LEAVECHATROOM_NOT":  20143,
	"CREATECHAR_REQ":     20151,
	"CREATECHAR_ANS":     20152,
	"CONTENTS_NOT":       20161,
	"OWNEDCHARACTER_NOT": 20162,
	"CURRENCY_NOT":       20163,
	"PLAYDUNGEON_REQ":    20171,
	"PLAYDUNGEON_ANS":    20172,
	"PLAYDUNGEON_NOT":    20173,
	"LEVELUPCHAR_REQ":    20181,
	"LEVELUPCHAR_ANS":    20182,
	"TIERUPCHAR_REQ":     20191,
	"TIERUPCHAR_ANS":     20192,
}

func (x MsgId) Enum() *MsgId {
	p := new(MsgId)
	*p = x
	return p
}
func (x MsgId) String() string {
	return proto.EnumName(MsgId_name, int32(x))
}
func (x *MsgId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgId_value, data, "MsgId")
	if err != nil {
		return err
	}
	*x = MsgId(value)
	return nil
}
func (MsgId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Error type
type ErrorCode int32

const (
	ErrorCode_ERR_SUCCESS         ErrorCode = 0
	ErrorCode_ERR_LOGIN_FAIL      ErrorCode = 1
	ErrorCode_ERR_ARGUMENT_FAIL   ErrorCode = 2
	ErrorCode_ERR_AUTHORITY_FAIL  ErrorCode = 3
	ErrorCode_ERR_SESSIONKEY_FAIL ErrorCode = 4
)

var ErrorCode_name = map[int32]string{
	0: "ERR_SUCCESS",
	1: "ERR_LOGIN_FAIL",
	2: "ERR_ARGUMENT_FAIL",
	3: "ERR_AUTHORITY_FAIL",
	4: "ERR_SESSIONKEY_FAIL",
}
var ErrorCode_value = map[string]int32{
	"ERR_SUCCESS":         0,
	"ERR_LOGIN_FAIL":      1,
	"ERR_ARGUMENT_FAIL":   2,
	"ERR_AUTHORITY_FAIL":  3,
	"ERR_SESSIONKEY_FAIL": 4,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// platform type
type PlatForm int32

const (
	PlatForm_IOS     PlatForm = 0
	PlatForm_ANDROID PlatForm = 1
)

var PlatForm_name = map[int32]string{
	0: "IOS",
	1: "ANDROID",
}
var PlatForm_value = map[string]int32{
	"IOS":     0,
	"ANDROID": 1,
}

func (x PlatForm) Enum() *PlatForm {
	p := new(PlatForm)
	*p = x
	return p
}
func (x PlatForm) String() string {
	return proto.EnumName(PlatForm_name, int32(x))
}
func (x *PlatForm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlatForm_value, data, "PlatForm")
	if err != nil {
		return err
	}
	*x = PlatForm(value)
	return nil
}
func (PlatForm) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// chat type
type ChatType int32

const (
	ChatType_CHAT_CHANNEL ChatType = 1
	ChatType_CHAT_GROUP   ChatType = 2
	ChatType_CHAT_ROOM    ChatType = 3
)

var ChatType_name = map[int32]string{
	1: "CHAT_CHANNEL",
	2: "CHAT_GROUP",
	3: "CHAT_ROOM",
}
var ChatType_value = map[string]int32{
	"CHAT_CHANNEL": 1,
	"CHAT_GROUP":   2,
	"CHAT_ROOM":    3,
}

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}
func (x ChatType) String() string {
	return proto.EnumName(ChatType_name, int32(x))
}
func (x *ChatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChatType_value, data, "ChatType")
	if err != nil {
		return err
	}
	*x = ChatType(value)
	return nil
}
func (ChatType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// contents type
type Contents__ContentType int32

const (
	// ingame
	Contents__CT_INGAME Contents__ContentType = 1
	// outgame
	Contents__CT_OUTGAME Contents__ContentType = 2
	// in app purchase
	Contents__CT_INAPPBUY Contents__ContentType = 3
)

var Contents__ContentType_name = map[int32]string{
	1: "CT_INGAME",
	2: "CT_OUTGAME",
	3: "CT_INAPPBUY",
}
var Contents__ContentType_value = map[string]int32{
	"CT_INGAME":   1,
	"CT_OUTGAME":  2,
	"CT_INAPPBUY": 3,
}

func (x Contents__ContentType) Enum() *Contents__ContentType {
	p := new(Contents__ContentType)
	*p = x
	return p
}
func (x Contents__ContentType) String() string {
	return proto.EnumName(Contents__ContentType_name, int32(x))
}
func (x *Contents__ContentType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Contents__ContentType_value, data, "Contents__ContentType")
	if err != nil {
		return err
	}
	*x = Contents__ContentType(value)
	return nil
}
func (Contents__ContentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type GMember__GradeType int32

const (
	GMember__Grade_1 GMember__GradeType = 1
	GMember__Grade_2 GMember__GradeType = 2
	GMember__Grade_3 GMember__GradeType = 3
	GMember__Grade_4 GMember__GradeType = 4
	GMember__Grade_5 GMember__GradeType = 5
)

var GMember__GradeType_name = map[int32]string{
	1: "Grade_1",
	2: "Grade_2",
	3: "Grade_3",
	4: "Grade_4",
	5: "Grade_5",
}
var GMember__GradeType_value = map[string]int32{
	"Grade_1": 1,
	"Grade_2": 2,
	"Grade_3": 3,
	"Grade_4": 4,
	"Grade_5": 5,
}

func (x GMember__GradeType) Enum() *GMember__GradeType {
	p := new(GMember__GradeType)
	*p = x
	return p
}
func (x GMember__GradeType) String() string {
	return proto.EnumName(GMember__GradeType_name, int32(x))
}
func (x *GMember__GradeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GMember__GradeType_value, data, "GMember__GradeType")
	if err != nil {
		return err
	}
	*x = GMember__GradeType(value)
	return nil
}
func (GMember__GradeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

type BattleData__AttackResult int32

const (
	BattleData__ALIVE BattleData__AttackResult = 0
	BattleData__DEAD  BattleData__AttackResult = 1
)

var BattleData__AttackResult_name = map[int32]string{
	0: "ALIVE",
	1: "DEAD",
}
var BattleData__AttackResult_value = map[string]int32{
	"ALIVE": 0,
	"DEAD":  1,
}

func (x BattleData__AttackResult) Enum() *BattleData__AttackResult {
	p := new(BattleData__AttackResult)
	*p = x
	return p
}
func (x BattleData__AttackResult) String() string {
	return proto.EnumName(BattleData__AttackResult_name, int32(x))
}
func (x *BattleData__AttackResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BattleData__AttackResult_value, data, "BattleData__AttackResult")
	if err != nil {
		return err
	}
	*x = BattleData__AttackResult(value)
	return nil
}
func (BattleData__AttackResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{9, 0} }

type BattleData__Team int32

const (
	BattleData__ALLY  BattleData__Team = 0
	BattleData__ENEMY BattleData__Team = 1
)

var BattleData__Team_name = map[int32]string{
	0: "ALLY",
	1: "ENEMY",
}
var BattleData__Team_value = map[string]int32{
	"ALLY":  0,
	"ENEMY": 1,
}

func (x BattleData__Team) Enum() *BattleData__Team {
	p := new(BattleData__Team)
	*p = x
	return p
}
func (x BattleData__Team) String() string {
	return proto.EnumName(BattleData__Team_name, int32(x))
}
func (x *BattleData__Team) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BattleData__Team_value, data, "BattleData__Team")
	if err != nil {
		return err
	}
	*x = BattleData__Team(value)
	return nil
}
func (BattleData__Team) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{9, 1} }

// character data
type CharData_ struct {
	Cid              *uint64 `protobuf:"varint,1,req,name=cid" json:"cid,omitempty"`
	Uid              *uint64 `protobuf:"varint,2,req,name=uid" json:"uid,omitempty"`
	SlotNo           *uint32 `protobuf:"varint,3,opt,name=slotNo" json:"slotNo,omitempty"`
	TypeNo           *uint32 `protobuf:"varint,4,opt,name=typeNo" json:"typeNo,omitempty"`
	Level            *uint32 `protobuf:"varint,5,opt,name=level" json:"level,omitempty"`
	Tier             *uint32 `protobuf:"varint,6,opt,name=tier" json:"tier,omitempty"`
	RegDate          *string `protobuf:"bytes,7,opt,name=regDate" json:"regDate,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CharData_) Reset()                    { *m = CharData_{} }
func (m *CharData_) String() string            { return proto.CompactTextString(m) }
func (*CharData_) ProtoMessage()               {}
func (*CharData_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CharData_) GetCid() uint64 {
	if m != nil && m.Cid != nil {
		return *m.Cid
	}
	return 0
}

func (m *CharData_) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *CharData_) GetSlotNo() uint32 {
	if m != nil && m.SlotNo != nil {
		return *m.SlotNo
	}
	return 0
}

func (m *CharData_) GetTypeNo() uint32 {
	if m != nil && m.TypeNo != nil {
		return *m.TypeNo
	}
	return 0
}

func (m *CharData_) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *CharData_) GetTier() uint32 {
	if m != nil && m.Tier != nil {
		return *m.Tier
	}
	return 0
}

func (m *CharData_) GetRegDate() string {
	if m != nil && m.RegDate != nil {
		return *m.RegDate
	}
	return ""
}

// user data
type UserData_ struct {
	Uid              *uint64      `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Name             *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Did              *string      `protobuf:"bytes,3,opt,name=did" json:"did,omitempty"`
	Platform         *PlatForm    `protobuf:"varint,4,opt,name=platform,enum=MSG.PlatForm" json:"platform,omitempty"`
	LoginDate        *string      `protobuf:"bytes,5,opt,name=loginDate" json:"loginDate,omitempty"`
	LogoutDate       *string      `protobuf:"bytes,6,opt,name=logoutDate" json:"logoutDate,omitempty"`
	RegDate          *string      `protobuf:"bytes,7,opt,name=regDate" json:"regDate,omitempty"`
	Vc1              *uint32      `protobuf:"varint,8,opt,name=vc1" json:"vc1,omitempty"`
	Vc2              *uint32      `protobuf:"varint,9,opt,name=vc2" json:"vc2,omitempty"`
	Vc3              *uint32      `protobuf:"varint,10,opt,name=vc3" json:"vc3,omitempty"`
	GroupName        *string      `protobuf:"bytes,11,opt,name=groupName" json:"groupName,omitempty"`
	Language         *string      `protobuf:"bytes,12,opt,name=language" json:"language,omitempty"`
	Chars            []*CharData_ `protobuf:"bytes,13,rep,name=chars" json:"chars,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *UserData_) Reset()                    { *m = UserData_{} }
func (m *UserData_) String() string            { return proto.CompactTextString(m) }
func (*UserData_) ProtoMessage()               {}
func (*UserData_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *UserData_) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *UserData_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UserData_) GetDid() string {
	if m != nil && m.Did != nil {
		return *m.Did
	}
	return ""
}

func (m *UserData_) GetPlatform() PlatForm {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return PlatForm_IOS
}

func (m *UserData_) GetLoginDate() string {
	if m != nil && m.LoginDate != nil {
		return *m.LoginDate
	}
	return ""
}

func (m *UserData_) GetLogoutDate() string {
	if m != nil && m.LogoutDate != nil {
		return *m.LogoutDate
	}
	return ""
}

func (m *UserData_) GetRegDate() string {
	if m != nil && m.RegDate != nil {
		return *m.RegDate
	}
	return ""
}

func (m *UserData_) GetVc1() uint32 {
	if m != nil && m.Vc1 != nil {
		return *m.Vc1
	}
	return 0
}

func (m *UserData_) GetVc2() uint32 {
	if m != nil && m.Vc2 != nil {
		return *m.Vc2
	}
	return 0
}

func (m *UserData_) GetVc3() uint32 {
	if m != nil && m.Vc3 != nil {
		return *m.Vc3
	}
	return 0
}

func (m *UserData_) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *UserData_) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *UserData_) GetChars() []*CharData_ {
	if m != nil {
		return m.Chars
	}
	return nil
}

// contents info data
type Contents_ struct {
	Type             *Contents__ContentType `protobuf:"varint,1,req,name=type,enum=MSG.Contents__ContentType" json:"type,omitempty"`
	Name             *string                `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Contents_) Reset()                    { *m = Contents_{} }
func (m *Contents_) String() string            { return proto.CompactTextString(m) }
func (*Contents_) ProtoMessage()               {}
func (*Contents_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Contents_) GetType() Contents__ContentType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Contents__CT_INGAME
}

func (m *Contents_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// group memebr
type GMember_ struct {
	Uid              *uint64             `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Name             *string             `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Grade            *GMember__GradeType `protobuf:"varint,3,req,name=grade,enum=MSG.GMember__GradeType" json:"grade,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GMember_) Reset()                    { *m = GMember_{} }
func (m *GMember_) String() string            { return proto.CompactTextString(m) }
func (*GMember_) ProtoMessage()               {}
func (*GMember_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GMember_) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *GMember_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GMember_) GetGrade() GMember__GradeType {
	if m != nil && m.Grade != nil {
		return *m.Grade
	}
	return GMember__Grade_1
}

// group
type Group_ struct {
	Gid              *uint64     `protobuf:"varint,1,req,name=gid" json:"gid,omitempty"`
	Name             *string     `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Country          *string     `protobuf:"bytes,3,req,name=country" json:"country,omitempty"`
	Leader           *uint64     `protobuf:"varint,4,req,name=leader" json:"leader,omitempty"`
	Limit            *uint32     `protobuf:"varint,5,req,name=limit" json:"limit,omitempty"`
	Members          []*GMember_ `protobuf:"bytes,6,rep,name=members" json:"members,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Group_) Reset()                    { *m = Group_{} }
func (m *Group_) String() string            { return proto.CompactTextString(m) }
func (*Group_) ProtoMessage()               {}
func (*Group_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Group_) GetGid() uint64 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *Group_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Group_) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *Group_) GetLeader() uint64 {
	if m != nil && m.Leader != nil {
		return *m.Leader
	}
	return 0
}

func (m *Group_) GetLimit() uint32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *Group_) GetMembers() []*GMember_ {
	if m != nil {
		return m.Members
	}
	return nil
}

// chat
type ChatData_ struct {
	Uid              *uint64 `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	GroupName        *string `protobuf:"bytes,3,req,name=groupName" json:"groupName,omitempty"`
	Language         *string `protobuf:"bytes,4,req,name=language" json:"language,omitempty"`
	Chat             *string `protobuf:"bytes,5,req,name=chat" json:"chat,omitempty"`
	RegDate          *uint64 `protobuf:"varint,6,req,name=regDate" json:"regDate,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatData_) Reset()                    { *m = ChatData_{} }
func (m *ChatData_) String() string            { return proto.CompactTextString(m) }
func (*ChatData_) ProtoMessage()               {}
func (*ChatData_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ChatData_) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *ChatData_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChatData_) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *ChatData_) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *ChatData_) GetChat() string {
	if m != nil && m.Chat != nil {
		return *m.Chat
	}
	return ""
}

func (m *ChatData_) GetRegDate() uint64 {
	if m != nil && m.RegDate != nil {
		return *m.RegDate
	}
	return 0
}

// chat room
type ChatRoom_ struct {
	Rid              *uint64      `protobuf:"varint,1,req,name=rid" json:"rid,omitempty"`
	Name             *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Chats            []*ChatData_ `protobuf:"bytes,3,rep,name=chats" json:"chats,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ChatRoom_) Reset()                    { *m = ChatRoom_{} }
func (m *ChatRoom_) String() string            { return proto.CompactTextString(m) }
func (*ChatRoom_) ProtoMessage()               {}
func (*ChatRoom_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ChatRoom_) GetRid() uint64 {
	if m != nil && m.Rid != nil {
		return *m.Rid
	}
	return 0
}

func (m *ChatRoom_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChatRoom_) GetChats() []*ChatData_ {
	if m != nil {
		return m.Chats
	}
	return nil
}

// chat channel
type ChatChannel_ struct {
	Cid              *uint64      `protobuf:"varint,1,req,name=cid" json:"cid,omitempty"`
	Name             *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Chats            []*ChatData_ `protobuf:"bytes,3,rep,name=chats" json:"chats,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ChatChannel_) Reset()                    { *m = ChatChannel_{} }
func (m *ChatChannel_) String() string            { return proto.CompactTextString(m) }
func (*ChatChannel_) ProtoMessage()               {}
func (*ChatChannel_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ChatChannel_) GetCid() uint64 {
	if m != nil && m.Cid != nil {
		return *m.Cid
	}
	return 0
}

func (m *ChatChannel_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChatChannel_) GetChats() []*ChatData_ {
	if m != nil {
		return m.Chats
	}
	return nil
}

// dungeon data
type DungeonData_ struct {
	No               *uint32      `protobuf:"varint,1,req,name=no" json:"no,omitempty"`
	Tier             *uint32      `protobuf:"varint,2,req,name=tier" json:"tier,omitempty"`
	Monsters         []*CharData_ `protobuf:"bytes,3,rep,name=monsters" json:"monsters,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DungeonData_) Reset()                    { *m = DungeonData_{} }
func (m *DungeonData_) String() string            { return proto.CompactTextString(m) }
func (*DungeonData_) ProtoMessage()               {}
func (*DungeonData_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *DungeonData_) GetNo() uint32 {
	if m != nil && m.No != nil {
		return *m.No
	}
	return 0
}

func (m *DungeonData_) GetTier() uint32 {
	if m != nil && m.Tier != nil {
		return *m.Tier
	}
	return 0
}

func (m *DungeonData_) GetMonsters() []*CharData_ {
	if m != nil {
		return m.Monsters
	}
	return nil
}

// battle data
type BattleData_ struct {
	SrcNo            *uint32               `protobuf:"varint,1,req,name=srcNo" json:"srcNo,omitempty"`
	Targets          []*BattleData__Attack `protobuf:"bytes,2,rep,name=targets" json:"targets,omitempty"`
	Team             *BattleData__Team     `protobuf:"varint,3,req,name=team,enum=MSG.BattleData__Team" json:"team,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *BattleData_) Reset()                    { *m = BattleData_{} }
func (m *BattleData_) String() string            { return proto.CompactTextString(m) }
func (*BattleData_) ProtoMessage()               {}
func (*BattleData_) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *BattleData_) GetSrcNo() uint32 {
	if m != nil && m.SrcNo != nil {
		return *m.SrcNo
	}
	return 0
}

func (m *BattleData_) GetTargets() []*BattleData__Attack {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *BattleData_) GetTeam() BattleData__Team {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return BattleData__ALLY
}

type BattleData__Attack struct {
	No               *uint32                   `protobuf:"varint,1,req,name=no" json:"no,omitempty"`
	Damage           *uint32                   `protobuf:"varint,2,req,name=damage" json:"damage,omitempty"`
	Result           *BattleData__AttackResult `protobuf:"varint,3,req,name=result,enum=MSG.BattleData__AttackResult" json:"result,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *BattleData__Attack) Reset()                    { *m = BattleData__Attack{} }
func (m *BattleData__Attack) String() string            { return proto.CompactTextString(m) }
func (*BattleData__Attack) ProtoMessage()               {}
func (*BattleData__Attack) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9, 0} }

func (m *BattleData__Attack) GetNo() uint32 {
	if m != nil && m.No != nil {
		return *m.No
	}
	return 0
}

func (m *BattleData__Attack) GetDamage() uint32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *BattleData__Attack) GetResult() BattleData__AttackResult {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return BattleData__ALIVE
}

func init() {
	proto.RegisterType((*CharData_)(nil), "MSG.CharData_")
	proto.RegisterType((*UserData_)(nil), "MSG.UserData_")
	proto.RegisterType((*Contents_)(nil), "MSG.Contents_")
	proto.RegisterType((*GMember_)(nil), "MSG.GMember_")
	proto.RegisterType((*Group_)(nil), "MSG.Group_")
	proto.RegisterType((*ChatData_)(nil), "MSG.ChatData_")
	proto.RegisterType((*ChatRoom_)(nil), "MSG.ChatRoom_")
	proto.RegisterType((*ChatChannel_)(nil), "MSG.ChatChannel_")
	proto.RegisterType((*DungeonData_)(nil), "MSG.DungeonData_")
	proto.RegisterType((*BattleData_)(nil), "MSG.BattleData_")
	proto.RegisterType((*BattleData__Attack)(nil), "MSG.BattleData_.Attack")
	proto.RegisterEnum("MSG.MsgId", MsgId_name, MsgId_value)
	proto.RegisterEnum("MSG.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("MSG.PlatForm", PlatForm_name, PlatForm_value)
	proto.RegisterEnum("MSG.ChatType", ChatType_name, ChatType_value)
	proto.RegisterEnum("MSG.Contents__ContentType", Contents__ContentType_name, Contents__ContentType_value)
	proto.RegisterEnum("MSG.GMember__GradeType", GMember__GradeType_name, GMember__GradeType_value)
	proto.RegisterEnum("MSG.BattleData__AttackResult", BattleData__AttackResult_name, BattleData__AttackResult_value)
	proto.RegisterEnum("MSG.BattleData__Team", BattleData__Team_name, BattleData__Team_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x8f, 0xdb, 0xc4,
	0x1b, 0xae, 0xf3, 0xed, 0x37, 0x9b, 0xd4, 0x9d, 0x5f, 0x3f, 0xac, 0xfe, 0x00, 0xad, 0x02, 0x12,
	0xdb, 0x95, 0x58, 0xa9, 0x5b, 0x7a, 0x42, 0x1c, 0x5c, 0x67, 0x9a, 0x5a, 0x24, 0xe3, 0x30, 0x71,
	0xb6, 0x5a, 0x0e, 0x44, 0x26, 0x31, 0xde, 0x88, 0xd8, 0x5e, 0x39, 0x93, 0x8a, 0xfe, 0x01, 0xd0,
	0x1b, 0x94, 0x63, 0x41, 0x11, 0x87, 0xc2, 0x8d, 0xc2, 0x11, 0xae, 0xc0, 0x11, 0x38, 0xc2, 0x09,
	0x09, 0xfe, 0x09, 0xbe, 0x39, 0xa0, 0x79, 0x27, 0x5f, 0x4b, 0xb2, 0x12, 0x70, 0x9b, 0xe7, 0x79,
	0x5e, 0x3f, 0xf3, 0xbc, 0x99, 0x79, 0x1d, 0xc3, 0x56, 0x3f, 0x89, 0xa2, 0x24, 0xde, 0x3b, 0x4e,
	0x13, 0x91, 0x90, 0x6c, 0xab, 0xd3, 0xa8, 0x3d, 0xd4, 0x40, 0xb7, 0x8f, 0xfc, 0xb4, 0xee, 0x0b,
	0xbf, 0x47, 0x0c, 0xc8, 0xf6, 0x87, 0x03, 0x53, 0xdb, 0xce, 0xec, 0xe4, 0xb8, 0x5c, 0x4a, 0x66,
	0x32, 0x1c, 0x98, 0x19, 0xc5, 0x4c, 0x86, 0x03, 0x72, 0x11, 0x0a, 0xe3, 0x51, 0x22, 0x58, 0x62,
	0x66, 0xb7, 0xb5, 0x9d, 0x0a, 0x9f, 0x21, 0xc9, 0x8b, 0xbb, 0xc7, 0x01, 0x4b, 0xcc, 0x9c, 0xe2,
	0x15, 0x22, 0xe7, 0x21, 0x3f, 0x0a, 0xee, 0x04, 0x23, 0x33, 0x8f, 0xb4, 0x02, 0x84, 0x40, 0x4e,
	0x0c, 0x83, 0xd4, 0x2c, 0x20, 0x89, 0x6b, 0x62, 0x42, 0x31, 0x0d, 0xc2, 0xba, 0x2f, 0x02, 0xb3,
	0xb8, 0xad, 0xed, 0xe8, 0x7c, 0x0e, 0x6b, 0xdf, 0x67, 0x40, 0xef, 0x8e, 0x83, 0x65, 0xca, 0xc9,
	0x32, 0xa5, 0xcc, 0x44, 0x20, 0x17, 0xfb, 0x51, 0x80, 0x31, 0x75, 0x8e, 0x6b, 0x59, 0x35, 0x18,
	0x0e, 0x30, 0xa4, 0xce, 0xe5, 0x92, 0x5c, 0x81, 0xd2, 0xf1, 0xc8, 0x17, 0xaf, 0x26, 0x69, 0x84,
	0x19, 0xab, 0xfb, 0x95, 0xbd, 0x56, 0xa7, 0xb1, 0xd7, 0x1e, 0xf9, 0xe2, 0x66, 0x92, 0x46, 0x7c,
	0x21, 0x93, 0xc7, 0x40, 0x1f, 0x25, 0xe1, 0x30, 0xc6, 0x30, 0x79, 0xb4, 0x58, 0x12, 0xe4, 0x09,
	0x80, 0x51, 0x12, 0x26, 0x13, 0x81, 0x72, 0x01, 0xe5, 0x15, 0xe6, 0xf4, 0x46, 0x64, 0xa8, 0x3b,
	0xfd, 0xab, 0x66, 0x09, 0xbb, 0x96, 0x4b, 0xc5, 0xec, 0x9b, 0xfa, 0x9c, 0xd9, 0x57, 0xcc, 0x35,
	0x13, 0xe6, 0xcc, 0x35, 0x99, 0x26, 0x4c, 0x93, 0xc9, 0x31, 0x93, 0x3d, 0x96, 0x55, 0x9a, 0x05,
	0x41, 0x2e, 0x43, 0x69, 0xe4, 0xc7, 0xe1, 0xc4, 0x0f, 0x03, 0x73, 0x0b, 0xc5, 0x05, 0x26, 0x4f,
	0x41, 0xbe, 0x7f, 0xe4, 0xa7, 0x63, 0xb3, 0xb2, 0x9d, 0xdd, 0x29, 0xef, 0x57, 0xb1, 0xdf, 0xc5,
	0x79, 0x73, 0x25, 0xd6, 0xde, 0x92, 0x97, 0x20, 0x89, 0x45, 0x10, 0x8b, 0x71, 0x8f, 0xec, 0x41,
	0x4e, 0x1e, 0x1d, 0xfe, 0xbe, 0xd5, 0xfd, 0xcb, 0xea, 0x91, 0xb9, 0x3a, 0x5f, 0x79, 0x77, 0x8f,
	0x03, 0x8e, 0x75, 0x9b, 0x7e, 0xfc, 0xda, 0xf3, 0x50, 0x5e, 0x29, 0x24, 0x15, 0xd0, 0x6d, 0xaf,
	0xe7, 0xb0, 0x86, 0xd5, 0xa2, 0x86, 0x46, 0xaa, 0x00, 0xb6, 0xd7, 0x73, 0xbb, 0x1e, 0xe2, 0x0c,
	0x39, 0x0b, 0x65, 0x94, 0xad, 0x76, 0xfb, 0x46, 0xf7, 0xd0, 0xc8, 0xd6, 0x1e, 0x69, 0x50, 0x6a,
	0xb4, 0x82, 0xe8, 0x95, 0x20, 0xfd, 0xa7, 0xc7, 0xfd, 0x0c, 0xe4, 0xc3, 0xd4, 0x1f, 0x04, 0x66,
	0x16, 0x63, 0x5f, 0xc2, 0xd8, 0x73, 0x8f, 0xbd, 0x86, 0x94, 0x30, 0xb3, 0xaa, 0xaa, 0x35, 0x41,
	0x5f, 0x70, 0xa4, 0x0c, 0x45, 0x04, 0xbd, 0xab, 0x86, 0xb6, 0x04, 0xfb, 0x46, 0x66, 0x09, 0xae,
	0x19, 0xd9, 0x25, 0x78, 0xd6, 0xc8, 0x2d, 0xc1, 0x75, 0x23, 0x5f, 0x7b, 0x5f, 0x83, 0x42, 0x43,
	0x1e, 0x08, 0xa6, 0x0d, 0x97, 0x69, 0xc3, 0x53, 0xd2, 0x9a, 0x50, 0xec, 0x27, 0x93, 0x58, 0xa4,
	0x77, 0x31, 0xaf, 0xce, 0xe7, 0x50, 0x8e, 0xd1, 0x28, 0xf0, 0x07, 0x41, 0x6a, 0xe6, 0xd0, 0x62,
	0x86, 0x70, 0x8c, 0x86, 0xd1, 0x50, 0x98, 0xf9, 0xed, 0x0c, 0x8e, 0x91, 0x04, 0xe4, 0x69, 0x28,
	0x46, 0xd8, 0xe2, 0xd8, 0x2c, 0xe0, 0x09, 0x57, 0x4e, 0xf4, 0xcd, 0xe7, 0x6a, 0xed, 0x81, 0x9a,
	0x73, 0xf1, 0x6f, 0x26, 0xe8, 0xc4, 0xb5, 0x53, 0x31, 0x4f, 0xb9, 0x76, 0x39, 0x14, 0x97, 0xd7,
	0x8e, 0x40, 0xae, 0x7f, 0xe4, 0xab, 0xac, 0x3a, 0xc7, 0xf5, 0xea, 0x50, 0x14, 0x70, 0xdf, 0xc5,
	0x74, 0xdf, 0x56, 0xd1, 0x78, 0x92, 0x44, 0x18, 0x2d, 0x5d, 0x46, 0x4b, 0x4f, 0x89, 0xa6, 0xee,
	0xb5, 0x18, 0x9b, 0xd9, 0x93, 0xf7, 0x5a, 0x2c, 0xef, 0xb5, 0x18, 0xd7, 0x5e, 0x82, 0x2d, 0xc9,
	0xd9, 0x47, 0x7e, 0x1c, 0x07, 0xa3, 0x4d, 0xaf, 0xb7, 0xff, 0xee, 0xfd, 0x32, 0x6c, 0xd5, 0x27,
	0x71, 0x18, 0x24, 0xb1, 0xfa, 0x49, 0xab, 0x90, 0x89, 0x13, 0xb4, 0xae, 0xf0, 0x4c, 0x9c, 0x2c,
	0x5e, 0x70, 0x19, 0x64, 0xd4, 0x0b, 0x6e, 0x17, 0x4a, 0x51, 0x12, 0x8f, 0x85, 0x3c, 0xae, 0xec,
	0xc6, 0x81, 0x5c, 0xe8, 0xb5, 0x8f, 0x32, 0x50, 0xbe, 0xe1, 0x0b, 0x31, 0x0a, 0x94, 0xff, 0x79,
	0xc8, 0x8f, 0xd3, 0x3e, 0x9b, 0x6f, 0xa1, 0x00, 0xb9, 0x0a, 0x45, 0xe1, 0xa7, 0x61, 0x20, 0xc6,
	0x66, 0x06, 0x0d, 0xd5, 0xbd, 0x5f, 0x79, 0x70, 0xcf, 0x12, 0xc2, 0xef, 0xbf, 0xc6, 0xe7, 0x75,
	0xe4, 0x0a, 0xe4, 0x44, 0xe0, 0x47, 0xb3, 0x39, 0xb9, 0xb0, 0x56, 0xef, 0x05, 0x7e, 0xc4, 0xb1,
	0xe4, 0x72, 0x08, 0x05, 0xf5, 0xf4, 0x5a, 0x77, 0x17, 0xa1, 0x30, 0xf0, 0x23, 0x79, 0xf4, 0xaa,
	0xbf, 0x19, 0x22, 0xd7, 0xa1, 0x90, 0x06, 0xe3, 0xc9, 0x48, 0xcc, 0xec, 0x1f, 0x3f, 0x2d, 0x0e,
	0x16, 0xf1, 0x59, 0x71, 0xed, 0x49, 0xd8, 0x5a, 0xe5, 0x89, 0x0e, 0x79, 0xab, 0xe9, 0x1c, 0x50,
	0xe3, 0x0c, 0x29, 0x41, 0xae, 0x4e, 0xad, 0xba, 0xa1, 0xd5, 0xfe, 0x0f, 0x39, 0x99, 0x4d, 0x32,
	0x56, 0xb3, 0x79, 0x68, 0x9c, 0x91, 0x65, 0x94, 0xd1, 0xd6, 0xa1, 0xa1, 0xed, 0xde, 0x2b, 0x40,
	0xbe, 0x35, 0x0e, 0x9d, 0x01, 0x29, 0x43, 0xc1, 0x6e, 0xba, 0x1d, 0x5a, 0x37, 0xee, 0x33, 0x52,
	0x05, 0xbd, 0xe9, 0x36, 0x1c, 0xd6, 0xe3, 0xf4, 0x45, 0xe3, 0xa7, 0x15, 0x6c, 0xb1, 0x8e, 0xf1,
	0x33, 0x23, 0x15, 0x28, 0xb5, 0x1d, 0xd6, 0x40, 0xf9, 0x17, 0x05, 0x5d, 0xd6, 0x40, 0xf5, 0x57,
	0x46, 0xce, 0x02, 0x70, 0xda, 0x70, 0x3a, 0x1e, 0xea, 0xbf, 0xad, 0x12, 0xb2, 0xe2, 0x77, 0x46,
	0x0c, 0x28, 0x1f, 0x50, 0xde, 0x71, 0x5c, 0xb5, 0xc3, 0x1f, 0x27, 0x18, 0x59, 0xf3, 0xa7, 0xdc,
	0xb3, 0x64, 0xdf, 0xb2, 0x94, 0xc7, 0x1b, 0x53, 0x6d, 0x81, 0xa5, 0xfc, 0xe6, 0x0a, 0x66, 0xae,
	0x67, 0xdc, 0x9b, 0x6a, 0xc4, 0x04, 0x62, 0x73, 0x6a, 0x79, 0x54, 0xb2, 0xdc, 0x75, 0x5b, 0xf8,
	0xe4, 0xdb, 0x1b, 0x15, 0xe9, 0x71, 0x7f, 0xa3, 0x22, 0xdd, 0xde, 0x51, 0x8a, 0xc3, 0x0e, 0x9c,
	0xbf, 0xb9, 0x3d, 0xd8, 0xa8, 0x48, 0xb7, 0x77, 0x37, 0x2a, 0xd2, 0xed, 0xbd, 0xa9, 0x46, 0x2e,
	0xc1, 0x39, 0xca, 0x3c, 0xca, 0x4f, 0x98, 0x3d, 0xdc, 0x24, 0x48, 0xaf, 0x0f, 0x36, 0x09, 0xd2,
	0xea, 0x43, 0x25, 0x34, 0xa9, 0x75, 0x70, 0x32, 0xd7, 0xa3, 0x4d, 0x82, 0xb4, 0xfa, 0x78, 0x93,
	0x20, 0xad, 0x3e, 0x99, 0x6a, 0xe4, 0x3c, 0x54, 0x17, 0xdd, 0x73, 0xf4, 0xf9, 0x74, 0x8d, 0x95,
	0x26, 0x9f, 0x4d, 0x35, 0x42, 0x60, 0xcb, 0x76, 0x99, 0x47, 0x99, 0xd7, 0xc1, 0xe7, 0x3f, 0x57,
	0xfd, 0xba, 0xb7, 0x19, 0xad, 0xcb, 0x42, 0xcb, 0xf6, 0x28, 0x47, 0xe5, 0x8b, 0x59, 0x75, 0x97,
	0x73, 0xca, 0xec, 0x43, 0xe4, 0xbe, 0x9c, 0x6a, 0xe4, 0x02, 0x9c, 0x6d, 0x37, 0xad, 0xc3, 0x7a,
	0x97, 0x35, 0xe8, 0xec, 0xdc, 0xbf, 0x5a, 0xa7, 0xe5, 0x7e, 0x5f, 0xaf, 0xd3, 0xd2, 0xe4, 0x1b,
	0x45, 0x37, 0xe9, 0x01, 0x6d, 0x76, 0xdb, 0x8b, 0xcc, 0xdf, 0xae, 0xd3, 0xd2, 0xe4, 0x3b, 0xd5,
	0x8a, 0xe7, 0x50, 0xbe, 0x52, 0xfc, 0xc3, 0x1a, 0x2b, 0x6b, 0x7f, 0x9c, 0x6a, 0xbb, 0xaf, 0x83,
	0x4e, 0xd3, 0x34, 0x49, 0xed, 0x64, 0x10, 0xc8, 0x7f, 0x56, 0xca, 0x79, 0xaf, 0xd3, 0xb5, 0x6d,
	0xda, 0xe9, 0x18, 0x67, 0x08, 0x81, 0xaa, 0x24, 0xd4, 0x10, 0xdc, 0xb4, 0x9c, 0xa6, 0x21, 0x37,
	0x3d, 0x27, 0x39, 0x8b, 0x37, 0xba, 0x2d, 0xca, 0x3c, 0x45, 0x67, 0xc8, 0x45, 0x20, 0x48, 0x77,
	0xbd, 0x5b, 0x2e, 0x77, 0xbc, 0x43, 0xc5, 0x67, 0xc9, 0x25, 0xf8, 0x1f, 0x7a, 0xd2, 0x8e, 0xbc,
	0xe5, 0x2f, 0xd0, 0x99, 0x90, 0xdb, 0xdd, 0x86, 0xd2, 0xfc, 0x53, 0x8a, 0x14, 0x21, 0xeb, 0xb8,
	0x72, 0xc3, 0x32, 0x14, 0x2d, 0x56, 0xe7, 0xae, 0x53, 0x37, 0xb4, 0xdd, 0xe7, 0xa0, 0x24, 0x5f,
	0xa4, 0xf8, 0xa7, 0x6b, 0xc0, 0x16, 0x5e, 0x7b, 0xfb, 0x96, 0xc5, 0x18, 0x6d, 0xce, 0x3e, 0x0b,
	0x24, 0xd3, 0xe0, 0x6e, 0xb7, 0x6d, 0x64, 0xf0, 0xab, 0x01, 0x07, 0xc7, 0x75, 0x5b, 0x46, 0xf6,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xa4, 0x55, 0x9c, 0xbe, 0x0a, 0x00, 0x00,
}
