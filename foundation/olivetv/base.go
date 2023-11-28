package olivetv

import (
	"github.com/violetpupil/gos/std/strings"
)

type base struct{}

func (b *base) Name() string {
	return "undefined"
}

func (b *base) RoomID(roomURL RoomURL) string {
	return strings.SplitLast(string(roomURL), "/")
}
