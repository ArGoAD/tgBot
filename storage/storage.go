package storage

import (
	e "awesomeProject/lib/error"
	"crypto/sha1"
	"fmt"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	isExcist(p *Page) (bool, error)
}

// the basic type of date, that which working Storage interface

type Page struct {
	URL      string
	UserName string
}

//Generate hash name for page

func (p Page) Hash() (string, error) {

	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil

}
