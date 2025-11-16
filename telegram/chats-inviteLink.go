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

func (chats *chats) newInviteLink() {
	chats.InviteLink = &chatsInviteLink{
		chats: chats,
	}
}

type (
	chatsInviteLink struct {
		chats *chats
	}
)

func (it *chatsInviteLink) Generate(chatId int64) *chatInviteLink {
	var link = object{}
	if it.chats.client.Load(&link, Object{"@type": "generateChatInviteLink", "chat_id": chatId}) {
		return getChatInviteLink(link)
	}
	return nil
}

func (it *chatsInviteLink) Check(inviteLink string) *chatInviteLinkInfo {
	var info = object{}
	if it.chats.client.Load(&info, Object{"@type": "checkChatInviteLink", "invite_link": inviteLink}) {
		return getChatInviteLinkInfo(info)
	}
	return nil
}

/*----------------------------------------/         items         /-----------*/
type (
	chatInviteLink     struct{ InviteLink string }
	chatInviteLinkInfo struct {
		ChatId        int64
		Type          *chatType
		Title         string
		Photo         *chatPhoto
		MemberCount   int32
		MemberUserIds []int32
		IsPublic      bool
	}
)

func getChatInviteLink(value object) *chatInviteLink {
	if CheckType(value, "chatInviteLink") {
		return &chatInviteLink{InviteLink: value.string("invite_link")}
	}
	return nil
}

func getChatInviteLinkInfo(value object) *chatInviteLinkInfo {
	if CheckType(value, "chatInviteLinkInfo") {
		return &chatInviteLinkInfo{
			ChatId:        value.int64("chat_id"),
			Type:          getChatType(value.Object("type")),
			Title:         value.string("title"),
			Photo:         getChatPhoto(value.Object("photo")),
			MemberCount:   value.int32("member_count"),
			MemberUserIds: value.Array("member_user_ids").int32(),
			IsPublic:      value.bool("is_public"),
		}
	}
	return nil
}
