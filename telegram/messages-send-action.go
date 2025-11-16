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

func (it *messagesSend) newChatAction() {
	it.ChatAction = &messagesSendChatAction{
		messagesSend: it,
	}
}

type (
	messagesSendChatAction struct {
		messagesSend *messagesSend
	}
)

func (it *messagesSendChatAction) Cancel(chatId int64) {
	it.send(chatId, "chatActionCancel")
}

func (it *messagesSendChatAction) ChoosingContact(chatId int64) {
	it.send(chatId, "chatActionChoosingContact")
}

func (it *messagesSendChatAction) ChoosingLocation(chatId int64) {
	it.send(chatId, "chatActionChoosingLocation")
}

func (it *messagesSendChatAction) RecordingVideo(chatId int64) {
	it.send(chatId, "chatActionRecordingVideo")
}

func (it *messagesSendChatAction) RecordingVideoNote(chatId int64) {
	it.send(chatId, "chatActionRecordingVideoNote")
}

func (it *messagesSendChatAction) RecordingVoiceNote(chatId int64) {
	it.send(chatId, "chatActionRecordingVoiceNote")
}

func (it *messagesSendChatAction) StartPlayingGame(chatId int64) {
	it.send(chatId, "chatActionStartPlayingGame")
}

func (it *messagesSendChatAction) Typing(chatId int64) {
	it.send(chatId, "chatActionTyping")
}

func (it *messagesSendChatAction) UploadingDocument(chatId int64) {
	it.send(chatId, "chatActionUploadingDocument")
}

func (it *messagesSendChatAction) UploadingPhoto(chatId int64) {
	it.send(chatId, "chatActionUploadingPhoto")
}

func (it *messagesSendChatAction) UploadingVideo(chatId int64) {
	it.send(chatId, "chatActionUploadingVideo")
}

func (it *messagesSendChatAction) UploadingVideoNote(chatId int64) {
	it.send(chatId, "chatActionUploadingVideoNote")
}

func (it *messagesSendChatAction) UploadingVoiceNote(chatId int64) {
	it.send(chatId, "chatActionUploadingVoiceNote")
}

/*----------------------------------------/         send         /-----------*/
func (it *messagesSendChatAction) send(chatId int64, action string) {
	it.messagesSend.messages.client.load(object{
		"@type":   "sendChatAction",
		"chat_id": chatId,
		"action":  object{"@type": action},
	})
}
