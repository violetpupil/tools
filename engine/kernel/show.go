package kernel

// Show 直播间
type Show struct {
	// 信息
	Platform     string `json:"platform"`      // 网站id
	RoomID       string `json:"room_id"`       // 房间id
	StreamerName string `json:"streamer_name"` // 网站名

	// 设置
	SaveDir   string `json:"save_dir"`   // 视频存放目录
	SplitRule string `json:"split_rule"` // 视频分段规则
}
