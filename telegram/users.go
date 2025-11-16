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

func (client *Client) newUsers() {
	client.Users = &users{
		client: client,
	}
}

type (
	users struct {
		client *Client
	}
)

func (users *users) Get(userId int64) *User {
	var user = object{}
	if users.client.Load(&user, Object{"@type": "getUser", "user_id": userId}) {
		return getUser(user)
	}
	return nil

}

func (users *users) GetFullInfo(userId int64) *UserFullInfo {
	var userFullInfo = object{}
	if users.client.Load(&userFullInfo, Object{"@type": "getUserFullInfo", "user_id": userId}) {
		return getUserFullInfo(userFullInfo)
	}
	return nil
}

func (users *users) GetProfilePhotos(userId int64, offset int, limit int) []*UserProfilePhoto {
	var all = object{}
	if users.client.Load(&all, Object{
		"@type":   "getUserProfilePhotos",
		"user_id": userId,
		"offset":  offset,
		"limit":   limit,
	}) {
		return convertList(all.Array("photos").object(), getUserProfilePhoto)
	}
	return nil
}

/*----------------------------------------/         items         /-----------*/
type (
	Users []*User
	User  struct {
		Id              int64
		FirstName       string
		LastName        string
		Username        string
		PhoneNumber     string
		Status          *userStatus
		ProfilePhoto    *profilePhoto
		IsContact       bool
		IsMutualContact bool
		IsVerified      bool
		IsSupport       bool
		HaveAccess      bool
		Type            *userType
		LanguageCode    string
	}
	userStatus struct {
		Empty     bool
		LastMonth bool
		LastWeek  bool
		Offline   bool
		Online    bool
		Recently  bool
	}
	profilePhoto struct {
		Id    int64
		Small *File
		Big   *File
	}
	userType struct {
		Bot     bool
		Deleted bool
		Regular bool
		Unknown bool
	}
	UserFullInfo struct {
		IsBlocked                       bool
		CanBeCalled                     bool
		HasPrivateCalls                 bool
		NeedPhoneNumberPrivacyException bool
		Bio                             string
		ShareText                       string
		GroupInCommonCount              int
		BotInfo                         *botInfo
	}
	botInfo struct {
		Description string
		Commands    []*botCommand
	}
	botCommand struct {
		Command     string
		Description string
	}
	UserProfilePhoto struct {
		Id        int64
		AddedDate int
		Photos    []*photo
	}
)

func getProfilePhoto(value object) *profilePhoto {
	if CheckType(value, "profilePhoto") {
		return &profilePhoto{
			Id:    value.int64("id"),
			Small: getFile(value.Object("small")),
			Big:   getFile(value.Object("big")),
		}
	}
	return nil
}

func getUserStatus(value object) *userStatus {
	if value != nil {
		return &userStatus{
			Empty:     value.string("@type") == "userStatusEmpty",
			LastMonth: value.string("@type") == "userStatusLastMonth",
			LastWeek:  value.string("@type") == "userStatusLastWeek",
			Offline:   value.string("@type") == "userStatusOffline",
			Online:    value.string("@type") == "userStatusOnline",
			Recently:  value.string("@type") == "userStatusRecently",
		}
	}
	return nil
}

func getUserType(value object) *userType {
	if value != nil {
		return &userType{
			Bot:     value.string("@type") == "userTypeBot",
			Deleted: value.string("@type") == "userTypeDeleted",
			Regular: value.string("@type") == "userTypeRegular",
			Unknown: value.string("@type") == "userTypeUnknown",
		}
	}
	return nil
}

func getUser(value object) *User {
	if CheckType(value, "user") {
		return &User{
			Id:              value.int64("id"),
			FirstName:       value.string("first_name"),
			LastName:        value.string("last_name"),
			Username:        value.string("username"),
			PhoneNumber:     value.string("phone_number"),
			Status:          getUserStatus(value.Object("status")),
			ProfilePhoto:    getProfilePhoto(value.Object("profile_photo")),
			IsContact:       value.bool("is_contact"),
			IsMutualContact: value.bool("is_mutual_contact"),
			IsVerified:      value.bool("is_verified"),
			IsSupport:       value.bool("is_support"),
			HaveAccess:      value.bool("have_access"),
			Type:            getUserType(value.Object("type")),
			LanguageCode:    value.string("language_code"),
		}
	}
	return nil
}

func getBotCommand(value object) *botCommand {
	if CheckType(value, "botCommand") {
		return &botCommand{
			Command:     value.string("command"),
			Description: value.string("description"),
		}
	}
	return nil
}

func getBotInfo(value object) *botInfo {
	if CheckType(value, "botInfo") {
		var commandsData = value.Array("commands")
		var commands = make([]*botCommand, len(commandsData))
		commandsData.Objects(func(index int, value object) {
			commands[index] = getBotCommand(value)
		})
		return &botInfo{
			Description: value.string("description"),
			Commands:    commands,
		}
	}
	return nil
}

func getUserFullInfo(value object) *UserFullInfo {
	if CheckType(value, "userFullInfo") {
		return &UserFullInfo{
			IsBlocked:                       value.bool("is_blocked"),
			CanBeCalled:                     value.bool("can_be_called"),
			HasPrivateCalls:                 value.bool("has_private_calls"),
			NeedPhoneNumberPrivacyException: value.bool("need_phone_number_privacy_exception"),
			Bio:                             value.string("bio"),
			ShareText:                       value.string("share_text"),
			GroupInCommonCount:              value.int("group_in_common_count"),
			BotInfo:                         getBotInfo(value.Object("bot_info")),
		}
	}
	return nil
}

func getUserProfilePhoto(value object) *UserProfilePhoto {
	if CheckType(value, "userProfilePhoto") {
		var photos = value.Array("sizes")
		var userProfilePhoto = &UserProfilePhoto{
			Id:        value.int64("id"),
			AddedDate: value.int("added_date"),
			Photos:    make([]*photo, len(photos)),
		}
		photos.Objects(func(index int, value object) {
			userProfilePhoto.Photos[index] = getPhoto(value)
		})
		return userProfilePhoto
	}
	return nil
}
