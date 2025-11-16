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

type (
	numeric interface {
		int | int64 | int32 | int16 | int8 |
			uint | uint64 | uint32 | uint16 | uint8 |
			float64 | float32
	}

	loadData   struct{ Data data }
	updateData struct {
		Update     object
		DeleteTime int64
		Next       *updateData
	}
	receivingClass struct {
		Type  string `json:"@type"`
		Extra string `json:"@extra"`
	}
	errorClass struct {
		Type    string `json:"@type"`
		Code    int64  `json:"code"`
		Message string `json:"message"`
	}
	class struct {
		Type string `json:"@type"`
	}
)
