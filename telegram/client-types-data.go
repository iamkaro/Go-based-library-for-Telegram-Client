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

import (
	"encoding/json"
	"errors"
	"reflect"
)

/*-------------------------|      data      |-------------------------*/
type data string

func (it data) extractTo(out interface{}) error {
	if reflect.ValueOf(out).IsValid() && (reflect.ValueOf(out).Type().Kind() == reflect.Ptr) {
		return json.Unmarshal([]byte(it), out)
	}
	return errors.New("please send a pointer")
}
