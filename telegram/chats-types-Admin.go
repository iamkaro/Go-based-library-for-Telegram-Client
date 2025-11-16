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

func getAdministrator(value object) *Administrator {
	if CheckType(value, "chatAdministrator") {
		return &Administrator{
			UserId:      value.int32("user_id"),
			CustomTitle: value.string("custom_title"),
			IsOwner:     value.bool("is_owner"),
		}
	}
	return nil
}

type (
	Administrator struct {
		UserId      int32
		CustomTitle string
		IsOwner     bool
	}
)
