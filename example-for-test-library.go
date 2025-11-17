package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/iamkaro/Go-based-library-for-Telegram-Client/telegram"
)

const (
	apiID   uint64 = 123456789 //your api ID
	apiHash string = "abc..."  //your api hash
)

var (
	client     *telegram.Client = nil
	clientPath                  = "client"
)

func main() {
	PrintMemUsage()
	//clientPath = get("clientPath: ")
	PrintMemUsage()
	client = telegram.NewClient(apiID, apiHash, clientPath)
	PrintMemUsage()
	if client == nil {
		fmt.Println("client == nil")
		return
	}

	auth()
	PrintMemUsage()
	loop()

}

/*-----------------------------------------------------------------------------*/
func get(name string) string {
	fmt.Print(name)
	var code string
	_, _ = fmt.Scanln(&code)
	return code
}
func getInt(name string) int64 {
	fmt.Print(name)
	var code string
	_, _ = fmt.Scanln(&code)
	n, e := strconv.ParseInt(code, 10, 64)
	if e == nil {
		return n
	} else {
		fmt.Print(e)
		return getInt(name)
	}
}
func out(value interface{}) bool {
	fmt.Printf("[%T]: \t %+v \n", value, value)
	if reflect.ValueOf(value).IsValid() &&
		(reflect.ValueOf(value).Type().Kind() == reflect.Ptr) &&
		(!reflect.ValueOf(value).IsNil()) {
		fmt.Println("show: ")
		var data = reflect.Indirect(reflect.ValueOf(value))
		var v reflect.Value
		for {
			var field = get("field: ")
			if field == "e" {
				return true
			} else {
				v = data.FieldByName(field)
				if v.IsValid() {
					if out(v.Interface()) {
						fmt.Printf("[%T]: \t %+v \n", value, value)
					}
				}
			}
		}
	}
	return false
}

func PrintMemUsage() {
	/** /
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %v b \tTotalAlloc: %v b \tSys: %v b \tNumGC: %v \n",
		m.Alloc, m.TotalAlloc, m.Sys, m.NumGC)
	/**/
}

/*-----------------------------------------------------------------------------*/
func auth() {
	fmt.Println("Auth: ")
	for state := client.Auth.State(); !state.OK; state = client.Auth.State() {
		if state.EnterPhone {
			client.Auth.EnterPhone(get("EnterPhone: "))
		} else if state.EnterCode {
			client.Auth.EnterCode(get("EnterCode: "))
		} else if state.EnterPassword != nil {
			client.Auth.EnterPassword(get("EnterPassword(" +
				state.EnterPassword.Hint + "): "))
		} else if state.RegistrationUser {
			client.Auth.RegistrationUser(get("name: "), get("lastName: "))
		} else {
			fmt.Println("Auth error !!!...")
			out(state)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("Auth ok ...")
}

/*-----------------------------------------------------------------------------*/
func loop() {
	fmt.Println("looping ...")
	for {
		var command = get("getCommand: ")
		switch command {
		case "onPrint":
			client.PrintAllReceivingData(true)
		case "offPrint":
			client.PrintAllReceivingData(false)
		case "State":
			out(client.Auth.State())
		case "Self.GetMe":
			out(client.Self.GetMe())
		case "Self.LogOut":
			client.Self.LogOut()
		case "Self.SetName":
			client.Self.SetName(get("name: "), get("lastName: "))
		case "Contacts.GetAll":
			out(client.Contacts.GetAll())
		case "Chats.Get":
			out(client.Chats.Get(getInt("id: ")))
		case "Chats.All":
			out(client.Chats.GetList.AllOfMainList(40))
		case "Chats.Open":
			client.Chats.Open(getInt("id: "))
		case "Chats.Users":
			out(client.Chats.GetList.Users(200))
		case "Chats.Bots":
			out(client.Chats.GetList.Bots(200))
		case "Chats.Channels":
			out(client.Chats.GetList.Channels(200))
		case "Chats.Groups":
			out(client.Chats.GetList.Groups(200))
		case "Chats.GetByUsername":
			out(client.Chats.GetByUsername(get("username: ")))
		case "Chats.Join.ByInviteLink":
			out(client.Chats.Join.ByInviteLink(get("link: ")))
		case "Chats.Join.ByChatId":
			client.Chats.Join.ByChatId(getInt("id: "))
		case "Chats.Leave":
			client.Chats.Leave(getInt("chatId: "))
		case "Chats.CheckLink":
			out(client.Chats.InviteLink.Check(get("link: ")))
		case "Messages.Send.Text":
			out(client.Messages.Send.Text(
				getInt("chatId: "),
				getInt("replayToId: "),
				false,
				get("text: "),
				true,
				true))
		case "Messages.Send.Document":
			out(client.Messages.Send.Document(
				getInt("chatId: "),
				getInt("replayToId: "),
				false,
				get("path: "),
				get("caption: ")))
		case "Messages.Send.Audio":
			out(client.Messages.Send.Audio(
				getInt("chatId: "),
				getInt("replayToId: "),
				false,
				get("path: "),
				0,
				get("title: "),
				get("caption: ")))
		case "Messages.Send.VoiceNote":
			out(client.Messages.Send.VoiceNote(
				getInt("chatId: "),
				getInt("replayToId: "),
				false,
				get("path: "),
				0,
				get("caption: ")))
		case "Messages.Send.VideoNote":
			out(client.Messages.Send.VideoNote(
				getInt("chatId: "),
				getInt("replayToId: "),
				false,
				get("path: "),
				0,
				0))
		case "Messages.Send.Photo":
			out(client.Messages.Send.Photo(
				getInt("chatId: "),
				getInt("replayToId: "),
				false,
				get("path: "),
				0, 0,
				get("caption: ")))
		case "Messages.Send.Video":
			out(client.Messages.Send.Video(
				getInt("chatId: "),
				getInt("replayToId: "),
				false,
				get("path: "),
				0, 0, 0, true,
				get("caption: ")))
		case "Messages.Get":
			out(client.Messages.Get(getInt("chatId: "), getInt("messageId: ")))
		case "Messages.Send.Count.UnreadMentions":
			out(client.Messages.Count.
				UnreadMentions(getInt("chatId: "), false))
		case "Messages.Send.Count.Mentions":
			out(client.Messages.Count.
				Mentions(getInt("chatId: "), false))
		case "Messages.Send.Search":
			out(client.Messages.Search.Full(
				getInt("chatId: "),
				get("query: "),
				int32(getInt("senderId")),
				0, 0, 40))
		case "off":
			client.Off()
		case "on":
			client.On()

		case "new":
			PrintMemUsage()
			var newClientPath = get("newClientPath: ")
			if newClientPath != "" {
				clientPath = newClientPath
				client = telegram.NewClient(apiID, apiHash, clientPath)
				if client == nil {
					fmt.Println("client is nil...")
					client = nil
				}
			}
			PrintMemUsage()

		case "Destroy":
			if client != nil {
				PrintMemUsage()
				client.Destroy()
				PrintMemUsage()
				client = nil
			}
			PrintMemUsage()
			time.Sleep(time.Second)
			PrintMemUsage()
			time.Sleep(time.Second)
			PrintMemUsage()
			time.Sleep(time.Second)
			PrintMemUsage()
			time.Sleep(time.Second)
			PrintMemUsage()
		case "move":
			var newClientPath = get("newClientPath: ")
			if newClientPath != "" {
				out(os.Rename(clientPath, newClientPath))
				clientPath = newClientPath
			}

		}
	}
}
