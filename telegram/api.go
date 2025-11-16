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

/*//------------------|   static library   |------------------\\
#cgo linux CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -L/usr/local/lib -L/usr/local/lib/td -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm
#include <stdlib.h>
#include <td/telegram/td_json_client.h>
*/
// import "C"

/*//------------------|   dynamic library   |------------------\\
#cgo linux CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -L/usr/local/lib/td -ltdjson -lstdc++ -lssl -lcrypto -ldl -lz -lm
#cgo darwin CFLAGS: -I/usr/local/include
#cgo darwin LDFLAGS: -L/usr/local/lib -ltdjson -lstdc++ -lssl -lcrypto -ldl -lz -lm
#cgo windows CFLAGS: -Ic:/td -Ic:/td/example/csharp/build
#cgo windows LDFLAGS: -Lc:/td/example/csharp/build/Release -ltdjson
#include <stdlib.h>
#include <td/telegram/td_json_client.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const timeout = C.double(60.0)

func newApi() *api {
	return &api{
		client: C.td_json_client_create(),
	}
}

type (
	api struct {
		client unsafe.Pointer
	}
)

func (it *api) send(Query string) {
	query := C.CString(Query)
	defer C.free(unsafe.Pointer(query))
	C.td_json_client_send(it.client, query)
}

func (it *api) receive() string {
	result := C.td_json_client_receive(it.client, timeout)
	/** / defer C.free(unsafe.Pointer(result)) /**/
	if result == nil {
		fmt.Println("api: error, update receiving timeout.")
		return ""
	}
	return C.GoString(result)
}

func (it *api) execute(Query string) string {
	query := C.CString(Query)
	defer C.free(unsafe.Pointer(query))
	result := C.td_json_client_execute(it.client, query)
	/** / defer C.free(unsafe.Pointer(result)) /**/
	if result == nil {
		fmt.Println("api: error, request can't be parsed.")
		return ""
	}
	return C.GoString(result)
}

func (it *api) destroy() {
	C.td_json_client_destroy(it.client)
}
