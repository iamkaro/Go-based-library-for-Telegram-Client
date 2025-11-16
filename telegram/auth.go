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

func (client *Client) newAuth() {
	client.Auth = &auth{
		client: client,
	}
}

type (
	auth struct {
		client *Client
	}
)

func (auth *auth) State() *authState {
	var state = object{}
	if auth.client.Load(&state, Object{"@type": "getAuthorizationState"}) {
		return getAuthState(state)
	}
	return nil
}

func (auth *auth) EnterPhone(phoneNumber string) {
	auth.client.load(object{
		"@type":        "setAuthenticationPhoneNumber",
		"phone_number": phoneNumber,
		"settings": object{
			"@type":                            "phoneNumberAuthenticationSettings",
			"allow_flash_call":                 false,
			"allow_missed_call":                false,
			"is_current_phone_number":          false,
			"has_unknown_phone_number":         false,
			"allow_sms_retriever_api":          false,
			"firebase_authentication_settings": nil,
			"authentication_tokens":            array{},
		},
	})
}

func (auth *auth) EnterCode(code string) {
	auth.client.load(object{
		"@type": "checkAuthenticationCode",
		"code":  code,
	})
}

func (auth *auth) EnterPassword(password string) {
	auth.client.load(object{
		"@type":    "checkAuthenticationPassword",
		"password": password,
	})
}

func (auth *auth) RegistrationUser(firstName string, lastName string) {
	auth.client.load(object{
		"@type":                "registerUser",
		"first_name":           firstName,
		"last_name":            lastName,
		"disable_notification": false,
	})
}

/*----------------------------------------/         items         /-----------*/
type (
	authPassword struct {
		Hint     string
		HasEmail bool
	}
	authState struct {
		Closed           bool
		EnterPhone       bool
		EnterCode        bool
		EnterPassword    *authPassword
		RegistrationUser bool
		OK               bool
	}
)

func getAuthState(value object) *authState {
	if value != nil {
		var state = &authState{
			Closed:           false,
			EnterPhone:       false,
			EnterCode:        false,
			EnterPassword:    nil,
			RegistrationUser: false,
			OK:               false,
		}
		switch value.string("@type") {
		case "authorizationStateWaitPhoneNumber":
			state.EnterPhone = true
		case "authorizationStateWaitCode":
			state.EnterCode = true
		case "authorizationStateWaitPassword":
			state.EnterPassword = &authPassword{
				Hint:     value.string("password_hint"),
				HasEmail: value.bool("has_recovery_email_address"),
			}
		case "authorizationStateWaitRegistration":
			state.RegistrationUser = true
		case "authorizationStateReady":
			state.OK = true
		default:
			/*
				authorizationStateWaitTdlibParameters
				authorizationStateWaitEncryptionKey
				authorizationStateClosed
				authorizationStateClosing
				authorizationStateLoggingOut
				authorizationStateWaitOtherDeviceConfirmation
			*/

			log("getAuthState", value)

			state.Closed = true
		}
		return state
	}
	return nil
}
