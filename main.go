package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/blueberryserver/tcpclient/MSG"
	"github.com/blueberryserver/tcpclient/contents"
	"github.com/blueberryserver/tcpserver/network"
)

const inputdelimiter = '\n'

func main() {
	logfile := "log_" + time.Now().Format("2006_01_02_15") + ".txt"
	fileLog, err := os.OpenFile(logfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer fileLog.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	mutiWriter := io.MultiWriter(fileLog, os.Stdout)
	log.SetOutput(mutiWriter)

	fmt.Println("start client")

	for true {
		var s string
		fmt.Scanf("%s", &s)
		if s == "con" {
			var count int
			fmt.Scanf("%d", &count)
			connect(count)
		} else if s == "login" {
			var name string
			fmt.Scanf("%s", &name)
			contents.Login(name)
		} else if s == "ping" {
			contents.Ping(" ")
		} else if s == "regist" {
			var name string
			fmt.Scanf("%s", &name)
			contents.Regist(name)
		} else if s == "close" {
			allclose()
			break
		} else if s == "recon" {
			reconnect()
		} else if s == "chat" {
			var t string
			fmt.Scan(&t)
			reader := bufio.NewReader(os.Stdin)
			schat, err := reader.ReadString(inputdelimiter)
			if err != nil {
				fmt.Println(err)
				return
			}
			// convert CRLF to LF
			schat = strings.Replace(schat, "\r\n", "", -1)
			contents.Chat(t, schat)
		} else if s == "create" {
			var name string
			fmt.Scanf("%s", &name)
			contents.CreatRoom(name)
		} else if s == "invite" {
			var name string
			var rid uint64
			var tname string
			var tuid uint64
			fmt.Scanf("%s %d %s %d", &name, &rid, &tname, &tuid)
			contents.InviteRoom(name, rid, tname, tuid)
		} else if s == "enter" {
			var name string
			var rid uint64
			fmt.Scanf("%s %d", &name, &rid)
			contents.EnterRoom(name, rid)
		} else if s == "leave" {
			var name string
			var rid uint64
			fmt.Scanf("%s %d", &name, &rid)
			contents.LeaveRoom(name, rid)
		} else if s == "char" {
			var charno uint32
			fmt.Scanf("%d", &charno)
			contents.CreateChar(charno)
		}
	}
}

//
func ClientStart() {
	log.Printf("client start\r\n")
	client := network.NewClient()

	// regist server handler
	client.AddMsgHandler(MSG.MsgId_value["LOGIN_ANS"], contents.GetHandlerLoginAns())
	client.AddMsgHandler(MSG.MsgId_value["PONG_ANS"], contents.GetHandlerPongAns())
	client.AddMsgHandler(MSG.MsgId_value["REGIST_ANS"], contents.GetHandlerRegistAns())

	client.AddMsgHandler(MSG.MsgId_value["CHAT_ANS"], contents.GetHandlerChatAns())
	client.AddMsgHandler(MSG.MsgId_value["CHAT_NOT"], contents.GetHandlerChatNot())

	client.AddMsgHandler(MSG.MsgId_value["CREATECHATROOM_ANS"], contents.GetHandlerCreateChatRoomAns())
	client.AddMsgHandler(MSG.MsgId_value["INVITECHATROOM_ANS"], contents.GetHandlerInviteChatRoomAns())
	client.AddMsgHandler(MSG.MsgId_value["ENTERCHATROOM_ANS"], contents.GetHandlerEnterChatRoomAns())
	client.AddMsgHandler(MSG.MsgId_value["ENTERCHATROOM_NOT"], contents.GetHandlerEnterChatRoomNot())
	client.AddMsgHandler(MSG.MsgId_value["LEAVECHATROOM_ANS"], contents.GetHandlerLeaveChatRoomAns())
	client.AddMsgHandler(MSG.MsgId_value["LEAVECHATROOM_NOT"], contents.GetHandlerLeaveChatRoomNot())

	client.AddMsgHandler(MSG.MsgId_value["CREATECHAR_ANS"], contents.GetHandlerCreateCharAns())

	client.AddMsgHandler(MSG.MsgId_value["CURRENCY_NOT"], contents.GetHandlerCurrencyNot())

	err := client.Connect("tcp", "13.124.76.58:12300")
	//err := client.Connect("tcp", "13.124.76.58:20000")
	if err != nil {
		fmt.Println(err)
	}

	contents.Clients[client.GetSession()] = client

	//time.Sleep(1 * time.Second)
}

func connect(count int) {
	for i := 0; i < count; i++ {
		ClientStart()
	}
}

func reconnect() {
	for _, c := range contents.Clients {
		if c != nil {
			err := c.Connect("tcp", "13.124.76.58:12300")
			//err := client.Connect("tcp", "13.124.76.58:20000")
			if err != nil {
				fmt.Println(err)
			}

			time.Sleep(1 * time.Millisecond)
		}
	}
}

func allclose() {
	for _, c := range contents.Clients {
		if c != nil {
			c.Close()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
