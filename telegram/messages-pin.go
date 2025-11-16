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

func (it *messages) newPin() {
	it.Pin = &messagesPin{
		messages: it,
	}
}

type (
	messagesPin struct {
		messages *messages
	}
)

func (it *messagesPin) Get(chatId int64) *Message {
	var message = object{}
	if it.messages.client.Load(&message, Object{"@type": "getChatPinnedMessage", "chat_id": chatId}) {
		return getMessage(message)
	}
	return nil
}

func (it *messagesPin) Set(chatId int64, messageId int64, disableNotification bool) {
	it.messages.client.load(object{
		"@type":                "pinChatMessage",
		"chat_id":              chatId,
		"message_id":           messageId,
		"disable_notification": disableNotification,
	})
}

func (it *messagesPin) Remove(chatId int64) {
	it.messages.client.load(object{
		"@type":   "unpinChatMessage",
		"chat_id": chatId,
	})
}
