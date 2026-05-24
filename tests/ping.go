// -- ping/ping.go --
// Пакет ping проверяет доступность URL.
package ping

import "net/http"

// Pinger проверяет доступность URL.
type Pinger struct {
	client *http.Client
}

// Ping запрашивает указанный URL.
// Возвращает true, если адрес доступен, и false в противном случае.
func (p Pinger) Ping(url string) bool {
	resp, err := p.client.Head(url)
	if err != nil {
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	return true
}
