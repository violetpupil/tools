package sites

import "olive/tv/model"

func init() {
	registerSite("douyin", &douyin{})
}

type douyin struct {
	base
}

// SimplifyRoomURL 简化直播间地址
func (s *douyin) SimplifyRoomURL(show *model.Show) error {
	err := s.setRoomID(show)
	if err != nil {
		return err
	}
	show.URL = "https://live.douyin.com/" + show.ID
	return nil
}

// Snap 抓取直播信息
func (s *douyin) Snap(show *model.Show) {}
