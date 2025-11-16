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

func (chats *chats) newGetList() {
	chats.GetList = &chatsGetList{
		chats: chats,
	}
}

type (
	chatsGetList struct {
		chats *chats
	}
)

func (it *chatsGetList) AllOfMainList(limit int32) []int64 {
	return it.getAll("chatListMain", limit)
}

func (it *chatsGetList) AllOfArchiveList(limit int32) []int64 {
	return it.getAll("chatListArchive", limit)
}

func (it *chatsGetList) Bots(limit int32) []int64 {
	return it.get("topChatCategoryBots", limit)
}

func (it *chatsGetList) Calls(limit int32) []int64 {
	return it.get("topChatCategoryCalls", limit)
}

func (it *chatsGetList) Channels(limit int32) []int64 {
	return it.get("topChatCategoryChannels", limit)
}

func (it *chatsGetList) ForwardChats(limit int32) []int64 {
	return it.get("topChatCategoryForwardChats", limit)
}

func (it *chatsGetList) Groups(limit int32) []int64 {
	return it.get("topChatCategoryGroups", limit)
}

func (it *chatsGetList) InlineBots(limit int32) []int64 {
	return it.get("topChatCategoryInlineBots", limit)
}

func (it *chatsGetList) Users(limit int32) []int64 {
	return it.get("topChatCategoryUsers", limit)
}

/*-------------------------/          get         /-----------*/
func (it *chatsGetList) get(category string, limit int32) []int64 {
	var chats = object{}
	if (it.chats.client.Load(&chats, Object{
		"@type":    "getTopChats",
		"category": object{"@type": category},
		"limit":    limit,
	})) {
		return chats.Array("chat_ids").int64()
	}
	return nil
}

func (it *chatsGetList) getAll(list string, limit int32) []int64 {
	var chats = object{}
	if (it.chats.client.Load(&chats, Object{
		"@type":     "getChats",
		"chat_list": object{"@type": list},
		"limit":     limit,
	})) {
		return chats.Array("chat_ids").int64()
	}
	return nil
}
