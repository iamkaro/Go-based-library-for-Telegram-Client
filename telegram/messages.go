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

func (client *Client) newMessages() {
	client.Messages = &messages{
		client: client,
	}
	/*--------*/
	client.Messages.newCount()
	client.Messages.newDelete()
	client.Messages.newSend()
	client.Messages.newEdit()
	client.Messages.newPin()
	client.Messages.newSearch()

}

type (
	messages struct {
		client *Client
		/*--------*/
		Count  *messagesCount
		Delete *messagesDelete
		Send   *messagesSend
		Edit   *messagesEdit
		Pin    *messagesPin
		Search *messagesSearch
	}
)

func (it *messages) Get(chatId int64, messageId int64) *Message {
	var message = object{}
	if it.client.Load(&message,
		Object{"@type": "getMessage", "chat_id": chatId, "message_id": messageId}) {
		return getMessage(message)
	}
	return nil
}

func (it *messages) GetLastBeforeDate(chatId int64, date int32) *Message {
	var message = object{}
	if it.client.Load(&message,
		Object{"@type": "getChatMessageByDate", "chat_id": chatId, "date": date}) {
		return getMessage(message)
	}
	return nil
}

func (it *messages) GetReplied(chatId int64, replyMessageId int64) *Message {
	var message = object{}
	if it.client.Load(&message,
		Object{"@type": "getRepliedMessage", "chat_id": chatId, "message_id": replyMessageId}) {
		return getMessage(message)
	}
	return nil
}

func (it *messages) GetLocally(chatId int64, messageId int64) *Message {
	var message = object{}
	if it.client.Load(&message,
		Object{"@type": "getMessageLocally", "chat_id": chatId, "message_id": messageId}) {
		return getMessage(message)
	}
	return nil

}

func (it *messages) GetList(chatId int64, messageIds []int64) Messages {
	var all = object{}
	if it.client.Load(&all,
		Object{"@type": "getMessages", "chat_id": chatId, "message_ids": messageIds}) {
		return convertList(all.Array("messages").object(), getMessage)
	}
	return nil
}

func (it *messages) GetChatHistory(chatId int64, fromMessageId int64,
	offset int32, limit int32, onlyLocal bool) Messages {
	var all = object{}
	if it.client.Load(&all, Object{
		"@type":           "getChatHistory",
		"chat_id":         chatId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
		"only_local":      onlyLocal,
	}) {
		return convertList(all.Array("messages").object(), getMessage)
	}
	return nil
}

func (it *messages) GetPublicLink(publicChatId int64, messageId int64, forAlbum bool) *PublicMessageLink {
	var messageLink = object{}
	if it.client.Load(&messageLink, Object{
		"@type":      "getPublicMessageLink",
		"chat_id":    publicChatId,
		"message_id": messageId,
		"for_album":  forAlbum,
	}) {
		return getPublicMessageLink(messageLink)
	}
	return nil
}

func (it *messages) SetDraft(chatId int64, replyToMessageId int64,
	text string, disableWebPagePreview, clearDraft bool) {
	it.client.load(object{
		"@type":   "setChatDraftMessage",
		"chat_id": chatId,
		"draft_message": object{
			"@type":               "draftMessage",
			"reply_to_message_id": replyToMessageId,
			"input_message_text": object{
				"@type":                    "inputMessageText",
				"text":                     object{"@type": "formattedText", "text": text},
				"disable_web_page_preview": disableWebPagePreview,
				"clear_draft":              clearDraft,
			},
		},
	})
}

func (it *messages) ReadAll(chatId int64) {
	it.client.load(object{
		"@type":   "readAllChatMentions",
		"chat_id": chatId,
	})
}

func (it *messages) View(chatId int64, messageIds []int64, forceRead bool) {
	it.client.load(object{
		"@type":       "viewMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"force_read":  forceRead,
	})
}

func (it *messages) OpenContent(chatId int64, messageId int64) {
	it.client.load(object{
		"@type":      "openMessageContent",
		"chat_id":    chatId,
		"message_id": messageId,
	})
}
