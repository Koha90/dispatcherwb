// Package storage keep storages methods
package storage

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/koha90/telegram-bot-api/lib/e"
)

// Storage interface with methods
type Storage interface {
	Save(ctx context.Context, p *Page) error
	PickRandom(ctx context.Context, username string) (*Page, error)
	Remove(ctx context.Context, p *Page) error
	IsExists(ctx context.Context, p *Page) (bool, error)
}

// ErrNoSavedPages variable for errors in this package
var ErrNoSavedPages = errors.New("no saved pages")

// Page ...
type Page struct {
	URL      string
	Username string
}

// Hash ...
func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.Username); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
