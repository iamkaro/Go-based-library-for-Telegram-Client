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

func (client *Client) load(request object) data {
	if client.on {
		var id = getID()
		request["@extra"] = id
		/*--------------*/
		client.loadsData[id] = &loadData{Data: ""}
		client.send(request)
		/*--------------*/
		var loops uint16 = 0
		for client.loadsData[id].Data == "" {
			loops += 1
			if loops > 800 {
				break
			}
			sleep(250)
		}
		defer delete(client.loadsData, id)
		/*--------------*/
		return client.loadsData[id].Data
	}
	return ""
}
