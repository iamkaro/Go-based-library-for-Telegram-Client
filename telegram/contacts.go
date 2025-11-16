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

func (client *Client) newContacts() {
	client.Contacts = &contacts{
		client: client,
	}
}

type (
	contacts struct {
		client *Client
	}
)

func (it *contacts) GetAll() (usersID []int64) {
	var ids = object{}
	if it.client.Load(&ids, Object{"@type": "getContacts"}) {
		return ids.Array("user_ids").int64()
	}
	return nil
}

func (it *contacts) Add(phone string, name string, lastname string) {
	it.client.load((&Contact{Phone: phone, Name: name, LastName: lastname}).get())
}

func (it *contacts) Remove(usersID ...int64) {
	it.client.load(object{
		"@type":    "removeContacts",
		"user_ids": usersID,
	})
}

func (it *contacts) Import(contacts ...*Contact) (usersID []int64) {
	var ids = object{}
	if it.client.Load(&ids, Object{
		"@type":    "importContacts",
		"contacts": Contacts(contacts).get(),
	}) {
		return ids.Array("user_ids").int64()
	}
	return []int64{}
}

func (it *contacts) ImportedGetCount() int {
	var count = object{}
	if it.client.load(object{"@type": "getImportedContactCount"}).extractTo(&count) == nil {
		return count.int("count")
	}
	return -1
}

func (it *contacts) ImportedClearAll() {
	it.client.load(object{"@type": "clearImportedContacts"})
}

/*----------------------------------------/         items         /-----------*/
type (
	Contacts []*Contact
	Contact  struct {
		Phone, Name, LastName string
	}
)

func (it Contacts) get() []object {
	var out = make([]object, len(it))
	for i := 0; i < len(it); i++ {
		out[i] = it[i].get()
	}
	return out
}

func (it *Contact) get() object {
	return object{
		"@type": "addContact",
		"contact": object{
			"@type":        "contact",
			"phone_number": it.Phone,
			"first_name":   it.Name,
			"last_name":    it.LastName,
			"vcard":        "",
			"user_id":      0,
		},
		"share_phone_number": false,
	}
}
