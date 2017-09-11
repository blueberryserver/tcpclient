package contents

import (
	"encoding/json"
	"log"

	"github.com/blueberryserver/tcpclient/MSG"
	"github.com/blueberryserver/tcpserver/network"
	"github.com/golang/protobuf/proto"
)

// client handler
//-------------------------------------------------------------------------------------

// session disconnection call function
func CloseHandler(session *network.Session) {
}

// ans login
type _LoginAns struct {
	msgID int32
}

// ans login
func GetHandlerLoginAns() _LoginAns {
	return _LoginAns{msgID: MSG.MsgId_value["LOGIN_ANS"]}
}

// type UserData struct {
// 	Uid        uint64       `json:"uid"`
// 	Name       string       `json:"name"`
// 	Did        string       `json:"did"`
// 	Platform   MSG.PlatForm `json:"platform"`
// 	LoginDate  time.Time    `json:"login_date"`
// 	LogoutDate time.Time    `json:"logout_date"`
// 	RegDate    time.Time    `json:"reg_date"`
// 	Vc1        uint         `json:"vc1"`
// 	Vc2        uint         `json:"vc2"`
// 	Vc3        uint         `json:"vc3"`
// }

// ans login
func (m _LoginAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.LoginAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(ans.Err.String())
	if *ans.Err != MSG.ErrorCode_ERR_SUCCESS {
		return true
	}

	log.Println("client login ans", ans.Data.String())

	SessionKeys[session] = ans.GetSessionKey()
	jData, _ := json.Marshal(ans.Data)
	log.Println("json:", string(jData), "sessionkey:", ans.GetSessionKey())
	return true
}

// ans login
type _PongAns struct {
	msgID int32
}

// ans login
func GetHandlerPongAns() _PongAns {
	return _PongAns{msgID: MSG.MsgId_value["PONG_ANS"]}
}

// ans login
func (m _PongAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.PongAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client pong ans", ans.Err.String())
	if *ans.Err != MSG.ErrorCode_ERR_SUCCESS {
		return true
	}

	return true
}

// ans login
type _RegistAns struct {
	msgID int32
}

// ans regist
func GetHandlerRegistAns() _RegistAns {
	return _RegistAns{msgID: MSG.MsgId_value["REGIST_ANS"]}
}

// ans regist
func (m _RegistAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.RegistAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client regist ans", ans.Err.String())
	if *ans.Err != MSG.ErrorCode_ERR_SUCCESS {
		return true
	}
	return true
}

type _ChatAns struct {
	msgID int32
}

// ans chat
func GetHandlerChatAns() _ChatAns {
	return _ChatAns{msgID: MSG.MsgId_value["CHAT_ANS"]}
}

func (m _ChatAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.ChatAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat ans", ans.Err.String())
	if *ans.Err != MSG.ErrorCode_ERR_SUCCESS {
		return true
	}

	return true
}

type _ChatNot struct {
	msgID int32
}

// not chat
func GetHandlerChatNot() _ChatNot {
	return _ChatNot{msgID: MSG.MsgId_value["CHAT_NOT"]}
}

func (m _ChatNot) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	not := &MSG.ChatNot{}
	err := proto.Unmarshal(data[:length], not)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat not", not.String())
	return true
}

//
type _CreateChatRoomAns struct {
	msgID int32
}

//
func GetHandlerCreateChatRoomAns() _CreateChatRoomAns {
	return _CreateChatRoomAns{msgID: MSG.MsgId_value["CREATECHATROOM_ANS"]}
}

//
func (m _CreateChatRoomAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.CreateChatRoomAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat ans", ans.Err.String())
	return true
}

//
type _EnterChatRoomAns struct {
	msgID int32
}

//
func GetHandlerEnterChatRoomAns() _EnterChatRoomAns {
	return _EnterChatRoomAns{msgID: MSG.MsgId_value["ENTERCHATROOM_ANS"]}
}

//
func (m _EnterChatRoomAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.EnterChatRoomAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat ans", ans.Err.String())
	return true
}

type _EnterChatRoomNot struct {
	msgID int32
}

// not chat
func GetHandlerEnterChatRoomNot() _EnterChatRoomNot {
	return _EnterChatRoomNot{msgID: MSG.MsgId_value["ENTERCHATROOM_NOT"]}
}

func (m _EnterChatRoomNot) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	not := &MSG.EnterChatRoomNot{}
	err := proto.Unmarshal(data[:length], not)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat not", not.String())
	return true
}

//
type _LeaveChatRoomAns struct {
	msgID int32
}

//
func GetHandlerLeaveChatRoomAns() _LeaveChatRoomAns {
	return _LeaveChatRoomAns{msgID: MSG.MsgId_value["LEAVECHATROOM_ANS"]}
}

//
func (m _LeaveChatRoomAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.LeaveChatRoomAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat ans", ans.Err.String())
	return true
}

type _LeaveChatRoomNot struct {
	msgID int32
}

// not chat
func GetHandlerLeaveChatRoomNot() _LeaveChatRoomNot {
	return _LeaveChatRoomNot{msgID: MSG.MsgId_value["LEAVECHATROOM_NOT"]}
}

func (m _LeaveChatRoomNot) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshalingc
	not := &MSG.LeaveChatRoomNot{}
	err := proto.Unmarshal(data[:length], not)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat not", not.String())
	return true
}

//
type _InviteChatRoomAns struct {
	msgID int32
}

//
func GetHandlerInviteChatRoomAns() _InviteChatRoomAns {
	return _InviteChatRoomAns{msgID: MSG.MsgId_value["INVITECHATROOM_ANS"]}
}

//
func (m _InviteChatRoomAns) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.InviteChatRoomAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat ans", ans.Err.String())
	return true
}

type _InviteChatRoomNot struct {
	msgID int32
}

// not chat
func GetHandlerInviteChatRoomNot() _InviteChatRoomNot {
	return _InviteChatRoomNot{msgID: MSG.MsgId_value["INVITECHATROOM_NOT"]}
}

func (m _InviteChatRoomNot) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	not := &MSG.InviteChatRoomNot{}
	err := proto.Unmarshal(data[:length], not)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client chat not", not.String())

	client := Clients[session]
	req := &MSG.EnterChatRoomReq{
		Rid:  not.Rid,
		Rkey: not.Rkey,
	}
	abuff, _ := proto.Marshal(req)
	client.SendPacket(MSG.MsgId_value["ENTERCHATROOM_REQ"], abuff, uint16(len(abuff)))
	return true
}

type _CreateCharAns struct {
	msgID int32
}

// ans create char
func GetHandlerCreateCharAns() _CreateCharAns {
	return _CreateCharAns{msgID: MSG.MsgId_value["CREATECHAR_ANS"]}
}

func (m _CreateCharAns) Execute(ession *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	ans := &MSG.CreateCharAns{}
	err := proto.Unmarshal(data[:length], ans)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client ans", ans.Err.String())
	if *ans.Err != MSG.ErrorCode_ERR_SUCCESS {
		return true
	}

	log.Println(ans.String())
	return true
}

type _CurrencyNot struct {
	msgID int32
}

// not chat
func GetHandlerCurrencyNot() _CurrencyNot {
	return _CurrencyNot{msgID: MSG.MsgId_value["CURRENCY_NOT"]}
}

func (m _CurrencyNot) Execute(session *network.Session, data []byte, length uint16) bool {
	// unmarshaling
	not := &MSG.CurrencyNot{}
	err := proto.Unmarshal(data[:length], not)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("client not", not.String())
	return true
}
