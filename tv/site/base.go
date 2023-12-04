package site

import "github.com/violetpupil/gos/std/strings"

type base struct{}

func (b *base) RoomID(roomURL string) string {
	return strings.SplitLast(roomURL, "/")
}
