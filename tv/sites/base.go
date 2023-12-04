package sites

import (
	"net/url"
	"olive/tv/model"

	"github.com/violetpupil/gos/std/strings"
)

type base struct {
	cookie string
}

// setRoomID 解析直播间 id 并设置
func (b *base) setRoomID(show *model.Show) error {
	u, err := url.Parse(show.URL)
	if err != nil {
		return err
	}
	path := strings.TrimSuffix(u.Path, "/")
	show.ID = strings.SplitLast(path, "/")
	return nil
}
