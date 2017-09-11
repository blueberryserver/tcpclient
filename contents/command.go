package contents

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/blueberryserver/tcpclient/MSG"
	"github.com/blueberryserver/tcpserver/network"
	"github.com/golang/protobuf/proto"
)

//
//var Clients = make([]*network.BlueClient, 1000)

//
var Clients = make(map[*network.Session]*network.BlueClient)

//
var SessionKeys = make(map[*network.Session]string)

//
type CmdFunc func(string)

//
var Cmds = make(map[string]CmdFunc)

//
func Init() {
	Cmds["Login"] = Login
	Cmds["Ping"] = Ping
	Cmds["Regist"] = Regist
	//Cmds["Chat"] = Chat
}

//
func Login(name string) {
	var i int
	for _, c := range Clients {
		if c != nil {
			namei := name + strconv.Itoa(i)
			req := &MSG.LoginReq{
				Name: &namei,
			}
			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["LOGIN_REQ"], abuff, uint16(len(abuff)))
		}
		i = i + 1
		time.Sleep(1 * time.Millisecond)
	}
}

//
func Ping(name string) {
	for s, c := range Clients {
		if c != nil {
			key := SessionKeys[s]
			req := &MSG.PingReq{
				SessionKey: &key,
			}

			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["PING_REQ"], abuff, uint16(len(abuff)))
		}
		time.Sleep(1 * time.Millisecond)
	}
}

//
func Regist(name string) {
	var i int
	for _, c := range Clients {
		if c != nil {
			fmt.Println(i, c)
			namei := name + strconv.Itoa(i)
			var strData = "{\"did\": \"df23789es\"," +
				"\"loginDate\": \"\"," +
				"\"logoutDate\": \"\"," +
				"\"name\": \"\"," +
				"\"platform\": \"ANDROID\"," +
				"\"regDate\": \"\"," +
				"\"uid\": 0," +
				"\"vc1\": 10000," +
				"\"vc2\": 20000," +
				"\"vc3\": 30000," +
				"\"groupName\": \"korea0\"," +
				"\"language\": \"kr\"}"

			data := &MSG.UserData_{}
			err := json.Unmarshal([]byte(strData), data)
			if err != nil {
				fmt.Println(err)
			}
			req := &MSG.RegistReq{
				Data: data,
			}
			now := time.Now().Format("2006_01_02_15")

			req.Data.Name = &namei
			req.Data.RegDate = &now

			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["REGIST_REQ"], abuff, uint16(len(abuff)))
			i = i + 1
			time.Sleep(50 * time.Millisecond)
		}

	}
}

//
func Chat(t string, schat string) {
	// var strData = "{\"uid\":10,\"name\":\"noom\",\"group_name\":\"korea0\"," +
	// 	"\"languge\":\"kr\"," +
	// 	"\"chat\":\"" + schat +
	// 	"\",\"reg_date\":1503639263}"
	var chatType MSG.ChatType
	if t == "ch" {
		chatType = *MSG.ChatType_CHAT_CHANNEL.Enum()
	} else {
		chatType = *MSG.ChatType_CHAT_ROOM.Enum()
	}

	for _, c := range Clients {
		if c != nil {
			//chat_type := MSG.ChatType_CHAT_CHANNEL
			req := &MSG.ChatReq{
				Msg:  &schat,
				Type: &chatType,
			}
			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["CHAT_REQ"], abuff, uint16(len(abuff)))
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//
func CreatRoom(name string) {
	greeting := "Hello Hi!!!"
	for _, c := range Clients {
		if c != nil {
			req := &MSG.CreateChatRoomReq{
				Name:     &name,
				Greeting: &greeting,
			}
			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["CREATECHATROOM_REQ"], abuff, uint16(len(abuff)))
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//
func InviteRoom(name string, rid uint64, targetName string, targetUID uint64) {
	for _, c := range Clients {
		if c != nil {
			req := &MSG.InviteChatRoomReq{
				Rid:        &rid,
				Rkey:       &name,
				TargetName: &targetName,
				TargetUid:  &targetUID,
			}
			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["INVITECHATROOM_REQ"], abuff, uint16(len(abuff)))
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//
func EnterRoom(name string, rid uint64) {
	for _, c := range Clients {
		if c != nil {
			req := &MSG.EnterChatRoomReq{
				Rid:  &rid,
				Rkey: &name,
			}
			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["ENTERCHATROOM_REQ"], abuff, uint16(len(abuff)))
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//
func LeaveRoom(name string, rid uint64) {
	for _, c := range Clients {
		if c != nil {
			req := &MSG.LeaveChatRoomReq{
				Rid:  &rid,
				Rkey: &name,
			}
			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["LEAVECHATROOM_REQ"], abuff, uint16(len(abuff)))
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//
func CreateChar(charNo uint32) {
	fmt.Println("charno: ", charNo)
	for _, c := range Clients {
		if c != nil {
			req := &MSG.CreateCharReq{
				CharNo: &charNo,
			}
			abuff, _ := proto.Marshal(req)
			c.SendPacket(MSG.MsgId_value["CREATECHAR_REQ"], abuff, uint16(len(abuff)))
			time.Sleep(50 * time.Millisecond)
		}
	}
}
