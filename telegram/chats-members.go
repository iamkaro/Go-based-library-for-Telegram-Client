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

func (chats *chats) newMembers() {
	chats.Members = &chatsMembers{
		chats: chats,
	}
	/*--------*/
	chats.Members.newSetStatus()

}

type (
	chatsMembers struct {
		chats *chats
		/*--------*/
		SetStatus *chatsMembersSetStatus
	}
)

func (it *chatsMembers) Get(chatId int64, userId int32) *Member {
	var member = object{}
	if it.chats.client.Load(&member, Object{
		"@type":   "getChatMember",
		"chat_id": chatId,
		"user_id": userId,
	}) {
		return getMember(member)
	}
	return nil
}

func (it *chatsMembers) GetAdministrators(SuperChatId int64, offset int32, limit int32) Members {
	return it.getList(SuperChatId, offset, limit, object{
		"@type": "supergroupMembersFilterAdministrators",
	})
}

func (it *chatsMembers) GetBanned(SuperChatId int64, offset int32, limit int32, query string) Members {
	return it.getList(SuperChatId, offset, limit, object{
		"@type": "supergroupMembersFilterBanned",
		"query": query,
	})
}

func (it *chatsMembers) GetBots(SuperChatId int64, offset int32, limit int32) Members {
	return it.getList(SuperChatId, offset, limit, object{
		"@type": "supergroupMembersFilterBots",
	})
}

func (it *chatsMembers) GetContacts(SuperChatId int64, offset int32, limit int32, query string) Members {
	return it.getList(SuperChatId, offset, limit, object{
		"@type": "supergroupMembersFilterContacts",
		"query": query,
	})
}

func (it *chatsMembers) GetRecentlyActive(SuperChatId int64, offset int32, limit int32) Members {
	return it.getList(SuperChatId, offset, limit, object{
		"@type": "supergroupMembersFilterRecent",
	})
}

func (it *chatsMembers) GetRestricted(SuperChatId int64, offset int32, limit int32, query string) Members {
	return it.getList(SuperChatId, offset, limit, object{
		"@type": "supergroupMembersFilterRestricted",
		"query": query,
	})
}

func (it *chatsMembers) GetSearch(SuperChatId int64, offset int32, limit int32, query string) Members {
	return it.getList(SuperChatId, offset, limit, object{
		"@type": "supergroupMembersFilterSearch",
		"query": query,
	})
}

func (it *chatsMembers) Add(chatId int64, userId int32, forwardLimit int32) {
	it.chats.client.load(object{
		"@type":         "addChatMember",
		"chat_id":       chatId,
		"user_id":       userId,
		"forward_limit": forwardLimit,
	})
}

func (it *chatsMembers) AddList(chatId int64, userIds []int32) {
	it.chats.client.load(object{
		"@type":    "addChatMembers",
		"chat_id":  chatId,
		"user_ids": userIds,
	})
}

/*----------------------------------------/         getList         /-----------*/
func (it *chatsMembers) getList(SuperChatId int64, offset int32, limit int32, filter object) Members {
	var Id = int64(0)
	var chat = it.chats.Get(SuperChatId)
	if chat != nil {
		if chat.Type.Channel != nil {
			Id = chat.Type.Channel.ChatId
		}
		if chat.Type.SuperGroup != nil {
			Id = chat.Type.SuperGroup.ChatId
		}
	}
	if Id != 0 {
		var members = object{}
		if it.chats.client.Load(&members, Object{
			"@type":         "getSupergroupMembers",
			"supergroup_id": Id,
			"filter":        filter,
			"offset":        offset,
			"limit":         limit,
		}) {
			return convertList(members.Array("members").object(), getMember)
		}
	}
	return Members{}
}
