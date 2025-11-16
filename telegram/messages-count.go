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

func (it *messages) newCount() {
	it.Count = &messagesCount{
		messages: it,
	}
}

type (
	messagesCount struct {
		messages *messages
	}
)

func (it *messagesCount) Animations(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterAnimation", returnLocal)
}

func (it *messagesCount) Audios(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterAudio", returnLocal)
}

func (it *messagesCount) Calls(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterCall", returnLocal)
}

func (it *messagesCount) ChatPhotos(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterChatPhoto", returnLocal)
}

func (it *messagesCount) Documents(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterDocument", returnLocal)
}

func (it *messagesCount) Photos(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterPhoto", returnLocal)
}

func (it *messagesCount) MissedCalls(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterMissedCall", returnLocal)
}

func (it *messagesCount) Mentions(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterMention", returnLocal)
}

func (it *messagesCount) PhotoAndVideos(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterPhotoAndVideo", returnLocal)
}

func (it *messagesCount) UnreadMentions(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterUnreadMention", returnLocal)
}

func (it *messagesCount) Urls(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterUrl", returnLocal)
}

func (it *messagesCount) Videos(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterVideo", returnLocal)
}

func (it *messagesCount) VideoNotes(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterVideoNote", returnLocal)
}

func (it *messagesCount) VoiceAndVideoNotes(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterVoiceAndVideoNote", returnLocal)
}

func (it *messagesCount) VoiceNotes(chatId int64, returnLocal bool) int32 {
	return it.count(chatId, "searchMessagesFilterVoiceNote", returnLocal)
}

/*----------------------------------------/         count         /-----------*/
func (it *messagesCount) count(chatId int64, filter string, returnLocal bool) int32 {
	var count = object{}
	if it.messages.client.Load(&count, Object{
		"@type":        "getChatMessageCount",
		"chat_id":      chatId,
		"filter":       object{"@type": filter},
		"return_local": returnLocal,
	}) {
		return count.int32("count")
	}
	return 0
}
