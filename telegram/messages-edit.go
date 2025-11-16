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

func (it *messages) newEdit() {
	it.Edit = &messagesEdit{
		messages: it,
	}
}

type (
	messagesEdit struct {
		messages *messages
	}
)

func (it *messagesEdit) Text(chatId int64, messageId int64,
	text string, disableWebPagePreview, clearDraft bool) *Message {
	var message = object{}
	if it.messages.client.Load(&message, Object{
		"@type":      "editMessageText",
		"chat_id":    chatId,
		"message_id": messageId,
		"input_message_content": object{
			"@type":                    "inputMessageText",
			"text":                     object{"@type": "formattedText", "text": text},
			"disable_web_page_preview": disableWebPagePreview,
			"clear_draft":              clearDraft,
		},
		/*"reply_markup": nil,*/
	}) {
		return getMessage(message)
	}
	return nil
}

func (it *messagesEdit) Caption(chatId int64, messageId int64, caption string) *Message {
	var message = object{}
	if it.messages.client.Load(&message, Object{
		"@type":      "editMessageCaption",
		"chat_id":    chatId,
		"message_id": messageId,
		"caption":    object{"@type": "formattedText", "text": caption},
		/*"reply_markup": nil,*/
	}) {
		return getMessage(message)
	}
	return nil
}

func (it *messagesEdit) Audio(chatId int64, messageId int64,
	path string, duration int32, title string, caption string) *Message {
	return it.editMedia(chatId, messageId, object{
		"@type":    "inputMessageAudio",
		"audio":    object{"@type": "inputFileLocal", "path": path},
		"duration": duration,
		"title":    title,
		"caption":  object{"@type": "formattedText", "text": caption},
		/*"album_cover_thumbnail": nil,*/
		/*"performer": "",*/
	})
}

func (it *messagesEdit) Document(chatId int64, messageId int64, path string, caption string) *Message {
	return it.editMedia(chatId, messageId, object{
		"@type":    "inputMessageDocument",
		"document": object{"@type": "inputFileLocal", "path": path},
		"caption":  object{"@type": "formattedText", "text": caption},
		/*"thumbnail": nil,*/
	})
}

func (it *messagesEdit) Photo(chatId int64, messageId int64,
	path string, width int32, height int32, caption string) *Message {
	return it.editMedia(chatId, messageId, object{
		"@type":   "inputMessagePhoto",
		"photo":   object{"@type": "inputFileLocal", "path": path},
		"width":   width,
		"height":  height,
		"caption": object{"@type": "formattedText", "text": caption},
		/*"thumbnail": nil,*/
		/*"added_sticker_file_ids": nil,*/
		/*"ttl": nil,*/
	})
}

func (it *messagesEdit) Video(chatId int64, messageId int64,
	path string, duration, width, height int32, supportsStreaming bool, caption string) *Message {
	return it.editMedia(chatId, messageId, object{
		"@type":              "inputMessageVideo",
		"video":              object{"@type": "inputFileLocal", "path": path},
		"duration":           duration,
		"width":              width,
		"height":             height,
		"supports_streaming": supportsStreaming,
		"caption":            object{"@type": "formattedText", "text": caption},
		/*"thumbnail": nil,*/
		/*"added_sticker_file_ids": nil,*/
		/*"ttl": nil,*/
	})
}

/*----------------------------------------/         editMedia         /-----------*/
func (it *messagesEdit) editMedia(chatId int64, messageId int64, messageContent object) *Message {
	var message = object{}
	if it.messages.client.Load(&message, Object{
		"@type":                 "editMessageMedia",
		"chat_id":               chatId,
		"message_id":            messageId,
		"input_message_content": messageContent,
		/*"reply_markup": nil,*/
	}) {
		return getMessage(message)
	}
	return nil
}
