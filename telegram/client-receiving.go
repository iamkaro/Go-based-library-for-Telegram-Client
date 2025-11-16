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

import "fmt"

func (client *Client) receiving() {
	client.isOn = true
	println("receiving start now...")
	var (
		response      data
		err           *errorClass
		responseClass *receivingClass
		loaded        *loadData
	)
	for (client.api != nil) && client.on {
		if response = client.receive(); response != "" {
			responseClass = &receivingClass{
				Type:  "",
				Extra: "",
			}
			if response.extractTo(responseClass) == nil {
				/*-------------*/
				loaded = client.loadsData[responseClass.Extra]
				if loaded != nil {
					loaded.Data = response
				}
				/*-------------*/
				switch responseClass.Type {
				case "error":
					if err = nil; (response.extractTo(&err) == nil) && (err != nil) {
						log("error", fmt.Sprint("(code: ", err.Code, ")  ", err.Message))
					} else {
						log("extract", response.extractTo(&err))
						log("error", response)
					}
				case "updateUser":
					client.Self.updateMe(response)
				case "updateMessageSendSucceeded":
					client.Messages.Send.update(response)
				case "updateMessageSendFailed":
					client.Messages.Send.update(response)
				}
				/*-------------*/
				if client.printAllReceivingData {
					println("\nReceived: ", response)
				}
			}
		}
		sleep(300)
	}
	client.isOn = false
	println("receiving exit...")
}
