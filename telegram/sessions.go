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

func (client *Client) newSessions() {
	client.Sessions = &sessions{
		client: client,
	}
}

type (
	sessions struct {
		client *Client
	}
)

func (it *sessions) GetAll() Sessions {
	var all = object{}
	if it.client.Load(&all, Object{"@type": "getActiveSessions"}) {
		return convertList(all.Array("sessions").object(), getSession)
	}
	return nil
}

func (it *sessions) Terminate(sessionId int64) {
	it.client.load(object{"@type": "terminateSession", "session_id": sessionId})
}

func (it *sessions) TerminateAllOther() {
	it.client.load(object{"@type": "terminateAllOtherSessions"})
}

/*----------------------------------------/         items         /-----------*/
type (
	Sessions []*Session
	Session  struct {
		Id                    int64
		IsCurrent             bool
		IsPasswordPending     bool
		ApiId                 int32
		ApplicationName       string
		ApplicationVersion    string
		IsOfficialApplication bool
		DeviceModel           string
		Platform              string
		SystemVersion         string
		LogInDate             int32
		LastActiveDate        int32
		Ip                    string
		Country               string
		Region                string
	}
)

func getSession(value object) *Session {
	if CheckType(value, "session") {
		return &Session{
			Id:                    value.int64("id"),
			IsCurrent:             value.bool("is_current"),
			IsPasswordPending:     value.bool("is_password_pending"),
			ApiId:                 value.int32("api_id"),
			ApplicationName:       value.string("application_name"),
			ApplicationVersion:    value.string("application_version"),
			IsOfficialApplication: value.bool("is_official_application"),
			DeviceModel:           value.string("device_model"),
			Platform:              value.string("platform"),
			SystemVersion:         value.string("system_version"),
			LogInDate:             value.int32("log_in_date"),
			LastActiveDate:        value.int32("last_active_date"),
			Ip:                    value.string("ip"),
			Country:               value.string("country"),
			Region:                value.string("region"),
		}
	}
	return nil
}
