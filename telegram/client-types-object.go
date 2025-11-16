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

/*-------------------------|      object      |-------------------------*/
type (
	Object map[string]any
	object map[string]interface{}
)

func (it object) toJson() string               { return toJson(it) }
func (it object) Object(index string) object   { return get[map[string]interface{}](it[index], object{}) }
func (it object) Array(index string) array     { return get[[]interface{}](it[index], array{}) }
func (it object) string(index string) string   { return get(it[index], "") }
func (it object) bool(index string) bool       { return get(it[index], false) }
func (it object) float64(index string) float64 { return get(it[index], float64(0)) }
func (it object) float32(index string) float32 { return float32(it.float64(index)) }
func (it object) int(index string) int         { return int(it.float64(index)) }
func (it object) int8(index string) int8       { return int8(it.float64(index)) }
func (it object) int16(index string) int16     { return int16(it.float64(index)) }
func (it object) int32(index string) int32     { return int32(it.float64(index)) }
func (it object) int64(index string) int64     { return int64(it.float64(index)) }
func (it object) byte(index string) byte       { return byte(it.float64(index)) }
func (it object) uint(index string) uint       { return uint(it.float64(index)) }
func (it object) uint8(index string) uint8     { return uint8(it.float64(index)) }
func (it object) uint16(index string) uint16   { return uint16(it.float64(index)) }
func (it object) uint32(index string) uint32   { return uint32(it.float64(index)) }
func (it object) uint64(index string) uint64   { return uint64(it.float64(index)) }
