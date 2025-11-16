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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10) +
		"_" + strconv.FormatInt(int64(rand.Int()), 10)
}

func sleep(millisecond time.Duration) {
	time.Sleep(millisecond * time.Millisecond)
}

func CheckType(value object, typeName string) bool {
	return (value != nil) && (value.string("@type") == typeName)
}

func log(name string, v interface{}) {
	fmt.Printf(name+" : %+v \n", v)
}

func get[T any](value interface{}, null T) T {
	if value != nil {
		if out, ok := value.(T); ok {
			return out
		}
	}
	return null
}

func convert[T numeric](list []float64) []T {
	var out = make([]T, len(list))
	for i := 0; i < len(list); i++ {
		out[i] = T(list[i])
	}
	return out
}

func parseList[T float64 | bool | string | object](list array, null T) []T {
	var out = make([]T, len(list))
	switch ptr := any(null); ptr.(type) {
	case object:
		for i := 0; i < len(list); i++ {
			out[i] = any(object(get[map[string]interface{}](list[i], ptr.(object)))).(T)
		}
		return out
	}
	for i := 0; i < len(list); i++ {
		out[i] = get(list[i], null)
	}
	return out
}

func convertList[T any](list []object, convert func(o object) T) []T {
	var out = make([]T, len(list))
	for i := 0; i < len(list); i++ {
		out[i] = convert(list[i])
	}
	return out
}

func toJson(value interface{}) string {
	out, err := json.Marshal(value)
	if err == nil {
		return string(out)
	}
	return ""
}

func decode(b64 string) string {
	var txt, _ = base64.StdEncoding.DecodeString(b64)
	return string(txt)
}
