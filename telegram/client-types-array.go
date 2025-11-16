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

/*-------------------------|      array      |-------------------------*/
type array []interface{}

func (it array) toJson() string     { return toJson(it) }
func (it array) object() []object   { return parseList[object](it, object{}) }
func (it array) string() []string   { return parseList[string](it, "") }
func (it array) bool() []bool       { return parseList[bool](it, false) }
func (it array) float64() []float64 { return parseList[float64](it, 0) }
func (it array) float32() []float32 { return convert[float32](it.float64()) }
func (it array) int() []int         { return convert[int](it.float64()) }
func (it array) int8() []int8       { return convert[int8](it.float64()) }
func (it array) int16() []int16     { return convert[int16](it.float64()) }
func (it array) int32() []int32     { return convert[int32](it.float64()) }
func (it array) int64() []int64     { return convert[int64](it.float64()) }
func (it array) byte() []byte       { return convert[byte](it.float64()) }
func (it array) uint() []uint       { return convert[uint](it.float64()) }
func (it array) uint8() []uint8     { return convert[uint8](it.float64()) }
func (it array) uint16() []uint16   { return convert[uint16](it.float64()) }
func (it array) uint32() []uint32   { return convert[uint32](it.float64()) }
func (it array) uint64() []uint64   { return convert[uint64](it.float64()) }
func (it array) Objects(handle func(index int, value object)) {
	for i := 0; i < len(it); i++ {
		handle(i, get(it[i], object{}))
	}
}
