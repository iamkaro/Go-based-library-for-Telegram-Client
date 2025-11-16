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

func getChat(value object) *Chat {
	if CheckType(value, "chat") {
		return &Chat{
			Id:                         value.int64("id"),
			Type:                       getChatType(value.Object("type")),
			ChatList:                   getChatList(value.Object("chat_list")),
			Title:                      value.string("title"),
			Photo:                      getChatPhoto(value.Object("photo")),
			Permissions:                getChatPermissions(value.Object("permissions")),
			LastMessage:                getMessage(value.Object("last_message")),
			Order:                      value.int64("order"),
			IsPinned:                   value.bool("is_pinned"),
			IsMarkedAsUnread:           value.bool("is_marked_as_unread"),
			IsSponsored:                value.bool("is_sponsored"),
			HasScheduledMessages:       value.bool("has_scheduled_messages"),
			CanBeDeletedOnlyForSelf:    value.bool("can_be_deleted_only_for_self"),
			CanBeDeletedForAllUsers:    value.bool("can_be_deleted_for_all_users"),
			DefaultDisableNotification: value.bool("default_disable_notification"),
			UnreadCount:                value.int32("unread_count"),
			UnreadCountMention:         value.int32("unread_mention_count"),
			LastReadMessageIdInbox:     value.int64("last_read_inbox_message_id"),
			LastReadMessageIdOutbox:    value.int64("last_read_outbox_message_id"),
			PinnedMessageId:            value.int64("pinned_message_id"),
			ReplyMarkupMessageId:       value.int64("reply_markup_message_id"),
		}
	}
	return nil
}

func getChatType(value object) *chatType {
	if value != nil {
		var chatType = &chatType{
			Group:      nil,
			SuperGroup: nil,
			Channel:    nil,
			Private:    nil,
			SecretChat: nil,
		}
		switch value.string("@type") {
		case "chatTypeBasicGroup":
			chatType.Group = &chatId{ChatId: value.int64("basic_group_id")}
		case "chatTypeSupergroup":
			if value.bool("is_channel") {
				chatType.Channel = &chatId{ChatId: value.int64("supergroup_id")}
			} else {
				chatType.SuperGroup = &chatId{ChatId: value.int64("supergroup_id")}
			}
		case "chatTypePrivate":
			chatType.Private = &userId{UserId: value.int64("user_id")}
		case "chatTypeSecret":
			chatType.SecretChat = &userId{UserId: value.int64("user_id")}
		}
		return chatType
	}
	return nil
}

func getChatList(value object) *chatList {
	if value != nil {
		return &chatList{
			Main:    value.string("@type") == "chatListMain",
			Archive: value.string("@type") == "chatListArchive",
		}
	}
	return nil
}

func getChatPhoto(value object) *chatPhoto {
	if CheckType(value, "chatPhoto") {
		return &chatPhoto{
			Small: getFile(value.Object("small")),
			Big:   getFile(value.Object("small")),
		}
	}
	return nil
}

func getChatPermissions(value object) *chatPermissions {
	if CheckType(value, "chatPermissions") {
		return &chatPermissions{
			CanSendMessages:       value.bool("can_send_messages"),
			CanSendMediaMessages:  value.bool("can_send_media_messages"),
			CanSendPolls:          value.bool("can_send_polls"),
			CanSendOtherMessages:  value.bool("can_send_other_messages"),
			CanAddWebPagePreviews: value.bool("can_add_web_page_previews"),
			CanChangeInfo:         value.bool("can_change_info"),
			CanInviteUsers:        value.bool("can_invite_users"),
			CanPinMessages:        value.bool("can_pin_messages"),
		}
	}
	return nil
}

type (
	userId   struct{ UserId int64 }
	chatId   struct{ ChatId int64 }
	ChatList map[int64]*Chat
	Chats    []*Chat
	Chat     struct {
		Id                         int64
		Type                       *chatType
		ChatList                   *chatList
		Title                      string
		Photo                      *chatPhoto
		Permissions                *chatPermissions
		LastMessage                *Message
		Order                      int64
		IsPinned                   bool
		IsMarkedAsUnread           bool
		IsSponsored                bool
		HasScheduledMessages       bool
		CanBeDeletedOnlyForSelf    bool
		CanBeDeletedForAllUsers    bool
		DefaultDisableNotification bool
		UnreadCount                int32
		UnreadCountMention         int32
		LastReadMessageIdInbox     int64
		LastReadMessageIdOutbox    int64
		PinnedMessageId            int64
		ReplyMarkupMessageId       int64
	}
	chatType struct {
		Group      *chatId
		SuperGroup *chatId
		Channel    *chatId
		Private    *userId
		SecretChat *userId
	}
	chatList struct {
		Main    bool
		Archive bool
	}
	chatPhoto struct {
		Small *File
		Big   *File
	}
	chatPermissions struct {
		CanSendMessages       bool
		CanSendMediaMessages  bool
		CanSendPolls          bool
		CanSendOtherMessages  bool
		CanAddWebPagePreviews bool
		CanChangeInfo         bool
		CanInviteUsers        bool
		CanPinMessages        bool
	}
)

func (it *chatType) String() string {
	switch {
	case it.Private != nil:
		return "private"
	case it.Group != nil:
		return "group"
	case it.SuperGroup != nil:
		return "super-group"
	case it.Channel != nil:
		return "channel"
	case it.SecretChat != nil:
		return "secret-chat"
	}
	return ""
}
