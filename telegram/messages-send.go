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
	"time"
)

func (it *messages) newSend() {
	it.Send = &messagesSend{
		messages: it,
		updates:  nil,
	}
	/*--------*/
	it.Send.newChatAction()

}

type (
	messagesSend struct {
		messages *messages
		updates  *updateData
		/*--------*/
		ChatAction *messagesSendChatAction
	}
)

func (it *messagesSend) BotStart(botUserId int32, chatId int64, parameter string) *Message {
	var message = object{}
	if it.messages.client.Load(&message, Object{
		"@type":       "sendBotStartMessage",
		"bot_user_id": botUserId,
		"chat_id":     chatId,
		"parameter":   parameter,
	}) {
		return getMessage(message)
	}
	return nil
}

func (it *messagesSend) ScreenshotTaken(chatId int64) {
	it.messages.client.load(object{
		"@type":   "sendChatScreenshotTakenNotification",
		"chat_id": chatId,
	})
}

func (it *messagesSend) Audio(chatId int64, replyToMessageId int64, disableNotification bool,
	path string, duration int32, title string, caption string) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type":    "inputMessageAudio",
		"audio":    object{"@type": "inputFileLocal", "path": path},
		"duration": duration,
		"title":    title,
		"caption":  object{"@type": "formattedText", "text": caption},
		/*"album_cover_thumbnail": nil,*/
		/*"performer": "",*/
	})
}

func (it *messagesSend) Contact(chatId int64, replyToMessageId int64, disableNotification bool,
	phoneNumber string, firstName string, lastName string) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type": "inputMessageContact",
		"contact": object{
			"@type":        "contact",
			"phone_number": phoneNumber,
			"first_name":   firstName,
			"last_name":    lastName,
			"vcard":        "",
			"user_id":      0,
		},
	})
}

func (it *messagesSend) Document(chatId int64, replyToMessageId int64,
	disableNotification bool, path string, caption string) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type":    "inputMessageDocument",
		"document": object{"@type": "inputFileLocal", "path": path},
		"caption":  object{"@type": "formattedText", "text": caption},
		/*"thumbnail": nil,*/
	})
}

func (it *messagesSend) Forwarded(chatId int64, replyToMessageId int64,
	disableNotification bool, fromChatId int64, messageId int64, inGameShare bool,
	sendCopy bool, removeCaption bool) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type":          "inputMessageForwarded",
		"from_chat_id":   fromChatId,
		"message_id":     messageId,
		"in_game_share":  inGameShare,
		"send_copy":      sendCopy,
		"remove_caption": removeCaption,
	})
}

func (it *messagesSend) Location(chatId int64, replyToMessageId int64,
	disableNotification bool, latitude float64, longitude float64) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type": "inputMessageLocation",
		"location": object{
			"@type":     "location",
			"latitude":  latitude,
			"longitude": longitude,
		},
		"live_period": 0,
	})
}

func (it *messagesSend) Photo(chatId int64, replyToMessageId int64,
	disableNotification bool, path string, width int32, height int32,
	caption string) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
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

func (it *messagesSend) PollRegular(chatId int64, replyToMessageId int64,
	disableNotification bool, question string, options []string, isAnonymous,
	multipleAnswers, isClosed bool) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type":        "inputMessagePoll",
		"question":     question,
		"options":      options,
		"is_anonymous": isAnonymous,
		"type": object{
			"@type":                  "pollTypeRegular",
			"allow_multiple_answers": multipleAnswers,
		},
		"is_closed": isClosed,
	})
}

func (it *messagesSend) Text(chatId int64, replyToMessageId int64,
	disableNotification bool, text string, disableWebPagePreview, clearDraft bool) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type":                "inputMessageText",
		"text":                 object{"@type": "formattedText", "text": text, "entities": array{}},
		"link_preview_options": object{"@type": "linkPreviewOptions", "is_disabled": disableWebPagePreview},
		"clear_draft":          clearDraft,
	})
}

func (it *messagesSend) Video(chatId int64, replyToMessageId int64,
	disableNotification bool, path string, duration, width, height int32,
	supportsStreaming bool, caption string) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
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

func (it *messagesSend) VideoNote(chatId int64, replyToMessageId int64,
	disableNotification bool, path string, duration, length int32) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type":      "inputMessageVideoNote",
		"video_note": object{"@type": "inputFileLocal", "path": path},
		"duration":   duration,
		"length":     length,
		/*"thumbnail": nil,*/
	})
}

func (it *messagesSend) VoiceNote(chatId int64, replyToMessageId int64,
	disableNotification bool, path string, duration int32, caption string) *Message {
	return it.send(chatId, replyToMessageId, disableNotification, object{
		"@type":      "inputMessageVoiceNote",
		"voice_note": object{"@type": "inputFileLocal", "path": path},
		"duration":   duration,
		"caption":    object{"@type": "formattedText", "text": caption},
		/*"waveform": "",*/
	})
}

/*----------------------------------------/         loadings         /-----------*/
func (it *messagesSend) send(chatId int64, replyToMessageId int64,
	disableNotification bool, messageContent object) *Message {
	var messageData = object{}
	if it.messages.client.Load(&messageData, Object{
		"@type":             "sendMessage",
		"chat_id":           chatId,
		"message_thread_id": 0,
		"reply_to": object{
			"@type":             "inputMessageReplyToMessage",
			"message_id":        replyToMessageId,
			"checklist_task_id": 0,
		},
		"options": object{
			"@type":                "messageSendOptions",
			"disable_notification": disableNotification,
		},
		"input_message_content": messageContent,
	}) {
		var state = messageData.Object("sending_state")
		if (state != nil) && (state.string("@type") == "messageSendingStatePending") {
			var loops uint16 = 0
			for it.updates == nil {
				loops += 1
				if loops > 800 {
					break
				}
				sleep(250)
			}
			/*-----------*/
			if it.updates != nil {
				var message = getMessage(messageData)
				var update = it.updates
				var ok = update.Update.int64("old_message_id") == message.Id
				for !ok {
					for (!ok) && (update.Next != nil) {
						update = update.Next
						ok = update.Update.int64("old_message_id") == message.Id
					}
					if !ok {
						loops += 1
						if loops > 800 {
							break
						}
						sleep(250)
					}
				}
				if ok && (update.Update.string("@type") == "updateMessageSendSucceeded") {
					return getMessage(update.Update.Object("message"))
				}
			}
			/*-----------*/
		}
	}
	return nil
}

func (it *messagesSend) update(data data) {
	var timeNow = time.Now().Unix()
	var index = it.updates
	for (index != nil) && (index.DeleteTime < timeNow) {
		index = index.Next
	}
	it.updates = index
	/*------------*/
	var update = &updateData{
		Update:     object{},
		DeleteTime: timeNow + 400,
		Next:       nil,
	}
	if data.extractTo(&(update.Update)) == nil {
		if index == nil {
			it.updates = update
		} else {
			for index.Next != nil {
				index = index.Next
			}
			index.Next = update
		}
	}
}
