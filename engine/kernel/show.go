package kernel

import (
	"fmt"
	"olive/engine/config"
)

// Show 直播间
type Show struct {
	// 信息
	Platform     string `json:"platform"`      // 网站 id
	RoomID       string `json:"room_id"`       // 房间 id
	StreamerName string `json:"streamer_name"` // 网站名
	ID           string `json:"show_id"`       // olive 直播间 id

	// 设置
	SaveDir   string `json:"save_dir"`   // 视频存放目录
	SplitRule string `json:"split_rule"` // 视频分段规则
}

// CheckAndFix 检查字段，没有设置的话，设置初始值
func (s *Show) CheckAndFix(cfg *config.Config) {
	if s.ID == "" {
		s.ID = fmt.Sprintf("%s-%s", s.Platform, s.RoomID)
	}
	if s.SaveDir == "" {
		s.SaveDir = cfg.SaveDir
	}
}