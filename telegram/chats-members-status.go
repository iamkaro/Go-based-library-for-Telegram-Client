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

func (it *chatsMembers) newSetStatus() {
	it.SetStatus = &chatsMembersSetStatus{
		chatsMembers: it,
	}
}

type (
	chatsMembersSetStatus struct {
		chatsMembers *chatsMembers
	}
)

func (it *chatsMembersSetStatus) Administrator(chatId int64, userId int32,
	customTitle string, canBeEdited, canChangeInfo, canPostMessages, canEditMessages,
	canDeleteMessages, canInviteUsers, canRestrictMembers, canPinMessages,
	canPromoteMembers bool) {
	it.set(chatId, userId, object{
		"@type":                "chatMemberStatusAdministrator",
		"custom_title":         customTitle,
		"can_be_edited":        canBeEdited,
		"can_change_info":      canChangeInfo,
		"can_post_messages":    canPostMessages,
		"can_edit_messages":    canEditMessages,
		"can_delete_messages":  canDeleteMessages,
		"can_invite_users":     canInviteUsers,
		"can_restrict_members": canRestrictMembers,
		"can_pin_messages":     canPinMessages,
		"can_promote_members":  canPromoteMembers,
	})
}

func (it *chatsMembersSetStatus) Banned(chatId int64, userId int32,
	untilDate int32) {
	it.set(chatId, userId, object{
		"@type":             "chatMemberStatusBanned",
		"banned_until_date": untilDate,
	})
}

func (it *chatsMembersSetStatus) Creator(chatId int64, userId int32,
	customTitle string) {
	it.set(chatId, userId, object{
		"@type":        "chatMemberStatusCreator",
		"custom_title": customTitle,
		/*"is_member": 0,*/
	})
}

func (it *chatsMembersSetStatus) Left(chatId int64, userId int32) {
	it.set(chatId, userId, object{
		"@type": "chatMemberStatusLeft",
	})
}

func (it *chatsMembersSetStatus) Member(chatId int64, userId int32) {
	it.set(chatId, userId, object{
		"@type": "chatMemberStatusMember",
	})
}

func (it *chatsMembersSetStatus) Restricted(chatId int64, userId int32,
	untilDate int32, canSendMessages, canSendMediaMessages, canSendPolls,
	canSendOtherMessages, canAddWebPagePreviews, canChangeInfo, canInviteUsers,
	canPinMessages bool) {
	it.set(chatId, userId, object{
		"@type":                 "chatMemberStatusRestricted",
		"restricted_until_date": untilDate,
		"permissions": object{
			"@type":                     "chatPermissions",
			"can_send_messages":         canSendMessages,
			"can_send_media_messages":   canSendMediaMessages,
			"can_send_polls":            canSendPolls,
			"can_send_other_messages":   canSendOtherMessages,
			"can_add_web_page_previews": canAddWebPagePreviews,
			"can_change_info":           canChangeInfo,
			"can_invite_users":          canInviteUsers,
			"can_pin_messages":          canPinMessages,
		},
	})
}

/*----------------------------------------/         set         /-----------*/
func (it *chatsMembersSetStatus) set(chatId int64, userId int32, status object) {
	it.chatsMembers.chats.client.load(object{
		"@type":   "setChatMemberStatus",
		"chat_id": chatId,
		"user_id": userId,
		"status":  status,
	})
}
