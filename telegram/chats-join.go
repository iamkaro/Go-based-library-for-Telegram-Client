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

func (chats *chats) newJoin() {
	chats.Join = &chatsJoin{
		chats: chats,
	}
}

type (
	chatsJoin struct {
		chats *chats
	}
)

func (it *chatsJoin) ByChatId(chatId int64) {
	it.chats.client.load(object{
		"@type":   "joinChat",
		"chat_id": chatId,
	})
}

func (it *chatsJoin) ByInviteLink(inviteLink string) *Chat {
	var chat = object{}
	if it.chats.client.Load(&chat, Object{"@type": "joinChatByInviteLink", "invite_link": inviteLink}) {
		return getChat(chat)
	}
	return nil
}
