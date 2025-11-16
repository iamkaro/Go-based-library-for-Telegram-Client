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

func (client *Client) newCallback() {
	client.Callback = &callback{
		client: client,
	}
}

type (
	callback struct {
		client *Client
	}
)

func (it *callback) Data(chatId, messageId int64, data []byte) *callbackQueryAnswer {
	return it.getCallbackQueryAnswer(chatId, messageId, object{
		"@type": "callbackQueryPayloadData",
		"data":  data, // bytes
	})
}

func (it *callback) DataWithPassword(chatId, messageId int64, password string, data []byte) *callbackQueryAnswer {
	return it.getCallbackQueryAnswer(chatId, messageId, object{
		"@type":    "callbackQueryPayloadDataWithPassword",
		"password": password,
		"data":     data, // bytes
	})
}

func (it *callback) Game(chatId, messageId int64, gameShortName string) *callbackQueryAnswer {
	return it.getCallbackQueryAnswer(chatId, messageId, object{
		"@type":           "callbackQueryPayloadGame",
		"game_short_name": gameShortName,
	})
}

/*----------------------------------------/         getCallbackQueryAnswer         /-----------*/
func (it *callback) getCallbackQueryAnswer(chatId, messageId int64, payload object) *callbackQueryAnswer {
	var answer = object{}
	if it.client.Load(&answer, Object{
		"@type":      "getCallbackQueryAnswer",
		"chat_id":    chatId,
		"message_id": messageId,
		"payload":    payload,
	}) {
		return getCallbackQueryAnswer(answer)
	}
	return nil
}

/*----------------------------------------/         items         /-----------*/
type (
	callbackQueryAnswer struct {
		Text      string
		ShowAlert bool
		Url       string
	}
)

func getCallbackQueryAnswer(value object) *callbackQueryAnswer {
	if CheckType(value, "callbackQueryAnswer") {
		return &callbackQueryAnswer{
			Text:      value.string("text"),
			ShowAlert: value.bool("show_alert"),
			Url:       value.string("url"),
		}
	}
	return nil
}
