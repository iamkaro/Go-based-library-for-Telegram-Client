/*!
 * I am Karo  😊👍
 *
 * Contact me:
 *     https://www.karo.link/
 *     https://github.com/iamkaro
 *     https://www.linkedin.com/in/iamkaro
 *
 * Go-based library for developing Telegram client applications
 * https://github.com/iamkaro/Go-based-library-for-Telegram-Client.git
 * Copyright © 2020 developed.
 */

package telegram

import (
	"math/rand"
	"time"
)

func NewClient(apiID uint64, apiHash string, rootFolder string) *Client {
	var (
		client = &Client{
			api:                   newApi(),
			isOn:                  true,
			on:                    true,
			loadsData:             map[string]*loadData{},
			printAllReceivingData: false,
			rootFolder:            rootFolder,
		}
	)

	if client.execute(object{"@type": "setLogVerbosityLevel", "new_verbosity_level": 0}) == "" {
		client.Destroy()
		return nil
	}

	client.newSelf()
	client.newAuth()
	client.newSessions()
	client.newFiles()
	client.newContacts()
	client.newUsers()
	client.newChats()
	client.newMessages()
	client.newCallback()
	client.newProxies()

	go client.receiving()

	var (
		state  *class = nil
		result data
	)

	for loops := byte(0); ; loops += 1 {
		if log("loop", loops); loops > 255 {
			println("load error ...")
			client.Destroy()
			return nil
		}

		sleep(time.Duration(2000 + rand.Intn(1000)))
		state = &class{Type: ""}
		_ = client.load(object{"@type": "getAuthorizationState"}).extractTo(state)
		log("state", state)
		if state.Type == "authorizationStateWaitTdlibParameters" {
			result = client.load(object{
				"@type":                   "setTdlibParameters",
				"use_test_dc":             false,
				"database_directory":      client.rootFolder,
				"files_directory":         client.rootFolder,
				"database_encryption_key": []byte("er1gh#yr46q3gy&rj$y2fe9mx@s0zaq2]l;wde,xsz"),
				"use_file_database":       true,
				"use_chat_info_database":  true,
				"use_message_database":    true,
				"use_secret_chats":        false,
				"api_id":                  apiID,
				"api_hash":                apiHash,
				"system_language_code":    "en",
				"device_model":            "Server",
				"system_version":          "1.0.0",
				"application_version":     "1.0.0",
			})
			continue
		}
		if state.Type == "authorizationStateWaitEncryptionKey" {
			log("checkDatabaseEncryptionKey", "send encryption_key.")
			continue
		}

		log("result", result)
		return client
	}
}

/*------------------------------------------------/          Client          /-----------*/

type (
	Clients []*Client
	Client  struct {
		api                   *api
		isOn                  bool
		on                    bool
		loadsData             map[string]*loadData
		printAllReceivingData bool
		rootFolder            string

		/*--------*/
		Self     *self
		Auth     *auth
		Sessions *sessions
		Files    *files
		Contacts *contacts
		Users    *users
		Chats    *chats
		Messages *messages
		Callback *callback
		Proxies  *proxies
	}
)

func (client *Client) send(request object)         { client.api.send(request.toJson()) }
func (client *Client) receive() data               { return data(client.api.receive()) }
func (client *Client) execute(request object) data { return data(client.api.execute(request.toJson())) }
func (client *Client) Destroy() {
	for client.on || client.isOn {
		client.Off()
		sleep(1000)
	}
	client.api.destroy()
	client.api = nil
	/*----------*/
	client.loadsData = nil
	client.Self = nil
	client.Auth = nil
	client.Sessions = nil
	client.Files = nil
	client.Contacts = nil
	client.Users = nil
	client.Chats = nil
	client.Messages = nil
}

/*------------------------------------------------/     client-control     /-----------*/

func (client *Client) PrintAllReceivingData(enable bool) { client.printAllReceivingData = enable }
func (client *Client) RootFolder() string                { return client.rootFolder }
func (client *Client) Off()                              { client.on = false }
func (client *Client) On() {
	sleep(time.Duration(rand.Intn(3077)))
	if !client.on {
		client.on = true
		sleep(3000)
		if !client.isOn {
			go client.receiving()
		}
	}
}
