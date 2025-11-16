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

func (chats *chats) newSetChatLists() {
	chats.SetChatLists = &chatsSetChatLists{
		chats: chats,
	}
}

type (
	chatsSetChatLists struct {
		chats *chats
	}
)

func (it *chatsSetChatLists) Main(chatId int64) {
	it.chats.client.load(object{
		"@type":     "setChatChatList",
		"chat_id":   chatId,
		"chat_list": object{"@type": "chatListMain"},
	})
}

func (it *chatsSetChatLists) Archive(chatId int64) {
	it.chats.client.load(object{
		"@type":     "setChatChatList",
		"chat_id":   chatId,
		"chat_list": object{"@type": "chatListArchive"},
	})
}
