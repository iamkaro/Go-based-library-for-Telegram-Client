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

func (client *Client) newSelf() {
	client.Self = &self{
		client: client,
		me:     nil,
	}
}

type (
	self struct {
		client *Client
		me     *User
	}
)

func (Self *self) GetMe() *User {
	if Self.me == nil {
		var me = object{}
		if Self.client.Load(&me, Object{"@type": "getMe"}) {
			Self.me = getUser(me)
		}
	}
	return Self.me
}

func (Self *self) updateMe(data data) {
	if Self.me != nil {
		var updateUser = object{}
		if data.extractTo(&updateUser) == nil {
			var me = getUser(updateUser.Object("user"))
			if (me != nil) && (Self.me.Id == me.Id) {
				Self.me = me
			}
		}
	}
}

func (Self *self) SetUsername(username string) bool {
	var (
		me    = Self.GetMe()
		check = &class{Type: ""}
		ok    = false
	)
	if me != nil {
		if Self.client.Load(check, Object{
			"@type":    "checkChatUsername",
			"chat_id":  me.Id,
			"username": username,
		}) {
			if check.Type == "checkChatUsernameResultOk" {
				Self.client.load(object{
					"@type":    "setUsername",
					"username": username,
				})
				ok = true
			}
		}
	}
	return ok
}

func (Self *self) SetName(firstName string, lastName string) {
	Self.client.load(object{
		"@type":      "setName",
		"first_name": firstName,
		"last_name":  lastName,
	})
}

func (Self *self) SetProfilePhoto(path string) {
	Self.client.load(object{
		"@type": "setProfilePhoto",
		"photo": object{
			"@type": "inputFileLocal",
			"path":  path,
		},
	})
}

func (Self *self) DeleteProfilePhoto(profilePhotoId int64) {
	Self.client.load(object{
		"@type":            "deleteProfilePhoto",
		"profile_photo_id": profilePhotoId,
	})
}

func (Self *self) SetBio(bio string) {
	Self.client.load(object{
		"@type": "setBio",
		"bio":   bio,
	})
}

func (Self *self) LogOut() {
	Self.client.load(object{"@type": "logOut"})
}
