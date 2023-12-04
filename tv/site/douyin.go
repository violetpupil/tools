package site

import "olive/tv/model"

type douyin struct {
	*base
}

// Snap 抓取直播信息
func (s *douyin) Snap(show *model.Show) {}
