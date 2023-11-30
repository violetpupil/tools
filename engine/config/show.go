package config

import (
	"fmt"
	"olive/foundation/olivetv"
)

type ID string

// Show 直播间配置
type Show struct {
	// 信息
	Platform     string `json:"platform"`      // 网站 id
	RoomID       string `json:"room_id"`       // 房间 id
	StreamerName string `json:"streamer_name"` // 网站名
	ID           ID     `json:"show_id"`       // olive 直播间 id

	// 配置
	SaveDir   string `json:"save_dir"`   // 视频存放目录
	SplitRule string `json:"split_rule"` // 视频分段规则
	Proxy     string `json:"proxy"`      // 代理
	Parser    string `json:"parser"`     // 视频流处理
}

func NewShow(roomURL, proxy string) (*Show, error) {
	tv, err := olivetv.NewWithURL(roomURL)
	if err != nil {
		return nil, err
	}

	show := &Show{
		StreamerName: tv.SiteName,
		Platform:     tv.SiteID,
		RoomID:       tv.RoomID,
		Proxy:        proxy,
	}
	return show, nil
}

// CheckAndFix 检查字段，没有设置的话，设置初始值
func (s *Show) CheckAndFix(cfg *Config) {
	if s.ID == "" {
		s.ID = ID(fmt.Sprintf("%s-%s", s.Platform, s.RoomID))
	}
	if s.Parser == "" {
		s.Parser = "flv"
	}
	if s.SaveDir == "" {
		s.SaveDir = cfg.SaveDir
	}
}
