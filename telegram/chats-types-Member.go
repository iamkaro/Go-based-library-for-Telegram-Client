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

func getMember(value object) *Member {
	if CheckType(value, "chatMember") {
		return &Member{
			UserId:         value.Object("member_id").int64("user_id"),
			InviterUserId:  value.int64("inviter_user_id"),
			JoinedChatDate: value.int32("joined_chat_date"),
			Status:         getMemberStatus(value.Object("status")),
			BotInfo:        getBotInfo(value.Object("bot_info")),
		}
	}
	return nil
}

func getMemberStatus(value object) *memberStatus {
	if value != nil {
		return &memberStatus{
			Administrator: getMemberStatusAdministrator(value),
			Banned:        getMemberStatusBanned(value),
			Creator:       getMemberStatusCreator(value),
			Restricted:    getMemberStatusRestricted(value),
			Left:          value.string("@type") == "chatMemberStatusLeft",
			Member:        value.string("@type") == "chatMemberStatusMember",
		}
	}
	return nil
}

func getMemberStatusAdministrator(value object) *memberStatusAdministrator {
	if CheckType(value, "chatMemberStatusAdministrator") {
		return &memberStatusAdministrator{
			CustomTitle:        value.string("custom_title"),
			CanBeEdited:        value.bool("can_be_edited"),
			CanChangeInfo:      value.bool("can_change_info"),
			CanPostMessages:    value.bool("can_post_messages"),
			CanEditMessages:    value.bool("can_edit_messages"),
			CanDeleteMessages:  value.bool("can_delete_messages"),
			CanInviteUsers:     value.bool("can_invite_users"),
			CanRestrictMembers: value.bool("can_restrict_members"),
			CanPinMessages:     value.bool("can_pin_messages"),
			CanPromoteMembers:  value.bool("can_promote_members"),
		}
	}
	return nil
}

func getMemberStatusBanned(value object) *memberStatusBanned {
	if CheckType(value, "chatMemberStatusBanned") {
		return &memberStatusBanned{UntilDate: value.int32("banned_until_date")}
	}
	return nil
}

func getMemberStatusCreator(value object) *memberStatusCreator {
	if CheckType(value, "chatMemberStatusCreator") {
		return &memberStatusCreator{
			CustomTitle: value.string("custom_title"),
			IsMember:    value.bool("is_member"),
		}
	}
	return nil
}

func getMemberStatusRestricted(value object) *memberStatusRestricted {
	if CheckType(value, "chatMemberStatusRestricted") {
		return &memberStatusRestricted{
			IsMember:    value.bool("is_member"),
			UntilDate:   value.int32("restricted_until_date"),
			Permissions: getChatPermissions(value.Object("permissions")),
		}
	}
	return nil
}

type (
	Members []*Member
	Member  struct {
		UserId         int64
		InviterUserId  int64
		JoinedChatDate int32
		Status         *memberStatus
		BotInfo        *botInfo
	}
	memberStatus struct {
		Administrator *memberStatusAdministrator
		Banned        *memberStatusBanned
		Creator       *memberStatusCreator
		Restricted    *memberStatusRestricted
		Left          bool
		Member        bool
	}
	memberStatusAdministrator struct {
		CustomTitle        string
		CanBeEdited        bool
		CanChangeInfo      bool
		CanPostMessages    bool
		CanEditMessages    bool
		CanDeleteMessages  bool
		CanInviteUsers     bool
		CanRestrictMembers bool
		CanPinMessages     bool
		CanPromoteMembers  bool
	}

	memberStatusBanned  struct{ UntilDate int32 }
	memberStatusCreator struct {
		CustomTitle string
		IsMember    bool
	}
	memberStatusRestricted struct {
		IsMember    bool
		UntilDate   int32
		Permissions *chatPermissions
	}
)
