package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Comment struct {
	ID       int
	Author   string
	Comments string
	ParentID int
}
