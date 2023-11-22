package olivetv

import "github.com/violetpupil/gos/std/strings"

func init() {
	registerSite("tiktok", &tiktok{})
}

type tiktok struct {
	base
}

func (this *tiktok) Name() string {
	return "tiktok"
}

func (this *tiktok) RoomID(roomURL RoomURL) string {
	s := strings.TrimSuffix(string(roomURL), "/live")
	s = strings.SplitLast(s, "@")
	return ""
}
