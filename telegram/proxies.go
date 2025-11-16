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

func (client *Client) newProxies() {
	client.Proxies = &proxies{
		client: client,
	}
}

type (
	proxies struct {
		client *Client
	}
)

func (it *proxies) GetAll() Proxies {
	var all = object{}
	if it.client.Load(&all, Object{"@type": "getProxies"}) {
		return convertList(all.Array("proxies").object(), getProxy)
	}
	return nil
}

func (it *proxies) AddMTProto(server string, port int, secret string, enable bool) *Proxy {
	return it.add(server, port, enable,
		object{"@type": "proxyTypeMtproto", "secret": secret})
}
func (it *proxies) AddSocks5(server string, port int, user, pass string, enable bool) *Proxy {
	return it.add(server, port, enable,
		object{"@type": "proxyTypeSocks5", "username": user, "password": pass})
}
func (it *proxies) AddHttp(server string, port int, user, pass string, httpOnly, enable bool) *Proxy {
	return it.add(server, port, enable,
		object{"@type": "proxyTypeHttp", "username": user, "password": pass, "http_only": httpOnly})
}

func (it *proxies) Disable() { it.client.load(object{"@type": "disableProxy"}) }
func (it *proxies) Enable(proxyId int64) {
	it.client.load(object{"@type": "enableProxy", "proxy_id": proxyId})
}
func (it *proxies) Remove(proxyId int64) {
	it.client.load(object{"@type": "removeProxy", "proxy_id": proxyId})
}

/*----------------------------------------/         add         /-----------*/
func (it *proxies) add(server string, port int, enable bool, Type object) *Proxy {
	var out = object{}
	if it.client.Load(&out, Object{
		"@type":  "addProxy",
		"server": server,
		"port":   port,
		"enable": enable,
		"type":   Type,
	}) {
		return getProxy(out)
	}
	return nil
}

/*----------------------------------------/         items         /-----------*/
type (
	Proxies []*Proxy
	Proxy   struct {
		Id           int64
		Server       string
		Port         int
		LastUsedDate int64
		IsEnabled    bool
		Type         *proxyType
	}
	proxyType struct {
		MTProto *proxyTypeMTProto
		Socks5  *proxyTypeSocks5
		Http    *proxyTypeHttp
	}
	proxyTypeMTProto struct{ Secret string }
	proxyTypeSocks5  struct{ Username, Password string }
	proxyTypeHttp    struct {
		Username, Password string
		HttpOnly           bool
	}
)

func getProxy(value object) *Proxy {
	if CheckType(value, "proxy") {
		return &Proxy{
			Id:           value.int64("id"),
			Server:       value.string("server"),
			Port:         value.int("port"),
			LastUsedDate: value.int64("last_used_date"),
			IsEnabled:    value.bool("is_enabled"),
			Type:         getProxyType(value.Object("type")),
		}
	}
	return nil
}

func getProxyType(value object) *proxyType {
	return &proxyType{
		MTProto: getProxyTypeMTProto(value),
		Socks5:  getProxyTypeSocks5(value),
		Http:    getProxyTypeHttp(value),
	}
}

func getProxyTypeMTProto(value object) *proxyTypeMTProto {
	if CheckType(value, "proxyTypeMTProto") {
		return &proxyTypeMTProto{
			Secret: value.string("secret"),
		}
	}
	return nil
}
func getProxyTypeSocks5(value object) *proxyTypeSocks5 {
	if CheckType(value, "proxyTypeSocks5") {
		return &proxyTypeSocks5{
			Username: value.string("username"),
			Password: value.string("password"),
		}
	}
	return nil
}
func getProxyTypeHttp(value object) *proxyTypeHttp {
	if CheckType(value, "proxyTypeHttp") {
		return &proxyTypeHttp{
			Username: value.string("username"),
			Password: value.string("password"),
			HttpOnly: value.bool("http_only"),
		}
	}
	return nil
}
