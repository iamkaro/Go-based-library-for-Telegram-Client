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

func (it *messages) newSearch() {
	it.Search = &messagesSearch{
		messages: it,
	}
}

type (
	messagesSearch struct {
		messages *messages
	}
)

func (it *messagesSearch) Full(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterEmpty")
}

func (it *messagesSearch) Animations(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterAnimation")
}

func (it *messagesSearch) Audios(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterAudio")
}

func (it *messagesSearch) Calls(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterCall")
}

func (it *messagesSearch) ChatPhotos(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterChatPhoto")
}

func (it *messagesSearch) Documents(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterDocument")
}

func (it *messagesSearch) Mentions(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterMention")
}

func (it *messagesSearch) MissedCalls(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterMissedCall")
}

func (it *messagesSearch) Photos(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterPhoto")
}

func (it *messagesSearch) PhotoAndVideos(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterPhotoAndVideo")
}

func (it *messagesSearch) UnreadMentions(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterUnreadMention")
}

func (it *messagesSearch) Urls(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterUrl")
}

func (it *messagesSearch) Videos(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterVideo")
}

func (it *messagesSearch) VideoNotes(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterVideoNote")
}

func (it *messagesSearch) VoiceAndVideoNotes(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterVoiceAndVideoNote")
}

func (it *messagesSearch) VoiceNotes(chatId int64, query string,
	senderUserId int32, fromMessageId int64, offset int32, limit int32) []*Message {
	return it.search(chatId, query, senderUserId, fromMessageId, offset, limit,
		"searchMessagesFilterVoiceNote")
}

/*----------------------------------------/         search         /-----------*/
func (it *messagesSearch) search(chatId int64, query string, senderUserId int32,
	fromMessageId int64, offset int32, limit int32, filter string) []*Message {
	var all = object{}
	if it.messages.client.Load(&all, Object{
		"@type":           "searchChatMessages",
		"chat_id":         chatId,
		"query":           query,
		"sender_user_id":  senderUserId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
		"filter":          object{"@type": filter},
	}) {
		return convertList(all.Array("messages").object(), getMessage)
	}
	return nil
}
