// Package telegram keep telegram clients and his entities
package telegram

// UpdatesResponse object checked recieve from telegram server
// if field "ok" is true, we have result
type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update object recieved from telegram server
type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

// IncomingMessage ...
type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}
