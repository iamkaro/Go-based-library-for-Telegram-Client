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

func (it *messages) newDelete() {
	it.Delete = &messagesDelete{
		messages: it,
	}
}

type (
	messagesDelete struct {
		messages *messages
	}
)

func (it *messagesDelete) List(chatId int64, messageIds []int64, forAll bool) {
	it.messages.client.load(object{
		"@type":       "deleteMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"revoke":      forAll,
	})
}

func (it *messagesDelete) AllFromUser(chatId int64, userId int32) {
	it.messages.client.load(object{
		"@type":   "deleteChatMessagesFromUser",
		"chat_id": chatId,
		"user_id": userId,
	})
}

func (it *messagesDelete) ChatHistory(chatId int64, forAll bool, removeChat bool) {
	it.messages.client.load(object{
		"@type":                 "deleteChatHistory",
		"chat_id":               chatId,
		"remove_from_chat_list": removeChat,
		"revoke":                forAll,
	})
}
