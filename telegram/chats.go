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

func (client *Client) newChats() {
	client.Chats = &chats{
		client: client,
	}
	/*--------*/
	client.Chats.newMembers()
	client.Chats.newSetChatLists()
	client.Chats.newJoin()
	client.Chats.newInviteLink()
	client.Chats.newGetList()

}

type (
	chats struct {
		client *Client
		/*--------*/
		Members      *chatsMembers
		SetChatLists *chatsSetChatLists
		Join         *chatsJoin
		InviteLink   *chatsInviteLink
		GetList      *chatsGetList
	}
)

func (chats *chats) Open(chatId int64) {
	chats.client.load(object{"@type": "openChat", "chat_id": chatId})
}
func (chats *chats) Close(chatId int64) {
	chats.client.load(object{"@type": "closeChat", "chat_id": chatId})
}

func (chats *chats) Leave(chatId int64) {
	chats.client.load(object{"@type": "leaveChat", "chat_id": chatId})
}

func (chats *chats) Get(chatId int64) *Chat {
	var chat = object{}
	if chats.client.Load(&chat, Object{"@type": "getChat", "chat_id": chatId}) {
		return getChat(chat)
	}
	return nil
}

func (chats *chats) GetByUsername(username string) *Chat {
	var chat = object{}
	if chats.client.Load(&chat, Object{"@type": "searchPublicChat", "username": username}) {
		return getChat(chat)
	}
	return nil
}

func (chats *chats) GetAdministrators(chatId int64) []*Administrator {
	var admins = object{}
	if (chats.client.Load(&admins, Object{"@type": "getChatAdministrators", "chat_id": chatId})) {
		return convertList(admins.Array("administrators").object(), getAdministrator)
	}
	return nil
}

func (chats *chats) Notification(chatId int64, disable bool) {
	chats.client.load(object{
		"@type":                        "toggleChatDefaultDisableNotification",
		"chat_id":                      chatId,
		"default_disable_notification": disable,
	})
}

func (chats *chats) SetState(chatId int64, isUnread bool) {
	chats.client.load(object{
		"@type":               "toggleChatIsMarkedAsUnread",
		"chat_id":             chatId,
		"is_marked_as_unread": isUnread,
	})
}

func (chats *chats) UpgradeToSuperGroup(basicGroupId int64) *Chat {
	var chat = object{}
	if chats.client.Load(&chat, Object{
		"@type":   "upgradeBasicGroupChatToSupergroupChat",
		"chat_id": basicGroupId,
	}) {
		return getChat(chat)
	}
	return nil
}

func (chats *chats) TransferOwner(chatId int64, toUserId int32, password string) {
	chats.client.load(object{
		"@type":    "transferChatOwnership",
		"chat_id":  chatId,
		"user_id":  toUserId,
		"password": password,
	})
}

func (chats *chats) SetDescription(chatId int64, description string) {
	chats.client.load(object{
		"@type":       "setChatDescription",
		"chat_id":     chatId,
		"description": description,
	})
}

func (chats *chats) SetLocation(chatId int64, latitude float64, longitude float64, address string) {
	chats.client.load(object{
		"@type":   "setChatLocation",
		"chat_id": chatId,
		"location": object{
			"@type": "chatLocation",
			"location": object{
				"@type":     "location",
				"latitude":  latitude,
				"longitude": longitude,
			},
			"address": address,
		},
	})
}

func (chats *chats) SetPermissions(chatId int64, canSendMessages, canSendMediaMessages,
	canSendPolls, canSendOtherMessages, canAddWebPagePreviews, canChangeInfo,
	canInviteUsers, canPinMessages bool) {
	chats.client.load(object{
		"@type":   "setChatPermissions",
		"chat_id": chatId,
		"permissions": object{
			"@type":                     "chatPermissions",
			"can_send_messages":         canSendMessages,
			"can_send_media_messages":   canSendMediaMessages,
			"can_send_polls":            canSendPolls,
			"can_send_other_messages":   canSendOtherMessages,
			"can_add_web_page_previews": canAddWebPagePreviews,
			"can_change_info":           canChangeInfo,
			"can_invite_users":          canInviteUsers,
			"can_pin_messages":          canPinMessages,
		},
	})
}

func (chats *chats) SetPhoto(chatId int64, path string) {
	chats.client.load(object{
		"@type":   "setChatPhoto",
		"chat_id": chatId,
		"photo": object{
			"@type": "inputFileLocal",
			"path":  path,
		},
	})
}

func (chats *chats) SetTitle(chatId int64, title string) {
	chats.client.load(object{
		"@type":   "setChatTitle",
		"chat_id": chatId,
		"title":   title,
	})
}

func (chats *chats) SetUsername(superChatId int64, username string) bool {
	var check = &class{Type: ""}
	if chats.client.Load(check, Object{
		"@type":    "checkChatUsername",
		"chat_id":  superChatId,
		"username": username,
	}) {
		if check.Type == "checkChatUsernameResultOk" {
			chats.client.load(object{
				"@type":         "setSupergroupUsername",
				"supergroup_id": superChatId,
				"username":      username,
			})
			return true
		}
	}
	return false
}
